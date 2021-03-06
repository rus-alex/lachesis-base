package packsdownloader

import (
	"sync"

	"github.com/Fantom-foundation/lachesis-base/gossip/fetcher"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
)

/*
 * PacksDownloader is a network agent, which is responsible for syncing events pack-by-pack.
 * It requests light pack infos with binary search, to find a lowest not connected pack.
 * Once lowest not connected pack is found, it requests full packs.
 * The full pack contains event hashes, which are re-directed to Fetcher.
 */

// PacksDownloader is responsible for accumulating pack announcements from various peers
// and scheduling them for retrieval.
type PacksDownloader struct {
	// Callbacks
	peerMisbehaviour PeerMisbehaviourFn
	fetcher          *fetcher.Fetcher
	onlyNotConnected OnlyNotConnectedFn

	config Config

	// State
	peers map[string]*PeerPacksDownloader

	peersMu    *sync.RWMutex
	terminated bool
}

type Config struct {
	MaxPeers int
}

// DefaultConfig returns default downloader config
func DefaultConfig() Config {
	return Config{6}
}

// DefaultConfig returns default downloader config for tests
func LiteConfig() Config {
	return Config{2}
}

// New creates a packs fetcher to retrieve events based on pack announcements.
func New(fetcher *fetcher.Fetcher, onlyNotConnected OnlyNotConnectedFn, peerMisbehaviour PeerMisbehaviourFn, config Config) *PacksDownloader {
	return &PacksDownloader{
		config:           config,
		fetcher:          fetcher,
		onlyNotConnected: onlyNotConnected,
		peerMisbehaviour: peerMisbehaviour,
		peers:            make(map[string]*PeerPacksDownloader),
		peersMu:          new(sync.RWMutex),
	}
}

type Peer struct {
	ID    string
	Epoch idx.Epoch

	RequestPackInfos RequestPackInfosFn
	RequestPack      RequestPackFn
}

// RegisterPeer injects a new download peer into the set of block source to be
// used for fetching hashes and blocks from.
func (d *PacksDownloader) RegisterPeer(peer Peer, myEpoch idx.Epoch) error {
	if peer.Epoch < myEpoch {
		// this peer is useless for syncing
		return d.UnregisterPeer(peer.ID)
	}

	d.peersMu.Lock()
	defer d.peersMu.Unlock()

	if d.terminated {
		return nil
	}

	if d.peers[peer.ID] != nil || len(d.peers) >= d.config.MaxPeers {
		return nil
	}

	d.peers[peer.ID] = newPeer(peer, myEpoch, d.fetcher, d.onlyNotConnected, d.peerMisbehaviour)
	d.peers[peer.ID].Start()

	return nil
}

func (d *PacksDownloader) OnNewEpoch(myEpoch idx.Epoch, peerEpoch func(string) idx.Epoch) {
	d.peersMu.Lock()
	defer d.peersMu.Unlock()

	newPeers := make(map[string]*PeerPacksDownloader)

	for peerID, peerDwnld := range d.peers {
		peerDwnld.Stop()

		if peerEpoch(peerID) >= myEpoch {
			// allocate new peer for the new epoch
			newPeerDwnld := newPeer(peerDwnld.peer, myEpoch, d.fetcher, d.onlyNotConnected, d.peerMisbehaviour)
			newPeerDwnld.Start()
			newPeers[peerID] = newPeerDwnld
		}
	}
	// wipe out old downloading state from prev. epoch
	d.peers = newPeers
}

func (d *PacksDownloader) Peer(peer string) *PeerPacksDownloader {
	d.peersMu.RLock()
	defer d.peersMu.RUnlock()

	return d.peers[peer]
}

func (d *PacksDownloader) PeersNum() int {
	d.peersMu.RLock()
	defer d.peersMu.RUnlock()

	return len(d.peers)
}

// UnregisterPeer removes a peer from the known list, preventing any action from
// the specified peer. An effort is also made to return any pending fetches into
// the queue.
func (d *PacksDownloader) UnregisterPeer(peer string) error {
	d.peersMu.Lock()
	defer d.peersMu.Unlock()

	if d.peers[peer] == nil {
		return nil
	}

	d.peers[peer].Stop()
	delete(d.peers, peer)
	return nil
}

// Terminate interrupts the downloader, canceling all pending operations.
// The downloader cannot be reused after calling Terminate.
func (d *PacksDownloader) Terminate() {
	d.peersMu.Lock()
	defer d.peersMu.Unlock()

	d.terminated = true
	for _, peerDownloader := range d.peers {
		peerDownloader.Stop()
	}
	d.peers = make(map[string]*PeerPacksDownloader)
}

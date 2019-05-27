package proxy

import (
	"context"
	"errors"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	"github.com/Fantom-foundation/go-lachesis/src/hash"
	"github.com/Fantom-foundation/go-lachesis/src/inter"
	"github.com/Fantom-foundation/go-lachesis/src/proxy/internal"
)

const (
	commandTimeout = 3 * time.Second
)

// grpcNodeProxy implements NodeProxy interface.
type grpcNodeProxy struct {
	conn   *grpc.ClientConn
	client internal.NodeClient
	logger *logrus.Logger
}

// NewGrpcNodeProxy initiates a NodeProxy-interface connected to remote node.
func NewGrpcNodeProxy(addr string, logger *logrus.Logger, opts ...grpc.DialOption) (NodeProxy, error) {
	if logger == nil {
		logger = logrus.New()
		logger.Level = logrus.DebugLevel
	}

	p := &grpcNodeProxy{
		logger: logger,
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	var err error
	p.conn, err = grpc.DialContext(ctx, addr,
		append(opts, grpc.WithInsecure(), grpc.WithBlock())...)
	if err != nil {
		return nil, err
	}

	p.client = internal.NewNodeClient(p.conn)

	return p, nil
}

/*
 * NodeProxy implementation:
 */

func (p *grpcNodeProxy) Close() {
	_ = p.conn.Close()
}

func (p *grpcNodeProxy) GetSelfID() (hash.Peer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), commandTimeout)
	defer cancel()

	resp, err := p.client.SelfID(ctx, &empty.Empty{})
	if err != nil {
		return hash.EmptyPeer, unwrapGrpcErr(err)
	}

	return hash.HexToPeer(resp.Hex), nil
}

func (p *grpcNodeProxy) GetBalanceOf(peer hash.Peer) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), commandTimeout)
	defer cancel()

	resp, err := p.client.BalanceOf(ctx, &internal.NodeID{
		Hex: peer.Hex(),
	})
	if err != nil {
		return 0, unwrapGrpcErr(err)
	}

	return resp.Amount, nil
}

func (p *grpcNodeProxy) SendTo(receiver hash.Peer, amount uint64) (hash.InternalTransaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), commandTimeout)
	defer cancel()

	req := internal.TransferRequest{
		Amount: amount,
		Receiver: &internal.NodeID{
			Hex: receiver.Hex(),
		},
	}

	resp, err := p.client.SendTo(ctx, &req)
	if err != nil {
		return hash.ZeroInternalTransaction, unwrapGrpcErr(err)
	}

	return hash.HexToInternalTransactionHash(resp.Hex), nil
}

func (p *grpcNodeProxy) GetTransaction(t hash.InternalTransaction) (*inter.InternalTransaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), commandTimeout)
	defer cancel()

	req := internal.TransactionRequest{
		Hex: t.Hex(),
	}

	resp, err := p.client.TransactionInfo(ctx, &req)
	if err != nil {
		return nil, unwrapGrpcErr(err)
	}

	return &inter.InternalTransaction{
		Amount:    resp.Amount,
		Confirmed: resp.Confirmed,
		Receiver:  hash.HexToPeer(resp.Receiver.Hex),
		Sender:    hash.HexToPeer(resp.Sender.Hex),
	}, nil
}

func unwrapGrpcErr(err error) error {
	st := status.Convert(err)
	return errors.New(st.Message())

}

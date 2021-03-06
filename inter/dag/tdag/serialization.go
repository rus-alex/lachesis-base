package tdag

import (
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/dag"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
)

type TestEventMarshaling struct {
	Epoch idx.Epoch
	Seq   idx.Event

	Frame  idx.Frame
	IsRoot bool

	Creator idx.ValidatorID

	Parents hash.Events

	Lamport idx.Lamport

	RawTime dag.RawTimestamp

	ID   hash.Event
	Name string
}

// EventToBytes serializes events
func (e *TestEvent) Bytes() []byte {
	b, _ := rlp.EncodeToBytes(&TestEventMarshaling{
		Epoch:   e.Epoch(),
		Seq:     e.Seq(),
		Frame:   e.Frame(),
		IsRoot:  e.IsRoot(),
		Creator: e.Creator(),
		Parents: e.Parents(),
		Lamport: e.Lamport(),
		RawTime: e.RawTime(),
		ID:      e.ID(),
		Name:    e.Name,
	})
	return b
}

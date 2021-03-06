package fetcher

import "github.com/Fantom-foundation/lachesis-base/inter/dag"

type LightCheck func(dag.Event) error

type HeavyCheck interface {
	Overloaded() bool
	Enqueue(events dag.Events, onValidated func(dag.Events, []error)) error
}

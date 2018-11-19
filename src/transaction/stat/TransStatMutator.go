package stat

const (
	Init = iota
	Commit
	Success
	Rollback
	Cancel
)

type TransStatMutator struct {
	stat int
}

func newTransStatMutor() *TransStatMutator {
	mutator := new(TransStatMutator)
	mutator.stat = Init
	return mutator
}
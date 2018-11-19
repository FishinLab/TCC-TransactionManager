package stat

type TransStat struct {
	mutator *TransStatMutator
}

func NewTransStat() *TransStat {
	stat := new(TransStat)
	stat.mutator = newTransStatMutor()
	return stat
}


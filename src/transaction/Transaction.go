package transaction

import (
	"../util"
	"./stat"
)

type Transaction struct {
	transID 	*util.TransID
	transStat 	*stat.TransStat
}

func NewTransaction() *Transaction {
	trx := new(Transaction)
	trx.transID = util.Get()
	trx.transStat = stat.NewTransStat()
	return trx
}

func Try() {}

func Confirm() {}

func Cancel() {}

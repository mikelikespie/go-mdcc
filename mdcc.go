package mdcc

import (
	"time"
)

func Execute(timeout time.Duration, tx Tx) {

}

type Tx interface {
	DoOperations()
	DoOnFailure()
	DoOnAccept()
	DoOnCommit(success bool)
	DoFinally(success bool, timeout bool)
	DoFinallyRemote(success bool, timeout bool)
}

type TxHandler struct {
	Operations    func()
	OnFailure     func()
	OnAccept      func()
	OnCommit      func(success bool)
	Finally       func(success bool, timeout bool)
	FinallyRemote func(success bool, timeout bool)
}

func (tx *TxHandler) DoOperations() {
	tx.Operations()
}

func (tx *TxHandler) DoOnFailure() {
	tx.OnFailure()
}

func (tx *TxHandler) DoOnAccept() {
	tx.OnAccept()
}

func (tx *TxHandler) DoOnCommit(success bool) {
	tx.OnCommit(success)
}

func (tx *TxHandler) DoFinally(success bool, timeout bool) {
	tx.DoFinally(success, timeout)
}

func (tx *TxHandler) DoFinallyRemote(success bool, timeout bool) {
	tx.DoFinallyRemote(success, timeout)
}

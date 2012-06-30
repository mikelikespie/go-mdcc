package main

import (
	"log"
)

type Tx interface {
	Timeout() int64;

	DoExecute();
	DoOnFailure();
	DoOnAccept();
	DoOnCommit(success bool);
	DoFinally(success bool, timeout bool);
	DoFinallyRemote(success bool, timeout bool);
}

type Txn struct {
	Execute func();
	OnFailure func();
	OnAccept func();
	OnCommit func(success bool);
	Finally func(success bool, timeout bool);
	FinallyRemote func(success bool, timeout bool);
}

func main () {
	log.Printf("Hello")
}

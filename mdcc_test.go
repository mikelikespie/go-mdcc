package mdcc

import (
	"testing"
	"time"
)

func TestFoo(t *testing.T) {
	Execute(time.Millisecond*600, &TxHandler{
		Operations: func() {

		},
		OnFailure: func() {

		},
		OnAccept: func() {

		},
		OnCommit: func(success bool) {

		},
		Finally: func(success bool, timeout bool) {

		},
		FinallyRemote: func(success bool, timeout bool) {

		},
	})
}

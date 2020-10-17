package poker

import (
	"io"
)

type CLI struct {
	store PlayerStore
	in    io.Reader
}

func (c CLI) playPoker() {
	c.store.RecordWin("Chris")
}

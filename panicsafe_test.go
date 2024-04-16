package panicsafe

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecover(t *testing.T) {
	var err = Recover(func() error { return nil })
	assert.Nil(t, err)
	err = Recover(func() error { return io.EOF })
	assert.ErrorIs(t, err, io.EOF)
	err = Recover(func() error { panic("lolo") })
	var perr *PanicError
	assert.ErrorAs(t, err, &perr)
}

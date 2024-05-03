package xerr_test

import (
	"errors"
	"testing"

	"github.com/lukaproject/xerr"
)

func TestAssert(t *testing.T) {
	t.Run("assert false", func(t *testing.T) {
		var err error = nil
		errormsg := "no equal"
		func() {
			defer xerr.Recover(&err)
			xerr.Assert(1 == 0, errors.New(errormsg))
		}()
		if err.Error() != errormsg {
			t.Fail()
		}
	})

	t.Run("assert true", func(t *testing.T) {
		var err error = nil
		errormsg := "no equal"
		func() {
			defer xerr.Recover(&err)
			xerr.Assert(1+1 == 2, errors.New(errormsg))
		}()
		if err != nil {
			t.Fail()
		}
	})
}

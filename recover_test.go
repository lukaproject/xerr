package xerr_test

import (
	"errors"
	"testing"

	"github.com/lukaproject/xerr"
)

func TestRecover(t *testing.T) {
	t.Run("test recover", func(t *testing.T) {
		var err error = nil
		func() {
			defer xerr.Recover(&err)
			xerr.Must0((errors.New(t.Name() + "1")))
			xerr.Must0((errors.New(t.Name() + "2")))
			xerr.Must0((errors.New(t.Name() + "3")))
		}()
		if err.Error() != t.Name()+"1" {
			t.Fatalf("%s is not equal with %s", err.Error(), t.Name()+"1")
		}
	})
}

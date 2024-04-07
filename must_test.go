package xerr_test

import (
	"errors"
	"testing"

	"github.com/lukaproject/xerr"
)

func TestMust(t *testing.T) {
	t.Run("test error panic ok", func(t *testing.T) {
		defer func() {
			if err := recover(); err == nil {
				t.Fatal(err)
			} else {
				t.Log(err)
			}
		}()
		hasErr := func() (int, error) {
			return 0, errors.New("test")
		}
		xerr.Must(hasErr())
	})
}

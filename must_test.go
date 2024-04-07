package xerr_test

import (
	"errors"
	"runtime/debug"
	"testing"

	"github.com/lukaproject/xerr"
)

func recoverT(t *testing.T) {
	if err := recover(); err == nil {
		t.Fatal(err)
	} else {
		t.Log(string(debug.Stack()))
		t.Log(err)
	}
}

func TestMust(t *testing.T) {
	t.Run("test error panic ok", func(t *testing.T) {
		defer recoverT(t)
		hasErr := func() (int, error) {
			return 0, errors.New("test")
		}
		xerr.Must(hasErr())
	})

	t.Run("test error no panic", func(t *testing.T) {
		hasErr := func() (int, error) {
			return 1, nil
		}
		v := xerr.Must(hasErr())
		if v != 1 {
			t.Fatalf("v != 1, v = %d", v)
		}
	})
}

func TestMust0(t *testing.T) {
	t.Run("test error panic ok", func(t *testing.T) {
		defer recoverT(t)
		hasErr := func() error {
			return errors.New("test")
		}
		xerr.Must0(hasErr())
	})
}

func TestMus2(t *testing.T) {
	t.Run("test error panic ok", func(t *testing.T) {
		defer recoverT(t)
		hasErr := func() (string, int, error) {
			return "", 1, errors.New("test")
		}
		xerr.Must2(hasErr())
	})

	t.Run("test error no panic", func(t *testing.T) {
		hasErr := func() (string, int, error) {
			return "123", 1, nil
		}
		str, val := xerr.Must2(hasErr())
		if str != "123" || val != 1 {
			t.Fatalf("str = [%s], val = [%d]", str, val)
		}
	})
}

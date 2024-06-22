package xerr_test

import (
	"errors"
	"runtime/debug"
	"sync"
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
		noErr := func() (int, error) {
			return 1, nil
		}
		v := xerr.Must(noErr())
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
	t.Run("test error no panic", func(t *testing.T) {
		hasErr := func() error {
			return nil
		}
		xerr.Must0(hasErr())
	})
}

func TestMust2(t *testing.T) {
	t.Run("test error panic ok", func(t *testing.T) {
		defer recoverT(t)
		hasErr := func() (string, int, error) {
			return "", 1, errors.New("test")
		}
		xerr.Must2(hasErr())
	})

	t.Run("test error no panic", func(t *testing.T) {
		noErr := func() (string, int, error) {
			return "123", 1, nil
		}
		str, val := xerr.Must2(noErr())
		if str != "123" || val != 1 {
			t.Fatalf("str = [%s], val = [%d]", str, val)
		}
	})
}

func TestMustOk(t *testing.T) {
	t.Run("test case ok", func(t *testing.T) {
		mp := &sync.Map{}
		mp.Store("1", 1)
		mp.Store("2", 2)
		mp.Store("3", 3)

		result := xerr.MustOk[int](mp.Load("1"))
		if result != 1 {
			t.Fatalf("result is not equal to 1")
		}
	})

	t.Run("test cast failed", func(t *testing.T) {
		defer func() {
			if err := recover(); err == nil {
				panic("expect cast failed")
			} else {
				if err.(error).Error() != "can't cast v[int]" {
					panic("unexpect error message")
				}
			}
		}()
		mp := &sync.Map{}
		mp.Store("1", 1)
		mp.Store("2", 2)
		mp.Store("3", 3)

		_ = xerr.MustOk[string](mp.Load("1"))
	})

	t.Run("test not ok", func(t *testing.T) {
		defer func() {
			if err := recover(); err == nil {
				panic("expect not ok failed")
			} else {
				if err.(error).Error() != "not ok" {
					panic("unexpect error message")
				}
			}
		}()
		mp := &sync.Map{}
		mp.Store("1", 1)
		mp.Store("2", 2)
		mp.Store("3", 3)
		_ = xerr.MustOk[int](mp.Load("4"))
	})
}

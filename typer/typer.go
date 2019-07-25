package typer

import (
	"errors"
	"runtime"
	"time"

	"github.com/micmonay/keybd_event"
)

type KeyPressCombo struct {
	pressShift bool
	KeyHex     int
}

type StringTyper struct {
	keybd_event.KeyBonding
	keys map[string]KeyPressCombo
}

func NewStringTyper(keys map[string]KeyPressCombo) (*StringTyper, error) {
	if keys == nil {
		return nil, errors.New("key mapping needed try typer.KeysOSXus")
	}
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}
	var err error
	st := &StringTyper{keys: keys}
	st.KeyBonding, err = keybd_event.NewKeyBonding()
	if err != nil {
		return nil, err
	}
	return st, nil
}

func (kt *StringTyper) TypeString(str string) error {
	for _, letter := range str {
		keyPressCombo, ok := kt.keys[string(letter)]
		if !ok {
			return errors.New(string(letter) + " not found in mapping")
		}
		kt.HasSHIFT(keyPressCombo.pressShift)
		kt.SetKeys(keyPressCombo.KeyHex)
		err := kt.Launching()
		if err != nil {
			return err
		}
	}
	return nil
}

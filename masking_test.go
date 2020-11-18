package masking

import (
	"encoding/json"
	"testing"
)

func TestCreditCard(t *testing.T) {
	var given CreditCard = "1234 5678 9012 3456"
	want := `"**** **** **** 3456"`

	get, err := json.Marshal(given)
	if err != nil {
		t.Error(err)
		return
	}

	if want != string(get) {
		t.Errorf("given %q want %q but got %q\n", given, want, get)
	}
}

func TestNationalID(t *testing.T) {
	var given CreditCard = "1234567890123"
	want := `"*********0123"`

	get, err := json.Marshal(given)
	if err != nil {
		t.Error(err)
		return
	}

	if want != string(get) {
		t.Errorf("given %q want %q but got %q\n", given, want, get)
	}
}

func TestEmailAddress(t *testing.T) {
	var given EmailAddress = "abc.def@gmail.com"
	want := `"***%***@gmail.com"`

	get, err := json.Marshal(given)
	if err != nil {
		t.Error(err)
		return
	}

	if want != string(get) {
		t.Errorf("given %q want %q but got %q\n", given, want, get)
	}
}

package masking

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCreditCard(t *testing.T) {
	var s CreditCard = "1234 5678 9012 3456"
	b, err := json.Marshal(s)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(b))
}

func TestNationalID(t *testing.T) {
	var s CreditCard = "1234567890123"
	b, err := json.Marshal(s)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(b))
}

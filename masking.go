package masking

import (
	"encoding/json"
	"strings"
)

type CreditCard string

func (c CreditCard) MarshalJSON() ([]byte, error) {
	r := strings.NewReplacer("0", "*", "1", "*", "2", "*", "3", "*", "4", "*", "5", "*", "6", "*", "7", "*", "8", "*", "9", "*")
	x := len(c) - 4

	return json.Marshal(r.Replace(string(c)[:x]) + string(c)[x:])
}

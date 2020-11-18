package masking

import (
	"encoding/json"
	"strings"
)

var creditCardReplacer = strings.NewReplacer("0", "*", "1", "*", "2", "*", "3", "*", "4", "*", "5", "*", "6", "*", "7", "*", "8", "*", "9", "*")

type CreditCard string

func (c CreditCard) MarshalJSON() ([]byte, error) {
	x := len(c) - 4

	return json.Marshal(creditCardReplacer.Replace(string(c)[:x]) + string(c)[x:])
}

var nationalIDReplacer = strings.NewReplacer("0", "*", "1", "*", "2", "*", "3", "*", "4", "*", "5", "*", "6", "*", "7", "*", "8", "*", "9", "*")

type NationalID string

func (id NationalID) MarshalJSON() ([]byte, error) {
	x := len(id) - 4

	return json.Marshal(nationalIDReplacer.Replace(string(id)[:x]) + string(id)[x:])
}

type EmailAddress string

var emailReplacer = strings.NewReplacer(
	"a", "*",
	"b", "*",
	"c", "*",
	"d", "*",
	"e", "*",
	"f", "*",
	"g", "*",
	"h", "*",
	"i", "*",
	"j", "*",
	"k", "*",
	"l", "*",
	"m", "*",
	"n", "*",
	"o", "*",
	"p", "*",
	"q", "*",
	"r", "*",
	"s", "*",
	"t", "*",
	"u", "*",
	"v", "*",
	"w", "*",
	"x", "*",
	"y", "*",
	"z", "*",
	"A", "*",
	"B", "*",
	"C", "*",
	"D", "*",
	"E", "*",
	"F", "*",
	"G", "*",
	"H", "*",
	"I", "*",
	"J", "*",
	"K", "*",
	"L", "*",
	"M", "*",
	"N", "*",
	"O", "*",
	"P", "*",
	"Q", "*",
	"R", "*",
	"S", "*",
	"T", "*",
	"U", "*",
	"V", "*",
	"W", "*",
	"X", "*",
	"Y", "*",
	"Z", "*",
	".", "%",
	"_", "%",
)

func (e EmailAddress) MarshalJSON() ([]byte, error) {
	x := strings.Index(string(e), "@")

	return json.Marshal(emailReplacer.Replace(string(e)[:x]) + string(e)[x:])
}

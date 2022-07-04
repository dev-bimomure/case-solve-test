package ConCurrency

import (
	"fmt"
	"testing"
)

func TestConvertCurrency(t *testing.T) {
	url := "https://api.frankfurter.app"

	jsonString := `[
		{"amount":15000.0, "currency":"IDR"},
		{"amount":3.1, "currency":"EUR"}
	]`
	jsonData := []byte(jsonString)

	result := convertCurrency(jsonData, url)
	fmt.Println(result)
}

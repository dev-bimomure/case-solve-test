package ConCurrency

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Valas struct {
	amount   float64
	currency string
	rate     map[string]interface{}
}

func convertCurrency(data []byte, URL string) []float64 {
	var tmp1 []interface{}
	var tmp2 map[string]interface{}
	var result []float64

	err := json.Unmarshal(data, &tmp1)
	if err != nil {
		panic(err)
	}
	for _, item := range tmp1 {
		val1 := getData(item)
		curr := getCurrency(val1.amount, val1.currency, URL)

		err := json.Unmarshal(curr, &tmp2)
		if err != nil {
			panic(err)
		}
		val2 := getData(tmp2)
		newFloat := getData(val2.rate)

		result = append(result, newFloat.amount)
	}

	return result
}

func getData(data interface{}) Valas {
	m := data.(map[string]interface{})
	valas := Valas{}
	if currency, ok := m["currency"].(string); ok {
		valas.currency = currency
	}
	if amount, ok := m["amount"].(float64); ok {
		valas.amount = amount
	}
	if rates, ok := m["rates"].(map[string]interface{}); ok {
		valas.rate = rates
	}
	if USD, ok := m["USD"].(float64); ok {
		valas.amount = USD
	}
	return valas
}

func getCurrency(amount float64, currency string, URL string) []byte {

	url := fmt.Sprintf("%s/latest?amount=%s&from=%s&to=USD", URL, strconv.FormatFloat(amount, 'f', 6, 64), currency)

	req, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	responseBody, err := io.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	result := responseBody

	return result
}

package currency

import "fmt"

func MapExample() {
	currencyCode := map[string]string{
		"USD": "US Dollar",
		"GBP": "Pound Sterling",
		"EUR": "Euro",
	}
	currency := "USD"
	currencyName := currencyCode[currency]
	fmt.Println("Currency name for currency code", currency, "is", currencyName)
}

func KeyExists() {
	currencyCode := map[string]string{
		"USD": "US Dollar",
		"GBP": "Pound Sterling",
		"EUR": "Euro",
	}

	currency := "INR"

	ccy, ok := currencyCode[currency]

	if ok {
		fmt.Println("Currency name for currency code", currency, "is", ccy)
	} else {
		fmt.Println("Currency code", currency, "not found")
	}
}

// in golang order of iteration is not guaranteed
func IterateMap() {
	currencyCode := map[string]string{
		"USD": "US Dollar",
		"GBP": "Pound Sterling",
		"EUR": "Euro",
	}

	for k, v := range currencyCode {
		fmt.Println("Currency code", k, "is", v)
	}
}

// deleting a non-existent key does not throw an error
func DeleteMap() {
	currencyCode := map[string]string{
		"USD": "US Dollar",
		"GBP": "Pound Sterling",
		"EUR": "Euro",
	}

	delete(currencyCode, "USD")
}

type Currency struct {
	name   string
	symbol string
}

func MapWithStruct() {
	curUSD := Currency{
		name:   "US Dollar",
		symbol: "$",
	}

	curEUR := Currency{
		name:   "Euro",
		symbol: "€",
	}

	currencyCode := map[string]Currency{
		"USD": curUSD,
		"EUR": curEUR,
	}

	for cyCode, cyInfo := range currencyCode {
		fmt.Println("Currency code", cyCode, "is", cyInfo.name, "and symbol is", cyInfo.symbol)
	}
}

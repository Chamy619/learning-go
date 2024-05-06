// Package money provides various utilities to make it easy to manage money.
package money

import "github.com/shopspring/decimal"

// Money representes the combination of an amount of money
// and the currency the money is in.
type Money struct {
	Value    decimal.Decimal
	Currency string
}

// Convert converts the value of one currency to another.
//
// It has two parameters: a Money instance with the value to convert,
// and a string that represents the currency to convert to. Convert returns
// the converted currency and any errors encountered from unknown or unconvertible
// currencies.
// If an error is returned, the Money instance is set to the zero value.
//
// Supported currencies are:
//
//	USD - US Dollar
//	CAD - Canadian Dollar
//	EUR - Euro
//	INR - Indian Rupee
//
// More information on exchange rates can be found
// at https://www.investopedia.com/terms/e/exchangerated.asp
func Convert(from Money, to string) (Money, error) {
	// 예제
	return from, nil
}

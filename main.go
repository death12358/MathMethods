package main

import (
	"fmt"

	"github.com/death12358/math-method/dot"
	decimalconversion "github.com/death12358/text-conversion/decimalconversion"
	"github.com/shopspring/decimal"
)

func Test() {
	decimalconversion.Float64([]decimal.Decimal{
		decimal.NewFromFloat(1),
		decimal.NewFromFloat(10)})
}
func main() {
	Test()
	fmt.Print("hello world")
	dot.Dot()
}

package task2

import (
	"fmt"
	"github.com/SpecialBrands/goexpression"
	"github.com/dhowden/numerus"
	"log"
	"regexp"
	"strings"
)

// RomanNumeralCalculate takes "XXIV + XI" and returns "XXXV"
func RomanNumeralCalculate(in string) string {
	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		log.Println(err)
		return "(parse error)"
	}
	sum := reg.ReplaceAllString(in, " ")
	roman := strings.Split(sum, " ")

	for _, num := range roman {
		iv, _ := numerus.Parse(num)
		in = strings.Replace(in, num, fmt.Sprintf("%d", iv), 1)
	}

	ans := goexpression.Eval(in, nil)

	return numerus.Numeral(ans).String()
}

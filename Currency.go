//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
package libcore
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
import (
	"fmt"
)
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
type Currency struct {
	Name                string
	Symbol              string
	CountDigitAfterDot  int
	CountCoinOfCurrency int
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
/**
 * @author Alexey Potehin <gnuplanet@gmail.com>, http://www.gnuplanet.online/doc/cv
 * @brief  get currency
 * @param  currencyName name of currency (ISO4217)
 * @return currency currency
 * @return err error
 */
func GetCurrency(currencyName string) (currency Currency, err error) {
	currencyList := []Currency {
		Currency{ "USD", "$",  2, 100  },
		Currency{ "TND", "DT", 3, 1000 },
		Currency{ "RUB", "₽",  2, 100  },
		Currency{ "KZT", "₸",  2, 100  } }


	for i := 0; i < len(currencyList); i++ {
		if currencyList[i].Name == currencyName {
			currency = currencyList[i]
			return
		}
	}


	err = fmt.Errorf("unknown currency ISO4217 name")
	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
/**
 * @author Alexey Potehin <gnuplanet@gmail.com>, http://www.gnuplanet.online/doc/cv
 * @brief  format amount, example: 1 -> 0.01 for USD
 * @param  amount amount of coins, example: 100 is 1 dollar
 * @param  currencyName name of currency (ISO4217)
 * @param  dotSign separator, example: .
 * @return amountFormated formated amount
 * @return err error
 */
func FormatAmount(amount, currencyName, dotSign string) (amountFormated string, err error) {
	var currency Currency
	currency, err = GetCurrency(currencyName)
	if err != nil {
		return
	}


	if IsUint(amount) == false {
		err = fmt.Errorf("amount is invalid")
		return
	}


	amountFormated = amount
	for {
		if len(amountFormated) > currency.CountDigitAfterDot {
			break
		}
		amountFormated = "0" + amountFormated
	}


	part1 := SubStr(amountFormated, 0, len(amountFormated) - currency.CountDigitAfterDot)
	part2 := SubStr(amountFormated, len(amountFormated) - currency.CountDigitAfterDot, currency.CountDigitAfterDot)


	amountFormated = part1 + dotSign + part2


	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func FormatAmountTest() (err error) {
	var amount string
	var amountFormated string
	var currencyName string


	amount = "0"
	currencyName = "RUB"
	amountFormated, err = FormatAmount(amount, currencyName, ".")
	if err != nil {
		return
	}
	if amountFormated != "0.00" {
		err = fmt.Errorf("broken test001")
		return
	}
	fmt.Printf("\"%s\" %s -> %s\n", currencyName, amount, amountFormated)


	amount = "1"
	currencyName = "RUB"
	amountFormated, err = FormatAmount(amount, currencyName, ".")
	if err != nil {
		return
	}
	if amountFormated != "0.01" {
		err = fmt.Errorf("broken test002")
		return
	}
	fmt.Printf("\"%s\" %s -> %s\n", currencyName, amount, amountFormated)


	amount = "10"
	currencyName = "RUB"
	amountFormated, err = FormatAmount(amount, currencyName, ".")
	if err != nil {
		return
	}
	if amountFormated != "0.10" {
		err = fmt.Errorf("broken test003")
		return
	}
	fmt.Printf("\"%s\" %s -> %s\n", currencyName, amount, amountFormated)


	amount = "100"
	currencyName = "RUB"
	amountFormated, err = FormatAmount(amount, currencyName, ".")
	if err != nil {
		return
	}
	if amountFormated != "1.00" {
		err = fmt.Errorf("broken test004")
		return
	}
	fmt.Printf("\"%s\" %s -> %s\n", currencyName, amount, amountFormated)


	amount = "1000"
	currencyName = "RUB"
	amountFormated, err = FormatAmount(amount, currencyName, ".")
	if err != nil {
		return
	}
	if amountFormated != "10.00" {
		err = fmt.Errorf("broken test005")
		return
	}
	fmt.Printf("\"%s\" %s -> %s\n", currencyName, amount, amountFormated)


	amount = "10000"
	currencyName = "RUB"
	amountFormated, err = FormatAmount(amount, currencyName, ".")
	if err != nil {
		return
	}
	if amountFormated != "100.00" {
		err = fmt.Errorf("broken test006")
		return
	}
	fmt.Printf("\"%s\" %s -> %s\n", currencyName, amount, amountFormated)


	amount = "100000"
	currencyName = "RUB"
	amountFormated, err = FormatAmount(amount, currencyName, ".")
	if err != nil {
		return
	}
	if amountFormated != "1000.00" {
		err = fmt.Errorf("broken test007")
		return
	}
	fmt.Printf("\"%s\" %s -> %s\n", currencyName, amount, amountFormated)


	amount = "0"
	currencyName = "TND"
	amountFormated, err = FormatAmount(amount, currencyName, ".")
	if err != nil {
		return
	}
	if amountFormated != "0.000" {
		err = fmt.Errorf("broken test008")
		return
	}
	fmt.Printf("\"%s\" %s -> %s\n", currencyName, amount, amountFormated)


	amount = "1"
	currencyName = "TND"
	amountFormated, err = FormatAmount(amount, currencyName, ".")
	if err != nil {
		return
	}
	if amountFormated != "0.001" {
		err = fmt.Errorf("broken test009")
		return
	}
	fmt.Printf("\"%s\" %s -> %s\n", currencyName, amount, amountFormated)


	amount = "10"
	currencyName = "TND"
	amountFormated, err = FormatAmount(amount, currencyName, ".")
	if err != nil {
		return
	}
	if amountFormated != "0.010" {
		err = fmt.Errorf("broken test010")
		return
	}
	fmt.Printf("\"%s\" %s -> %s\n", currencyName, amount, amountFormated)


	amount = "100"
	currencyName = "TND"
	amountFormated, err = FormatAmount(amount, currencyName, ".")
	if err != nil {
		return
	}
	if amountFormated != "0.100" {
		err = fmt.Errorf("broken test011")
		return
	}
	fmt.Printf("\"%s\" %s -> %s\n", currencyName, amount, amountFormated)


	amount = "1000"
	currencyName = "TND"
	amountFormated, err = FormatAmount(amount, currencyName, ".")
	if err != nil {
		return
	}
	if amountFormated != "1.000" {
		err = fmt.Errorf("broken test012")
		return
	}
	fmt.Printf("\"%s\" %s -> %s\n", currencyName, amount, amountFormated)


	amount = "10000"
	currencyName = "TND"
	amountFormated, err = FormatAmount(amount, currencyName, ".")
	if err != nil {
		return
	}
	if amountFormated != "10.000" {
		err = fmt.Errorf("broken test013")
		return
	}
	fmt.Printf("\"%s\" %s -> %s\n", currencyName, amount, amountFormated)


	amount = "100000"
	currencyName = "TND"
	amountFormated, err = FormatAmount(amount, currencyName, ".")
	if err != nil {
		return
	}
	if amountFormated != "100.000" {
		err = fmt.Errorf("broken test014")
		return
	}
	fmt.Printf("\"%s\" %s -> %s\n", currencyName, amount, amountFormated)


	return
}
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
/*
func main() {
	var err error


	err = FormatAmountTest()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}


	fmt.Printf("done\n")
	os.Exit(0)
}
*/
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//

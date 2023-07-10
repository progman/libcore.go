//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
package libcore
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
import (
	"fmt"
)
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//
type Currency struct {
	Name                string
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
		Currency{ "USD", 2, 100 },
		Currency{ "TND", 3, 1000 },
		Currency{ "RUB", 2, 100 },
		Currency{ "KZT", 2, 100 } }

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

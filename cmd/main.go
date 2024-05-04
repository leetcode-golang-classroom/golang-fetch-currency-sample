package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/leetcode-golang-classroom/golang-fetch-currency-sample/internal/currency"
)

func main() {
	ce := &currency.MyCurrencyExchange{
		Currencies: make(map[string]currency.Currency),
	}

	err := ce.FetchAllCurrencies()
	if err != nil {
		panic(err)
	}
	wg := sync.WaitGroup{}
	startTime := time.Now()
	go func() {
		for {
			ce.Lock()
			usd, ok := ce.Currencies["usd"]
			ce.Unlock()
			if ok {
				fmt.Println("USD", usd.Rates)
			}
		}
	}()
	for code := range ce.Currencies {
		wg.Add(1)
		go func(code string) {
			rates, err := currency.FetchCurrencyRates(code)
			if err != nil {
				panic(err)
			}
			ce.Lock()
			ce.Currencies[code] = currency.Currency{
				Code:  code,
				Name:  ce.Currencies[code].Name,
				Rates: rates,
			}
			ce.Unlock()
			wg.Done()
		}(code)
	}
	wg.Wait()
	fmt.Println("==============Result ===========")

	for _, curr := range ce.Currencies {
		fmt.Printf("%s (%s): %d rates\n", curr.Name, curr.Code, len(curr.Rates))
	}
	fmt.Println("==============================")
	fmt.Println("Time taken:", time.Since(startTime))
}

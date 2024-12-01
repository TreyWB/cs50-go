package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/dustin/go-humanize"
)

type CoinDeskResponse struct {
	BPI struct {
		USD struct {
			Rate_Float float64 `json:"rate_float"`
		} `json:"USD"`
	} `json:"bpi"`
}

var errArg = fmt.Errorf("invalid argument")
var errApi = fmt.Errorf("failed to get current BTC price")

func main() {
	amount, err := getArgs()
	if err != nil {
		fmt.Println(err)
		return
	}

	url := "https://api.coindesk.com/v1/bpi/currentprice.json"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(errors.Join(errApi, fmt.Errorf("cannot access CoinDesk API")))
		return
	}
	defer resp.Body.Close()

	var result CoinDeskResponse

	decoder := json.NewDecoder(resp.Body)
	decode_err := decoder.Decode(&result)
	if decode_err != nil {
		fmt.Println(errors.Join(errApi, fmt.Errorf("failed to decode CoinDesk API response")))
		return
	}

	price := result.BPI.USD.Rate_Float
	output := price * amount

	fmt.Printf("$%s", humanize.CommafWithDigits(output, 2))
}

func getArgs() (float64, error) {
	if len(os.Args) != 2 {
		return 0, errors.Join(errArg, fmt.Errorf("missing command-line argument"))
	}

	input := os.Args[1]
	amount, _ := strconv.ParseFloat(input, 64)

	if amount <= 0 {
		return 0, errors.Join(errArg, fmt.Errorf("amount must be a positive number"))
	}

	return amount, nil
}

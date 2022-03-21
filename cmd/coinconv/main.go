package main

import (
	"fmt"
	"os"
	"strconv"

	"coinconv/internal/converter"
)

var help = "usage: coinconv <sum> <from> <to>\nwhere sum is number"

func main() {
	errCloseFunc := func(err error) {
		fmt.Println(err)
		fmt.Println(help)
		os.Exit(1)
	}

	args, err := parseArgs(os.Args)
	if err != nil {
		errCloseFunc(err)
	}

	srv := converter.NewService()
	res, err := srv.PriceConversion(args.Sum, args.FromCurrency, args.ToCurrency)
	if err != nil {
		errCloseFunc(err)
	}

	fmt.Println(res)
}

type Args struct {
	Sum          float64
	FromCurrency string
	ToCurrency   string
}

func parseArgs(args []string) (Args, error) {
	if len(args) != 4 {
		return Args{}, fmt.Errorf("invalid count arguments")
	}

	sum, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return Args{}, err
	}

	return Args{
		Sum:          sum,
		FromCurrency: args[2],
		ToCurrency:   args[3],
	}, nil
}

package main

import (
	"fmt"

	gohousekeeping "github.com/arigatosimarmata/go-housekeeping"
)

func main() {
	list_hk := []gohousekeeping.HousekeepingModel{
		{
			Directory: "./logs/",
			Filename:  "*.log",
		},
		{
			Directory: "./cache/202309*",
			Filename:  "",
		},
	}
	err := gohousekeeping.Housekeeping(list_hk, false)
	if err != nil {
		fmt.Println(err)
	}
}

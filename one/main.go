package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(0)
	}
	fmt.Println(time)
}

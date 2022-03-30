package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var decode bool
	flag.BoolVar(&decode, "d", false, "-d: transforms timestamp into date")
	flag.Parse()

	if !decode {
		printUnixSecs()
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		fmt.Println(time.Unix(num, 0))
	}
}

func printUnixSecs() {
	fmt.Printf("%d", time.Now().Unix())
}

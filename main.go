/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/alidevhere/fake-gen/cmd"
)

func main() {
	cmd.Execute()
	// ReadConv()
}

func ReadConv() {
	f, err := os.Open("males.csv")

	if err != nil {

		log.Fatal(err)
	}

	r := csv.NewReader(f)

	for {

		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		var s string
		for _, value := range record {

			k := strings.Split(value, " ")
			s += fmt.Sprintf(`,"%s","%s"`, k[0], k[1])
		}
		fmt.Printf("%s", s)
	}
}

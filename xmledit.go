package main

import (
	. "./library/ovf_struct"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}
	ovaManifest, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println(err)
	}

	defer ovaManifest.Close()

	byteValue, _ := ioutil.ReadAll(ovaManifest)

	var virtualSystem CVirtualSystem__ovf

	xml.Unmarshall(byteValue, &virtualSystem)

	for _, v := range virtualSystem {
		fmt.Println(v)
	}
}

func usage() {
	fmt.Println(os.Args[0], ": file to edit")
}

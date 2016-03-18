package main

import (
	"encoding/xml"
	"os"
	"fmt"
	"io/ioutil"
	"encoding/csv"
	"bufio"
)

type Resources struct {
	XMLName xml.Name	`xml:"resources"`
	Strings []StringRes	`xml:"string"`
}

type StringRes struct {
	Name string		`xml:"name,attr"`
	Value string		`xml:",chardata"`
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}
	var outputFile string
	inputFile := os.Args[1]
	if len(os.Args) > 2 {
		outputFile = os.Args[2]
	} else {
		outputFile = inputFile + ".csv"
	}
	fmt.Printf("Input: %s, output: %s\n", inputFile, outputFile)

	// read XML
	res := Resources{}
	xmlContent,_ := ioutil.ReadFile(inputFile)
	var err = xml.Unmarshal(xmlContent, &res)
	if err != nil { panic(err) }

	strings := res.Strings
	fmt.Printf("Found: %d string resources!\n", len(strings))

	// write CSV
	f, err := os.Create(outputFile)
	if err != nil { panic(err) }
	w := csv.NewWriter(bufio.NewWriter(f))
	for _, i := range strings {
		err = w.Write([]string{i.Name, i.Value})
		if err != nil { panic(err) }
	}

	w.Flush()

}

func printHelp()  {
	fmt.Println("Arguments:")
	fmt.Println("inputFile - the path to the Android String XML file")
	fmt.Println("outputFile - the path to the output file (optional). If missed than name of input will be used by appending .csv suffix")
}

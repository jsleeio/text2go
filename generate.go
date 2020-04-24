package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func generate(packagename, prefix string, raw []byte) string {
	preamble := fmt.Sprintf("package %s\n\nimport \"encoding/base64\"\n\nconst (\n\t%sEncodedText = `\n%s\n`\n)\n",
		packagename, prefix, base64.StdEncoding.EncodeToString(raw))
	decoder := fmt.Sprintf("func %sText() string {\n\tdata,_ := base64.StdEncoding.DecodeString(%sEncodedText)\n\treturn string(data)\n}\n", prefix, prefix)
	return preamble + decoder
}

func main() {
	var inFile, outFile, packageName, prefix string
	flag.StringVar(&inFile, "input", "", "input file to encode")
	flag.StringVar(&outFile, "output", "", "output .go file to write")
	flag.StringVar(&packageName, "package", "main", "Go package name to write text into")
	flag.StringVar(&prefix, "prefix", "", "Go name prefix for generated variables/functions, eg. package.{prefix}EncodedText and package.{prefix}Text()")
	flag.Parse()
	raw, err := ioutil.ReadFile(inFile)
	if err != nil {
		fmt.Printf("unable to read input file: %v\n", err)
		os.Exit(1)
	}
	if err = ioutil.WriteFile(outFile, []byte(generate(packageName, prefix, raw)), 0644); err != nil {
		fmt.Printf("unable to write output file: %v\n", err)
		os.Exit(1)
	}
}

# text2go

A trivial tool for embedded data into Golang binaries. Data is base64-encoded,
and a helper function generated to return the decoded form.

## tldr

```
$ ./text2go -input=Dockerfile -output=/dev/stdout -prefix=Dockerfile
package main

import "encoding/base64"

const (
	DockerfileEncodedText = `
RlJPTSBzY3JhdGNoCkNPUFkgdGV4dDJnbyAgL3RleHQyZ28KRU5UUllQT0lOVCBbIi90ZXh0MmdvIl0KYQ==
`
)
func DockerfileText() string {
	data,_ := base64.StdEncoding.DecodeString(DockerfileEncodedText)
	return string(data)
}
```

## options

```
$ ./text2go -help
Usage of ./text2go:
  -input string
    	input file to encode
  -output string
    	output .go file to write
  -package string
    	Go package name to write text into (default "main")
  -prefix string
    	Go name prefix for generated variables/functions, eg. package.{prefix}EncodedText and package.{prefix}Text()
```


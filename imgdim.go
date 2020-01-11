package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

var in string

func init() {
	flag.StringVar(&in, "in", "", "input file")
}

func main() {
	flag.Parse()

	reader := os.Stdin
	if in != "" {
		file, err := os.Open(in)
		if err != nil {
			fatal(err.Error())
		}
		reader = file
	}

	config, format, err := image.DecodeConfig(reader)
	if err != nil {
		fatal(err.Error())
	}

	fmt.Println("format\twidth\theight")
	fmt.Println(fmt.Sprintf("%s\t%d\t%d", format, config.Width, config.Height))
}

func fatal(a ...interface{}) {
	fmt.Fprintln(os.Stderr)
	os.Exit(1)
}

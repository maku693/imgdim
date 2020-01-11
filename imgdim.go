package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func main() {
	config, format, err := image.DecodeConfig(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Println("format\twidth\theight")
	fmt.Println(fmt.Sprintf("%s\t%d\t%d", format, config.Width, config.Height))
}

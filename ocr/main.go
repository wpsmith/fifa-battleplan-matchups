package main

import (
	"fmt"
	"github.com/otiai10/gosseract"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("/Users/travis.smith/Projects/GO/src/github.com/wpsmith/fifa-battleplan-matchups/ocr/IMG_6996.PNG")
	text, _ := client.Text()
	fmt.Println(text)
	// Hello, World!
}

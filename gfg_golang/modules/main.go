package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("Hi from Ayush")

	color.Cyan("Prints text in cyan.")

	// a newline will be appended automatically
	color.Blue("Prints %s in blue.", "text")

	// More default foreground colors..
	color.Red("We have red")
	color.Yellow("Yellow color too!")
	color.Magenta("And many others ..")

	// Hi-intensity colors
	color.HiGreen("Bright green color.")
	color.HiBlack("Bright black means gray..")
	color.HiWhite("Shiny white color!")
}

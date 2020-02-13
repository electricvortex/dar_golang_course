package eighth

import (
	"fmt"
	"github.com/fatih/color"
)

func main() {
	PrintHello()
}

func PrintHello() {
	fmt.Println("Hello World")
	color.Yellow("%v", "Yello color")
}


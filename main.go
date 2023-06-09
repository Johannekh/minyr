package main


import (

	

	
	"bufio"
	"fmt"
	"os"
"github.com/Johannekh/minyr/yr"
)


func main() {
var input string

scanner := bufio.NewScanner(os.Stdin)



for {
	fmt.Println("Velg convert, average eller exit")
	if !scanner.Scan() {
		break
	}
	input = scanner.Text()
	switch input {
	case "convert":
	yr.Konvertering()
	case "average":
	fmt.Println("Velg enhet for gjennomsnittstemperatur (c/f):")
	if !scanner.Scan() {
	break
	}

	enhet := scanner.Text()
	switch enhet {
	case "c":
	yr.GjennomsnittAvCelsius()
	case "f":
	yr.GjennomsnittAvFahr()
	default:
	fmt.Println("Ugyldig valg.")

	}

 	case "exit":
        fmt.Println("forvell:)")
        return
		}
	}
}

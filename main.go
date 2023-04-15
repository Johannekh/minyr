package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Johannekh/funtemps/conv"
)

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Venligst velg convert, average eller exit:")
		scanner.Scan()
		input = scanner.Text()
		if input == "exit" {
			fmt.Println("exit")
			os.Exit(0)
		} else if input == "convert" {

			_, err := os.Stat("kjevik-tempfahr-20220318-20230318.csv")
			if err == nil {
				src, err := os.Open("kjevik-tempfahr-20220318-20230318.csv")
					if err != nil {
    				log.Fatal(err)
					}
				defer src.Close()
				fmt.Println("Fil kjevik-tempfahr-20220318-20230318.csv finnes allerede. Vil du generere filen på nytt? (j/n)")
				scanner.Scan()
				input = scanner.Text()
				if input == "n" {
				  break
} else if input == "average" {
    var src *os.File
    fmt.Println("Beregner gjennomsnittstemperatur for hele perioden.")
    src, err := os.Open("kjevik-tempfahr-20220318-20230318.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer src.Close()

    var buffer []byte
    var linebuf []byte // nil
    buffer = make([]byte, 1)
    bytesCount := 0
    tempSum := 0.0
    for {
        _, err := src.Read(buffer)
        if err != nil && err != io.EOF {
            log.Fatal(err)
        }

        bytesCount++
        if buffer[0] == 0x0A {
            elementArray := strings.Split(string(linebuf), ";")
            if len(elementArray) > 3 {
                celsius := elementArray[3]
                celsiusFloat, err := strconv.ParseFloat(celsius, 64)
                if err != nil {
                    log.Fatal(err)
                }
                fahr := conv.CelsiusToFarhrenheit(celsiusFloat)
                fahrString := strconv.FormatFloat(fahr, 'f', -1, 64)
                temp, err := strconv.ParseFloat(fahrString, 64)
                if err != nil {
                    log.Fatal(err)
                }
                tempSum += temp
            }
            linebuf = nil
        } else {
            linebuf = append(linebuf, buffer[0])
        }
        if err == io.EOF {
            break
        }
    }

    fmt.Println("Gjennomsnittstemperaturen er:")
    for {
        fmt.Println("Vil du ha temperaturen i grader Celsius eller Fahrenheit? (c/f)")
        scanner.Scan()
        input = scanner.Text()
        if input == "c" {
            avg := tempSum / float64(bytesCount)
            fmt.Printf("%.2f grader Celsius\n", avg)
            break
        } else if input == "f" {
            avg := conv.FarhenheitToCelsius(tempSum / float64(bytesCount))
            fmt.Printf("%.2f grader Fahrenheit\n", avg)
            break
        } else {
            fmt.Println("Ugyldig valg. Prøv igjen.")
        }
    }
} 
}
}


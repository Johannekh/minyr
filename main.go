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
	"github.com/Johannekh/minyr/yr"

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
        } // Åpne filen for konvertering
		src, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
		if err != nil {
			log.Fatal(err)
		}
		defer src.Close()
		
		// Opprett den nye filen for konverterte data
		dst, err := os.Create("kjevik-temp-fahrenheit-20220318-20230318.csv")
		if err != nil {
			log.Fatal(err)
		}
		defer dst.Close()
		
		// Lag en bufio Scanner for å lese data fra kildelinjen
		scanner := bufio.NewScanner(src)
		
		// Lag en bufio Writer for å skrive konverterte data til destinasjonslinjen
		writer := bufio.NewWriter(dst)
		
		// Skanne kildelinjen og konverter data
		for scanner.Scan() {
			line := scanner.Text()
			convertedLine, err := yr.CelsiusToFahrenheitLine(line)
			if err != nil {
				log.Fatal(err)
			}
			_, err = fmt.Fprintln(writer, convertedLine)
			if err != nil {
				log.Fatal(err)
			}
		}
		
		// Flush bufferen for å sikre at alt er skrevet til destinasjonslinjen
		err = writer.Flush()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Konvertering fullført!")
		
            // funksjon som åpner fil, leser linjer, gjør endringer og lagrer nye linjer i en ny fil
            // flere else-if setninger
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
				var response string
				fmt.Println("Vil du ha temperaturen i grader Celsius eller Fahrenheit? (c/f)")
				scanner.Scan()
				response = scanner.Text()
				if response == "c" {
					avg := tempSum / float64(bytesCount)
					fmt.Printf("%.2f grader Celsius\n", avg)
					break
				} else if response == "f" {
					tempAvg, err := strconv.ParseFloat(strconv.FormatFloat(tempSum/float64(bytesCount), 'f', -1, 64), 64)
				if err != nil {
    				log.Fatal(err)
				}
					avg := conv.FarhenheitToCelsius(tempAvg)

					fmt.Printf("%.2f grader Fahrenheit\n", avg)
					break
				} else {
					fmt.Println("Ugyldig valg. Prøv igjen.")
				}
			}

		}
	}
}

			


package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/Johannekh/funtemps/conv"
	"github.com/Johannekh/minyr/yr"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Velkommen til minyr!")

	for {
		fmt.Println("Hva vil du gjøre?")
		fmt.Println("1. Konvertere temperaturer")
		fmt.Println("2. Gjennomsnittstemperatur")
		fmt.Println("3. Avslutt")

		choice, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("Feil ved lesing av input:", err)
			continue
		}

		switch string(choice) {
		case "1":
			err := convertTemperatures(reader)
			if err != nil {
				fmt.Println("Feil ved konvertering:", err)
			}
		case "2":
			printAverageTemperature(reader)
		case "3":
			fmt.Println("Ha det bra!")
			return
		default:
			fmt.Println("Ugyldig valg")
		}
	}
}

func convertTemperatures(reader *bufio.Reader) error {
	fmt.Println("Konvertering av temperaturer pågår...")

	// Les inn temperaturdata fra fil
	temperatures, err := yr.ReadTemperaturesFromFile("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		return fmt.Errorf("feil ved lesing av temperaturer fra fil: %v", err)
	}

	// Konverter temperaturer til Fahrenheit
	for i, temp := range temperatures {
		fahr := conv.CtoF(temp)
		temperatures[i] = fahr
	}

	// Skriv temperaturer til fil
	err = yr.WriteTemperaturesToFile("kjevik-tempfahr-20220318-20230318.csv", temperatures)
	if err != nil {
		return fmt.Errorf("feil ved skriving av temperaturer til fil: %v", err)
	}

	fmt.Println("Konvertering fullført!")
	return nil
}

func printAverageTemperature(reader *bufio.Reader) {
	fmt.Println("Hvilken temperaturskala vil du ha gjennomsnittet i? (c/f)")
	scale, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("Feil ved lesing av input:", err)
		return
	}

	// Les inn temperaturdata fra fil
	temperatures, err := yr.ReadTemperaturesFromFile("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		fmt.Println("Feil ved lesing av temperaturer fra fil:", err)
		return
	}

	// Beregn gjennomsnittstemperatur
	var sum float64
	for _, temp := range temperatures {
		if string(scale) == "f" {
			temp = conv.CtoF(temp)
		}
		sum += temp
	}
	average := sum / float64(len(temperatures))

	// Skriv ut gjennomsnittstemperatur
	if string(scale) == "c" {
		fmt.Printf("Gjennomsnittstemperaturen er %.1f grader Celsius\n", average)
	} else {
		fmt.Printf("Gjennomsnittstemperaturen er %.1f grader Fahrenheit\n", average)
	}
}



package yr

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	//"github.com/uia-worker/misc/conv"
)


func CelsiusToFahrenheitFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		convertedLine, err := CelsiusToFahrenheitLine(line)
		if err != nil {
			return nil, err
		}
		lines = append(lines, convertedLine)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func CelsiusToFahrenheitLine(line string) (string, error) {
	fields := strings.Split(line, ";")
	if len(fields) != 3 {
		return "", errors.New("invalid input line format")
	}

	celsius, err := strconv.ParseFloat(fields[2], 64)
	if err != nil {
		return "", err
	}

	fahrenheit := (celsius * 9 / 5) + 32

	return fmt.Sprintf("%s;%s;%.1f", fields[0], fields[1], fahrenheit), nil
}

func main() {
	var cmd = &cobra.Command{
		Use:   "minyr",
		Short: "A brief description of your application",
		Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application.`,
		Run: func(cmd *cobra.Command, args []string) {
			inputFilename := "kjevik-temp-celsius-20220318-20230318.csv"
			outputFilename := "kjevik-tempfahr-20220318-20230318.csv"

			lines, err := CelsiusToFahrenheitFile(inputFilename)
			if err != nil {
				log.Fatal(err)
			}

			outputFile, err := os.Create(outputFilename)
			if err != nil {
				log.Fatal(err)
			}
			defer outputFile.Close()

			writer := bufio.NewWriter(outputFile)
			for i, line := range lines {
				if i == 0 {
					writer.WriteString(line + "\n")
				} else {
					writer.WriteString(line + "\n")
				}
			}

			writer.WriteString("Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Johanne Haakenstad")
			writer.Flush()
		},
	}

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


package yr

import (
    "bufio"
    "encoding/csv"
    "fmt"
    "io"
    "os"
    "strconv"

    "github.com/Johannekh/funtemps/conv"
)

// ConvertTemperatures leser en csv-fil med temperaturer i Celsius og
// konverterer dem til Fahrenheit. De konverterte verdiene skrives til en ny
// csv-fil med samme format som originalen.
func ConvertTemperatures(inputFile, outputFile string) error {
    // Åpne input-filen for lesing
    inFile, err := os.Open(inputFile)
    if err != nil {
        return fmt.Errorf("failed to open input file: %w", err)
    }
    defer inFile.Close()

    // Åpne output-filen for skriving
    outFile, err := os.Create(outputFile)
    if err != nil {
        return fmt.Errorf("failed to create output file: %w", err)
    }
    defer outFile.Close()

    // Opprett en CSV-reader for input-filen
    reader := csv.NewReader(bufio.NewReader(inFile))
    reader.Comma = ';'

    // Opprett en CSV-writer for output-filen
    writer := csv.NewWriter(bufio.NewWriter(outFile))
    writer.Comma = ';'
    defer writer.Flush()

    // Les hver linje i input-filen og konverter temperaturene til Fahrenheit
    isFirstLine := true
    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            return fmt.Errorf("failed to read input file: %w", err)
        }

        // Skriv den første linjen uendret til output-filen
        if isFirstLine {
            err = writer.Write(record)
            if err != nil {
                return fmt.Errorf("failed to write first line to output file: %w", err)
            }
            isFirstLine = false
            continue
        }

        // Konverter temperaturen til Fahrenheit
        tempCelsius, err := strconv.ParseFloat(record[1], 64)
        if err != nil {
            return fmt.Errorf("failed to parse temperature from input file: %w", err)
        }
        tempFahrenheit := conv.CelsiusToFahrenheit(tempCelsius)

        // Skriv konvertert linje til output-filen
        record[1] = strconv.FormatFloat(tempFahrenheit, 'f', 1, 64)
        err = writer.Write(record)
        if err != nil {
            return fmt.Errorf("failed to write converted line to output file: %w", err)
        }
    }

    // Skriv siste linjen til output-filen
    lastLine := fmt.Sprintf("Data er basert på gyldig data (per %s) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Johanne Haakenstad", "18.03.2023")
    err = writer.Write([]string{lastLine})
    if err != nil {
        return fmt.Errorf("failed to write last line to output file: %w", err)
    }

    return nil
}

// CalculateAverageTemperature beregner gjennomsnittstemperaturen fra en csv-fil
// med temperaturer i Celsius. Gjennomsnittstemperaturen kan returneres i
// Celsius eller Fahrenheit.
func CalculateAverageTemperature(inputFile string, outputUnit string) (float64, error) {
    // Åpne input-filen for lesing
   
file, err := os.Open(inputFile)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Les temperaturene fra filen og legg dem til i en slice
	var temperatures []float64
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}

		temp, err := strconv.ParseFloat(strings.Replace(record[1], ",", ".", -1), 64)
		if err != nil {
			return 0, err
		}
		temperatures = append(temperatures, temp)
	}

	// Beregn gjennomsnittstemperaturen i Celsius
	var avgCelsius float64
	for _, t := range temperatures {
		avgCelsius += t
	}
	avgCelsius /= float64(len(temperatures))

	// Konverter gjennomsnittstemperaturen til Fahrenheit hvis ønskelig
	var avgTemp float64
	if outputUnit == "c" {
		avgTemp = avgCelsius
	} else {
		avgTemp = conv.CelsiusToFahrenheit(avgCelsius)
	}

	return avgTemp, nil
}

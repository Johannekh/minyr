package yr


import (
"github.com/Johannekh/funtemps/conv"
        "strconv"
	"bufio"
        "strings"
        "os"
        "fmt"
	"log"
)

func LesAntallLinjerFil(filename string) int{
src, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
if err != nil {

        log.Fatal(err)

}
defer src.Close()
scanner := bufio.NewScanner(src)
antall := 0
for scanner.Scan() {
antall++
}

return antall
}

func KonverteringAvLinjer() {
	regenererFil := SjekkOmFahrFilEksisterer()
	if !regenererFil {
	return
	}

//Apner input filen
inputFil, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
if err != nil {
	log.Fatal(err)
}
defer inputFil.Close()
//Lager output filen
outputFil, err := os.Create("kjevik-temp-fahr-20220318-20230318.csv")
if err != nil {
	log.Fatal(err)
}
defer outputFil.Close()
outputWriter := bufio.NewWriter(outputFil)
inputScanner := bufio.NewScanner(inputFil)

for inputScanner.Scan() {
linje := inputScanner.Text()
konvertertLinje := ProsesserLinjer(linje)
_, err := outputWriter.WriteString(konvertertLinje + "\n")
if err != nil {

	log.Fatal(err)

}

}

err = outputWriter.Flush()
if err != nil {
	log.Fatal(err)
}
fmt.Println("Konverteringen er ferdig")


}


func ProsesserLinjer(linje string) string {
if strings.Contains(linje, "Navn;Stasjon;Tid(norsk normaltid);Lufttemperatur") {
return linje
}



if strings.Contains(linje, "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;") {
linje = strings.Replace(linje, "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;", "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET); endringen er gjort av Johanne Haakenstad", 1)

return linje

}
elementer := strings.Split(linje, ";")
celsiusStr := elementer[len(elementer)-1]
celsius, err := strconv.ParseFloat(celsiusStr, 64)
if err != nil {
log.Fatal(err)
}


fahrenheit := conv.CelsiusToFahrenheit(celsius)
elementer[len(elementer)-1] = fmt.Sprintf("%.1f", fahrenheit)
konvertertLinje := strings.Join(elementer, ";")
return konvertertLinje
}

func GjennomsnittsBeregningCelsius() float64 {
fil, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
if err != nil {
log.Fatal(err)
}
defer fil.Close()

var adderteTemperaturer float64
var antallTemperaturer int

scanner := bufio.NewScanner(fil)
for scanner.Scan() {
	linje := scanner.Text()
	elementer := strings.Split(linje, ";")
	if len(elementer) >= 4 {
	temperatur, err := strconv.ParseFloat(elementer[3], 64)
	if err == nil {
	adderteTemperaturer += temperatur
	antallTemperaturer++
}
}
}

	if antallTemperaturer > 0 {
	gjennomsnittsTemperatur :=  adderteTemperaturer / float64(antallTemperaturer)
	fmt.Printf ("Gjennomsnittstemperaturen i celsius : %.2f\n", gjennomsnittsTemperatur)

}
	return 0
}

func GjennomsnittsBeregningFahr() float64 {
fil, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
if err != nil {
	log.Fatal(err)
}
defer fil.Close()


var adderteTemperaturer float64
var antallTemperaturer int
scanner := bufio.NewScanner(fil)
for scanner.Scan() {
linje := scanner.Text()
elementer := strings.Split(linje, ";")
if len (elementer) >= 4 {
temperatur, err := strconv.ParseFloat(elementer[3], 64)
if err == nil {

	adderteTemperaturer += temperatur
	antallTemperaturer++
}
}
}

if antallTemperaturer > 0 {
	gjennomsnittsTemperatur := adderteTemperaturer / float64(antallTemperaturer)
	fahrenheitGjennomsnitt := conv.CelsiusToFahrenheit(gjennomsnittsTemperatur)
	fmt.Printf ("Gjennomsnittstemperaturen i fahrenhei : %.2f\n", fahrenheitGjennomsnitt)
	return fahrenheitGjennomsnitt
}
return 0

}

func SjekkOmFahrFilEksisterer() bool {
if _, err := os.Stat("kjevik-temp-fahr-20220318-20230318.csv"); err == nil {
	fmt.Println("Filen eksisterer allerede, generere en ny fil? (j/n)")
	var regenerer string
	fmt.Scanln(&regenerer)
	if strings.ToLower(regenerer) == "j" || strings.ToLower(regenerer) == "J" {
	err := os.Remove("kjevik-temp-fahr-20220318-20230318.csv")
		if err != nil {
			log.Fatal(err)
			}
	return true
	} else if strings.ToLower(regenerer) == "n" || strings.ToLower(regenerer) == "N" {
		fmt.Println("Avbrutt")

	return false
	} else {

	fmt.Println("Ugyldig svar. Skriv 'j' for ja eller 'n' for nei.")

	return false
	}
	}
	return true

	}

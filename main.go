pacage main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"

    "github.com/Johannekh/minyr/yr"
)

func main() {
	src, err := os.Open("/home/Johannehaakenstad/minyr/kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	var linebuf []byte
	buffer := make([]byte, 1)
	bytesCount := 0
	for {
		_, err := src.Read(buffer)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		bytesCount++
		if buffer[0] == 0x0A {
			log.Println(string(linebuf))
			// Her
			elementArray := strings.Split(string(linebuf), ";")
			if len(elementArray) > 3 {
				celsius := elementArray[3]
				fahr := conv.CelsiusToFahrenheit(celsius)
				log.Println(fahr)
			}
			linebuf = nil
		} else {
			linebuf = append(linebuf, buffer[0])
		}

		if err == io.EOF {
			break
		}
	}
}
              


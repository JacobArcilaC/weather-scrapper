package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Enter city name to get current weather:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(getCityCurrentWeather(scanner.Text()))
}

func getCityCurrentWeather(city string) string {
	var geoDecoderURI string = fmt.Sprintf("https://api.weatherstack.com/current?query=%s&access_key=b11af9136360749ad009888d7290f659", city)
	response, err := http.Get(geoDecoderURI)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode == 200 {
		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return string(body)
		}
	}
	return ""
}

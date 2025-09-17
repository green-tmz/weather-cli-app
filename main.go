package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Читаем API-ключ из переменной окружения.
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: Please set the OPENWEATHER_API_KEY environment variable.")
		os.Exit(1)
	}

	// Проверяем, передано ли название города в аргументах командной строки.
	var city string
	if len(os.Args) > 1 {
		city = strings.Join(os.Args[1:], " ")
	} else {
		// Запрашиваем у пользователя название города, если аргумент не указан.
		fmt.Print("Enter the city name: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		city = strings.TrimSpace(input)
	}

	// Получаем данные о погоде.
	weather, err := getWeatherData(city, apiKey)
	if err != nil {
		fmt.Println("Error fetching weather data:", err)
		return
	}

	// Отображаем данные о погоде.
	displayWeatherData(weather)
}

// displayWeatherData выводит информацию о погоде в удобном для пользователя формате.
func displayWeatherData(weather *WeatherData) {
	fmt.Printf("\nWeather for %s:\n", weather.Name)
	if len(weather.Weather) > 0 {
		fmt.Printf("Description: %s\n", strings.Title(weather.Weather[0].Description))
	}
	fmt.Printf("Temperature: %.2f°C\n", weather.Main.Temp)
	fmt.Printf("Humidity: %d%%\n\n", weather.Main.Humidity)
}

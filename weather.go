package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// WeatherData — структура под JSON-ответ OpenWeatherMap.
type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

// getWeatherData получает данные о погоде для указанного города.
func getWeatherData(city string, apiKey string) (*WeatherData, error) {
	// Формируем URL API. Используем метрическую систему (градусы Цельсия).
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

	// Отправляем HTTP-запрос.
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем, что запрос прошёл успешно.
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("error: received status code %d: %s", resp.StatusCode, string(bodyBytes))
	}

	// Читаем тело ответа.
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	// Декодируем JSON в структуру WeatherData.
	var weather WeatherData
	err = json.Unmarshal(bodyBytes, &weather)
	if err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %v", err)
	}

	return &weather, nil
}

CLI-приложение прогноза погоды

Подготовка к запуску:
1.  Зарегестрироваться на сайте https://openweathermap.org/
2.  Получить api ключ
3.  Установка fpi ключа в переменную окружения:
    Для Linux/Mac:
    export OPENWEATHER_API_KEY=your_api_key_here
    Для Windows (Command Prompt):
    set OPENWEATHER_API_KEY=your_api_key_here

Запуск:
./weather-app Казань
или
go run *.go Kazan

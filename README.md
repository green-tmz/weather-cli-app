# CLI-приложение прогноза погоды

### Подготовка к запуску:
1.  Зарегестрироваться на сайте https://openweathermap.org/
2.  Получить api ключ
3.  Установка api ключа в переменную окружения:

    **для Linux/Mac:**
    ```bash
      export OPENWEATHER_API_KEY=your_api_key_here
    ```
    **для Windows (Command Prompt):**
    ```bash
      set OPENWEATHER_API_KEY=your_api_key_here
     ```
### Запуск:
```bash
    ./weather-app Казань
```
или
```bash
    go run *.go Kazan
```

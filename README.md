# CLI-приложение прогноза погоды

### Подготовка к запуску:
1.  Зарегестрироваться на сайте https://openweathermap.org/
2.  Получить api ключ
3.  Установка api ключа в переменную окружения:
    Для Linux/Mac:
   4. ```bash
      export OPENWEATHER_API_KEY=your_api_key_here
      ```
    
       Для Windows (Command Prompt):
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

# HTTP Cheaker with Dynamic Worker Pool

Лёгкое консольное приложение на Go для параллельной проверки HTTP-ссылок с динамически настраиваемым пулом воркеров.

## Возможности

* Чтение URL из `stdin` и отправка запросов через пул горутин.
* Динамическое добавление и удаление воркеров во время работы.
* Вывод статусов ответов (`200 OK`, `404 Not Found` и т.д.).
* Простое управление через команды: `add`, `remove`, `status`, `quit`.
* Набор модульных тестов для базовой проверки логики пула.

## Структура проекта

```text
http-cheaker/
├── go.mod           # модуль Go
├── main.go          # точка входа, чтение команд и URL
├── models.go        # реализация пула воркеров (Pool)
├── worker.go        # функция-воркер для обработки задач
```

## Требования

* Go 1.20 или новее

## Установка и запуск

1. Склонируйте репозиторий:

   ```bash
   git clone https://github.com/Skotchmaster/http-cheaker
   cd http-cheaker
   ```
2. Инициализируйте зависимости:

   ```bash
   go mod tidy
   ```
3. Сборка бинарника:

   ```bash
   go build -o http-cheaker
   ```
4. Запуск приложения:

   ```bash
   ./http-cheaker -workers=5
   ```

   По умолчанию пул запускает 5 воркеров. Флаг `-workers` можно опустить.

## Использование

После запуска введите:

* **URL** (например):

  ```text
  https://www.google.com
  https://httpbin.org/status/418
  ```
* **Команды управления**:

  * `add`    — добавить нового воркера
  * `remove` — убрать одного воркера
  * `status` — отобразить текущее число воркеров
  * `quit`   — корректно завершить работу

Пример сессии:

```text
$ ./http-cheaker -workers=3
Added worker #0 (total: 1)
Added worker #1 (total: 2)
Added worker #2 (total: 3)
Введите URL или команду (add, remove, status, quit):
https://jsonplaceholder.typicode.com/posts/1
[Worker #0] https://jsonplaceholder.typicode.com/posts/1 → 200 OK
status
Active workers: 3
add
Added worker #3 (total: 4)
remove
Removed worker #1 (remaining: 3)
quit
[Worker #0] shutting down
[Worker #2] shutting down
[Worker #3] shutting down
```

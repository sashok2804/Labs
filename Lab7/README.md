
---

# README

## Описание проекта

Этот проект реализован на двух языках программирования: Go и C#. Оба варианта программы выполняют одну и ту же логику, с возможностью передачи параметра через флаг `-t`. Параметр определяет выполнение определенной задачи программы и принимает значения от 1 до 6.

### Доступные версии:
- **Go**: Требуется Go версии 1.23 и выше.

## Требования

- **Go**: Операционная система: Windows/Linux/macOS, Go версии 1.23 и выше.
- **.NET**: Операционная система: Windows/Linux/macOS, .NET SDK версии 6.0 и выше.

## Запуск программы

### Запуск на Go

- запуск сервка осуществляется из папки самого проекта при помощи команды: 
```bash
go run .\serv1.go
```
- запуск сервера происходит на порту 12345
- при запуске сервера в терминале вы увидите что сервер был запущен
## Как работать с сервером
- для работы например первого задания вам нужно запустит первый сервер (все задания помечены номерами) а также клиента, клиента запускается также как и сервер   
```bash
go run .\p2.go
```
-после запуска сервера и клиента вы можете написать смс с клиента и сервер получит смс от клиента всё также в терминале и выдаст ответ клиенту также в терминале

## Запуск разлиых заданий и проверка их
- 1 задание: запускаем сервер и клиент по инструкции
- 2 задание: запускаем сервер и клиент по инструкции
- 3 аналогично 1му заданию ток сервер поддерживает несколько клиентов
- 4 запускаем только сервер, открываем страничку ```bash http://localhost:12345/hello ```, в  новом терминале прописываем ```bash curl http://localhost:12345/hello ``` это гет запрос также пост зарос ```bash Invoke-WebRequest -Uri "http://localhost:12345/data" -Method Post -Body '{"message":"Hello, server!"}' -ContentType "application/json" -Headers @{"Content-Type"="application/json"} ``` в терминале будет показано что сервер получил оба запроса и обработал его
- 5 запускаем сервер, открываем новый терминал, прописываем ```bash curl http://localhost:12345/another ``` и видим что сервер при получении запроса отозвался что что получен гет запрос
- 6 запускам всё также сервер под его номером, теперь нужно запустить клиента с веб сокетом ```bash go run ./soc.go ``` запускаем 2 таких клиента с разных терминалов и отправляем смс с одного и сервер пересылает его всем клиентам котооыре подключены к серверу, то есть второму нашему клиенту
- ps важно после каждого использования сервера выключасть его, если этого не сделать порт будет занят и вы не сможете запустить другой сервер на этом порту 
### Пример для Go

```bash
go run .\Lab1\Go\main.go -t 6
```

### Запуск на C# (.NET)

Для запуска программы на C# используйте команду:

```bash
dotnet run --project .\Lab<num>\C#\Task\Task.csproj -t <value>
```

- `<num>` — номер лабораторной работы.
где `<value>` — это значение от 1 до 6, которое определяет поведение программы.

### Пример для .NET

```bash
dotnet run --project .\Lab1\C#\Task\Task.csproj -t 6
```

## Задания

- `1`: Создание TCP-сервера:
-	Реализуйте простой TCP-сервер, который слушает указанный порт и принимает входящие соединения.
-	Сервер должен считывать сообщения от клиента и выводить их на экран.
-	По завершении работы клиенту отправляется ответ с подтверждением получения сообщения.
- `2`: Реализация TCP-клиента:
-	Разработайте TCP-клиента, который подключается к вашему серверу.
-	Клиент должен отправлять сообщение, введённое пользователем, и ожидать ответа.
-	После получения ответа от сервера клиент завершает соединение.
- `3`: Асинхронная обработка клиентских соединений:
-	Добавьте в сервер многопоточную обработку нескольких клиентских соединений.
-	Используйте горутины для обработки каждого нового соединения.
-	Реализуйте механизм graceful shutdown: сервер должен корректно завершать все активные соединения при остановке.
- `4`: Создание HTTP-сервера:
-	Реализуйте базовый HTTP-сервер с обработкой простейших GET и POST запросов.
-	Сервер должен поддерживать два пути:
-	GET /hello — возвращает приветственное сообщение.
-	POST /data — принимает данные в формате JSON и выводит их содержимое в консоль.
- `5`: Добавление маршрутизации и middleware:
-	Реализуйте обработку нескольких маршрутов и добавьте middleware для логирования входящих запросов.
- `6`: Middleware должен логировать метод, URL, и время выполнения каждого запроса.
-	Веб-сокеты:
-	Реализуйте сервер на основе веб-сокетов для чата.
-	Клиенты должны подключаться к серверу, отправлять и получать сообщения.
-	Сервер должен поддерживать несколько клиентов и рассылать им сообщения, отправленные любым подключённым клиентом.

## Ошибки

- Если параметр `-t` не указан, программа выполнит задание 1.
- Если передано значение вне диапазона от 1 до 6, программа также выполнит задание 1.
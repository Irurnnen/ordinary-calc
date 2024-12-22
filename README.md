# ordinary-calc

Ordinary-calc это простая библиотека и HTTP сервер написанная на Go для вычисления математических выражений. Данный проект может работать в двух режимах: библиотека и HTTP проект

## Возможности

- Вычисление простых математических выражений

## Как использовать проект как библиотеку

### Установка

Для использования данной библиотеки в своем коде нужно скачать в ваш проект библиотеку с помощью команды:

```bash
go get -u github.com/Irurnnen/ordinary-calc/pkg/calc
```

### Пример кода

```golang
package main

import (
	"fmt"

	"github.com/Irurnnen/ordinary-calc/pkg/calc"
)

func main() {
	expression := "2+2*2"

	result, err := calc.Calc(expression)
	if err != nil {
		panic("Error while calculating expression: " + err.Error())
	}
	fmt.Printf("%s = %f", expression, result)
}
```

## Как использовать как HTTP сервер

На данный момент есть несколько вариантов запуска HTTP сервера: bare-metal, docker и несколько режимов сборки: debug и release. 

### Разница между debug и release

HTTP сервер имеет несколько режимов сборки для возможности простой разработки проекта и простого запуска проекта в прод. Разница между режимами сборки это доступ к онлайн документации по пути /swagger/ (об этом будет ниже) для более простого доступа к функционалу проекта. 

### Запуск bare-metal

#### Сборка проекта

Для сборки проекта понадобится Golang.

Сборка проекта проходит в 2 этапа:

1. Генерирование файлов документации (нужно только для debug версии):

    Для генерации документации используется [swag-go](https://github.com/swaggo/swag). Для его установки воспользуйтесь командой:

    ```bash
    go install github.com/swaggo/swag/cmd/swag@latest
    ```

    Для генерации документации воспользуйтесь командой:

    ```bash
    swag init --generalInfo ./cmd/main.go --output ./docs
    ```

2. Сборка проекта:

    Сборка проекта запускается с помощью команды. Важно, вместо ${BUILD_MODE} поставьте вариант сборки проекта (доступно: debug, release):

    ```bash
    go build --tags ${BUILD_MODE} -o ./ordinary-calc.exe ./cmd/
    ```

Также нужно создать в environment переменную PORT для выбора на каком порте запустится программа

В Bash
```bash
PORT=8080
```

В PowerShell
```powershell
$PORT=8080
```

В CMD
```cmd
set PORT=8080
```

Осталось только запустить проект. соответственно для этого запускаем его:
```bash
./ordinary-calc.exe
```

### Запуск docker-cli

Также есть вариант запустить проект используя docker, [установка docker](https://docs.docker.com/engine/install/).

```bash
# Сборка проекта, при запуске замените ${BUILD_MODE} на один из вариантов сборки
# проекта (возможные варианты: debug, release)
docker build -t ordinary-calc -f .\docker\Dockerfile --build-arg BUILD_MODE=${BUILD_MODE} .

# Запуск проекта с помощью docker cli. Для смены порта измените все 3 значения в команде
docker run -p 8080:8080 -e PORT=8080 ordinary-calc
```

### Запуск docker-compose

Также есть вариант запуска HTTP сервера используя docker-compose файл (предпочтительный вариант). Соответственно для запуска через docker-compose потребуется [установить docker](https://docs.docker.com/engine/install/).

Для запуска нужно создать файл `.env`, где будут хранится настройки HTTP сервера. Для создания есть пример файла `.env.example` скопируйте пример в файл `.env`:

```bash
cp .env.example .env
```

Далее после создания файла `.env` можно собирать и запускать проект

```bash
# Сборка проекта
docker compose build
# Запуск проекта
docker compose up
```

### О проекте

Данный проект имеет один endpoint по пути `/api/v1/calculate` с помощью которого можно вычислить математическое выражение. Пользователь может отправить по данному пути POST запрос с телом

```json
{
    "expression": "2 + 2 * 2"
}
```
В ответ пользователь получает ответ с данным телом и кодом 200, если выражение было посчитано без ошибок:

```json
{
    "result": 6
}
```

В случае если выражение имело ошибку будет отправлен HTTP-ответ с кодом 422 и телом:

```json
{
    "error": "Пример ошибки"
}
```

Также в случае ошибки на стороне сервера будет отправлен код HTTP-ответ с кодом 500 и с телом:

```json
{
    "error": "Internal server error"
}
```

## Swagger

Для удобства тестирования API можно воспользоваться Swagger [(Что такое Swagger)](https://practicum.yandex.ru/blog/chto-takoe-swagger/). Доступ к swagger можно получить 2-мя путями: через файл и через debug версию проекта

### Доступ по файлу

Для получения доступа к Swagger можно воспользоваться Плагинами IDE, Онлайн редакторами Swagger, Postman и другими способами. Файл с документацией находится по пути `docs/swagger.json`

Пример для VS Code: Плагин [Swagger Viewer](https://marketplace.visualstudio.com/items?itemName=Arjun.swagger-viewer)

### Доступ в debug версии

Как было написано ранее в debug версии есть swagger встроенный в проект, он доступен по пути `/swagger/index.html` в котором открывается интерактивное меню для отправки запросов на сервер.

## Curl

Получить доступ к API можно используя curl. Пример команды:

```bash
# Работает Linux и в cmd (в Windows)
curl -X POST http://localhost:8080/api/v1/calculate \
    --header 'Content-Type: application/json' \
    --data '{
        "expression": "2+2*2"
    }'
```
## Варианты ошибок

Далее будут описаны все ошибки что заложены в программу

- `Expression has extra characters` - в математическом выражении есть символы, что соответствуют маске `[^0-9\.+\-*\/()^\s]`.

- `Expression has unpaired brackets"` - в математическом выражении есть непарные скобочки.

- `Expression has wrong bracket order` - в математическом выражении неправильный порядок скобочек.

- `Expression has multiple operands` - в математическом выражении несколько операндов идут друг за другом.

- `Expression has multiple sequential numbers` - в математическом выражении несколько чисел идут друг за другом.

- `Expression has zero by division` - при вычислении математического выражения было произведено действие деление на ноль.

- `Expression has operand at the beginning or at the end` - в начале или в конце математического выражения стоит операнд.

- `Expression is empty` - математическое выражение не задано

- `Internal server error` - неизвестная ошибка в программе (лучше написать об этом в Issues)

## Roadmap

- [ ] Добавление Github Actions
- [ ] Добавление шаблонов для Github
- [ ] Добавление Kubernetes
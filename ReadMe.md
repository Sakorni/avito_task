# Avito_task

Данный проект является реализацией тестового задания на стажёра backend-разработки в Avito.
(https://github.com/avito-tech/autumn-2021-intern-assignment)

Для корректной работы сервиса требуется **MYSQL** БД (дамп находится в директории `database/`)

Вместе с этим требуется заполнить корректными данными `.env` файлы, находящиеся в директории `/config`

Документация API - https://documenter.getpostman.com/view/11575297/UVeDsSRK

UPD1: Выполнил так же дополнительное задание №1. Для получения текущего курса валют используется ресурс
(https://exchangeratesapi.io/). Для корректного функционирования сервиса необходимо задать переменную среды
`CURRENCY_API_KEY` через средства ОС, либо добавив соответствующий файл .env в папку `/config`
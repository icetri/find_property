# find_property

----
Backend для тг-бота который парсит объявления 
о продаже квартир с циана и выдает объявления по фильтрам клиенту.
----

    Для корректного запуска необходимо создать базу данных (user=user123 password=123 dbname=find_property) и сделать миграцию из папки ./schema,
    а также скорректировать config-local.yaml в папке ./config добавив токен бота и ваш чат айди.

## Запуск: $ go run cmd/find_property/main.go --config-path="config/config-local.yaml"

По желанию прописать General API аннотация swagger в main.go и сделать swag init -g cmd/find_property/main.go

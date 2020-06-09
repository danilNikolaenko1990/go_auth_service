**Сервис авторизации.** 

Клонируем репозиторий в рабочую папку go.
В системе должен быть установлен docker-compose и docker

Для запуска базы заходим в папку build и делаем `docker-compose up &`

После того, как база стартовала, делаем go build и запускаем сервис:
`./auth_service http --server-port=8080 --db-user=admin --db-password=admin --db-host=localhost --db-port=5432 --db-name=postgres`

Сервис имеет два метода. 

Запрос для регистрации

`curl --location --request POST 'http://localhost:8080/register' \
 --header 'Content-Type: text/plain' \
 --data-raw '{
 	"login":"testLogin",
 	"email":"testemail@test.ru",
 	"password":"querty123",
 	"phone":"11398724911"
 }'`

 пример ответа: 
 
 `{"registered":true,"error_message":"","error_code":""}`
 
 Запрос для логина
 
 `curl --location --request POST 'http://localhost:8080/login' \
  --header 'Content-Type: text/plain' \
  --data-raw '{
  	"login":"testLogin",
  	"password":"querty123"
  }'`
  
  Пример ответа: 
  `{
      "logged": true,
      "error_message": ""
  }`

**Что умеет**
* Валидация входящих параметров при регистрации и логине 
* Проверка на наличие в системе юзера с уже имеющимися логином, телефоном или email
* Бизнес логика закрыта unit тестами
* пароли хранятся в виде хешей

**Где сознательно срезал углы в угоду скорости реализации**
* di контейнер для проекта такого уровня сознательно не стал использовать.
* Логгирование. Вместо этого прокидываю ошибки наружу, сделав допущение, что это внутренний сервис.
* Человекочитаемые ошибки валидации (считаем, что это внутренний сервис)
* Масштаб покрытия unit тестами
* Полная докеризация проекта (отталкивался от отсутствия необходимости вообще докеризировать по ТЗ)
* конфигурирование через файл\env переменные
* не стал внедрять механизм миграций, схема бд инициализируется средствами докер-контейнера (см. docker-compose)

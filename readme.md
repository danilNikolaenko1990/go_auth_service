Test task for one of the companies.
Quest text:
Implement the service in the Golang programming language.
The task of the service is to register, authorize users.
The service must have two HTTP POST methods:
* Login 
* Register

Use PostgreSQL database

Required fields for registration:
* Login
* Email
* password
* phone number

Implementation details are at the discretion of the developer.

**Auth Service.** 

Clone repository to go work folder.
docker-compose and docker must be installed

run: 
`docker-compose  -f docker-compose.yml up`


Service has two methods: 

Registration query:

`curl --location --request POST 'http://localhost:8080/register' \
 --header 'Content-Type: text/plain' \
 --data-raw '{
 	"login":"testLogin",
 	"email":"testemail@test.ru",
 	"password":"querty123",
 	"phone":"11398724911"
 }'`

 example of the answer: 
 
 `{"registered":true,"error_message":"","error_code":""}`
 
 login query
 
 `curl --location --request POST 'http://localhost:8080/login' \
  --header 'Content-Type: text/plain' \
  --data-raw '{
  	"login":"testLogin",
  	"password":"querty123"
  }'`

example of the answer:
`{
      "logged": true,
      "error_message": ""
  }`

**Features**
* Validation params during registration and login 
* Checking existing user with same login, phone number or email
* Business logic covered with unit tests
* Password stored as hashes

**Made it easier to make it faster to code**
* DI container wasn't use
* Throwing errors outside without logger
* Less unit tests
* Without migration system like goose because db init made with docker container tools (see ./build/postgres/schema.sql:/docker-entrypoint-initdb.d/10-init.sql in docker-compose) 

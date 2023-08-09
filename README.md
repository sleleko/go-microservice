# go-microservice
Simple microservice example based on Golang

## How to run?

1) First you should install Go into your operation system
2) run service 1 ```go run service1.go```
3) run service 2 ```go run service2.go```
4) run gateway ```go run gateway.go```

## How to test result ?

1) Run (for example via curl or your browser) http://localhost:8080/service1 (You should see message from service1 called via microservice gateway)
2) Run (for example via curl or your browser) http://localhost:8080/service2 (You should see message from service2 called via microservice gateway)
3) Run (for example via curl or your browser) http://localhost:8081/ (You should see message from service1 called via service1 directly)
4) Run (for example via curl or your browser) http://localhost:8082/ (You should see message from service1 called via service2 directly)

## How to run in Docker ?
1) git clone this project
2) For start: go to project dir and run ```docker-compose up --build```
3) For stop: ```docker-compose down```
   

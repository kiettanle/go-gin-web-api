# go-gin-web-api
Creating Golang REST API with Gin-Gonic Web Framework &amp; MongoDB

1. Create new go.mod to save dependencies

```
go mod init <Module Name>
```

2. Then install gin framework
```
 go get github.com/gin-gonic/gin
```
3. Install go mongodb driver
```
go get go.mongodb.org/mongo-driver
```
4. Create folder controllers, services, models
```
mkdir controllers services models
```
5. Create main.go
```
touch main.go
```
6. Create user model
```
touch models/user.go
```
7. Create user service
```
touch services/user.service.go
```
8. Create user service implementation
```
touch services/user.service.impl.go
```
9. Create user controller
```
touch controllers/user.controller.go
```
10. Create main.go
```
touch main.go
```
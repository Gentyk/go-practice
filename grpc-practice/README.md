go mod init example/web-service-gin/go-practice/grpc-practice
go mod tidy


1. установка зависимоcтей

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

linux
```
export PATH="$PATH:$(go env GOPATH)/bin"
```
windows
# Если установлен winget (Windows 10/11)
winget install Google.Protobuf```


# Получаем текущий GOPATH
$goPath = go env GOPATH

# Добавляем в переменную Path пользователя
$userPath = [Environment]::GetEnvironmentVariable("Path", "User")
[Environment]::SetEnvironmentVariable("Path", "$userPath;$goPath\bin", "User")
```

2. создание и заполнение example.proto
3. компиляция

```
protoc -I .\api\proto `
        --go_out=pkg\api --go_opt=paths=source_relative `
        --go-grpc_out=pkg\api --go-grpc_opt=paths=source_relative `
        .\api\proto\example.proto
```

---
4. создать grpc сервер и собрать
```
go build .\cmd\server\
```

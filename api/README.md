## работа с gin

короткая дока
https://go.dev/doc/tutorial/web-service-gin

`go mod init example/web-service-gin`

для получения зависимостей
`go get .`

запуск
`go run .`

дергать ручки через 
`curl http://localhost:8080/albums`
или
```
curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```
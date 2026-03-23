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

---------------------------------------------------------------
bench
```
 go test -bench=Benchmar*
```
1. -bench — запускает бенчмарки, соответствующие регулярному выражению
2. -benchmem — добавить в вывод статистику по использованию памяти:
- B/op — байты, выделенные на операцию;
- allocs/op — количество аллокаций памяти на операцию.
3. -benchtime — установить минимальное время выполнения бенчмарка.
4. -count — количество запусков каждого бенчмарка (для получения стабильных результатов
5. -cpu — указать список значений GOMAXPROCS для тестирования.
10. -timeout — максимальное время выполнения теста/бенчмарка.
11. -o — скомпилировать тесты в отдельный исполняемый файл (без запуска).
```
 go test -bench=Bench* -o my_benchmarks.test
 ./my_benchmarks.test -test.bench
```
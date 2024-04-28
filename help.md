migrate command

```
    migrate create -ext sql -dir sql/migration -seq init_schema
```

generate table into struct (tapi ini cara standart karena tetap manual buat functionnya)

```
    go install github.com/fraenky8/tables-to-go@master

    tables-to-go -v -of ./sql/dto -t mysql -h 127.0.0.1 -d bwa -u root -p secret -pn dto
```

generate with generatedb.go with package "gorm.io/gen"

```
    go run generatedb.go
```

disable error docker with golang
```
    sudo snap install golangci-lint
    
    //check code program golang untuk build di docker atau manual
    GOFLAGS=-buildvcs=false golangci-lint run -v

```

auto reload untuk golang sama seperti nodemon
````
    go install github.com/cosmtrek/air@latest
````
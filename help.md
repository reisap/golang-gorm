migrate command

```
    migrate create -ext sql -dir sql/migration -seq init_schema
```

generate table into struct (but this is standard because still needed manual process function)

```
    go install github.com/fraenky8/tables-to-go@master

    tables-to-go -v -of ./sql/dto -t mysql -h 127.0.0.1 -d bwa -u root -p secret -pn dto
```

generate with generatedb.go with package "gorm.io/gen"

```
    go run generatedb.go
```

linter golang for disable error docker build golang

```
    sudo snap install golangci-lint

    //check code program golang for build in docker or manual
    GOFLAGS=-buildvcs=false golangci-lint run -v

```

auto reload for golang looks like nodemon

```
    go install github.com/cosmtrek/air@latest
    air init //generate edit .air.toml

    //edit file in .air.toml if using golang >=1.22
    cmd = "go build -buildvcs=false -o ./tmp/main ."
```

check json empty in golang

```
    err := json.Unmarshal([]byte(`{}`), &jsonData1)
```

different between save and update in golang with gorm library

```
    //untuk proses save/create/insert pada gorm menggunakan method/fungsi create contohnya
    err := db.Create(&user).Error

    //untuk proses update/edit pada golang menggunakan gorm contohnya
    err := db.Save(&user).Error
```

Swagger

```
    go install github.com/swaggo/swag/cmd/swag@latest
    swag init
```

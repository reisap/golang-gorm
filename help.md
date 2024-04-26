migrate command

```
migrate create -ext sql -dir sql/migration -seq init_schema
```

generate table into struct

```
go install github.com/fraenky8/tables-to-go@master

tables-to-go -v -of ./dto -t mysql -h 127.0.0.1 -d bwa -u root -p secret -pn dto -suf _entity

```
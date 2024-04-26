app:
	docker-compose up
stopapp:
	docker-compose down
migrateup:
	migrate -path sql/migration -database "mysql://root:secret@tcp(localhost:3306)/bwa?x-tls-insecure-skip-verify=false" -verbose up
migratedown:
	migrate -path sql/migration -database "mysql://root:secret@tcp(localhost:3306)/bwa?x-tls-insecure-skip-verify=false" -verbose down
test :
	go test -v -cover ./...
.PHONY: app stopapp migrateup migratedown test
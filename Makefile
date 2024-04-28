start:
	docker-compose up
stop:
	docker-compose down
migrate-up:
	migrate -path sql/migration -database "mysql://root:secret@tcp(localhost:3306)/bwa?x-tls-insecure-skip-verify=false" -verbose up
migrate-down:
	migrate -path sql/migration -database "mysql://root:secret@tcp(localhost:3306)/bwa?x-tls-insecure-skip-verify=false" -verbose down
test :
	go test -v -cover ./test/...
.PHONY: start stop migrate-up migrate-down test
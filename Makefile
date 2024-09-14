pbgen:
	cd src && protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.

dup:
	docker-compose -f docker-compose.yaml up

urun:
	cd src/updates && PORT=3001 REDIS_ADDR=localhost:6379 go run main.go

crun:
	cd src/cdc && REDIS_ADDR=localhost:6379 MYSQL_DSN="user:pass123@tcp(localhost:3306)/db" go run main.go
pbgen:
	cd src && protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.

dup:
	docker-compose -f docker-compose.yaml up

ddown:
	docker-compose -f docker-compose.yaml down

updaterun:
	cd src/updates && PORT=3001 REDIS_ADDR=localhost:6379 go run main.go

cdcrun:
	cd src/cdc && REDIS_ADDR=localhost:6379 MYSQL_DSN="user:pass123@tcp(localhost:3306)/db" go run main.go

userrun:
	cd src/user && PORT=3002 MYSQL_DSN="user:pass123@tcp(localhost:3306)/db" go run main.go

apigwrun:
	cd src/apigw && PORT=3003 UPDATES_ADDR=localhost:3001 USER_ADDR=localhost:3002 REDIS_ADDR=localhost:6379 BASE_URL=localhost:8080 GOOGLE_CLIENT_ID=CLIENT_ID_HERE GOOGLE_CLIENT_SECRET=CLIENT_SECRET_HERE SERVER_SECRET=helloworld go run main.go
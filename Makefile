# Code gen
pbgen:
	cd src && protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.

# Run applications
runupdates:
	cd src/updates && PORT=3001 REDIS_ADDR=localhost:6379 go run main.go

runitem:
	cd src/item && REDIS_ADDR=localhost:6379 MYSQL_DSN="user:pass123@tcp(localhost:3306)/db" go run main.go

runuser:
	cd src/user && PORT=3002 MYSQL_DSN="user:pass123@tcp(localhost:3306)/db" go run main.go

runbusiness:
	cd src/business && MYSQL_DSN="user:pass123@tcp(localhost:3306)/db" go run main.go

runapigw:
	cd src/apigw && PORT=3003 UPDATES_ADDR=localhost:3001 USER_ADDR=localhost:3002 REDIS_ADDR=localhost:6379 BASE_URL=localhost:8080 GOOGLE_CLIENT_ID=CLIENT_ID_HERE GOOGLE_CLIENT_SECRET=CLIENT_SECRET_HERE SERVER_SECRET=helloworld go run main.go

# Docker env
envup:
	docker-compose -f docker-compose.yaml up --build

envdown:
	docker-compose -f docker-compose.yaml down
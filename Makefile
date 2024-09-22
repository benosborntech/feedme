ENV_FILE = .env

# Code gen
pbgen:
	cd src && protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.

# Docker env
envup:
	docker-compose -f docker-compose.yaml up

envdown:
	docker-compose -f docker-compose.yaml down

# Run applications
runupdates:
	cd src/updates && PORT=3001 REDIS_ADDR=localhost:6379 go run main.go

runcdc:
	cd src/cdc && REDIS_ADDR=localhost:6379 MYSQL_DSN="user:pass123@tcp(localhost:3306)/db" go run main.go

runuser:
	cd src/user && PORT=3002 MYSQL_DSN="user:pass123@tcp(localhost:3306)/db" go run main.go

runmain:
	cd src/main && MYSQL_DSN="user:pass123@tcp(localhost:3306)/db" go run main.go

runapigw:
	cd src/apigw && PORT=3003 UPDATES_ADDR=localhost:3001 USER_ADDR=localhost:3002 REDIS_ADDR=localhost:6379 BASE_URL=localhost:8080 GOOGLE_CLIENT_ID=CLIENT_ID_HERE GOOGLE_CLIENT_SECRET=CLIENT_SECRET_HERE SERVER_SECRET=helloworld go run main.go

# Build applications
buildupdates:
	eval $$(minikube -p minikube docker-env) && docker build -t feedme/updates -f src/updates/Dockerfile ./src

buildcdc:
	eval $$(minikube -p minikube docker-env) && docker build -t feedme/cdc -f src/cdc/Dockerfile ./src

builduser:
	eval $$(minikube -p minikube docker-env) && docker build -t feedme/user -f src/user/Dockerfile ./src

buildmain:
	eval $$(minikube -p minikube docker-env) && docker build -t feedme/main -f src/main/Dockerfile ./src

buildapigw:
	eval $$(minikube -p minikube docker-env) && docker build -t feedme/apigw -f src/apigw/Dockerfile ./src

build: buildupdates buildcdc builduser buildmain buildapigw
	echo "build complete"

# Local kubernetes
localkubestart:
	minikube start --addons=ingress

localkubeconfig:
	kubectl config view > ./.kube/config

localkubeinit:
	source $(ENV_FILE) && cd terraform/local && terraform apply -var="google_client_id=$$GOOGLE_CLIENT_ID" -var="google_client_secret=$$GOOGLE_CLIENT_SECRET" -var="server_secret=helloworld" -var="mysql_dsn=user:pass123@tcp(localhost:3306)/db"

localkubesetup: localkubestart localkubeconfig localkubeinit
	echo "local kube setup complete"

localkubeapply:
	kubectl apply -f ./manifests --prune -n dev --selector app=feedme

localkubeupdate: build localkubeapply
	echo "local kube start complete"
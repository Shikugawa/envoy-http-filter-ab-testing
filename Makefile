build-green:
	GOOS=linux CGO_ENABLED=0 go build -o ./server/green/dist/server-green -i ./server/green/main.go

build-blue:
	GOOS=linux CGO_ENABLED=0 go build -o ./server/blue/dist/server-blue -i ./server/blue/main.go

build-blue-image:
	docker build ./server/blue -t ayamaruyama/envoy-test-blue

build-green-image:
	docker build ./server/green -t ayamaruyama/envoy-test-green

push-blue-image:
	docker push ayamaruyama/envoy-test-blue

push-green-image:
	docker push ayamaruyama/envoy-test-green

all:
	make build-blue
	make build-green
	make build-blue-image
	make build-green-image
	make push-blue-image
	make push-green-image

apply-all:
	kubectl apply -f ./kubernetes/tier-mysql.yaml
	kubectl apply -f ./kubernetes/tier-redis.yaml	
	kubectl apply -f ./kubernetes/tier-backend.yaml

build:
	make build-blue
	make build-green
	make build-blue-image
	make build-green-image
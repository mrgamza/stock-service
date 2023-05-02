NAME=go

build:
	docker build --no-cache -t ${NAME}-web -f Dockerfile-nginx .
	docker build --no-cache -t ${NAME}-service .

deploy: build
	docker-compose up -d

up: down
	docker-compose up -d

run: build
	docker-compose up -d

down:
	docker-compose down
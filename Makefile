.PHONY: analytics-up analytics-down blogs-up blogs-down users-up users-down infra-up infra-down up-all down-all

analytics-up:
	docker-compose -f analytics-service/docker-compose.yml up

analytics-down:
	docker-compose -f analytics-service/docker-compose.yml down

analytics-build:
	docker-compose -f analytics-service/docker.compose.yml build

blogs-up:
	docker-compose -f blogs-service/docker-compose.yml up

blogs-down:
	docker-compose -f blogs-service/docker-compose.yml down

blogs-build:
	docker-compose -f blogs-service/docker-compose.yml build

users-up:
	docker-compose -f users-service/docker-compose.yml up

users-down:
	docker-compose -f users-service/docker-compose.yml down

users-build:
	docker-compose -f users-service/docker-compose.yml build

infra-up:
	docker-compose -f deployments/docker-compose-queues.yml up

infra-down:
	docker-compose -f deployments/docker-compose-queues.yml down

infra-build:
	docker-compose -f deployments/docker-compose-queues.yml build

up-all:
	make -j3 blogs-up users-up infra-up

down-all:
	make -j blogs-down users-down infra-down

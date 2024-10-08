.PHONY: analytics-up analytics-down blogs-up blogs-down users-up users-down kafka-up kafka-down rabbit-up rabbit-down up-all down-all

api-gateway-up:
	docker-compose -f api-gateway-service/docker-compose.yml up

api-gateway-down:
	docker-compose -f api-gateway-service/docker-compose.yml down

api-gateway-build:
	docker-compose -f api-gateway-service/docker-compose.yml build

analytics-up:
	docker-compose -f analytics-service/docker-compose.yml up

analytics-down:
	docker-compose -f analytics-service/docker-compose.yml down

analytics-build:
	docker-compose -f analytics-service/docker-compose.yml build

blogs-up:
	docker-compose -f blogs-service/docker-compose.yml up

blogs-down:
	docker-compose -f blogs-service/docker-compose.yml down

blogs-build:
	docker-compose -f blogs-service/docker-compose.yml build

comments-up:
	docker-compose -f comments-service/docker-compose.yml up

comments-down:
	docker-compose -f comments-service/docker-compose.yml down

comments-build:
	docker-compose -f comments-service/docker-compose.yml build

users-up:
	docker-compose -f users-service/docker-compose.yml up

users-down:
	docker-compose -f users-service/docker-compose.yml down

users-build:
	docker-compose -f users-service/docker-compose.yml build

kafka-up:
	docker-compose -f deployments/docker-compose-kafka.yml up

kafka-down:
	docker-compose -f deployments/docker-compose-kafka.yml down

kafka-build:
	docker-compose -f deployments/docker-compose-kafka.yml build

rabbit-up:
	docker-compose -f deployments/docker-compose-rabbitmq.yml up

rabbit-down:
	docker-compose -f deployments/docker-compose-rabbitmq.yml down

rabbit-build:
	docker-compose -f deployments/docker-compose-rabbitmq.yml build

mon-up:
	docker-compose -f deployments/docker-compose-monitoring.yml up

mon-down:
	docker-compose -f deployments/docker-compose-monitoring.yml down

mon-build:
	docker-compose -f deployments/docker-compose-monitoring.yml build

up-all:
	make -j3 blogs-up users-up kafka-up rabbit-up

down-all:
	make -j blogs-down users-down kafka-down rabbit-down

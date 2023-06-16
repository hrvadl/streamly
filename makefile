dcup:
	docker compose -f ./config/docker-compose.kafka.yml -f ./config/docker-compose.user.yml -f ./config/docker-compose.auth.yml -f ./config/docker-compose.gateway.yml up -d

dcdown:
	docker compose -f ./config/docker-compose.kafka.yml -f ./config/docker-compose.user.yml -f ./config/docker-compose.auth.yml -f ./config/docker-compose.gateway.yml down
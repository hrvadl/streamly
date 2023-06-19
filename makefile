dcup:
	docker compose -f ./config/docker-compose.kafka.yml -f ./config/docker-compose.user.yml -f ./config/docker-compose.auth.yml -f ./config/docker-compose.gateway.yml -f ./config/docker-compose.email.yml up -d

dcdown:
	docker compose -f ./config/docker-compose.kafka.yml -f ./config/docker-compose.user.yml -f ./config/docker-compose.auth.yml -f ./config/docker-compose.gateway.yml -f ./config/docker-compose.email.yml down

dcdevup:
	docker compose -f ./config/docker-compose.dev.yml up -d

dcdevdown:
	docker compose -f ./config/docker-compose.dev.yml down

dcup:
	docker compose -f ./config/docker-compose.kafka.yml -f ./config/docker-compose.user.yml up -d

dcdown:
	docker compose -f ./config/docker-compose.kafka.yml -f ./config/docker-compose.user.yml down 
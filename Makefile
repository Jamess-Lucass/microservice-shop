.PHONY: compose-up
compose-up:
	docker compose -f ./basket-service/docker-compose.yml -f ./catalog-service/docker-compose.yml -f ./order-service/docker-compose.yml -f ./email-service/docker-compose.yml up

.PHONY: compose-down
compose-down:
	docker compose -f ./basket-service/docker-compose.yml -f ./catalog-service/docker-compose.yml -f ./order-service/docker-compose.yml -f ./email-service/docker-compose.yml down
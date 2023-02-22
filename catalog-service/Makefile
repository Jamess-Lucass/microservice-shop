.PHONY: run
run:
	dotenv -- go run ./cmd

.PHONY: format
format:
	go fmt ./...
	
.PHONY: compose
compose:
ifdef SERVICE
	docker compose up -d $(SERVICE)
else
	docker compose up -d
endif

.PHONY: compose-build
compose-build:
ifdef SERVICE
	docker compose up -d $(SERVICE) --build
else
	docker compose up -d --build
endif

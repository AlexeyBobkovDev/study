include .env
export 

service-run:
	@export conn_string=$(CONN_STRING) && \
	go run main.go

some-target:
	export env_var=hello && \
	printenv env_var

migrate-up:
	migrate -path migrations -database $(CONN_STRING) up

migrate-down:
	migrate -path migrations -database $(CONN_STRING) down

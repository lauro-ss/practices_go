.PHONY: all

all: build run

BIN = './bin'
build:
	go build -o "./bin" ./cmd/api/main.go

buildSeed:
	go build -o "./bin" ./cmd/seed/main.go

run:
	./bin/main

runw: build
	./bin/main.exe

seed: buildSeed run

create:
ifdef fn
	migrate create -ext sql -dir ./internal/data/migration -tz America/Sao_Paulo $(fn)
else
	@echo Provide a file name ex: 'make create fn=name'
endif

up:
	migrate -path ./internal/data/migration -database "postgresql://app:123456@localhost:5432/appanimal?sslmode=disable" -verbose up

down:
	migrate -path ./internal/data/migration -database "postgresql://app:123456@localhost:5432/appanimal?sslmode=disable" -verbose down

down_n:
	migrate -path ./internal/data/migration -database "postgresql://app:123456@localhost:5432/appanimal?sslmode=disable" -verbose down $(n)

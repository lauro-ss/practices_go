.PHONY: all

all: build run

BIN = './bin'
build:
	go build -o "./bin" main.go

run:
	./bin/main

create:
ifdef fn
	migrate create -ext sql -dir ./migration -tz America/Sao_Paulo $(fn)
else
	@echo Provide a file name ex: 'make create fn=name'
endif

up:
	migrate -path ./migration -database "postgresql://app:123456@localhost:5432/appanimal?sslmode=disable" -verbose up

down:
	migrate -path ./migration -database "postgresql://app:123456@localhost:5432/appanimal?sslmode=disable" -verbose down

down_n:
	migrate -path ./migration -database "postgresql://app:123456@localhost:5432/appanimal?sslmode=disable" -verbose down $(n)

force_n:
	migrate -path ./migration -database "postgresql://app:123456@localhost:5432/appanimal?sslmode=disable" -verbose force $(n)
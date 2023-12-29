make create:
ifdef fn
	migrate create -ext sql -dir ./migration -tz America/Sao_Paulo $(fn)
else
	@echo Provide a file name ex: 'make create fn=name'
endif
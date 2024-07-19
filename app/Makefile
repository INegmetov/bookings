run:
	 go run .\cmd\web\

build:
	go build -o ..\bin\booking.exe .\cmd\web\

## up: starts all containers in the background without forcing build
up:
	@echo Starting Docker images...
	docker-compose up -d
	@echo Docker images started!

down:
	@echo Stopping docker compose...
	docker-compose down
	@echo Done!
run:
	env PORTS_FILE_PATH=.data/ports.json go run ./src/main.go
test:
	chmod a+x ./tests.sh && ./tests.sh
run-docker:
	docker-compose up -d
stop-docker:
	docker-compose down
sec:
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	gosec ./...
docs:
	cd src && go get . && swag init && cd -
	@echo "a folder named 'docs' was generated inside the src containing the open API yml/json file. to check the docs run the api and visit http://localhost:8080/swagger/index.html"
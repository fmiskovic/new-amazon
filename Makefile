build: # build the go code
	@echo "building app..."
	@go build -o bin/app ./cmd/app
	@echo "app build has finished."

run: # run server
	@echo "running go app..."
	@./bin/app serve

db: # db migration related commands, like init, migrate, status, rollback...
	@./bin/app db $(cmd)

run-db: # run posgres db in docker with default configs
	@echo "starting go-db..."
	@docker run --name new-amz -e POSTGRES_PASSWORD=dbadmin -e POSTGRES_USER=dbadmin -e PGDATA=/var/lib/postgresql/data -e POSTGRES_DB=new-amz --volume=/var/lib/postgresql/data -p 5432:5432 -d postgres
	@echo "posgres go-db started."

stop-db: # stop posgres db in docker
	@echo "stopping go-db..."
	@docker stop new-amz
	@docker rm new-amz
	@echo "posgres go-db stopped."

clean: # delete app build
	@rm -rf bin

test: # run tests
	@go test -v ./...

race: # check race conditions
	@go test -v ./... --race

cover: # check test coverage
	@go test -cover ./...

mockery: # install mockery
	@go install github.com/vektra/mockery/v2@v2.42.2

mocks: # generate mocks
	@mockery #--all --with-expecter --keeptree --inpackage

fmt: # format go code
	@go fmt ./...

lint: # check go lint
	@echo "golangci-lint run..."
	@golangci-lint run --timeout 5m
	@echo "done"

.PHONY: all build run test clean

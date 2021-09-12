.PHONE: lint build start test swagger wire clean

lint:
	@golangci-lint run ./..

build:
	./build/build.sh

start: 
	@CONF_FILE_PATH=${PWD}/configs/dev.yaml go run main.go 

test:
	@mkdir -p coverage
	@CONF_FILE_PATH=${PWD}/configs/test.yaml go test -v -coverprofile=coverage/cover.out ./..

swagger:
	@swag init --parseDependency --generalInfo ./main.go --output ./internal/app/swagger

wire:
	@wire gen ./internal/app

clean:
	rm -rf coverage dist
SERVER_DIR=server
CLIENT_DIR=client 

.PHONY: install dev tidy server swag server-build

install:
	cd $(CLIENT_DIR) && npm install  

dev:
	cd $(CLIENT_DIR) && npm run dev

tidy:
	cd $(SERVER_DIR) && go mod tidy

server:
	cd $(SERVER_DIR) && air

swag:
	cd $(SERVER_DIR) && swag init -d . -g cmd/api/main.go --parseDependency --parseInternal -o docs

server-build:
	cd $(SERVER_DIR) && go build -o ./cmd/api/main.go
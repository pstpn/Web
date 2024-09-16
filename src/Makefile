go-tests-template:
	gotests -all -w <path>

mockery-template:
	mockery --all --inpackage

go-test-command:
	go test -v ./...

gen-swag-v1:
	swag init --parseDependency  --parseInternal --parseDepth 1 -g main.go -dir ./backend/internal/controller/v1/http --instanceName v1 -o ./backend/docs/v1

gen-swag-v2:
	swag init --parseDependency  --parseInternal --parseDepth 1 -g main.go -dir ./backend/internal/controller/v2/http --instanceName v2 -o ./backend/docs/v2

gen-graphql-v1:
	cd ./backend/internal/controller/v1/graphql && go run github.com/99designs/gqlgen generate

gen-graphql-v2:
	cd ./backend/internal/controller/v2/graphql && go run github.com/99designs/gqlgen generate

gen-graphql: gen-graphql-v1 gen-graphql-v2

convert-swag-to-3.0:
	cd backend && ./docs/to3.sh

fmt-swag:
	swag f

swagger: fmt-swag gen-swag-v1 gen-swag-v2 convert-swag-to-3.0

back: gen-graphql swagger
	cd backend && go run cmd/main.go

front:
	cd frontend/course && npm run serve

.PHONY: go-tests-template mockery-template go-test-command gen-swag-v1 gen-swag-v2 gen-graphql gen-graphql-v1 gen-graphql-v2 convert-swag-to-3.0 fmt-swag swagger back front
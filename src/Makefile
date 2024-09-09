go-tests-template:
	gotests -all -w <path>

mockery-template:
	mockery --all --inpackage

go-test-command:
	go test -v ./...

gen-swag:
	swag init -g ./cmd/main.go -o ./docs

gen-graphql:
	cd internal/controller/graphql && go run github.com/99designs/gqlgen generate

convert-swag-to-3.0:
	./docs/to3.sh

fmt-swag:
	swag f

swagger: fmt-swag gen-swag convert-swag-to-3.0

run: swagger
	go run cmd/main.go

.PHONY: go-tests-template mockery-template go-test-command gen-swag gen-graphql convert-swag-to-3.0 fmt-swag swagger run
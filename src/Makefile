go-tests_template:
	gotests -all -w <path>

mockery_template:
	mockery --all --inpackage

go-test_command:
	go test -v ./...

gen-swag:
	swag init -g ./cmd/main.go -o ./docs

convert-swag-to-3.0:
	./docs/to3.sh

fmt-swag:
	swag f

swagger: fmt-swag gen-swag convert-swag-to-3.0

run: swagger
	go run cmd/main.go

.PHONY: go-tests_template mockery_template go-test_command swag-gen convert-swag-to-3.0 swagger fmt-swag run
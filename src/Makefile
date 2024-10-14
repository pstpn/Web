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

test-front:
	cd frontend/course && npm run test

run-docker:
	docker compose up -d

rm-docker:
	docker compose down && docker image rm src-backend src-backend-mirror:latest src-backend-ro-1:latest src-backend-ro-2:latest

ab1:
	ab -n 10000 -c 100 -m GET \
		-H "Accept-Encoding: gzip,deflate" \
	  	-H "accept: application/json" \
	  	-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIyIiwiaWF0IjoxNzI3NzI4NzU3fQ.rLpyEuHAN8dZoQjBnPSnkSXGaKV1cXcRestfINzc0FU" \
	  	"http://localhost:8081/api/v2/profile"

ab3:
	ab -n 10000 -c 100 -m GET \
		-H "Accept-Encoding: gzip,deflate" \
	  	-H "accept: application/json" \
	  	-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIyIiwiaWF0IjoxNzI3NzI4NzU3fQ.rLpyEuHAN8dZoQjBnPSnkSXGaKV1cXcRestfINzc0FU" \
	  	"http://127.0.0.1/api/v2/profile"

.PHONY: go-tests-template mockery-template go-test-command gen-swag-v1 gen-swag-v2 gen-graphql gen-graphql-v1 gen-graphql-v2 convert-swag-to-3.0 fmt-swag swagger back front test-front run-docker rm-docker ab1 ab3
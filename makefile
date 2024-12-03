
COVERFILE = coverage.out

test:
	go test ./... -coverprofile=$(COVERFILE)

coverage:
	go tool cover -func=$(COVERFILE)

app:
	go run ./cmd/app/main.go
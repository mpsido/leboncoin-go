run:
	PORT=8080 go run main.go

install-linter:
	go get golang.org/x/tools/cmd/goimports

lint:
	goimports -w main.go

dep: 
	dep ensure -v

test:
	go test

linttest:
	goimports -w main_test.go

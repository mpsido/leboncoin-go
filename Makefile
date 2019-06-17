run:
	PORT=8080 go run main.go

install-dev:
	go get golang.org/x/tools/cmd/goimports
	go get github.com/golang/dep/cmd/dep

format:
	goimports -w main.go

dep: 
	dep ensure -v

test:
	go test

format-test:
	goimports -w main_test.go

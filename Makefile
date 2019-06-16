run:
	PORT=8080 go run main.go

lint:
	goimports -w main.go

dep: 
	dep ensure -v
test:
	go test

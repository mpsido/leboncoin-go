run:
	go run main.go

lint:
	goimports -w main.go

dep: 
	dep ensure -v
run:
	go run *.go

test: 
	go test -cover -race `go list ./...`

test_co:
	go test -v ./checkout

test_pc:
	go test -v ./checkout
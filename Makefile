run-tests:
	go test ./... -cover

doc:
# if godoc is not present, install: go install golang.org/x/tools/cmd/godoc@latest
	godoc -http=:8080

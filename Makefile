all: build exec
def: build exec_def

build:
	go run cmd/codegen/main.go
	gofmt -w internal/syntaxtree/*.go
	go build cmd/exec/main.go

exec:
	./main $(ARG)

exec_def:
	./main example/input.txt
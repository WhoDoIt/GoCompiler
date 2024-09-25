all: build exec
def: build exec_def

build:
	go build cmd/exec/main.go

exec:
	./main $(ARG)

exec_def:
	./main example/input.txt
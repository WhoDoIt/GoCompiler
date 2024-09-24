all: build exec
def: build exec_def

build:
	go build cmd/exec/main.go

exec:
	./main.exe $(ARG)

exec_def:
	./main.exe example/input.txt
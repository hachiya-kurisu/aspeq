all: aspeq

again: clean all

aspeq: aspeq.go cmd/aspeq/main.go
	go build -o aspeq cmd/aspeq/main.go

clean:
	rm -f aspeq

test:
	go test -cover

push:
	got send
	git push github

fmt:
	gofmt -s -w *.go
	gofmt -s -w cmd/aspeq/main.go

release: push
	git push github --tags

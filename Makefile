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
	gofmt -s -w *.go cmd/*/main.go

README.md: README.gmi
	sisyphus -f markdown <README.gmi >README.md

doc: README.md

release: push
	git push github --tags

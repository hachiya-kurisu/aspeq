all: aspeq slasher

again: clean all

aspeq: aspeq.go cmd/aspeq/main.go
	go build -o aspeq cmd/aspeq/main.go

slasher: aspeq.go cmd/slasher/main.go
	go build -o slasher cmd/slasher/main.go

clean:
	rm -f aspeq slasher

test:
	go test -cover

push:
	got send
	git push github

fmt:
	gofmt -s -w *.go cmd/*/main.go

cover:
	go test -coverprofile=cover.out
	go tool cover -html cover.out

README.md: README.gmi
	sisyphus -f markdown <README.gmi >README.md

doc: README.md

release: push
	git push github --tags
	got send -T

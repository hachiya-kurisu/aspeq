all: aspeq

again: clean all

aspeq: aspeq.go cmd/aspeq.go
	@go build cmd/aspeq.go

clean:
	rm -f aspeq

push:
	got send
	git push github

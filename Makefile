aspeq: aspeq.go cmd/aspeq.go
	@go build cmd/aspeq.go

push:
	got send
	git push github

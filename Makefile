.PHONY: gen

gen:
	@go run cmd/builder/main.go
	@goimports -w .

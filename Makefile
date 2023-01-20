.PHONY:generate
generate:
	mockery --all --dir ./example4 --output ./example4/mocks

.PHONY:test
test:
	go test ./...
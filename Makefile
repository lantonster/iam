# wire 依赖注入
.PHONY: wire
wire:
	cd cmd/wire && wire

# gen 生成 GORM 代码
.PHONY: gen
gen:
	go run cmd/gen/main.go

# swag 文档
.PHONY: swag
swag:
	swag fmt
	swag init -g internal/router/router.go --parseDependency -o docs/api

# mock 生成 mock 代码
.PHONY: mock
mock:
	sh scripts/mock_gen.sh internal/repo test/repo repo
	sh scripts/mock_gen.sh internal/service test/service service

# test 测试
.PHONY: test
test:
	go test ./... -v -coverprofile=test/coverage.out
	go tool cover -html=test/coverage.out
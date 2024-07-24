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
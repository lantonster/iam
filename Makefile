# wire 依赖注入
.PHONY: wire
wire:
	cd cmd/wire && wire

# gen 生成 GORM 代码
.PHONY: gen
gen:
	go run cmd/gen/main.go
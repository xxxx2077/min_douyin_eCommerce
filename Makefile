# Build

.PHONY: gen-client
gen-client: ## gen client code of {svc}. example: make gen-client svc=product
	@cd rpc_gen && cwgo client --type RPC --service ${svc} --module github.com/xxxx2077/min_douyin_eCommerce/rpc_gen  -I ../idl  --idl ../idl/${svc}.proto

.PHONY: gen-server
gen-server: ## gen service code of {svc}. example: make gen-server svc=product
	@cd app/${svc} && cwgo server --type RPC --service ${svc} --module github.com/xxxx2077/min_douyin_eCommerce/app/${svc} --pass "-use github.com/xxxx2077/min_douyin_eCommerce/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/${svc}.proto

.PHONY: tidy
tidy: ## run `go mod tidy` for all go module
	@scripts/tidy.sh
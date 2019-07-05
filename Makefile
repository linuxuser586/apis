.PHONY: apis
apis: bins
	@./scripts/generate-grpc.sh

.PHONY: protos
bins: 
	@./scripts/download-protobuf.sh
	@./scripts/download-mockgen.sh

REPO_URL  := https://github.com/merzzzl/accessibility-bridge.git
REPO_DIR  := accessibility-bridge
PROTO_SRC := $(REPO_DIR)/app/src/main/proto
GO_OUT    := .
GO_PKG    := github.com/merzzzl/accessibility-bridge-go

.PHONY: all clone update proto clean

all: proto

clone:
	@git clone $(REPO_URL) $(REPO_DIR) 2>/dev/null || true

update: clone
	@cd $(REPO_DIR) && git pull --ff-only

proto: clone
	@mkdir -p $(GO_OUT)
	@echo ">> generating Go from proto into $(GO_OUT)"
	@PROTO_FILES="$$(ls $(PROTO_SRC)/*.proto)"; \
		M_FLAGS=""; \
		for f in $$PROTO_FILES; do \
			b=$$(basename $$f); \
			M_FLAGS="$$M_FLAGS --go_opt=M$$b=$(GO_PKG) --go-grpc_opt=M$$b=$(GO_PKG)"; \
		done; \
		protoc -I $(PROTO_SRC) \
			--go_out=$(GO_OUT) --go_opt=paths=source_relative $$M_FLAGS \
			--go-grpc_out=$(GO_OUT) --go-grpc_opt=paths=source_relative $$M_FLAGS \
			$$PROTO_FILES
	@go mod tidy

clean:
	@rm -rf $(REPO_DIR) $(GO_OUT)

.PHONY: start build build-release build-linux-static build-linux-minimal clean pack

NOW = $(shell date -u '+%Y%m%d%I%M%S')

RELEASE_VERSION = v1.0.0

APP 		= gowallet
APP_BIN  	= ${APP}
RELEASE_ROOT 	= release
RELEASE_APP 	= release/${APP}
GIT_COUNT 	= $(shell git rev-list --all --count)
GIT_HASH        = $(shell git rev-parse --short HEAD)
RELEASE_TAG     = $(RELEASE_VERSION).$(GIT_COUNT).$(GIT_HASH)

all: start

build:
	@go build -ldflags "-w -s -X github.com/spark8899/gowallet/cmd.Version=$(RELEASE_TAG) -X github.com/spark8899/gowallet/cmd.GitCommit=$(GIT_HASH) -X github.com/spark8899/gowallet/cmd.BuildTime=$(NOW)" -trimpath -o $(APP_BIN) main.go
	@echo "Build complete: $(APP_BIN)"
	@ls -lh $(APP_BIN)

build-release: build
	@echo "Attempting UPX compression..."
	@if command -v upx >/dev/null 2>&1; then \
		cp $(APP_BIN) $(APP_BIN).bak; \
		upx --best --lzma $(APP_BIN) && echo "✅ UPX compression successful" || (mv $(APP_BIN).bak $(APP_BIN) && echo "⚠️  UPX compression failed, using uncompressed binary"); \
		rm -f $(APP_BIN).bak; \
		ls -lh $(APP_BIN); \
	else \
		echo "⚠️  UPX not installed. Install with: brew install upx (macOS) or sudo apt-get install upx-ucl (Linux)"; \
		echo "📦 Using standard optimized binary (8MB)"; \
	fi

build-linux-static:
	@echo "Building static Linux binary..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
		-ldflags="-s -w -extldflags '-static' \
		-X github.com/spark8899/gowallet/cmd.Version=$(RELEASE_TAG) \
		-X github.com/spark8899/gowallet/cmd.GitCommit=$(GIT_HASH) \
		-X github.com/spark8899/gowallet/cmd.BuildTime=$(NOW)" \
		-trimpath \
		-o $(APP_BIN)_linux_amd64 main.go
	@echo "✅ Linux static build complete"
	@ls -lh $(APP_BIN)_linux_amd64

build-linux-minimal: build-linux-static
	@echo "Optimizing binary size..."
	@if command -v strip >/dev/null 2>&1; then \
		strip --strip-all $(APP_BIN)_linux_amd64 2>/dev/null && echo "✅ Stripped debug symbols" || echo "⚠️  Strip not available"; \
	fi
	@if command -v upx >/dev/null 2>&1; then \
		upx --best --lzma $(APP_BIN)_linux_amd64 && echo "✅ UPX compression complete"; \
	else \
		echo "⚠️  UPX not installed. Install with: sudo apt-get install upx-ucl"; \
	fi
	@echo "📦 Final optimized size:"
	@ls -lh $(APP_BIN)_linux_amd64

test:
	cd ./internal/app/test && go test -v

clean:
	rm -rf data release $(SERVER_BIN) internal/app/test/data cmd/${APP}/data $(APP_BIN)_linux_*

pack: build
	rm -rf $(RELEASE_ROOT) && mkdir -p $(RELEASE_APP)
	cp -r $(APP_BIN) $(APP_BIN)
	cd $(RELEASE_ROOT) && tar -cvf $(APP).tar ${APP} && rm -rf ${APP}

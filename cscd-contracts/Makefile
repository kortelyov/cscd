# Переменные
PROTO_DIR = proto
GEN_GO_DIR = pkg/contracts
PACKAGE_NAME = github.com/kortelyov/cscd/cscd-contracts/pkg/contracts

# Список инструментов
PROTOC = $(shell message which protoc 2> /dev/null)
PROTOC_GEN_GO = $(shell which protoc-gen-go 2> /dev/null)

.PHONY: all gen setup check clean

all: gen

# 1. Генерация кода
gen: check
	@echo "==> Generating Go code from Proto..."
	@mkdir -p $(GEN_GO_DIR)
	protoc -I=$(PROTO_DIR) \
		--go_out=$(GEN_GO_DIR) \
		--go_opt=paths=source_relative \
		$(PROTO_DIR)/*.proto
	@echo "==> Done!"

# 2. Проверка наличия тулзов
check:
ifndef PROTOC
	$(error "protoc is not installed. Run 'make setup' to install")
endif
ifndef PROTOC_GEN_GO
	$(error "protoc-gen-go is not installed. Run 'make setup' to install")
endif

# 3. Установка окружения (только для Mac)
setup:
	@echo "==> Checking for Homebrew..."
	@which brew > /dev/null || (echo "Homebrew not found! Install it first: brew.sh" && exit 1)

	@echo "==> Installing protobuf (protoc)..."
	@brew install protobuf

	@echo "==> Installing protoc-gen-go..."
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

	@echo "==> All tools installed!"
	@echo "Make sure your PATH includes \$$GOPATH/bin (usually \$$HOME/go/bin)"

# 4. Очистка сгенерированного кода
clean:
	@rm -rf $(GEN_GO_DIR)
	@echo "==> Cleaned up generated files"
.PHONY: build run dev test lint fmt clean migrate migrate-dev container-up container-down setup

BINARY_NAME=main
BUILD_DIR=bin
MAIN_PATH=./cmd/server

build:
		@echo "ビルド中..."
		@mkdir -p $(BUILD_DIR)
		@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
		@echo "ビルド完了: $(BUILD_DIR)/$(BINARY_NAME)"

run: build
		@echo "アプリケーションを実行中..."
		@./$(BUILD_DIR)/$(BINARY_NAME)

dev:
		@echo "ホットリロード開発サーバーを起動中..."
		@air

test:
		@echo "テストを実行中..."
		@go test -v ./...

lint:
		@echo "リンターを実行中..."
		@go vet ./...
		@gofmt -s -l .

fmt:
		@echo "コードをフォーマット中..."
		@go fmt ./...

clean:
		@echo "クリーンアップ中..."
		@rm -rf $(BUILD_DIR)
		@echo "クリーンアップ完了"

deps:
		@echo "依存関係を取得中..."
		@go mod tidy

setup: deps
		@echo ""
		@echo "✅ 初期設定が完了しました！"
		@echo ""
		@echo "次のコマンドで開発サーバーを起動できます:"
		@echo "  make dev"
		@echo ""
		@echo "アクセスURL: http://localhost:8080"

# ヘルプを表示
help:
		@echo "利用可能なコマンド:"
		@echo "  setup         - 🚀 初期設定（開発環境の完全セットアップ）"
		@echo "  build         - アプリケーションをビルド"
		@echo "  run           - アプリケーションを実行"
		@echo "  dev           - ホットリロード開発サーバーを起動（air使用）"
		@echo "  test          - テストを実行"
		@echo "  lint          - リンターを実行"
		@echo "  fmt           - コードをフォーマット"
		@echo "  clean         - ビルド成果物をクリーンアップ"
		@echo "  deps          - 依存関係を取得"
		@echo "  help          - このヘルプを表示"

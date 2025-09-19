# PostHog Go Test Project Makefile

# 变量定义
BINARY_NAME=posthog-go-test
BINARY_UNIX=$(BINARY_NAME)_unix
MAIN_FILE=main.go

# 默认目标
.PHONY: all
all: build

# 构建项目
.PHONY: build
build:
	@echo "构建项目..."
	go build -o $(BINARY_NAME) $(MAIN_FILE)
	@echo "构建完成: $(BINARY_NAME)"

# 构建Linux版本
.PHONY: build-linux
build-linux:
	@echo "构建Linux版本..."
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_UNIX) $(MAIN_FILE)
	@echo "构建完成: $(BINARY_UNIX)"

# 运行项目
.PHONY: run
run:
	@echo "运行项目..."
	go run $(MAIN_FILE)

# 清理构建文件
.PHONY: clean
clean:
	@echo "清理构建文件..."
	go clean
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	@echo "清理完成"

# 安装依赖
.PHONY: deps
deps:
	@echo "安装依赖..."
	go mod download
	go mod tidy
	@echo "依赖安装完成"

# 更新依赖
.PHONY: update-deps
update-deps:
	@echo "更新依赖..."
	go get -u ./...
	go mod tidy
	@echo "依赖更新完成"

# 运行测试
.PHONY: test
test:
	@echo "运行测试..."
	go test ./...

# 运行测试并显示覆盖率
.PHONY: test-coverage
test-coverage:
	@echo "运行测试并生成覆盖率报告..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "覆盖率报告已生成: coverage.html"

# 代码格式化
.PHONY: fmt
fmt:
	@echo "格式化代码..."
	go fmt ./...
	@echo "代码格式化完成"

# 代码检查
.PHONY: vet
vet:
	@echo "运行代码检查..."
	go vet ./...
	@echo "代码检查完成"

# 代码检查（使用golangci-lint）
.PHONY: lint
lint:
	@echo "运行linter..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint未安装，跳过linter检查"; \
	fi

# 安装开发工具
.PHONY: install-tools
install-tools:
	@echo "安装开发工具..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "开发工具安装完成"

# 开发模式（自动重新构建和运行）
.PHONY: dev
dev:
	@echo "启动开发模式..."
	@if command -v air >/dev/null 2>&1; then \
		air; \
	else \
		echo "air未安装，使用go run代替"; \
		go run $(MAIN_FILE); \
	fi

# 安装air（热重载工具）
.PHONY: install-air
install-air:
	@echo "安装air热重载工具..."
	go install github.com/cosmtrek/air@latest
	@echo "air安装完成"

# 压力测试相关命令
.PHONY: stress-test
stress-test:
	@echo "运行压力测试..."
	@echo "请提供API密钥: make stress-test API_KEY=your_api_key_here"
	@if [ -z "$(API_KEY)" ]; then \
		echo "错误: 请提供API密钥"; \
		echo "用法: make stress-test API_KEY=your_api_key_here"; \
		exit 1; \
	fi
	@echo "使用API密钥: $(shell echo $(API_KEY) | cut -c1-4)***"
	go run $(MAIN_FILE) -key $(API_KEY)

.PHONY: stress-test-custom
stress-test-custom:
	@echo "运行自定义压力测试..."
	@echo "请提供API密钥: make stress-test-custom API_KEY=your_api_key_here CONCURRENT=20 REQUESTS=1000"
	@if [ -z "$(API_KEY)" ]; then \
		echo "错误: 请提供API密钥"; \
		echo "用法: make stress-test-custom API_KEY=your_api_key_here CONCURRENT=20 REQUESTS=1000"; \
		exit 1; \
	fi
	@echo "使用API密钥: $(shell echo $(API_KEY) | cut -c1-4)***"
	@echo "并发数: $(or $(CONCURRENT),10)"
	@echo "请求数: $(or $(REQUESTS),100)"
	go run $(MAIN_FILE) -key $(API_KEY) -concurrent $(or $(CONCURRENT),10) -requests $(or $(REQUESTS),100)


.PHONY: stress-test-help
stress-test-help:
	@echo "显示压力测试帮助信息..."
	go run $(MAIN_FILE) -help

# 显示帮助信息
.PHONY: help
help:
	@echo "可用的命令:"
	@echo "  build          - 构建项目"
	@echo "  build-linux    - 构建Linux版本"
	@echo "  run            - 运行项目"
	@echo "  clean          - 清理构建文件"
	@echo "  deps           - 安装依赖"
	@echo "  update-deps    - 更新依赖"
	@echo "  test           - 运行测试"
	@echo "  test-coverage  - 运行测试并生成覆盖率报告"
	@echo "  fmt            - 格式化代码"
	@echo "  vet            - 代码检查"
	@echo "  lint           - 运行linter"
	@echo "  install-tools  - 安装开发工具"
	@echo "  dev            - 开发模式（热重载）"
	@echo "  install-air    - 安装air热重载工具"
	@echo ""
	@echo "压力测试命令:"
	@echo "  stress-test        - 运行基本压力测试 (自托管端点, 需要API_KEY)"
	@echo "  stress-test-custom - 运行自定义压力测试 (需要API_KEY, 可选CONCURRENT, REQUESTS)"
	@echo "  stress-test-help   - 显示压力测试帮助信息"
	@echo ""
	@echo "  help           - 显示此帮助信息"


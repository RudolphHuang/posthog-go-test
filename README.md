# PostHog Go Test

这是一个使用PostHog Go SDK的测试项目，用于发送事件数据到PostHog分析平台。

## 功能特性

- 使用PostHog Go SDK发送事件
- 支持自定义事件和用户ID
- 配置化的PostHog端点

## 环境要求

- Go 1.19 或更高版本
- 有效的PostHog项目API密钥

## 快速开始

### 1. 克隆项目

```bash
git clone <repository-url>
cd posthog-go-test
```

### 2. 安装依赖

```bash
make deps
```

### 3. 配置PostHog

在 `main.go` 文件中更新以下配置：

```go
// 替换为你的PostHog项目API密钥
client, _ := posthog.NewWithConfig("YOUR_API_KEY", posthog.Config{Endpoint: "https://us.i.posthog.com"})
```

### 4. 运行项目

```bash
make run
```

或者直接运行：

```bash
go run main.go
```

## 可用命令

使用 `make help` 查看所有可用命令：

```bash
make help
```

### 常用命令

- `make build` - 构建项目
- `make run` - 运行项目
- `make test` - 运行测试
- `make clean` - 清理构建文件
- `make fmt` - 格式化代码
- `make vet` - 代码检查

### 开发工具

- `make install-tools` - 安装开发工具（golangci-lint等）
- `make install-air` - 安装热重载工具
- `make dev` - 开发模式（支持热重载）

## 项目结构

```
posthog-go-test/
├── main.go          # 主程序文件
├── go.mod           # Go模块文件
├── go.sum           # 依赖校验文件
├── Makefile         # 构建脚本
├── .gitignore       # Git忽略文件
└── README.md        # 项目说明
```

## 配置说明

### PostHog配置

项目使用PostHog Go SDK，主要配置项：

- **API密钥**: 在PostHog项目设置中获取
- **端点**: 根据你的PostHog实例选择正确的端点
  - 美国: `https://us.i.posthog.com`
  - 欧盟: `https://eu.i.posthog.com`
  - 自托管: `https://your-instance.posthog.com`

### 事件配置

在 `main.go` 中可以自定义：

- `DistinctId`: 用户唯一标识符
- `Event`: 事件名称
- 其他事件属性

## 开发指南

### 代码规范

- 使用 `make fmt` 格式化代码
- 使用 `make vet` 检查代码
- 使用 `make lint` 运行linter（需要先安装golangci-lint）

### 测试

```bash
# 运行所有测试
make test

# 运行测试并生成覆盖率报告
make test-coverage
```

### 热重载开发

安装air工具后可以使用热重载：

```bash
make install-air
make dev
```

## 故障排除

### 常见问题

1. **API密钥无效**: 检查PostHog项目设置中的API密钥
2. **网络连接问题**: 确认PostHog端点URL正确
3. **依赖问题**: 运行 `make deps` 重新安装依赖

### 调试

启用详细日志：

```go
client, _ := posthog.NewWithConfig("YOUR_API_KEY", posthog.Config{
    Endpoint: "https://us.i.posthog.com",
    Logger:   posthog.NewDefaultLogger(),
})
```

## 许可证

MIT License

## 贡献

欢迎提交Issue和Pull Request！
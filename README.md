# PostHog Go 压力测试工具

这是一个使用PostHog Go SDK的压力测试工具，用于对PostHog分析平台进行性能测试和压力测试。

## 功能特性

- 🚀 **高性能压力测试**: 支持多并发请求，测试PostHog API性能
- 📊 **详细统计报告**: 提供成功率、RPS、耗时等详细统计信息
- 🔧 **灵活配置**: 支持自定义并发数、请求数、事件名称等参数
- 🌍 **多端点支持**: 支持美国、欧盟等不同PostHog端点
- 💻 **命令行界面**: 简单易用的命令行参数配置
- 📈 **实时监控**: 实时显示测试进度和结果

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

### 3. 获取PostHog API密钥

1. 登录你的PostHog账户
2. 进入项目设置
3. 复制项目API密钥（格式类似：`phc_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx`）

### 4. 运行压力测试

#### 基本用法

```bash
# 使用默认参数运行压力测试
go run main.go -key phc_your_api_key_here
```

#### 使用Makefile（推荐）

```bash
# 基本压力测试
make stress-test API_KEY=phc_your_api_key_here

# 自定义参数压力测试
make stress-test-custom API_KEY=phc_your_api_key_here CONCURRENT=20 REQUESTS=1000

# 欧盟端点压力测试
make stress-test-eu API_KEY=phc_your_api_key_here
```

## 命令行用法

### 基本语法

```bash
go run main.go -key <API_KEY> [选项]
```

### 必需参数

- `-key string`: PostHog API密钥（必需）

### 可选参数

- `-endpoint string`: PostHog端点URL（默认: `https://your-posthog-instance.com`）
- `-concurrent int`: 并发数（默认: 10）
- `-requests int`: 总请求数（默认: 100）
- `-event string`: 事件名称（默认: `stress-test-event`）
- `-user string`: 用户ID（默认: `stress-test-user`）
- `-help, -h`: 显示帮助信息

### 使用示例

```bash
# 基本压力测试
go run main.go -key phc_your_api_key_here

# 高并发压力测试
go run main.go -key phc_your_api_key_here -concurrent 50 -requests 5000

# 使用自定义端点
go run main.go -key phc_your_api_key_here -endpoint https://your-custom-posthog.com

# 自定义事件和用户
go run main.go -key phc_your_api_key_here -event "custom-event" -user "test-user-123"

# 查看帮助信息
go run main.go -help
```

### 输出示例

```
开始压力测试...
API密钥: phc_***QYO
端点: https://your-posthog-instance.com
并发数: 10
总请求数: 100
事件名称: stress-test-event
用户ID: stress-test-user
----------------------------------------
----------------------------------------
压力测试完成!
总请求数: 100
成功请求: 98
失败请求: 2
成功率: 98.00%
总耗时: 2.345s
平均RPS: 42.64
```

## 可用命令

使用 `make help` 查看所有可用命令：

```bash
make help
```

### 压力测试命令

- `make stress-test API_KEY=your_key` - 运行基本压力测试（自托管端点）
- `make stress-test-custom API_KEY=your_key CONCURRENT=20 REQUESTS=1000` - 运行自定义压力测试
- `make stress-test-help` - 显示压力测试帮助信息

### 开发命令

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

项目使用PostHog Go SDK，默认使用自托管端点：

- **默认端点**: `https://your-posthog-instance.com`
- **自定义端点**: 可通过 `-endpoint` 参数指定任意PostHog实例地址

### 压力测试参数

- **并发数**: 控制同时发送请求的协程数量，影响测试强度
- **请求数**: 总共发送的事件数量
- **事件名称**: 发送的事件类型，便于在PostHog中识别
- **用户ID**: 事件关联的用户标识符

### 性能调优建议

1. **并发数设置**:
   - 小规模测试: 10-50
   - 中等规模测试: 50-200
   - 大规模测试: 200-1000

2. **请求数设置**:
   - 快速验证: 100-1000
   - 性能测试: 1000-10000
   - 压力测试: 10000+

3. **网络考虑**:
   - 本地网络: 可以设置较高并发
   - 跨地区网络: 建议降低并发数

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

1. **API密钥无效**: 
   - 检查PostHog项目设置中的API密钥
   - 确认API密钥格式正确（以`phc_`开头）

2. **网络连接问题**: 
   - 确认PostHog端点URL正确
   - 检查网络连接和防火墙设置

3. **高失败率**:
   - 降低并发数
   - 检查网络延迟
   - 确认PostHog服务状态

4. **依赖问题**: 
   - 运行 `make deps` 重新安装依赖
   - 检查Go版本兼容性

### 性能优化

1. **提高成功率**:
   - 适当降低并发数
   - 增加重试机制
   - 使用更稳定的网络环境

2. **提高吞吐量**:
   - 增加并发数（在可接受范围内）
   - 优化网络配置
   - 使用更近的PostHog端点

### 调试模式

如需启用详细日志，可以修改代码：

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
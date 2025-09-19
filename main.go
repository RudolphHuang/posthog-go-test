package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/posthog/posthog-go"
)

// 配置结构体
type Config struct {
	APIKey     string
	Endpoint   string
	Concurrent int
	Requests   int
	EventName  string
	UserID     string
	Help       bool
}

// 压力测试结果
type TestResult struct {
	TotalRequests   int
	SuccessRequests int
	FailedRequests  int
	Duration        time.Duration
	RPS             float64
}

func main() {
	config := parseFlags()

	if config.Help {
		showHelp()
		return
	}

	if config.APIKey == "" {
		fmt.Println("错误: 必须提供API密钥")
		showHelp()
		os.Exit(1)
	}

	// 创建PostHog客户端
	client, err := posthog.NewWithConfig(config.APIKey, posthog.Config{
		Endpoint: config.Endpoint,
	})
	if err != nil {
		log.Fatalf("创建PostHog客户端失败: %v", err)
	}
	defer client.Close()

	fmt.Printf("开始压力测试...\n")
	fmt.Printf("API密钥: %s\n", maskAPIKey(config.APIKey))
	fmt.Printf("端点: %s\n", config.Endpoint)
	fmt.Printf("并发数: %d\n", config.Concurrent)
	fmt.Printf("总请求数: %d\n", config.Requests)
	fmt.Printf("事件名称: %s\n", config.EventName)
	fmt.Printf("用户ID: %s\n", config.UserID)
	fmt.Println("----------------------------------------")

	// 执行压力测试
	result := runStressTest(client, config)

	// 输出结果
	printResults(result)
}

// 解析命令行参数
func parseFlags() *Config {
	config := &Config{}

	flag.StringVar(&config.APIKey, "key", "", "PostHog API密钥 (必需)")
	flag.StringVar(&config.Endpoint, "endpoint", "https://your-posthog-instance.com", "PostHog端点URL")
	flag.IntVar(&config.Concurrent, "concurrent", 10, "并发数")
	flag.IntVar(&config.Requests, "requests", 100, "总请求数")
	flag.StringVar(&config.EventName, "event", "stress-test-event", "事件名称")
	flag.StringVar(&config.UserID, "user", "stress-test-user", "用户ID")
	flag.BoolVar(&config.Help, "help", false, "显示帮助信息")
	flag.BoolVar(&config.Help, "h", false, "显示帮助信息")

	flag.Parse()

	return config
}

// 显示帮助信息
func showHelp() {
	fmt.Println("PostHog Go 压力测试工具")
	fmt.Println("")
	fmt.Println("用法:")
	fmt.Println("  go run main.go -key <API_KEY> [选项]")
	fmt.Println("")
	fmt.Println("必需参数:")
	fmt.Println("  -key string")
	fmt.Println("        PostHog API密钥")
	fmt.Println("")
	fmt.Println("可选参数:")
	fmt.Println("  -endpoint string")
	fmt.Println("        PostHog端点URL (默认: https://your-posthog-instance.com)")
	fmt.Println("  -concurrent int")
	fmt.Println("        并发数 (默认: 10)")
	fmt.Println("  -requests int")
	fmt.Println("        总请求数 (默认: 100)")
	fmt.Println("  -event string")
	fmt.Println("        事件名称 (默认: stress-test-event)")
	fmt.Println("  -user string")
	fmt.Println("        用户ID (默认: stress-test-user)")
	fmt.Println("  -help, -h")
	fmt.Println("        显示帮助信息")
	fmt.Println("")
	fmt.Println("示例:")
	fmt.Println("  # 基本用法")
	fmt.Println("  go run main.go -key phc_your_api_key_here")
	fmt.Println("")
	fmt.Println("  # 自定义参数")
	fmt.Println("  go run main.go -key phc_your_api_key_here -concurrent 20 -requests 1000")
	fmt.Println("")
	fmt.Println("  # 使用自定义端点")
	fmt.Println("  go run main.go -key phc_your_api_key_here -endpoint https://your-custom-posthog.com")
}

// 执行压力测试
func runStressTest(client posthog.Client, config *Config) *TestResult {
	startTime := time.Now()

	// 创建通道来分发任务
	taskChan := make(chan int, config.Requests)
	resultChan := make(chan bool, config.Requests)

	// 启动工作协程
	var wg sync.WaitGroup
	for i := 0; i < config.Concurrent; i++ {
		wg.Add(1)
		go worker(client, config, taskChan, resultChan, &wg)
	}

	// 发送任务
	go func() {
		for i := 0; i < config.Requests; i++ {
			taskChan <- i
		}
		close(taskChan)
	}()

	// 等待所有工作协程完成
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// 收集结果
	successCount := 0
	failedCount := 0
	for success := range resultChan {
		if success {
			successCount++
		} else {
			failedCount++
		}
	}

	duration := time.Since(startTime)
	rps := float64(config.Requests) / duration.Seconds()

	return &TestResult{
		TotalRequests:   config.Requests,
		SuccessRequests: successCount,
		FailedRequests:  failedCount,
		Duration:        duration,
		RPS:             rps,
	}
}

// 工作协程
func worker(client posthog.Client, config *Config, taskChan <-chan int, resultChan chan<- bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for range taskChan {
		success := sendEvent(client, config)
		resultChan <- success
	}
}

// 发送单个事件
func sendEvent(client posthog.Client, config *Config) bool {
	err := client.Enqueue(posthog.Capture{
		DistinctId: config.UserID,
		Event:      config.EventName,
		Properties: map[string]interface{}{
			"timestamp": time.Now().Unix(),
			"test_type": "stress_test",
		},
	})

	return err == nil
}

// 输出测试结果
func printResults(result *TestResult) {
	fmt.Println("----------------------------------------")
	fmt.Println("压力测试完成!")
	fmt.Printf("总请求数: %d\n", result.TotalRequests)
	fmt.Printf("成功请求: %d\n", result.SuccessRequests)
	fmt.Printf("失败请求: %d\n", result.FailedRequests)
	fmt.Printf("成功率: %.2f%%\n", float64(result.SuccessRequests)/float64(result.TotalRequests)*100)
	fmt.Printf("总耗时: %v\n", result.Duration)
	fmt.Printf("平均RPS: %.2f\n", result.RPS)
}

// 掩码API密钥用于显示
func maskAPIKey(key string) string {
	if len(key) <= 8 {
		return "***"
	}
	return key[:4] + "***" + key[len(key)-4:]
}

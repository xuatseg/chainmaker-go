# 网络性能测试工具

本工具用于测试长安链节点间的网络性能，包括带宽、延迟和稳定性等指标。

## 文件结构

```
net_performance_test/
└── main.go    # 网络性能测试实现
```

## 功能特性

### 1. 测试指标
- **带宽测试**：测量节点间最大数据传输速率
- **延迟测试**：测量消息往返时间(RTT)
- **稳定性测试**：长时间传输稳定性
- **并发测试**：多连接并发性能

### 2. 测试模式
- 点对点测试
- 星型拓扑测试
- 全网状测试
- 自定义拓扑测试

## 使用指南

### 1. 基本用法

```bash
# 启动服务端模式
go run main.go -mode server -port 8080

# 启动客户端测试
go run main.go -mode client -server 192.168.1.100:8080 -duration 60s
```

### 2. 参数说明

| 参数 | 说明 |
|------|------|
| -mode | 运行模式(server/client) |
| -port | 服务监听端口 |
| -server | 服务端地址 |
| -duration | 测试持续时间 |
| -size | 测试数据包大小 |
| -threads | 并发线程数 |

## 实现原理

### 1. 测试架构

```go
type NetworkTester struct {
    Mode       string // "server" or "client"
    ServerAddr string
    Port       int
    // ...其他配置
}
```

### 2. 测试流程

```
1. 建立TCP连接
2. 协商测试参数
3. 执行测试循环
4. 收集性能数据
5. 生成测试报告
```

## 测试报告

示例输出:
```
=== 网络测试结果 ===
测试时长: 60.0s
平均延迟: 12.3ms
最大延迟: 45.6ms
最小延迟: 8.7ms
吞吐量: 85.2 MB/s
丢包率: 0.05%
```

## 测试场景

### 1. 基础性能测试
```bash
# 测试基础带宽
go run main.go -mode client -server node1:8080 -test bandwidth

# 测试延迟
go run main.go -mode client -server node1:8080 -test latency
```

### 2. 压力测试
```bash
# 多线程压力测试
go run main.go -mode client -server node1:8080 -threads 16 -duration 300s
```

## 注意事项

1. 环境准备
   - 确保网络连通性
   - 关闭防火墙限制
   - 准备足够的测试时间

2. 测试建议
   - 从简单测试开始
   - 逐步增加负载
   - 记录测试环境状态

3. 结果分析
   - 关注异常波动
   - 比较历史数据
   - 考虑网络环境因素

## 扩展功能

1. 高级测试
   - 支持UDP协议测试
   - 添加TLS加密测试
   - 支持多协议比较

2. 可视化
   - 生成实时图表
   - 保存测试数据
   - 历史趋势分析

3. 自动化
   - 定时测试
   - 异常告警
   - 集成监控系统

# gRPC 客户端测试工具

本工具用于测试长安链节点的 gRPC 接口功能，支持发送交易、查询状态等操作。

## 文件结构

```
grpc_client/
└── grpc_client.go    # gRPC 客户端实现
```

## 功能特性

### 1. 支持的操作
- 发送交易
- 查询交易状态
- 获取区块信息
- 订阅链上事件
- 调用智能合约

### 2. 测试模式
- 单次请求测试
- 持续压力测试
- 并发请求测试
- 异常场景测试

## 使用指南

### 1. 基本用法

```bash
# 查询区块高度
go run grpc_client.go -cmd getBlock -height 100

# 发送测试交易
go run grpc_client.go -cmd sendTx -count 100

# 压力测试
go run grpc_client.go -cmd stress -tps 1000 -duration 60s
```

### 2. 参数说明

| 参数 | 说明 |
|------|------|
| -cmd | 测试命令 |
| -endpoint | gRPC 服务地址 |
| -tps | 目标TPS(压力测试) |
| -duration | 测试持续时间 |
| -concurrent | 并发数 |

## 实现原理

### 1. 客户端架构

```go
type GRPCClient struct {
    conn       *grpc.ClientConn
    chainClient pb.ChainServiceClient
    // ...其他客户端
}
```

### 2. 请求处理流程

```go
1. 建立gRPC连接
2. 构造请求参数
3. 发送gRPC请求
4. 处理响应结果
5. 输出测试报告
```

## 测试场景

### 1. 功能测试
```go
// 测试区块查询
func TestGetBlock(t *testing.T) {
    client := NewClient("localhost:12301")
    block, err := client.GetBlock(100)
    assert.NoError(t, err)
    assert.Equal(t, uint64(100), block.Height)
}
```

### 2. 性能测试
```go
// 测试交易发送性能
func BenchmarkSendTx(b *testing.B) {
    client := NewClient("localhost:12301")
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        client.SendTestTx()
    }
}
```

## 测试报告

示例输出:
```
=== 测试结果 ===
请求总数: 1000
成功请求: 998 (99.8%)
平均延迟: 23.5ms
最大延迟: 102ms
吞吐量: 850 TPS
```

## 注意事项

1. 环境配置
   - 确保节点gRPC服务已启动
   - 检查网络连接
   - 准备测试账户

2. 测试建议
   - 从低压力逐步增加
   - 监控节点状态
   - 记录测试日志

3. 安全提示
   - 不要使用生产环境
   - 保护私钥安全
   - 及时清理测试数据

## 扩展功能

1. 自定义测试
   - 支持自定义交易
   - 添加新的测试用例

2. 结果分析
   - 生成可视化报告
   - 保存原始数据

3. 自动化测试
   - CI集成
   - 定期回归测试

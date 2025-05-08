# RPC Server 模块

RPC 服务模块是长安链对外提供服务接口的核心组件，负责处理客户端请求、管理订阅服务和提供区块链交互接口。

## 目录结构

```
rpcserver/
├── api_service.go                      # API 服务实现
├── archive_service.go                  # 归档服务实现
├── gas.go                             # Gas 计费相关
├── middleware.go                       # 中间件实现
├── rateLimiter/                       # 限流器实现
├── rpc_server.go                      # RPC 服务器核心
├── subscribe_service.go               # 订阅服务基础
├── subscribe_service_block.go         # 区块订阅服务
├── subscribe_service_contract_event.go # 合约事件订阅
├── subscribe_service_tx.go            # 交易订阅服务
├── tx_result_dispatcher.go            # 交易结果分发
└── utils.go                          # 工具函数
```

## 架构设计

```
                    +----------------+
                    |  RPC Server   |
                    +-------+-------+
                            |
           +---------------+---------------+
           |               |               |
    +------+-----+  +-----+------+  +-----+------+
    |    API     |  | Subscribe  |  |  Archive   |
    |  Service   |  |  Service   |  |  Service   |
    +------------+  +------------+  +------------+
           |               |               |
    +------+-----+  +-----+------+  +-----+------+
    | Middleware  |  |   Rate     |  |    TX     |
    |            |  |  Limiter   |  | Dispatcher |
    +------------+  +------------+  +------------+
```

## 核心组件

### 1. RPC Server
- 服务器配置管理
- 请求路由分发
- 连接管理
- 服务生命周期

### 2. API Service
- 链操作接口
- 交易处理
- 查询服务
- 系统管理

### 3. Subscribe Service
- 区块订阅
- 交易订阅
- 事件订阅
- 实时通知

## 主要接口

### 1. 链操作接口
```go
type ChainService interface {
    // 发送交易
    SendTransaction(tx *Transaction) (*TxResponse, error)
    
    // 查询交易
    GetTransaction(txId string) (*TransactionInfo, error)
    
    // 查询区块
    GetBlock(height uint64) (*Block, error)
    
    // 查询链信息
    GetChainInfo() (*ChainInfo, error)
}
```

### 2. 订阅接口
```go
type SubscribeService interface {
    // 订阅区块
    SubscribeBlock(req *SubscribeRequest) (<-chan *Block, error)
    
    // 订阅交易
    SubscribeTransaction(req *SubscribeRequest) (<-chan *Transaction, error)
    
    // 订阅合约事件
    SubscribeContractEvent(req *SubscribeRequest) (<-chan *ContractEvent, error)
}
```

## 中间件功能

### 1. 认证中间件
```go
// 身份认证
func AuthMiddleware(handler Handler) Handler

// 权限验证
func AccessControlMiddleware(handler Handler) Handler
```

### 2. 限流中间件
```go
// 请求限流
func RateLimitMiddleware(limit rate.Limit) Middleware

// 并发控制
func ConcurrencyLimitMiddleware(max int) Middleware
```

## 配置说明

```yaml
rpcserver:
  # 服务配置
  server:
    port: 12301
    max_connections: 1000
    timeout: "30s"
    
  # 限流配置
  rate_limit:
    enabled: true
    requests_per_second: 100
    burst_size: 200
    
  # 订阅配置
  subscribe:
    buffer_size: 1000
    max_subscribers: 100
    
  # TLS配置
  tls:
    enabled: true
    cert_file: "/path/to/cert"
    key_file: "/path/to/key"
```

## 使用示例

### 1. 启动服务器
```go
config := &RPCConfig{...}
server := NewRPCServer(config)
server.Start()
```

### 2. 注册服务
```go
// 注册API服务
server.RegisterService("chain", NewChainService())

// 注册订阅服务
server.RegisterService("subscribe", NewSubscribeService())
```

## 安全特性

### 1. 访问控制
- TLS 加密
- 身份认证
- 权限验证

### 2. 流量控制
- 请求限流
- 并发控制
- 超时处理

### 3. 资源保护
- 内存限制
- 连接控制
- 队列管理

## 性能优化

### 1. 请求处理
- 异步处理
- 批量处理
- 缓存优化

### 2. 订阅优化
- 消息缓冲
- 订阅过滤
- 高效分发

### 3. 连接管理
- 连接池
- 心跳检测
- 优雅关闭

## 监控指标

### 1. 服务指标
- QPS
- 响应时间
- 错误率
- 连接数

### 2. 订阅指标
- 订阅者数量
- 消息积压量
- 处理延迟

### 3. 资源指标
- CPU使用率
- 内存使用
- 网络流量

## 调试功能

### 1. 日志记录
```go
logger.Debug("request processing",
    "method", req.Method,
    "client", req.ClientID,
    "latency", latency,
)
```

### 2. 性能分析
```go
metrics.RecordLatency("rpc_request", latency)
metrics.IncrementCounter("rpc_error", 1)
```

## 常见问题

### 1. 连接问题
- 连接超时
- 连接断开
- 握手失败

### 2. 性能问题
- 请求堆积
- 响应延迟
- 内存溢出

### 3. 订阅问题
- 消息丢失
- 订阅超时
- 重复消息

## 最佳实践

1. 服务配置
   - 合理设置限流
   - 配置超时时间
   - 启用TLS加密

2. 错误处理
   - 优雅降级
   - 重试机制
   - 错误恢复

3. 监控告警
   - 设置关键指标
   - 配置告警阈值
   - 及时响应

## 注意事项

1. 安全建议
   - 定期更新证书
   - 检查访问控制
   - 监控异常访问

2. 性能建议
   - 控制并发数
   - 优化处理逻辑
   - 合理使用缓存

3. 运维建议
   - 定期检查日志
   - 监控系统资源
   - 及时更新配置

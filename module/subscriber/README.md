# Subscriber 模块

订阅模块是长安链的事件通知系统，负责管理区块链事件的发布和订阅，支持区块、交易和合约事件的实时通知。

## 目录结构

```
subscriber/
├── feed.go           # 订阅源实现
├── subscriber.go     # 订阅者管理
└── model/           # 数据模型定义
```

## 架构设计

```
                   +------------------+
                   |    Subscriber   |
                   +--------+--------+
                            |
              +------------+-----------+
              |                       |
      +-------+-------+     +--------+--------+
      |    Feed      |     |    Model       |
      +------+-------+     +----------------+
             |
    +--------+--------+
    |  Subscribers   |
    +---------------+
```

## 核心组件

### 1. Feed
- 事件源管理
- 消息分发
- 订阅管理
- 缓冲控制

### 2. Subscriber
- 订阅者管理
- 事件过滤
- 消息投递
- 状态维护

### 3. Model
- 数据结构定义
- 事件类型
- 消息格式

## 事件类型

### 1. 区块事件
```go
type BlockEvent struct {
    Height    uint64
    BlockHash string
    TxCount   int
    Timestamp int64
}
```

### 2. 交易事件
```go
type TransactionEvent struct {
    TxId      string
    Status    string
    Timestamp int64
    Result    []byte
}
```

### 3. 合约事件
```go
type ContractEvent struct {
    ContractName string
    EventName    string
    Payload     []byte
}
```

## 主要接口

### 1. Feed 接口
```go
type Feed interface {
    // 发布事件
    Publish(event interface{}) error
    
    // 订阅事件
    Subscribe(filter Filter) (Subscription, error)
    
    // 取消订阅
    Unsubscribe(id string) error
}
```

### 2. Subscriber 接口
```go
type Subscriber interface {
    // 处理事件
    OnEvent(event interface{}) error
    
    // 获取订阅ID
    GetID() string
    
    // 获取订阅状态
    GetStatus() Status
}
```

## 配置说明

```yaml
subscriber:
  # 订阅配置
  config:
    buffer_size: 1000
    max_subscribers: 100
    timeout: "30s"
    
  # 事件配置
  event:
    batch_size: 100
    retry_count: 3
    retry_interval: "1s"
    
  # 过滤配置
  filter:
    enabled: true
    rules: []
```

## 使用示例

### 1. 创建订阅
```go
feed := NewFeed(config)
filter := &EventFilter{
    Type: "block",
    Height: ">1000",
}

sub, err := feed.Subscribe(filter)
if err != nil {
    // 处理错误
}
```

### 2. 处理事件
```go
go func() {
    for event := range sub.Events() {
        // 处理事件
        ProcessEvent(event)
    }
}()
```

## 性能优化

### 1. 消息处理
- 批量处理
- 异步投递
- 缓冲管理

### 2. 订阅管理
- 高效过滤
- 智能路由
- 负载均衡

### 3. 内存优化
- 对象池
- 消息压缩
- 垃圾回收

## 监控指标

### 1. 订阅指标
- 活跃订阅数
- 消息积压量
- 处理延迟

### 2. 性能指标
- 消息吞吐量
- 投递成功率
- 响应时间

### 3. 资源指标
- 内存使用
- CPU使用
- 网络流量

## 调试功能

### 1. 日志记录
```go
logger.Debug("event processing",
    "type", event.Type,
    "subscriber", sub.GetID(),
    "latency", latency,
)
```

### 2. 状态监控
```go
stats := feed.GetStats()
health := feed.CheckHealth()
```

## 错误处理

### 1. 订阅错误
- 连接断开
- 超时处理
- 重试机制

### 2. 投递错误
- 消息丢失
- 顺序错乱
- 重复投递

### 3. 资源错误
- 内存溢出
- 队列满载
- 连接耗尽

## 最佳实践

1. 订阅管理
   - 合理过滤
   - 及时清理
   - 监控状态

2. 性能优化
   - 批量处理
   - 异步操作
   - 资源控制

3. 错误处理
   - 优雅降级
   - 自动恢复
   - 日志记录

## 安全考虑

### 1. 访问控制
- 身份验证
- 权限检查
- 连接限制

### 2. 数据安全
- 消息加密
- 完整性校验
- 防重放攻击

### 3. 资源保护
- 限流控制
- 黑名单机制
- 超时管理

## 常见问题

### 1. 性能问题
- 消息堆积
- 处理延迟
- 内存增长

### 2. 可靠性问题
- 消息丢失
- 断开重连
- 顺序保证

### 3. 扩展性问题
- 订阅者过多
- 消息量激增
- 资源不足

## 注意事项

1. 开发建议
   - 异常处理
   - 性能优化
   - 资源管理

2. 运维建议
   - 监控告警
   - 容量规划
   - 日志分析

3. 使用建议
   - 合理过滤
   - 控制并发
   - 及时清理

## 扩展功能

### 1. 消息过滤
- 主题过滤
- 内容过滤
- 自定义过滤

### 2. 消息转换
- 格式转换
- 协议适配
- 数据压缩

### 3. 监控管理
- 性能监控
- 状态查询
- 统计分析

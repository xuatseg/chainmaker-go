# Net 模块

网络模块是长安链的核心组件之一，负责节点间的网络通信、消息传递和连接管理。

## 目录结构

```
net/
├── net_factory.go           # 网络工厂实现
├── net_logger.go           # 网络日志
├── net_msg_creator.go      # 消息创建器
├── net_service.go          # 网络服务实现
├── net_service_factory.go  # 服务工厂
├── net_service_options.go  # 服务配置选项
├── v220_compat.go         # v2.2.0 兼容性实现
├── testdata/              # 测试数据
└── *_test.go              # 测试文件
```

## 架构设计

```
                    +----------------+
                    |  NetFactory   |
                    +-------+-------+
                            |
                 +---------++---------+
                 |                    |
          +------+-------+    +------+-------+
          | NetService   |    | MsgCreator   |
          +------+-------+    +--------------+
                 |
        +--------+--------+
        |               |
   +----+----+    +----+----+
   | Send    |    | Receive |
   +---------+    +---------+
```

## 核心组件

### 1. NetFactory
- 创建网络服务实例
- 管理网络配置
- 提供服务生命周期管理

### 2. NetService
- 处理网络连接
- 管理消息收发
- 维护节点状态
- 提供网络接口

### 3. MsgCreator
- 创建网络消息
- 消息序列化
- 消息验证

## 主要功能

### 1. 消息处理
```go
// 消息接口
type Message interface {
    GetType() MessageType
    GetPayload() []byte
    GetFrom() string
    GetTo() string
}

// 消息处理
func (s *NetService) HandleMessage(msg Message) error
```

### 2. 连接管理
```go
// 连接接口
type Connection interface {
    Send(msg Message) error
    Close() error
    GetPeerInfo() *PeerInfo
}

// 连接管理
func (s *NetService) Connect(address string) error
```

### 3. 节点发现
```go
// 节点发现接口
type Discovery interface {
    Start() error
    Stop() error
    AddPeer(peer *PeerInfo)
    RemovePeer(id string)
}
```

## 配置说明

### 1. 网络配置
```yaml
net:
  # 监听地址
  listen_addr: ":8080"
  
  # 连接配置
  connection:
    max_connections: 100
    timeout: 5s
    keep_alive: true
    
  # 消息配置
  message:
    max_size: 4MB
    timeout: 3s
    
  # 发现配置
  discovery:
    mode: "static"
    bootstrap_nodes: []
```

## 使用示例

### 1. 创建网络服务
```go
factory := net.NewNetFactory(config)
service := factory.CreateNetService()
service.Start()
```

### 2. 发送消息
```go
msg := msgCreator.CreateMessage(payload)
service.SendMessage(msg)
```

### 3. 注册消息处理器
```go
service.RegisterHandler(msgType, func(msg Message) error {
    // 处理消息
    return nil
})
```

## 性能优化

### 1. 连接池管理
- 复用连接
- 连接限流
- 超时控制

### 2. 消息处理优化
- 消息批处理
- 异步处理
- 优先级队列

### 3. 网络优化
- 压缩传输
- 多路复用
- 流量控制

## 监控指标

### 1. 连接指标
- 活跃连接数
- 连接建立率
- 连接错误率

### 2. 消息指标
- 消息吞吐量
- 消息延迟
- 消息丢失率

### 3. 网络指标
- 带宽使用率
- 网络延迟
- 错误率

## 安全特性

### 1. 传输安全
- TLS 加密
- 证书验证
- 消息签名

### 2. 访问控制
- 节点认证
- 权限验证
- 黑名单机制

### 3. 攻击防护
- DoS 防护
- 消息过滤
- 流量控制

## 调试指南

### 1. 日志配置
```go
logger.SetLevel(log.DebugLevel)
logger.SetOutput(os.Stdout)
```

### 2. 监控检查
```go
metrics := service.GetMetrics()
status := service.GetStatus()
```

### 3. 连接诊断
```go
peers := service.GetPeers()
for _, peer := range peers {
    info := peer.GetInfo()
    // 分析连接状态
}
```

## 常见问题

### 1. 连接问题
- 检查网络配置
- 验证证书有效性
- 检查防火墙设置

### 2. 性能问题
- 优化连接池
- 调整消息大小
- 控制并发连接

### 3. 消息问题
- 检查消息格式
- 验证消息签名
- 分析处理延迟

## 最佳实践

1. 网络配置
   - 合理设置连接数
   - 配置适当超时
   - 启用 Keep-Alive

2. 消息处理
   - 异步处理大消息
   - 实现消息重试
   - 添加消息确认

3. 监控告警
   - 监控关键指标
   - 设置告警阈值
   - 定期检查日志

## 版本兼容

### v2.2.0 兼容性
- 支持旧版本消息
- 兼容性处理
- 平滑升级

## 注意事项

1. 资源管理
   - 及时关闭连接
   - 控制内存使用
   - 清理过期消息

2. 错误处理
   - 优雅降级
   - 重试机制
   - 日志记录

3. 安全建议
   - 定期更新证书
   - 检查访问控制
   - 监控异常行为

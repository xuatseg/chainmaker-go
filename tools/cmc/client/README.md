# 客户端核心实现

本模块提供长安链的客户端核心功能，包括链配置管理、合约调用和交易处理等。

## 功能架构

```
Client Core
├── Chain Config
│   ├── Member Management
│   ├── Consensus Config
│   └── Gas Config
├── Contract
│   ├── System Contracts
│   └── User Contracts
└── Transaction
    ├── Creation
    ├── Signing
    └── Submission
```

## 核心功能

### 1. 链配置管理
- 成员管理
- 共识参数配置
- 权限管理
- Gas费用设置

### 2. 合约操作
- 系统合约调用
- 用户合约部署
- 合约升级
- 合约查询

### 3. 交易处理
- 交易构造
- 多签支持
- 交易提交
- 状态查询

## 使用示例

### 1. 链配置更新
```go
client := NewChainClient(config)
err := client.UpdateChainConfig(
    WithBlockInterval(2000),
    WithGasLimit(5000000),
)
```

### 2. 合约调用
```go
resp, err := client.InvokeContract(
    "my_contract",
    "transfer",
    WithArgs("to", "bob", "amount", "100"),
    WithTimeout(30*time.Second),
)
```

## 开发接口

### 1. 链配置接口
```go
type ChainConfig interface {
    UpdateBlockInterval(int) error
    AddConsensusNode(Node) error
    UpdateGasConfig(GasConfig) error
}
```

### 2. 合约接口
```go
type Contract interface {
    Deploy(ContractSpec) (string, error)
    Invoke(name, method string, opts ...CallOption) (*Response, error)
    Upgrade(name string, spec ContractSpec) error
}
```

## 注意事项

1. 并发安全
   - 客户端非并发安全
   - 建议每个goroutine创建独立client

2. 错误处理
   - 检查所有error返回值
   - 使用IsRetryableError判断可重试错误

3. 资源管理
   - 及时关闭客户端
   - 控制连接池大小
   - 合理设置超时

## 扩展开发

1. 添加新功能
   - 实现新的链配置选项
   - 支持新的合约特性

2. 性能优化
   - 连接复用
   - 批量请求
   - 缓存机制

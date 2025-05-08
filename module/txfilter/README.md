# TxFilter 模块

交易过滤模块是长安链的安全防护组件，负责验证交易的合法性、防止重复交易和恶意攻击，确保只有合规的交易能够进入交易池。

## 目录结构

```
txfilter/
├── birdnest/                # 鸟巢算法实现
├── filtercommon/           # 通用过滤组件
├── filterdefault/          # 默认过滤器实现
├── map/                    # 映射过滤器
├── shardingbirdsnest/      # 分片鸟巢过滤器
├── tx_filter_factory.go    # 过滤器工厂
└── tx_filter_factory_test.go # 工厂测试
```

## 架构设计

```
                    +------------------+
                    |  FilterFactory   |
                    +--------+---------+
                             |
              +--------------+---------------+
              |                              |
      +-------+--------+           +--------+--------+
      |  BirdNest     |           |  DefaultFilter  |
      |   Filter      |           |                 |
      +-------+--------+           +----------------+
              |
     +--------+---------+
     |                  |
+----+----+      +------+-----+
| Sharding |      | MapFilter |
| BirdNest |      |           |
+---------+      +-----------+
```

## 核心组件

### 1. FilterFactory
- 创建过滤器实例
- 管理过滤器配置
- 提供工厂方法

### 2. BirdNestFilter
- 基于鸟巢算法的过滤器
- 高效重复交易检测
- 概率型数据结构

### 3. DefaultFilter
- 基础交易验证
- 签名检查
- 格式验证

## 过滤策略

### 1. 基础验证
- 交易格式检查
- 签名验证
- 时间戳验证

### 2. 重复交易检测
- 交易哈希检查
- 鸟巢算法检测
- 分片检测

### 3. 高级过滤
- 恶意交易识别
- 频率限制
- 模式匹配

## 主要接口

### 1. 过滤器接口
```go
type TxFilter interface {
    // 验证交易
    Validate(tx *Transaction) error
    
    // 添加交易
    Add(tx *Transaction) error
    
    // 检查是否存在
    Contains(txId string) (bool, error)
    
    // 获取过滤器信息
    GetInfo() *FilterInfo
}
```

### 2. 工厂接口
```go
type FilterFactory interface {
    // 创建过滤器
    Create(config *FilterConfig) (TxFilter, error)
    
    // 获取过滤器类型
    GetFilterType() string
    
    // 更新配置
    UpdateConfig(config *FilterConfig) error
}
```

## 配置说明

```yaml
txfilter:
  # 过滤器配置
  filter:
    type: "birdnest"  # birdnest/default/map/sharding
    capacity: 1000000
    false_positive_rate: 0.001
    
  # 鸟巢配置
  birdnest:
    shard_count: 4
    bits_per_item: 10
    max_num_items: 100000
    
  # 验证配置
  validation:
    check_signature: true
    check_timestamp: true
    check_duplicate: true
```

## 使用示例

### 1. 创建过滤器
```go
factory := NewFilterFactory()
filter, err := factory.Create(config)
if err != nil {
    // 处理错误
}
```

### 2. 验证交易
```go
// 验证交易
if err := filter.Validate(tx); err != nil {
    // 处理无效交易
}

// 添加交易到过滤器
if err := filter.Add(tx); err != nil {
    // 处理错误
}
```

## 性能优化

### 1. 内存优化
- 概率型数据结构
- 分片存储
- 压缩存储

### 2. 查询优化
- 并行查询
- 批量操作
- 缓存优化

### 3. 算法优化
- 高效哈希
- 位图压缩
- 智能分片

## 监控指标

### 1. 过滤指标
- 交易过滤数
- 误判率
- 过滤延迟

### 2. 性能指标
- 查询吞吐量
- 添加速度
- 内存使用

### 3. 资源指标
- 内存占用
- CPU使用
- 存储使用

## 调试功能

### 1. 日志记录
```go
logger.Debug("tx filter operation",
    "tx_id", tx.ID,
    "operation", "validate",
    "result", result,
)
```

### 2. 状态检查
```go
info := filter.GetInfo()
stats := filter.GetStats()
```

## 常见问题

### 1. 误判问题
- 调整误判率
- 增加容量
- 优化哈希

### 2. 性能问题
- 增加分片
- 优化配置
- 升级硬件

### 3. 一致性问题
- 定期同步
- 数据校验
- 恢复机制

## 最佳实践

1. 配置优化
   - 合理设置容量
   - 平衡误判率
   - 适当分片

2. 性能优化
   - 批量操作
   - 并行处理
   - 资源控制

3. 安全建议
   - 定期更新
   - 监控异常
   - 防御增强

## 安全特性

### 1. 防重复交易
- 高效检测
- 低误判率
- 快速验证

### 2. 防恶意攻击
- 频率限制
- 模式识别
- 动态过滤

### 3. 数据保护
- 加密存储
- 访问控制
- 安全清理

## 注意事项

1. 开发建议
   - 完整测试
   - 性能评估
   - 边界处理

2. 运维建议
   - 监控告警
   - 容量规划
   - 定期维护

3. 使用建议
   - 合理配置
   - 及时更新
   - 异常处理

# 链数据查询工具

本模块提供长安链数据的查询功能，支持区块、交易和合约状态的查询。

## 功能特性

### 1. 区块查询
- 按高度查询区块
- 按哈希查询区块
- 区块头信息
- 区块体详情

### 2. 交易查询
- 交易基本信息
- 交易收据
- 交易输入输出
- 交易事件日志

### 3. 合约查询
- 合约代码查询
- 合约状态查询
- 合约ABI查询
- 合约事件查询

## 使用指南

### 1. 查询区块
```bash
cmc query block -height 1000
```

### 2. 查询交易
```bash
cmc query tx -hash 0x123...
```

### 3. 查询合约
```bash
cmc query contract -name my_token -method balanceOf -params '{"owner":"0x123"}'
```

## 查询选项

### 1. 输出格式
- JSON格式
- 表格格式
- 原始二进制
- 自定义模板

### 2. 过滤条件
- 时间范围
- 区块高度范围
- 交易类型
- 合约地址

## 实现原理

### 1. 查询流程
```
1. 解析查询请求
2. 访问链数据
3. 应用过滤条件
4. 格式化输出
```

### 2. 缓存优化
- LRU缓存区块头
- 批量查询优化
- 并行查询

## 注意事项

1. 性能考虑
   - 避免全表扫描
   - 使用索引字段
   - 限制返回结果

2. 数据一致性
   - 检查确认数
   - 验证默克尔证明
   - 注意分叉情况

3. 隐私保护
   - 敏感数据过滤
   - 访问权限控制
   - 查询日志审计

## 开发接口

### 1. Go API
```go
// 创建查询客户端
client := query.NewClient(config)

// 查询区块
block, err := client.GetBlock(height)

// 查询交易
tx, err := client.GetTransaction(hash)
```

### 2. 扩展开发
1. 添加新查询类型
2. 支持更多存储后端
3. 优化查询性能

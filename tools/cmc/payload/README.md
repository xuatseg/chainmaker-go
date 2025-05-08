# 交易负载处理

本模块负责长安链交易负载的构造、签名和序列化操作。

## 核心功能

### 1. 负载构造
- 交易创建
- 合约调用负载
- 系统操作负载
- 多签请求构造

### 2. 签名处理
- 单签生成
- 多签聚合
- 签名验证
- 签名解析

### 3. 序列化
- JSON 编码/解码
- Protobuf 序列化
- 二进制格式处理
- 十六进制转换

## 使用指南

### 1. 创建交易
```bash
cmc payload create \
  -method transfer \
  -params '{"to":"0x123","amount":100}' \
  -out tx.json
```

### 2. 签名交易
```bash
cmc payload sign \
  -in tx.json \
  -key user.key \
  -out tx_signed.json
```

### 3. 序列化交易
```bash
cmc payload serialize \
  -in tx_signed.json \
  -out tx_raw.bin
```

## 数据结构

### 1. 交易负载格式
```json
{
  "version": "1.0",
  "contract": "my_contract",
  "method": "transfer",
  "params": {
    "to": "0x123",
    "amount": 100
  },
  "signatures": []
}
```

### 2. 签名格式
```
{
  "signer": "user1",
  "algorithm": "SM2",
  "signature": "304502...",
  "timestamp": 1234567890
}
```

## 开发接口

### 1. Go API
```go
// 创建负载
payload := NewPayload("my_contract", "transfer", params)

// 添加签名
err := payload.Sign(userKey)

// 序列化
data, err := payload.Serialize()
```

### 2. 扩展开发
1. 支持新的负载类型
2. 添加自定义签名方案
3. 优化序列化性能

## 注意事项

1. 安全考虑
   - 验证输入参数
   - 保护私钥安全
   - 检查签名有效性

2. 性能优化
   - 批量化处理
   - 内存复用
   - 并行签名

3. 兼容性
   - 版本控制
   - 向后兼容
   - 格式验证

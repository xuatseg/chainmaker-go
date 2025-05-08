# 地址计算测试工具

本工具用于验证长安链地址计算算法的正确性，支持多种地址格式的计算和验证。

## 文件结构

```
calculate_addr/
└── calculate_addr.go    # 地址计算实现
```

## 功能说明

### 1. 支持的地址类型
- 普通账户地址
- 合约地址
- 多签地址
- 特殊地址格式

### 2. 计算功能
- 从公钥计算地址
- 从私钥推导地址
- 地址格式转换
- 地址校验和验证

## 使用方式

### 1. 命令行使用

```bash
# 计算地址
go run calculate_addr.go -pubkey [公钥]

# 示例
go run calculate_addr.go -pubkey 04123456789abcdef...
```

### 2. 参数说明

| 参数 | 说明 |
|------|------|
| -pubkey | 输入的公钥(16进制格式) |
| -type   | 地址类型(account/contract) |
| -format | 输出格式(base58/bech32/hex) |
| -check  | 验证地址有效性 |

## 实现原理

### 1. 地址生成流程

```go
1. 输入公钥或私钥
2. 进行哈希计算(SHA3-256)
3. 添加版本前缀
4. 计算校验和
5. 编码为指定格式
```

### 2. 关键算法

```go
// 核心计算函数
func CalculateAddress(pubKey []byte, addrType AddrType) (string, error) {
    hash := sha3.Sum256(pubKey)
    payload := append([]byte{byte(addrType)}, hash[:]...)
    checksum := checksum(payload)
    return encode(append(payload, checksum...))
}
```

## 测试用例

### 1. 标准测试向量

```go
// 测试标准公钥的地址计算
func TestStandardAddress(t *testing.T) {
    pubKey := hexDecode("041234...")
    expected := "chain1qxyz..."
    actual, _ := CalculateAddress(pubKey, Account)
    assert.Equal(t, expected, actual)
}
```

### 2. 边界测试

```go
// 测试空公钥
func TestEmptyKey(t *testing.T) {
    _, err := CalculateAddress([]byte{}, Account)
    assert.Error(t, err)
}
```

## 注意事项

1. 输入要求
   - 公钥必须是有效的椭圆曲线点
   - 私钥需要保密处理

2. 输出验证
   - 建议使用官方工具交叉验证
   - 检查校验和

3. 安全提示
   - 不要在测试代码中硬编码私钥
   - 使用测试专用账户

## 扩展功能

1. 批量测试
   - 从文件读取测试向量
   - 自动化验证

2. 性能测试
   - 地址计算速度
   - 并发性能

3. 格式转换
   - 不同编码格式互转
   - 兼容性测试

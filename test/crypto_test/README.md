# 密码学测试

本目录包含长安链密码学功能的测试代码，验证核心加密算法的正确性和性能。

## 目录结构

```
crypto_test/
├── asym.go     # 非对称加密测试
├── crypto.go   # 密码学通用测试
├── hash.go     # 哈希算法测试
└── sym.go      # 对称加密测试
```

## 测试范围

### 1. 非对称加密 (asym.go)
- 国密SM2算法
- RSA算法
- 密钥生成
- 签名验证
- 加密解密

### 2. 对称加密 (sym.go)
- 国密SM4算法
- AES算法
- 加密模式测试
- 性能基准

### 3. 哈希算法 (hash.go)
- 国密SM3算法
- SHA系列算法
- 哈希一致性
- 碰撞测试

## 测试准备

### 1. 环境要求
- Go 1.16+
- 国密算法支持
- 测试密钥对

### 2. 测试数据
- 预生成的测试密钥
- 标准测试向量
- 随机测试数据

## 运行测试

### 1. 运行全部测试
```bash
go test -v ./...
```

### 2. 运行特定测试
```bash
# 只运行SM2测试
go test -v -run TestSM2

# 带性能分析
go test -bench . -benchmem
```

## 测试用例示例

### 1. SM2签名验证
```go
func TestSM2SignVerify(t *testing.T) {
    priv, pub := generateSM2KeyPair()
    msg := []byte("test message")
    
    // 签名
    sig, err := SignSM2(priv, msg)
    assert.NoError(t, err)
    
    // 验证
    valid := VerifySM2(pub, msg, sig)
    assert.True(t, valid)
}
```

### 2. SM4加密解密
```go
func TestSM4EncryptDecrypt(t *testing.T) {
    key := randomBytes(16)
    plaintext := []byte("test data")
    
    // 加密
    ciphertext, err := EncryptSM4(key, plaintext)
    assert.NoError(t, err)
    
    // 解密
    decrypted, err := DecryptSM4(key, ciphertext)
    assert.NoError(t, err)
    assert.Equal(t, plaintext, decrypted)
}
```

## 性能基准

### 1. 哈希算法性能
```go
func BenchmarkSM3(b *testing.B) {
    data := make([]byte, 1024)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        SumSM3(data)
    }
}
```

### 2. 签名算法性能
```go
func BenchmarkSM2Sign(b *testing.B) {
    priv, _ := generateSM2KeyPair()
    msg := []byte("benchmark data")
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        SignSM2(priv, msg)
    }
}
```

## 测试数据管理

### 1. 标准测试向量
- NIST标准测试数据
- 国密标准测试数据
- 边界条件测试数据

### 2. 随机测试
- 随机生成密钥
- 随机生成消息
- 随机数据加密

## 注意事项

1. 密钥安全
   - 测试密钥仅用于开发
   - 不要使用生产密钥

2. 性能测试
   - 在稳定环境下运行
   - 多次运行取平均值

3. 算法合规
   - 验证算法实现标准
   - 检查官方测试向量

## 扩展测试

1. 兼容性测试
   - 不同语言实现互验
   - 版本兼容性

2. 安全性测试
   - 侧信道分析
   - 故障注入

3. 并发测试
   - 多线程加密
   - 并发密钥生成

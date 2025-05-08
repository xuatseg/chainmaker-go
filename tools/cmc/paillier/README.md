# Paillier 加密工具

本模块实现Paillier同态加密算法，支持长安链中的隐私计算场景。

## 功能特性

### 1. 密钥管理
- 密钥对生成
- 密钥安全存储
- 密钥参数导出

### 2. 加密运算
- 数值加密
- 同态加法
- 标量乘法
- 零知识证明

### 3. 高级功能
- 门限加密
- 安全多方计算
- 隐私数据聚合

## 使用指南

### 1. 生成密钥
```bash
cmc paillier genkey -bits 2048 -out keypair.pem
```

### 2. 数据加密
```bash
cmc paillier encrypt \
  -pubkey pub.pem \
  -value 100 \
  -out encrypted.dat
```

### 3. 同态运算
```bash
cmc paillier compute \
  -enc1 a.enc \
  -enc2 b.enc \
  -op add \
  -out sum.enc
```

## 技术实现

### 1. 算法参数
- 密钥长度：2048位
- 素数生成：Miller-Rabin测试
- 加密随机数：安全随机源

### 2. 性能优化
- 预计算加速
- 批处理操作
- 并行计算

## 安全注意事项

1. 密钥安全
   - 私钥加密存储
   - 最小权限访问
   - 定期密钥轮换

2. 参数验证
   - 验证公钥参数
   - 检查随机数质量
   - 防范参数攻击

3. 性能权衡
   - 密钥长度选择
   - 预计算空间
   - 批处理大小

## 开发接口

### 1. Go API
```go
// 生成密钥对
priv, pub := paillier.GenerateKey(rand.Reader, 2048)

// 加密数据
ciphertext, err := paillier.Encrypt(pub, big.NewInt(100))

// 同态加法
sum := paillier.Add(pub, c1, c2)
```

### 2. 扩展开发
1. 支持门限加密
2. 添加零知识证明
3. 优化大数运算

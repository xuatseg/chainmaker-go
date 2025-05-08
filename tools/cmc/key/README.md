# 密钥管理工具

本模块提供长安链中的密钥生成、导出和管理功能，支持多种加密算法。

## 功能特性

### 1. 密钥生成
- 生成SM2密钥对
- 生成RSA密钥对
- 生成ED25519密钥
- 生成国密算法密钥

### 2. 密钥操作
- 公钥导出
- 私钥加密存储
- 密钥格式转换
- 密钥信息查看

### 3. 密钥派生
- 分层确定性密钥派生
- 基于口令的密钥派生
- 多签密钥共享

## 使用指南

### 1. 生成密钥对
```bash
# 生成SM2密钥对
cmc key generate -algo sm2 -out sm2.key

# 生成RSA密钥
cmc key generate -algo rsa -bits 2048 -out rsa.key
```

### 2. 导出公钥
```bash
cmc key export -in sm2.key -type pub -out sm2.pub
```

### 3. 密钥信息
```bash
cmc key info -in sm2.key
```

## 密钥格式

### 1. 私钥格式
```
-----BEGIN PRIVATE KEY-----
算法标识: SM2
创建时间: 2023-01-01
加密方式: AES-256-CBC
-----END PRIVATE KEY-----
```

### 2. 公钥格式
```
-----BEGIN PUBLIC KEY-----
算法: SM2
曲线: sm2p256v1
指纹: SHA256:xxxx
-----END PUBLIC KEY-----
```

## 安全注意事项

1. 私钥存储
   - 使用强密码加密
   - 硬件安全模块保护
   - 访问权限控制

2. 密钥传输
   - 安全通道传输
   - 临时密钥对
   - 密钥哈希验证

3. 密钥生命周期
   - 定期轮换
   - 安全销毁
   - 使用记录审计

## 开发接口

### 1. Go API
```go
// 生成密钥对
key, err := key.Generate(algorithm.SM2, nil)

// 导出公钥
pubKey := key.Public()

// 加载密钥
key, err := key.Load("sm2.key", "password")
```

### 2. 扩展开发
1. 添加新算法支持
2. 实现硬件安全模块集成
3. 增强密钥派生功能

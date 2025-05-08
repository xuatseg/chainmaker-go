# 公钥管理工具

本模块提供长安链公钥的格式转换、验证和提取功能。

## 功能特性

### 1. 公钥操作
- 从私钥提取公钥
- 公钥格式转换(PEM/DER/HEX)
- 公钥指纹生成
- 公钥信息查看

### 2. 公钥验证
- 验证公钥有效性
- 检查公钥算法
- 验证公钥参数

### 3. 地址派生
- 从公钥生成地址
- 地址格式转换
- 地址校验和验证

## 使用指南

### 1. 提取公钥
```bash
cmc pubkey extract -key priv.key -out pub.key
```

### 2. 公钥转换
```bash
cmc pubkey convert -in pub.pem -fmt der -out pub.der
```

### 3. 地址生成
```bash
cmc pubkey address -pub pub.key -algo sm2 -out addr.txt
```

## 实现原理

### 1. 公钥格式
```
-----BEGIN PUBLIC KEY-----
算法: SM2
曲线参数: sm2p256v1
编码格式: ASN.1 DER
-----END PUBLIC KEY-----
```

### 2. 地址生成流程
```
1. 公钥序列化
2. Keccak256哈希
3. 取最后20字节
4. 添加校验和
5. Base58编码
```

## 安全注意事项

1. 公钥验证
   - 检查曲线参数
   - 验证生成点
   - 防范无效曲线攻击

2. 地址安全
   - 校验和验证
   - 格式标准化
   - 防混淆处理

3. 密钥保护
   - 保护私钥文件
   - 安全删除临时文件
   - 访问权限控制

## 开发接口

### 1. Go API
```go
// 从私钥提取公钥
pub, err := pubkey.FromPrivate(privateKey)

// 生成地址
addr := pubkey.ToAddress(pub, ChainTypeMainnet)

// 验证公钥
valid := pubkey.Validate(pub)
```

### 2. 扩展开发
1. 支持新算法
2. 添加更多格式
3. 增强验证功能

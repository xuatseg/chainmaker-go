# 分层身份加密 (HIBE) 工具

本模块实现基于身份的层级加密方案，支持长安链中的细粒度数据权限控制。

## 核心功能

### 1. 密钥生成
- 主密钥生成
- 用户私钥派生
- 密钥托管服务

### 2. 加密解密
- 基于身份的加密
- 层级数据解密
- 批量加密操作

### 3. 权限管理
- 访问策略定义
- 属性基加密
- 细粒度访问控制

## 使用指南

### 1. 初始化HIBE系统
```bash
cmc hibe setup \
  -params hibe_params.json \
  -master_key master.key
```

### 2. 派生用户密钥
```bash
cmc hibe keygen \
  -id alice@org1 \
  -master master.key \
  -out alice.key
```

### 3. 数据加密
```bash
cmc hibe encrypt \
  -msg sensitive_data.txt \
  -policy "org1.department1" \
  -out encrypted.hibe
```

## 技术实现

### 1. 算法基础
- 基于双线性配对
- 使用BN曲线
- 安全参数：256位

### 2. 密钥层级
```
根主密钥
  │
  ├─ 组织层级 (org1)
  │    │
  │    ├─ 部门层级 (dept1)
  │    │
  │    └─ 角色层级 (role1)
  │
  └─ 组织层级 (org2)
```

## 安全注意事项

1. 主密钥保护
   - 硬件安全模块存储
   - 多因素保护
   - 定期轮换

2. 密钥分发
   - 安全传输通道
   - 最小权限分配
   - 密钥生命周期管理

3. 性能考虑
   - 深层级加密开销
   - 批量操作优化
   - 缓存常用密钥

## 开发接口

### 1. Go API
```go
// 初始化系统
params, masterKey := hibe.Setup(256)

// 派生密钥
userKey := hibe.KeyDerive(masterKey, "org1.dept1.user1")

// 加密数据
ciphertext := hibe.Encrypt(params, "org1.dept1", message)
```

### 2. 扩展开发
1. 支持新的加密策略
2. 添加密钥撤销功能
3. 优化层级查询效率

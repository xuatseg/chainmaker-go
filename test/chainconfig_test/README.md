# 链配置测试

本目录包含长安链配置管理的单元测试，验证链配置的各个组件功能是否符合预期。

## 目录结构

```
chainconfig_test/
├── block_contract_test.go    # 区块合约测试
├── cert_manage_test.go      # 证书管理测试
├── chain_config_test.go     # 链配置测试
├── contract_manage_test.go  # 合约管理测试
├── multisign_test.go        # 多签配置测试
├── bytes_code.go           # 辅助函数
└── common.go               # 公共测试代码
```

## 测试范围

### 1. 链配置测试
- 配置加载验证
- 配置更新测试
- 配置持久化测试
- 版本兼容性测试

### 2. 证书管理测试
- CA证书验证
- 节点证书管理
- 证书撤销列表
- 证书过期处理

### 3. 合约管理测试
- 系统合约配置
- 合约权限管理
- 合约升级测试
- 合约冻结解冻

## 测试准备

### 1. 环境要求
- Go 1.16+
- 测试证书文件
- 配置文件模板

### 2. 配置文件
测试使用的配置文件位于 `testdata/chainconfig` 目录，包括：
- `chainconfig.yml` - 链基础配置
- `certs/` - 测试证书
- `contracts/` - 测试合约

## 运行测试

### 1. 运行全部测试
```bash
go test -v ./...
```

### 2. 运行特定测试
```bash
# 只运行证书管理测试
go test -v -run TestCertManage

# 带覆盖率统计
go test -cover -v ./...
```

## 测试用例示例

### 1. 链配置更新测试
```go
func TestChainConfigUpdate(t *testing.T) {
    // 初始化配置
    config := loadTestConfig()
    
    // 修改配置
    newConfig := config.Copy()
    newConfig.BlockInterval = 2000
    
    // 验证更新
    err := UpdateChainConfig(config, newConfig)
    assert.NoError(t, err)
    assert.Equal(t, 2000, config.BlockInterval)
}
```

### 2. 证书撤销测试
```go
func TestCertRevocation(t *testing.T) {
    // 初始化证书管理器
    mgr := NewCertManager(testCerts)
    
    // 撤销证书
    cert := testCerts[0]
    err := mgr.Revoke(cert.SerialNumber)
    assert.NoError(t, err)
    
    // 验证撤销状态
    revoked, err := mgr.IsRevoked(cert.SerialNumber)
    assert.NoError(t, err)
    assert.True(t, revoked)
}
```

## 测试数据管理

### 1. 测试证书
- `ca.crt` - 根CA证书
- `node1.crt` - 节点1证书
- `user1.crt` - 用户1证书

### 2. 测试合约
- `system_contract` - 系统合约
- `test_contract` - 测试合约

## 注意事项

1. 证书安全
   - 测试证书仅用于开发环境
   - 不要使用生产证书

2. 测试隔离
   - 每个测试用例应独立
   - 清理测试数据

3. 测试覆盖
   - 验证正常流程
   - 测试异常情况
   - 边界条件测试

## 扩展测试

1. 性能测试
   - 配置加载速度
   - 证书验证性能

2. 兼容性测试
   - 旧版本配置兼容
   - 证书格式兼容

3. 并发测试
   - 并发配置更新
   - 并发证书验证

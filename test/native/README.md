# 原生合约测试

本目录包含长安链原生合约的测试代码，主要验证多签合约等系统合约的功能和性能。

## 目录结构

```
native/
├── common.go          # 测试公共函数
└── multisign_test.go   # 多签合约测试
```

## 测试范围

### 1. 多签合约测试
- 多签账户创建
- 多签交易提交
- 签名收集验证
- 合约执行控制
- 权限管理测试

## 测试准备

### 1. 环境要求
- 已部署原生合约的链节点
- 测试用账户和密钥
- 测试数据准备

### 2. 测试数据
- `testdata/multisign` - 多签测试数据
- `testdata/accounts` - 测试账户

## 运行测试

### 1. 运行全部测试
```bash
go test -v ./...
```

### 2. 运行特定测试
```bash
# 只运行多签测试
go test -v -run TestMultiSign

# 带覆盖率统计
go test -cover -v ./...
```

## 测试用例示例

### 1. 多签账户创建
```go
func TestCreateMultiSignAccount(t *testing.T) {
    // 准备测试账户
    accounts := prepareTestAccounts(3)
    
    // 创建多签账户
    contract := GetMultiSignContract()
    result, err := contract.CreateAccount(accounts, 2)
    
    assert.NoError(t, err)
    assert.NotEmpty(t, result.AccountAddress)
}
```

### 2. 多签交易验证
```go
func TestMultiSignTransaction(t *testing.T) {
    // 准备多签账户
    account := createTestMultiSignAccount()
    
    // 发起交易
    tx := prepareTestTransaction()
    txId := submitTransaction(tx)
    
    // 收集签名
    sign1 := signTransaction(account.Members[0], txId)
    sign2 := signTransaction(account.Members[1], txId)
    
    // 验证执行
    result := executeMultiSignTx(account, txId, []string{sign1, sign2})
    assert.True(t, result.Success)
}
```

## 测试数据管理

### 1. 账户配置
- 单签账户
- 多签账户(2/3)
- 多签账户(3/5)

### 2. 交易类型
- 资产转移
- 合约调用
- 账户管理

## 注意事项

1. 测试隔离
   - 每个测试用例使用独立账户
   - 清理测试产生的数据

2. 异常测试
   - 签名不足情况
   - 无效签名
   - 重复签名

3. 性能考虑
   - 多签收集超时
   - 大量签名验证

## 扩展测试

1. 性能测试
   - 多签账户创建速度
   - 签名收集效率

2. 压力测试
   - 并发多签请求
   - 大量多签账户管理

3. 安全测试
   - 签名伪造测试
   - 权限绕过测试

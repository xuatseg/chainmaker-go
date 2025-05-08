# EVM 合约场景测试

本目录包含长安链 EVM 兼容合约的完整测试场景，验证以太坊虚拟机合约的功能和性能。

## 目录结构

```
scenario1_evm/
├── README.md         # 测试说明文档
├── chain1.py         # 链1 EVM测试脚本
├── chain2.py         # 链2 EVM测试脚本
├── chain3.py         # 链3 EVM测试脚本
└── testcase.py       # EVM测试用例实现
```

## 测试场景

### 1. 测试目标
- 验证 EVM 兼容性
- 测试 Solidity 合约功能
- 评估 EVM 执行性能
- 检查 gas 计费机制

### 2. 测试范围
- ERC20 代币合约
- ERC721 NFT 合约
- 复杂业务合约
- 跨合约调用

## 环境准备

### 1. 测试网络配置
```yaml
evm_chains:
  - name: chain1
    evm_version: "0.8.0"
    gas_limit: 8000000
  - name: chain2  
    evm_version: "0.7.0"
    gas_limit: 5000000
  - name: chain3
    evm_version: "0.6.0"
    gas_limit: 3000000
```

### 2. 合约准备
- 编译 Solidity 合约
- 准备测试账户
- 预分配测试代币

## 运行测试

### 1. 部署合约
```bash
python3 chain1.py deploy
python3 chain2.py deploy  
python3 chain3.py deploy
```

### 2. 执行测试
```bash
python3 testcase.py --evm --all
```

### 3. 可选参数
| 参数 | 说明 |
|------|------|
| --chain | 指定测试链 |
| --contract | 测试特定合约 |
| --gas | 显示gas消耗 |
| --debug | 调试模式 |

## 测试用例

### 1. ERC20 测试
```python
def test_erc20_transfer():
    # 初始化
    token = ERC20Contract('MyToken')
    alice = Account('alice')
    bob = Account('bob')
    
    # 转账测试
    tx = token.transfer(alice, bob, 100)
    assert tx.status == 'success'
    assert token.balanceOf(bob) == 100
```

### 2. Gas 消耗测试  
```python
def test_gas_consumption():
    contract = deploy_complex_contract()
    gas_used = []
    
    for i in range(10):
        tx = contract.execute(i)
        gas_used.append(tx.gas_used)
    
    # 验证gas消耗稳定
    assert max(gas_used) - min(gas_used) < 1000
```

## 测试报告

示例输出:
```
=== EVM测试报告 ===
测试合约: 5
测试用例: 42
通过率: 95.2%
平均Gas消耗: 128,456
最大Gas消耗: 523,789
异常交易: 2
```

## 注意事项

1. 合约版本
   - 明确 Solidity 版本
   - 检查编译器兼容性
   - 验证 ABI 编码

2. Gas 配置
   - 合理设置 gas limit
   - 监控 gas 消耗
   - 优化合约代码

3. 测试顺序
   - 先功能测试
   - 再性能测试  
   - 最后边界测试

## 扩展测试

1. 性能测试
   - 合约执行速度
   - 状态读写性能
   - 并发处理能力

2. 安全测试
   - 重入攻击测试
   - 溢出测试
   - 权限控制测试

3. 调试工具
   - 交易回放
   - 状态追踪
   - 事件日志

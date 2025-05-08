# 原生合约场景测试

本目录包含长安链原生合约的完整场景测试，模拟真实业务场景下的合约调用和验证。

## 目录结构

```
scenario0_native/
├── README.md         # 测试说明文档
├── chain1.py         # 链1测试脚本
├── chain2.py         # 链2测试脚本
├── chain3.py         # 链3测试脚本
└── testcase.py       # 测试用例实现
```

## 测试场景

### 1. 测试目标
- 验证原生合约功能完整性
- 测试跨链合约调用
- 评估合约执行性能
- 检查异常处理能力

### 2. 测试范围
- 资产转移场景
- 数据存证场景
- 权限管理场景
- 跨链互操作场景

## 环境准备

### 1. 测试网络配置
```yaml
chains:
  - name: chain1
    type: solo
    nodes: 1
  - name: chain2
    type: raft
    nodes: 3
  - name: chain3
    type: tbft
    nodes: 4
```

### 2. 账户准备
- 初始化测试账户
- 分配初始资金
- 配置访问权限

## 运行测试

### 1. 启动测试网络
```bash
python3 chain1.py start
python3 chain2.py start
python3 chain3.py start
```

### 2. 执行测试用例
```bash
python3 testcase.py --all
```

### 3. 可选参数
| 参数 | 说明 |
|------|------|
| --chain | 指定测试链 |
| --case | 运行特定用例 |
| --debug | 调试模式 |
| --report | 生成报告 |

## 测试用例

### 1. 资产转移测试
```python
def test_asset_transfer():
    # 初始化账户
    alice = Account('alice')
    bob = Account('bob')
    
    # 执行转账
    tx = alice.transfer(bob, 100)
    
    # 验证结果
    assert tx.status == 'confirmed'
    assert bob.balance == 100
```

### 2. 跨链调用测试
```python
def test_cross_chain():
    # 链1存证
    chain1.store('key1', 'value1')
    
    # 链2验证
    value = chain2.verify('key1')
    assert value == 'value1'
```

## 测试报告

示例输出:
```
=== 测试结果 ===
测试场景: 原生合约基础测试
测试用例数: 23
通过用例: 22
失败用例: 1
覆盖率: 95.6%
平均执行时间: 1.2s
```

## 注意事项

1. 测试顺序
   - 先单链测试
   - 再跨链测试
   - 最后压力测试

2. 环境隔离
   - 使用独立测试网络
   - 清理测试数据
   - 避免环境污染

3. 结果验证
   - 检查链上数据
   - 验证状态一致性
   - 核对交易日志

## 扩展测试

1. 性能测试
   - 增加并发用户
   - 延长测试时间
   - 监控资源使用

2. 安全测试
   - 异常输入测试
   - 边界条件测试
   - 权限验证测试

3. 自动化
   - 定时执行
   - 自动报告
   - 异常告警

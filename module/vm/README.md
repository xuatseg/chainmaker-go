# VM (Virtual Machine) 模块

虚拟机模块是长安链的核心组件之一，负责智能合约的执行环境管理和合约生命周期管理。

## 目录结构

```
vm/
└── vm_provider.go    # 虚拟机服务提供者接口定义
```

## 架构设计

```
                    +----------------+
                    |  VMProvider   |
                    +-------+-------+
                            |
                 +---------++---------+
                 |                    |
          +------+-------+    +------+-------+
          | ContractMgr  |    | RuntimeMgr   |
          +------+-------+    +------+-------+
                 |                    |
        +--------+--------+   +------+-------+
        | Contract Life   |   | Execution    |
        | Cycle           |   | Environment  |
        +-----------------+   +--------------+
```

## 核心接口

### VMProvider 接口
```go
type VMProvider interface {
    // 初始化虚拟机
    Init() error
    
    // 启动虚拟机服务
    Start() error
    
    // 停止虚拟机服务
    Stop() error
    
    // 执行合约
    InvokeContract(contract *Contract) (*Result, error)
    
    // 部署合约
    DeployContract(contract *Contract) error
    
    // 升级合约
    UpgradeContract(contract *Contract) error
}
```

## 支持的合约类型

1. 原生合约 (Native)
   - 系统内置合约
   - 高性能执行
   - 直接访问系统资源

2. WASM 合约
   - WebAssembly 合约
   - 跨平台支持
   - 安全沙箱环境

3. EVM 合约
   - 以太坊虚拟机合约
   - 兼容 Solidity
   - 生态系统丰富

4. Docker 合约
   - 容器化执行环境
   - 隔离性好
   - 资源可控

## 合约生命周期

### 1. 部署阶段
```
1. 验证合约代码
2. 编译合约（如需要）
3. 创建运行环境
4. 初始化状态
5. 存储合约信息
```

### 2. 调用阶段
```
1. 加载合约
2. 准备执行环境
3. 参数验证
4. 执行合约代码
5. 更新状态
6. 返回结果
```

### 3. 升级阶段
```
1. 验证升级权限
2. 保存旧版本
3. 更新合约代码
4. 迁移状态（如需要）
5. 更新版本信息
```

## 配置说明

```yaml
vm:
  # 虚拟机类型配置
  providers:
    - type: "native"
      enabled: true
    - type: "wasm"
      enabled: true
      runtime: "wasmer"
    - type: "evm"
      enabled: true
      version: "1.9.24"
    - type: "docker"
      enabled: true
      
  # 资源限制
  limits:
    memory: "1GB"
    cpu: "2"
    execution_timeout: "30s"
    
  # 安全配置
  security:
    enable_sandbox: true
    allow_network: false
    allow_file_system: false
```

## 使用示例

### 1. 部署合约
```go
contract := &Contract{
    Name:    "MyContract",
    Version: "1.0",
    Code:    contractCode,
    Type:    "wasm",
}

err := vmProvider.DeployContract(contract)
```

### 2. 调用合约
```go
result, err := vmProvider.InvokeContract(&Contract{
    Name:    "MyContract",
    Method:  "transfer",
    Params:  []string{"addr1", "addr2", "100"},
})
```

## 安全特性

### 1. 沙箱隔离
- 资源隔离
- 网络隔离
- 文件系统隔离

### 2. 访问控制
- 合约权限管理
- API 访问限制
- 资源使用限制

### 3. 代码安全
- 代码静态分析
- 运行时检查
- 漏洞防护

## 性能优化

### 1. 执行优化
- JIT 编译
- 代码缓存
- 并行执行

### 2. 资源管理
- 内存池
- 协程池
- 缓存管理

## 监控指标

### 1. 执行指标
- 合约执行时间
- 资源使用情况
- 错误率统计

### 2. 性能指标
- TPS (每秒交易数)
- 响应时间
- 资源利用率

## 调试功能

### 1. 日志记录
```go
logger.Debug("contract execution",
    "contract", contract.Name,
    "method", contract.Method,
    "gas_used", gasUsed,
)
```

### 2. 状态检查
```go
status := vmProvider.GetStatus()
metrics := vmProvider.GetMetrics()
```

## 常见问题

### 1. 执行超时
- 检查超时配置
- 优化合约代码
- 增加资源配置

### 2. 内存溢出
- 检查内存限制
- 优化内存使用
- 清理未使用资源

### 3. 合约错误
- 检查合约代码
- 验证参数格式
- 查看错误日志

## 最佳实践

1. 合约开发
   - 遵循最佳实践
   - 优化代码性能
   - 做好错误处理

2. 资源配置
   - 合理设置限制
   - 监控资源使用
   - 定期优化配置

3. 安全防护
   - 启用安全特性
   - 定期安全审计
   - 及时更新修复

## 注意事项

1. 开发建议
   - 测试覆盖完整
   - 考虑兼容性
   - 文档完善

2. 运维建议
   - 监控告警
   - 定期备份
   - 版本管理

3. 安全建议
   - 权限最小化
   - 定期安全检查
   - 漏洞响应机制

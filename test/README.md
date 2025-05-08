# ChainMaker 测试目录

本目录包含了长安链（ChainMaker）的完整测试套件，涵盖了单元测试、集成测试、性能测试和场景测试等多个维度。

## 目录结构

```
test/
├── bench/                  # 基准测试
├── calculate_addr/         # 地址计算测试
├── chain1/                # 链1测试
├── chain2/                # 链2测试
├── chain3/                # 链3测试
├── chainconfig_test/      # 链配置测试
├── common/                # 通用测试组件
├── config/                # 测试配置
├── crypto_test/           # 密码学测试
├── grpc_client/          # gRPC客户端测试
├── native/               # 原生合约测试
├── net_performance_test/ # 网络性能测试
├── performance_pipeline/ # 性能管道测试
├── scenario0_native/     # 原生场景测试
├── scenario1_evm/        # EVM场景测试
├── scenario2_rust/       # Rust场景测试
├── scenario3_dockergo/   # Docker Go场景测试
├── scenario4_wasmer_sql/ # Wasmer SQL场景测试
├── send_proposal_request*/ # 提案请求测试系列
├── subscribe_test_tool/  # 订阅测试工具
├── testdata/            # 测试数据
├── tps/                 # TPS测试
├── tx_duplicate*/       # 交易重复相关测试
├── txfilter/           # 交易过滤器测试
├── utils/              # 测试工具
└── wasm/               # WebAssembly测试
```

## 测试类型说明

### 1. 性能测试
- **bench/**
  - 基准测试套件
  - 测量关键组件性能
  - 性能指标收集

- **net_performance_test/**
  - 网络性能测试
  - 带宽测试
  - 延迟测试
  - 吞吐量测试

- **tps/**
  - 交易每秒处理能力测试
  - 系统负载测试
  - 性能瓶颈分析

### 2. 功能测试
- **chainconfig_test/**
  - 链配置功能测试
  - 配置更新测试
  - 参数验证测试

- **crypto_test/**
  - 密码学功能测试
  - 加密解密测试
  - 签名验证测试

- **native/**
  - 原生合约功能测试
  - 接口测试
  - 异常处理测试

### 3. 场景测试
- **scenario0_native/**
  - 原生合约场景测试
  - 完整业务流程测试

- **scenario1_evm/**
  - 以太坊虚拟机场景测试
  - 智能合约部署和执行测试

- **scenario2_rust/**
  - Rust合约场景测试
  - Rust合约生命周期测试

- **scenario3_dockergo/**
  - Docker Go环境测试
  - 容器化部署测试

- **scenario4_wasmer_sql/**
  - Wasmer SQL场景测试
  - 数据库操作测试

### 4. 工具测试
- **grpc_client/**
  - gRPC客户端功能测试
  - 接口调用测试
  - 错误处理测试

- **subscribe_test_tool/**
  - 订阅功能测试
  - 事件监听测试
  - 消息处理测试

## 测试执行指南

### 1. 准备环境
```bash
# 安装依赖
cd test
go mod tidy

# 准备测试数据
./prepare_testdata.sh
```

### 2. 运行测试

#### 单个测试套件
```bash
# 运行特定测试
go test ./scenario1_evm/...

# 运行性能测试
go test ./bench/... -bench=.
```

#### 全套测试
```bash
# 运行所有测试
go test ./...

# 带覆盖率的测试
go test ./... -coverprofile=coverage.out
```

### 3. 性能测试

```bash
# 运行TPS测试
cd tps
./run_tps_test.sh

# 运行网络性能测试
cd net_performance_test
./run_network_test.sh
```

## 测试配置

### 1. 环境配置
- 配置文件位置：`config/test_config.yaml`
- 测试网络设置
- 节点配置
- 测试参数

### 2. 测试数据
- 测试数据位置：`testdata/`
- 示例合约
- 测试证书
- 模拟数据

## 最佳实践

1. 测试编写规范
   - 遵循表格驱动测试模式
   - 包含正向和反向测试用例
   - 添加详细测试注释
   - 确保测试的独立性

2. 性能测试注意事项
   - 在稳定环境下进行测试
   - 多次运行取平均值
   - 监控系统资源使用
   - 记录测试环境信息

3. 测试维护
   - 定期更新测试用例
   - 删除过时测试
   - 优化测试性能
   - 保持测试代码整洁

## 故障排除

1. 测试失败分析
   - 检查测试日志
   - 验证测试环境
   - 检查依赖版本
   - 隔离失败用例

2. 性能问题
   - 分析性能瓶颈
   - 检查资源使用
   - 优化测试配置
   - 调整测试参数

## 贡献指南

1. 添加新测试
   - 遵循现有测试结构
   - 添加详细文档
   - 包含测试用例说明
   - 确保测试可重复

2. 更新测试
   - 保持向后兼容
   - 更新相关文档
   - 验证所有场景
   - 添加变更说明

## 注意事项

1. 测试环境
   - 确保环境一致性
   - 避免环境污染
   - 清理测试数据

2. 资源管理
   - 及时释放资源
   - 控制测试规模
   - 监控资源使用

3. 数据安全
   - 保护测试数据
   - 避免敏感信息泄露
   - 定期清理测试数据

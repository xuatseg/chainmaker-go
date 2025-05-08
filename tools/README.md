# ChainMaker 工具集

本目录包含了长安链（ChainMaker）的各种实用工具，用于链码管理、区块扫描和数据库操作等功能。

## 工具列表

### CMC (ChainMaker CLI)
位置：`./cmc`

ChainMaker 命令行工具，提供了与区块链交互的完整功能集。

主要功能：
- 链码管理（部署、升级、调用）
- 合约操作
- 交易查询
- 区块查询
- 证书管理
- 权限管理

### Scanner
位置：`./scanner`

区块链数据扫描工具，用于分析和监控区块链数据。

主要功能：
- 区块扫描
- 交易分析
- 数据统计
- 状态监控
- 异常检测

### Simple-LevelDB
位置：`./simple-leveldb`

LevelDB 数据库管理工具，用于直接操作和查看区块链的底层数据。

主要功能：
- 数据库查询
- 键值对操作
- 数据导入导出
- 数据库维护
- 性能分析

## 使用指南

### CMC 工具使用

1. 基本命令格式：
```bash
./cmc [command] [options]
```

2. 常用命令：
```bash
# 查看帮助
./cmc --help

# 查看版本
./cmc version

# 部署链码
./cmc deploy [contract_name] [version] [path]

# 调用合约
./cmc invoke [contract_name] [method] [params]

# 查询交易
./cmc query tx [tx_id]
```

### Scanner 工具使用

1. 启动扫描：
```bash
./scanner start --config [config_file]
```

2. 常用选项：
```bash
# 指定扫描起始区块
--start-block [number]

# 指定扫描结束区块
--end-block [number]

# 设置扫描间隔
--interval [seconds]

# 输出格式设置
--output [format]
```

### Simple-LevelDB 工具使用

1. 基本操作：
```bash
# 打开数据库
./simple-leveldb open [db_path]

# 查询键值
./simple-leveldb get [key]

# 遍历数据
./simple-leveldb scan [start_key] [end_key]
```

2. 高级功能：
```bash
# 数据导出
./simple-leveldb export [output_file]

# 数据导入
./simple-leveldb import [input_file]

# 压缩数据库
./simple-leveldb compact
```

## 配置说明

### CMC 配置
- 配置文件位置：`./cmc/config.yaml`
- 主要配置项：
  - 链接信息
  - 证书配置
  - 默认参数

### Scanner 配置
- 配置文件位置：`./scanner/config.yaml`
- 主要配置项：
  - 扫描参数
  - 输出设置
  - 告警规则

### Simple-LevelDB 配置
- 配置文件位置：`./simple-leveldb/config.yaml`
- 主要配置项：
  - 数据库路径
  - 操作参数
  - 性能设置

## 最佳实践

1. 工具使用建议
   - 在测试环境验证命令
   - 保存重要数据备份
   - 记录操作日志

2. 安全注意事项
   - 妥善保管密钥文件
   - 控制访问权限
   - 定期更新工具版本

3. 性能优化
   - 合理设置扫描间隔
   - 优化查询参数
   - 定期维护数据库

## 故障排除

1. CMC 常见问题
   - 连接失败
   - 权限错误
   - 参数错误

2. Scanner 常见问题
   - 扫描卡住
   - 数据不同步
   - 内存占用过高

3. Simple-LevelDB 常见问题
   - 数据库损坏
   - 性能下降
   - 空间不足

## 开发指南

1. 工具扩展
   - 遵循接口规范
   - 保持向后兼容
   - 添加单元测试

2. 文档维护
   - 更新使用说明
   - 添加示例代码
   - 记录更新日志

3. 代码规范
   - 遵循项目编码规范
   - 做好错误处理
   - 添加适当注释

## 支持和反馈

如遇到问题：
1. 查看工具的帮助文档
2. 检查错误日志
3. 通过 Issue 系统报告问题
4. 寻求社区支持

## 注意事项

1. 使用前备份
   - 备份重要数据
   - 记录当前配置

2. 权限控制
   - 限制工具访问
   - 保护敏感数据

3. 资源管理
   - 监控资源使用
   - 及时清理日志
   - 定期维护数据

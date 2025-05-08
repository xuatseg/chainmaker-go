# ChainMaker 主程序目录

本目录包含了长安链（ChainMaker）的主程序入口和核心组件注册。

## 目录结构

```
main/
├── arm/                   # ARM架构相关文件
├── cmd/                   # 命令行工具
├── component_registry.go  # 组件注册
├── libwasmer.dylib       # Wasmer动态库 (MacOS)
├── libwasmer_runtime_c_api.so # Wasmer运行时库 (Linux)
├── main.go               # 主程序入口
├── pprof_readme.md       # 性能分析说明
└── prebuilt/             # 预编译文件
```

## 主要文件说明

### main.go
主程序入口文件，负责：
- 初始化系统配置
- 启动核心服务
- 注册必要组件
- 处理启动参数
- 配置日志系统

### component_registry.go
组件注册管理，负责：
- 注册各类组件
- 管理组件生命周期
- 处理组件依赖
- 提供组件查询接口

## 命令行工具

### cmd 目录
包含各种命令行工具，用于：
- 节点管理
- 配置更新
- 状态查询
- 调试工具

## WebAssembly 支持

### Wasmer 相关文件
- `libwasmer.dylib` (MacOS)
- `libwasmer_runtime_c_api.so` (Linux)

这些文件提供了 WebAssembly 运行时支持，用于：
- 智能合约执行
- WASM 模块加载
- 运行时环境管理

## 性能分析

详见 `pprof_readme.md` 文件，包含：
- 性能分析工具使用说明
- 分析数据收集方法
- 常见性能问题解决方案

## 使用指南

### 1. 编译

```bash
# 编译主程序
go build -o chainmaker main.go

# 指定平台编译
GOOS=linux GOARCH=amd64 go build -o chainmaker main.go
```

### 2. 运行

```bash
# 启动节点
./chainmaker start

# 指定配置文件启动
./chainmaker start --config /path/to/config.yml

# 启动调试模式
./chainmaker start --debug
```

### 3. 配置

主要配置项：
- 日志级别
- 监听端口
- 数据存储
- 共识参数
- P2P网络

## 开发指南

### 1. 添加新组件

1. 在 `component_registry.go` 中注册：
```go
func init() {
    RegisterComponent("component_name", NewComponent)
}
```

2. 实现组件接口：
```go
type Component interface {
    Init() error
    Start() error
    Stop() error
}
```

### 2. 修改启动流程

1. 更新 `main.go`：
```go
func main() {
    // 初始化配置
    config.Init()
    
    // 启动组件
    startComponents()
    
    // 等待信号
    waitForSignal()
}
```

## 调试指南

### 1. 性能分析

```bash
# 启用 pprof
./chainmaker start --pprof

# 查看 CPU 分析
go tool pprof http://localhost:6060/debug/pprof/profile

# 查看内存分析
go tool pprof http://localhost:6060/debug/pprof/heap
```

### 2. 日志调试

```bash
# 设置日志级别
./chainmaker start --log-level debug

# 输出日志到文件
./chainmaker start --log-file chain.log
```

## 部署说明

### 1. 系统要求
- Go 1.16+
- GCC 编译器
- 足够的磁盘空间
- 推荐 8GB+ 内存

### 2. 依赖检查
- 检查动态库
- 验证系统限制
- 确认端口可用
- 检查文件权限

### 3. 生产环境配置
- 配置系统服务
- 设置自动启动
- 配置日志轮转
- 设置监控告警

## 常见问题

### 1. 启动失败
- 检查配置文件
- 验证端口占用
- 确认权限设置
- 查看错误日志

### 2. 性能问题
- 使用 pprof 分析
- 检查系统资源
- 优化配置参数
- 更新系统限制

### 3. 组件错误
- 检查组件注册
- 验证依赖关系
- 确认版本兼容
- 查看组件日志

## 注意事项

1. 安全建议
   - 定期更新依赖
   - 及时修复漏洞
   - 控制访问权限
   - 加密敏感数据

2. 性能优化
   - 合理配置资源
   - 优化启动参数
   - 定期清理日志
   - 监控系统负载

3. 维护建议
   - 定期备份数据
   - 更新配置文件
   - 检查系统健康
   - 记录变更日志

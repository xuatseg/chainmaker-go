# Core 模块

Core 模块是长安链的核心实现模块，负责协调和管理各个核心组件的运行。

## 目录结构

```
core/
├── cache/          # 缓存实现
├── common/         # 公共组件
├── maxbftmode/     # MaxBFT 模式实现
├── provider/       # 核心服务提供者
├── syncmode/       # 同步模式实现
├── core_factory.go        # 核心工厂实现
└── core_provider_registrar.go  # 提供者注册器
```

## 核心组件关系图

```
                    +----------------+
                    |  CoreFactory  |
                    +-------+-------+
                            |
                 +---------++---------+
                 |                    |
          +------+-------+    +------+-------+
          |   Provider   |    |  Registrar   |
          +------+-------+    +--------------+
                 |
        +--------+--------+
        |               |
   +----+----+    +----+----+
   | MaxBFT  |    |  Sync   |
   +---------+    +---------+
```

## 组件说明

### 1. CoreFactory
- 负责创建和管理核心组件
- 提供组件生命周期管理
- 处理组件依赖关系

### 2. Provider
- 提供核心服务接口
- 实现各类服务功能
- 管理服务状态

### 3. Cache
- 实现缓存机制
- 提供数据快速访问
- 管理缓存生命周期

### 4. MaxBFT 模式
- 实现 MaxBFT 共识
- 处理共识流程
- 维护共识状态

### 5. Sync 模式
- 实现区块同步
- 处理节点同步
- 维护同步状态

## 主要流程

### 1. 初始化流程
```
1. CoreFactory 初始化
2. 注册核心服务提供者
3. 创建缓存实例
4. 初始化运行模式（MaxBFT/Sync）
5. 启动核心服务
```

### 2. 服务提供流程
```
1. 接收服务请求
2. Provider 处理请求
3. 调用相应模式处理
4. 返回处理结果
```

### 3. 模式切换流程
```
1. 检测切换条件
2. 暂停当前模式
3. 切换到新模式
4. 恢复服务处理
```

## 配置说明

### 1. 核心配置
```yaml
core:
  provider:
    type: default
  cache:
    size: 1000
  mode:
    type: maxbft
```

### 2. 模式配置
```yaml
maxbft:
  timeout: 10s
  batch_size: 500

sync:
  interval: 5s
  batch_size: 1000
```

## 开发指南

### 1. 添加新的服务提供者
```go
type NewProvider struct {
    // 实现 Provider 接口
}

func init() {
    RegisterProvider("new_provider", NewProvider)
}
```

### 2. 扩展运行模式
```go
type NewMode struct {
    // 实现 Mode 接口
}

func init() {
    RegisterMode("new_mode", NewMode)
}
```

## 使用示例

### 1. 创建核心实例
```go
factory := core.NewCoreFactory(config)
core := factory.Create()
```

### 2. 使用服务
```go
provider := core.GetProvider()
result := provider.HandleRequest(req)
```

## 注意事项

1. 线程安全
   - 确保并发访问安全
   - 使用适当的锁机制
   - 避免死锁情况

2. 性能优化
   - 合理使用缓存
   - 优化关键路径
   - 控制资源使用

3. 错误处理
   - 优雅处理异常
   - 提供错误恢复机制
   - 记录详细日志

## 调试指南

### 1. 日志级别
```go
logger.SetLevel(log.DebugLevel)
```

### 2. 性能分析
```go
pprof.StartCPUProfile(f)
defer pprof.StopCPUProfile()
```

### 3. 状态检查
```go
status := core.GetStatus()
metrics := core.GetMetrics()
```

## 常见问题

1. 服务启动失败
   - 检查配置正确性
   - 验证依赖服务
   - 查看错误日志

2. 性能问题
   - 检查缓存使用
   - 分析处理延迟
   - 优化关键代码

3. 模式切换问题
   - 验证切换条件
   - 检查状态一致性
   - 分析切换日志

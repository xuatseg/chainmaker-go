# Snapshot 模块

快照模块是长安链的核心组件之一，负责管理区块链状态的快照，支持状态回滚、数据恢复和证据管理等功能。

## 目录结构

```
snapshot/
├── shard_map.go                    # 分片映射管理
├── snapshot_evidence.go            # 快照证据实现
├── snapshot_factory.go             # 快照工厂实现
├── snapshot_impl.go                # 快照核心实现
├── snapshot_manager.go             # 快照管理器
├── snapshot_manager_evidence.go    # 证据管理器
└── snapshot_mananger_delegate.go   # 管理器代理
```

## 架构设计

```
                   +------------------+
                   | SnapshotFactory |
                   +--------+--------+
                            |
              +------------+-----------+
              |                       |
      +-------+-------+     +--------+--------+
      |  Snapshot    |     |    Manager     |
      |    Impl     |     |               |
      +------+-------+     +--------+-------+
             |                      |
    +--------+--------+   +--------+--------+
    |   Evidence    |    |  Shard Map     |
    |    Manager    |    |    Manager     |
    +---------------+    +----------------+
```

## 核心组件

### 1. SnapshotFactory
- 创建快照实例
- 管理快照配置
- 提供工厂方法

### 2. SnapshotManager
- 快照生命周期管理
- 状态管理
- 回滚控制

### 3. EvidenceManager
- 证据收集
- 证据验证
- 证据存储

## 主要接口

### 1. Snapshot 接口
```go
type Snapshot interface {
    // 获取快照版本
    GetVersion() uint64
    
    // 获取键值对
    Get(key []byte) ([]byte, error)
    
    // 设置键值对
    Set(key []byte, value []byte) error
    
    // 删除键值对
    Delete(key []byte) error
    
    // 提交快照
    Commit() error
    
    // 回滚快照
    Rollback() error
}
```

### 2. Manager 接口
```go
type SnapshotManager interface {
    // 创建快照
    CreateSnapshot(height uint64) (Snapshot, error)
    
    // 获取快照
    GetSnapshot(version uint64) (Snapshot, error)
    
    // 删除快照
    DeleteSnapshot(version uint64) error
    
    // 获取最新快照
    GetLatestSnapshot() (Snapshot, error)
}
```

## 快照生命周期

### 1. 创建阶段
```
1. 确定快照版本
2. 复制当前状态
3. 初始化快照数据
4. 记录元数据
```

### 2. 使用阶段
```
1. 读取状态数据
2. 修改状态数据
3. 验证数据一致性
4. 管理临时状态
```

### 3. 提交阶段
```
1. 验证变更
2. 持久化数据
3. 更新索引
4. 清理临时数据
```

## 配置说明

```yaml
snapshot:
  # 快照配置
  config:
    max_retain_size: 100
    auto_cleanup: true
    cleanup_interval: "1h"
    
  # 存储配置
  storage:
    path: "/path/to/snapshots"
    compression: true
    
  # 证据配置
  evidence:
    retain_height: 10000
    verify_level: "strict"
```

## 使用示例

### 1. 创建快照
```go
factory := NewSnapshotFactory(config)
manager := factory.CreateManager()

snapshot, err := manager.CreateSnapshot(height)
if err != nil {
    // 处理错误
}
```

### 2. 使用快照
```go
// 读取数据
value, err := snapshot.Get(key)

// 修改数据
err = snapshot.Set(key, newValue)

// 提交变更
err = snapshot.Commit()
```

## 性能优化

### 1. 存储优化
- 增量存储
- 数据压缩
- 批量操作

### 2. 内存优化
- 缓存管理
- 分片存储
- 垃圾回收

### 3. 并发优化
- 读写分离
- 并发控制
- 锁优化

## 监控指标

### 1. 快照指标
- 快照数量
- 快照大小
- 创建时间
- 提交延迟

### 2. 性能指标
- 读写吞吐量
- 响应时间
- 内存使用
- 磁盘使用

### 3. 错误指标
- 创建失败率
- 提交失败率
- 回滚次数

## 调试功能

### 1. 日志记录
```go
logger.Debug("snapshot operation",
    "version", snapshot.GetVersion(),
    "operation", "commit",
    "time_cost", timeCost,
)
```

### 2. 状态检查
```go
stats := manager.GetStats()
health := manager.CheckHealth()
```

## 常见问题

### 1. 存储问题
- 空间不足
- 数据损坏
- 索引异常

### 2. 性能问题
- 创建慢
- 提交延迟
- 内存溢出

### 3. 一致性问题
- 数据不一致
- 版本冲突
- 回滚失败

## 最佳实践

1. 快照管理
   - 定期清理
   - 合理保留
   - 监控容量

2. 性能优化
   - 批量操作
   - 异步提交
   - 合理配置

3. 数据安全
   - 定期备份
   - 验证完整性
   - 安全清理

## 注意事项

1. 存储建议
   - 预留足够空间
   - 定期维护
   - 监控状态

2. 性能建议
   - 控制快照大小
   - 优化访问模式
   - 合理使用缓存

3. 安全建议
   - 权限控制
   - 数据加密
   - 安全删除

## 错误处理

### 1. 创建错误
- 空间检查
- 权限验证
- 版本冲突

### 2. 操作错误
- 数据验证
- 并发控制
- 超时处理

### 3. 恢复策略
- 自动重试
- 回滚机制
- 错误通知

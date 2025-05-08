# 并行测试工具

本工具提供对长安链的并行压力测试能力，支持高并发交易提交和查询。

## 功能特性

### 1. 测试模式
- 合约创建压力测试
- 合约调用并发测试
- 混合读写测试
- 持久化稳定性测试

### 2. 监控指标
- TPS实时统计
- 成功率监控
- 延迟分布
- 资源使用率

## 使用指南

### 1. 基本命令
```bash
# 启动并行调用测试
cmc parallel invoke \
  -c 10 \          # 10个并发worker
  -n 10000 \       # 总共10000次调用
  -contract my_contract \
  -method transfer

# 带参数调用
cmc parallel invoke \
  -params '{"to":"0x123","amount":100}' \
  -rate 500        # 500 TPS限速
```

### 2. 测试报告
```
=== 测试结果 ===
Duration:   60.0s
Workers:    10
Total TX:   58234
Success:    58001 (99.6%)
Avg TPS:    966.7
Avg Latency: 103.2ms
Max Latency: 452ms
CPU Usage:  85%
Memory:     2.3GB
```

## 实现架构

### 1. 核心组件
- Worker Pool: 并发工作池
- Rate Limiter: TPS控制
- Monitor: 实时监控
- Reporter: 结果统计

### 2. 数据流
```
Worker -> TX Generator -> Chain Client -> Monitor
                        ↓
                  Result Collector
```

## 开发扩展

### 1. 添加新测试类型
1. 实现Generator接口
2. 注册到测试工厂
3. 添加CLI参数支持

### 2. 自定义指标
- 添加新的监控指标
- 实现自定义报告格式
- 支持外部监控系统集成

## 注意事项

1. 测试安全
   - 使用测试专用账户
   - 避免生产环境误操作
   - 清理测试数据

2. 资源控制
   - 合理设置并发数
   - 监控系统资源
   - 避免OOM

3. 结果分析
   - 关注成功率
   - 分析延迟分布
   - 检查错误模式

# 交互式控制台

本模块提供长安链的交互式命令行控制台，支持便捷的链操作和调试功能。

## 功能特性

### 1. 交互模式
- 命令补全
- 历史记录
- 上下文感知帮助
- 多行输入支持

### 2. 调试功能
- 交易构造器
- 合约调试器
- 状态检查器
- 事件监视器

### 3. 开发辅助
- 脚本执行
- 测试数据生成
- 性能分析
- 模拟器集成

## 使用指南

### 1. 启动控制台
```bash
cmc console -config chainmaker.yaml
```

### 2. 常用命令
```bash
# 查询区块
>> getBlock(100)

# 调用合约
>> contract.invoke("my_contract", "transfer", {to:"0x123", amount:100})

# 查看帮助
>> help()
```

## 功能扩展

### 1. 自定义命令
1. 实现Command接口
2. 注册到命令列表
3. 添加帮助文档

### 2. 主题定制
- 颜色方案
- 提示符定制
- 输出格式

## 开发接口

### 1. Go API
```go
// 创建控制台
console := NewConsole(config)

// 注册命令
console.Register("mycmd", myCommand{})

// 运行交互循环
console.Run()
```

### 2. 插件开发
1. 添加语法高亮
2. 集成调试器
3. 支持更多REPL特性

## 注意事项

1. 安全建议
   - 限制敏感操作
   - 命令历史加密
   - 会话超时设置

2. 性能优化
   - 响应式设计
   - 后台加载
   - 结果缓存

3. 用户体验
   - 清晰的错误提示
   - 上下文帮助
   - 可发现的特性

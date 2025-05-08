# 命令行工具辅助功能

本目录包含 cmc 命令行工具的辅助功能和实用程序。

## 功能列表

### 1. 基础转换
- Base64 与 Hex 互转
- JSON 格式美化
- 时间格式转换

### 2. 数据验证
- 签名验证
- 地址校验
- 格式检查

### 3. 工具集成
- 管道操作支持
- 批量处理
- 脚本集成

## 使用示例

### 1. 格式转换
```bash
# Hex 转 Base64
cmc base64-to-hex 48656c6c6f

# Base64 转 Hex
cmc hex-to-base64 SGVsbG8=
```

### 2. 数据验证
```bash
# 验证签名
cmc verify-sig -data test.txt -sig signature.bin -pubkey pub.pem
```

## 开发接口

### 1. 核心函数
```go
// Base64 转 Hex
func Base64ToHex(base64Str string) (string, error)

// Hex 转 Base64 
func HexToBase64(hexStr string) (string, error)
```

### 2. 扩展开发
1. 添加新的转换功能
2. 支持更多数据格式
3. 集成验证工具链

## 注意事项

1. 数据安全
   - 敏感信息处理
   - 内存清理
   - 临时文件管理

2. 性能考虑
   - 大文件处理
   - 批量操作优化
   - 并发安全

3. 错误处理
   - 格式错误检测
   - 详细错误信息
   - 恢复机制

# 零知识证明工具 (Bulletproofs)

本模块实现基于Bulletproofs的零知识证明功能，支持范围证明和隐私交易验证。

## 核心功能

### 1. 范围证明
- 生成数值的范围证明
- 验证证明的有效性
- 支持自定义范围

### 2. 隐私交易
- 隐藏交易金额
- 隐藏交易方
- 可验证的隐私交易

## 使用示例

### 1. 生成范围证明
```bash
cmc bulletproofs genproof \
  -value 100 \
  -min 0 \
  -max 1000 \
  -output proof.json
```

### 2. 验证证明
```bash
cmc bulletproofs verify \
  -proof proof.json \
  -commitment comm.txt
```

## 实现原理

### 1. 技术基础
- 基于椭圆曲线密码学
- 使用Pedersen承诺
- 内积证明优化

### 2. 证明流程
1. 生成承诺
2. 创建证明
3. 验证证明
4. 打开承诺

## 开发接口

### 1. Go API
```go
// 生成范围证明
proof, err := bulletproofs.Prove(value, min, max)

// 验证证明
valid := bulletproofs.Verify(proof, commitment) 
```

### 2. CLI命令
| 命令 | 说明 |
|------|------|
| genproof | 生成范围证明 |
| verify | 验证证明 |
| genopening | 生成承诺打开 |

## 注意事项

1. 参数安全
   - 使用安全随机数
   - 保护私密参数

2. 性能考虑
   - 大范围证明较耗时
   - 批量验证优化

3. 兼容性
   - 验证曲线参数
   - 检查证明版本

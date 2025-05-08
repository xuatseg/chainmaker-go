# Access Control 模块

访问控制模块是长安链的安全基础设施，负责身份认证、权限管理和访问策略控制。

## 目录结构

```
accesscontrol/
├── ac_factory.go                    # 访问控制工厂
├── ac_provider_registrar.go         # 提供者注册器
├── ac_server.go                     # 访问控制服务器
├── cert_ac.go                       # 证书访问控制
├── cert_ac_subscriber.go            # 证书访问控制订阅者
├── cert_member.go                   # 证书成员管理
├── interface_base.go                # 基础接口定义
├── interface_inner.go               # 内部接口定义
├── organization.go                  # 组织管理
├── permissioned_pk_ac.go           # 许可公钥访问控制
├── public_pk_ac.go                 # 公共公钥访问控制
├── policies.go                     # 策略定义
├── policy.go                       # 策略实现
├── principal.go                    # 主体管理
├── provider.go                     # 服务提供者
└── utils.go                        # 工具函数
```

## 架构设计

```
                    +----------------+
                    |  ACFactory    |
                    +-------+-------+
                            |
              +------------+-----------+
              |                       |
      +-------+-------+     +--------+--------+
      |  ACProvider   |     |  ACRegistrar   |
      +-------+-------+     +----------------+
              |
     +--------+---------+
     |                  |
+----+----+      +-----+-----+
| CertAC  |      |  PKAC     |
+---------+      +-----------+
```

## 核心组件

### 1. 访问控制工厂 (ACFactory)
- 创建访问控制实例
- 管理访问控制配置
- 提供工厂方法接口

### 2. 访问控制提供者 (ACProvider)
- 实现访问控制逻辑
- 管理认证和授权
- 处理访问请求

### 3. 证书访问控制 (CertAC)
- 基于证书的身份认证
- 证书链验证
- 证书撤销管理

### 4. 公钥访问控制 (PKAC)
- 基于公钥的身份认证
- 权限管理
- 访问控制列表

## 主要接口

### 1. 基础接口
```go
type AccessControlProvider interface {
    // 验证签名
    VerifySignature(msg []byte, sig []byte, member Member) error
    
    // 验证证书
    VerifyCertification(cert *x509.Certificate) error
    
    // 获取成员信息
    GetMember(id string) (Member, error)
}
```

### 2. 策略接口
```go
type Policy interface {
    // 评估访问请求
    Evaluate(principal Principal, resource string) bool
    
    // 更新策略
    Update(policy *PolicyConfig) error
    
    // 获取策略信息
    GetInfo() *PolicyInfo
}
```

## 认证流程

### 1. 证书认证
```
1. 接收证书
2. 验证证书链
3. 检查证书状态
4. 验证签名
5. 授权访问
```

### 2. 公钥认证
```
1. 接收公钥
2. 验证签名
3. 检查权限
4. 授权访问
```

## 配置说明

```yaml
accesscontrol:
  # 认证方式
  auth_type: "cert"  # cert/pk
  
  # 证书配置
  cert:
    root_ca: "/path/to/ca.crt"
    chain_trust: true
    crl_check: true
    
  # 策略配置
  policy:
    rule_file: "/path/to/rules.yaml"
    default_action: "deny"
    
  # 成员配置
  member:
    cache_size: 1000
    update_interval: "10m"
```

## 策略定义

### 1. 规则格式
```yaml
rules:
  - name: "admin_access"
    principal: "admin"
    resources: ["chain.*", "config.*"]
    action: "allow"
    
  - name: "user_access"
    principal: "user"
    resources: ["chain.query"]
    action: "allow"
```

## 使用示例

### 1. 验证访问权限
```go
// 创建访问控制实例
ac := acFactory.CreateAccessControl(config)

// 验证访问权限
if err := ac.VerifyAccess(member, resource); err != nil {
    // 处理访问拒绝
}
```

### 2. 管理策略
```go
// 更新策略
policy := &PolicyConfig{
    Rules: []Rule{...},
}
ac.UpdatePolicy(policy)
```

## 安全特性

### 1. 身份认证
- 多级证书链验证
- CRL 检查
- 签名验证

### 2. 访问控制
- 细粒度权限控制
- 动态策略更新
- 审计日志

### 3. 安全防护
- 防重放攻击
- 时间戳验证
- 会话管理

## 性能优化

### 1. 缓存优化
- 证书缓存
- 成员信息缓存
- 策略缓存

### 2. 验证优化
- 并行验证
- 批量处理
- 预验证

## 监控指标

### 1. 认证指标
- 认证请求数
- 认证成功率
- 认证延迟

### 2. 策略指标
- 策略评估数
- 策略命中率
- 策略更新数

## 调试功能

### 1. 日志记录
```go
logger.Debug("access verification",
    "member", member.ID,
    "resource", resource,
    "result", result,
)
```

### 2. 状态检查
```go
status := ac.GetStatus()
metrics := ac.GetMetrics()
```

## 常见问题

### 1. 证书问题
- 证书链验证失败
- 证书过期
- CRL 更新失败

### 2. 权限问题
- 策略配置错误
- 权限不足
- 资源未定义

## 最佳实践

1. 证书管理
   - 定期更新证书
   - 及时更新 CRL
   - 安全存储私钥

2. 策略管理
   - 最小权限原则
   - 定期审查策略
   - 记录策略变更

3. 安全建议
   - 启用审计日志
   - 监控异常访问
   - 定期安全评估

## 注意事项

1. 开发建议
   - 完整错误处理
   - 安全日志记录
   - 性能优化

2. 运维建议
   - 证书生命周期管理
   - 策略定期审查
   - 监控告警配置

3. 安全建议
   - 私钥保护
   - 访问控制审计
   - 安全更新管理
# Access Control 
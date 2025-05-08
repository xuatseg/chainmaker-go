# ChainMaker 监控系统

本目录包含了长安链（ChainMaker）的监控系统配置文件和相关组件。监控系统基于 Prometheus + Grafana 构建，提供了全面的性能监控和可视化功能。

## 目录结构

```
monitor/
├── dashboard.json      # Grafana 仪表盘配置
├── docker-compose.yml  # Docker 编排文件
├── grafana.ini        # Grafana 配置文件
├── grpc.json          # gRPC 监控配置
├── index.html         # 监控首页
├── influx.sql         # InfluxDB 初始化脚本
├── mysql.sql          # MySQL 初始化脚本
└── prometheus.yml     # Prometheus 配置文件
```

## 快速开始

### 1. 启动监控系统

```bash
# 启动所有监控组件
docker-compose up -d

# 检查服务状态
docker-compose ps
```

### 2. 访问监控界面

- Grafana: http://localhost:3000
- Prometheus: http://localhost:9090

## 组件说明

### 1. Prometheus
- 配置文件：`prometheus.yml`
- 主要功能：
  - 指标收集
  - 数据存储
  - 告警规则
  - 服务发现

配置示例：
```yaml
global:
  scrape_interval: 15s
  evaluation_interval: 15s

scrape_configs:
  - job_name: 'chainmaker'
    static_configs:
      - targets: ['localhost:8080']
```

### 2. Grafana
- 配置文件：`grafana.ini`
- 仪表盘：`dashboard.json`
- 主要功能：
  - 数据可视化
  - 告警配置
  - 用户管理
  - 权限控制

### 3. 数据库
- InfluxDB 脚本：`influx.sql`
- MySQL 脚本：`mysql.sql`
- 用途：
  - 性能数据存储
  - 历史数据查询
  - 趋势分析

## 监控指标

### 1. 系统指标
- CPU 使用率
- 内存使用情况
- 磁盘 I/O
- 网络流量

### 2. 链性能指标
- TPS (每秒交易数)
- 区块生成时间
- 交易确认时间
- 共识延迟

### 3. 网络指标
- P2P 连接状态
- 消息传播延迟
- 带宽使用情况
- 节点响应时间

### 4. 合约指标
- 合约调用次数
- 执行时间
- 资源消耗
- 错误率

## 告警配置

### 1. 告警规则
- 系统资源告警
- 性能阈值告警
- 节点状态告警
- 错误率告警

### 2. 通知方式
- 邮件通知
- Webhook
- Slack
- 企业微信

## 配置说明

### 1. Docker 配置
```yaml
# docker-compose.yml
version: '3'
services:
  prometheus:
    image: prom/prometheus
    volumes:
      - ./prometheus.yml:/etc/prometheus/prometheus.yml
    ports:
      - "9090:9090"

  grafana:
    image: grafana/grafana
    volumes:
      - ./grafana.ini:/etc/grafana/grafana.ini
    ports:
      - "3000:3000"
```

### 2. Grafana 配置
```ini
# grafana.ini
[server]
http_port = 3000

[security]
admin_user = admin
admin_password = admin

[auth.anonymous]
enabled = false
```

### 3. 数据库配置
```sql
-- influx.sql
CREATE DATABASE chainmaker;
CREATE RETENTION POLICY "30days" ON "chainmaker" DURATION 30d REPLICATION 1 DEFAULT;
```

## 运维指南

### 1. 日常维护
- 检查服务状态
- 清理历史数据
- 更新配置文件
- 备份重要数据

### 2. 故障处理
- 服务无响应
- 数据异常
- 告警误报
- 性能问题

### 3. 扩容升级
- 添加新节点
- 更新组件版本
- 调整资源配置
- 优化性能

## 最佳实践

1. 监控部署
   - 合理规划资源
   - 配置高可用
   - 做好备份
   - 定期维护

2. 告警配置
   - 设置合理阈值
   - 避免告警风暴
   - 分级处理
   - 定期review

3. 数据管理
   - 合理设置保留期
   - 定期归档
   - 容量规划
   - 性能优化

## 常见问题

1. 数据收集问题
   - 检查网络连接
   - 验证配置正确性
   - 查看服务日志
   - 确认权限设置

2. 展示问题
   - 刷新数据源
   - 清理浏览器缓存
   - 检查时间范围
   - 验证查询语句

3. 性能问题
   - 优化查询
   - 调整采集间隔
   - 清理无用数据
   - 增加资源配置

## 安全建议

1. 访问控制
   - 启用认证
   - 设置强密码
   - 限制访问IP
   - 定期修改密码

2. 数据安全
   - 加密敏感数据
   - 定期备份
   - 访问审计
   - 数据脱敏

3. 网络安全
   - 使用HTTPS
   - 配置防火墙
   - 限制端口访问
   - 启用SSL/TLS

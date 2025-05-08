# ChainMaker 脚本目录

本目录包含了长安链（ChainMaker）的各类脚本，用于项目的构建、部署、测试和维护。

## 目录结构

```
scripts/
├── bin/                    # 可执行文件目录
├── docker/                 # Docker 相关脚本
├── service/               # 服务相关脚本
├── test/                  # 测试相关脚本
└── *.sh                   # 各类功能脚本
```

## 主要脚本说明

### 构建和发布脚本
- **build_release.sh**
  - 用途：构建发布版本
  - 功能：编译源码、打包资源、生成发布包
  - 使用：`./build_release.sh`

### 环境准备脚本
- **prepare.sh**
  - 用途：准备标准环境
  - 功能：安装依赖、初始化配置
  - 使用：`./prepare.sh`

- **prepare_pk.sh**
  - 用途：准备公钥环境
  - 功能：配置公钥相关环境
  - 使用：`./prepare_pk.sh`

- **prepare_pwk.sh**
  - 用途：准备公私钥环境
  - 功能：配置公私钥相关环境
  - 使用：`./prepare_pwk.sh`

### 集群管理脚本
- **cluster_quick_start.sh**
  - 用途：快速启动集群
  - 功能：初始化并启动节点集群
  - 使用：`./cluster_quick_start.sh`

- **cluster_quick_stop.sh**
  - 用途：快速停止集群
  - 功能：安全停止所有节点
  - 使用：`./cluster_quick_stop.sh`

- **range_cluster_quick_start.sh**
  - 用途：范围式启动集群
  - 功能：按指定范围启动节点
  - 使用：`./range_cluster_quick_start.sh [start_index] [end_index]`

### 依赖管理脚本
- **gomod_update.sh**
  - 用途：更新 Go 模块依赖
  - 功能：更新和整理 go.mod 文件
  - 使用：`./gomod_update.sh`

### 测试脚本
- **ut_cover.sh**
  - 用途：运行单元测试并生成覆盖率报告
  - 功能：执行测试、统计覆盖率
  - 使用：`./ut_cover.sh`

## 子目录说明

### bin 目录
- 存放编译生成的可执行文件
- 包含各类工具和程序

### docker 目录
- Docker 相关配置和脚本
- 容器化部署所需文件

### service 目录
- 服务管理相关脚本
- 系统服务配置文件

### test 目录
- 测试相关脚本
- 测试用例和测试数据

## 使用指南

### 首次使用

1. 准备环境：
```bash
./prepare.sh
```

2. 构建项目：
```bash
./build_release.sh
```

3. 启动集群：
```bash
./cluster_quick_start.sh
```

### 日常开发

1. 更新依赖：
```bash
./gomod_update.sh
```

2. 运行测试：
```bash
./ut_cover.sh
```

### 部署维护

1. 停止集群：
```bash
./cluster_quick_stop.sh
```

2. 范围启动：
```bash
./range_cluster_quick_start.sh 1 4  # 启动节点1到4
```

## 注意事项

1. 执行权限
   - 确保脚本具有执行权限
   - 必要时使用 `chmod +x script.sh`

2. 环境要求
   - 确保已安装必要的依赖
   - 检查系统环境变量配置

3. 执行顺序
   - 遵循推荐的执行顺序
   - 注意脚本间的依赖关系

4. 错误处理
   - 查看日志输出
   - 检查错误信息
   - 按需清理临时文件

## 故障排除

1. 脚本执行失败
   - 检查执行权限
   - 验证环境配置
   - 查看详细日志

2. 集群启动问题
   - 检查端口占用
   - 验证配置文件
   - 确认节点状态

3. 测试覆盖率问题
   - 确保测试环境完整
   - 检查测试依赖
   - 验证测试用例

## 贡献指南

1. 脚本开发规范
   - 添加适当的注释
   - 包含使用说明
   - 做好错误处理
   - 保持代码风格一致

2. 测试要求
   - 提供测试用例
   - 验证各种场景
   - 确保向后兼容

3. 文档更新
   - 更新相关文档
   - 添加使用示例
   - 说明注意事项

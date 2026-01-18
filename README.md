# goblog
这是一个前后端分离的项目，使用后端go+gin+gorm，使用vue3作为前端
## 项目技术栈介绍

### 后端

#### gin

#### MySQL

### 前端

#### vue3

## 开发流程

### 读取配置

1.配置使用规范
2.配置只在启动阶段加载一次
3.不使用全局 Config
4.不在业务层 import bootstrap
5.通过构造函数显式传递依赖
6.bootstrap 只做“初始化与组装”

yaml 写入配置文件，然后通过config/config.go解析配置，在main.go中通过返回值使用！
```yaml
mysql:
  host: 127.0.0.1
  port: 3306
  dbname: dev_db
  username: root
  password: rootpassword
  log_level: dev

logger:
  level: info
  prefix: "[DEV] "
  director: log
  show_line: true
  log_in_console: true

system:
  host: "0.0.0.0"
  port: 8080
  env: dev

```


```go
package bootstrap

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Mysql  Mysql  `yaml:"mysql"`
	Logger Logger `yaml:"logger"`
	System System `yaml:"system"`
}

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"`
}

type Logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"show_line"`
	LogInConsole bool   `yaml:"log_in_console"`
}

type System struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
}

func LoadConfig(path string) (*Config, error) {

	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config file failed: %w", err)
	}

	var c Config
	if err := yaml.Unmarshal(b, &c); err != nil {
		return nil, fmt.Errorf("unmarshal config file failed: %w", err)
	}

	return &c, nil
}
```

## go_structure_standard
**go包名的规范**：全小写、无下划线、见名知义、简洁不冗余，比如utils/valid是合规的，global合规但「慎用」，这是 Go 官方明确要求的。
```
.
├── cmd/                        # 程序入口目录
│   └── server/
│       └── main.go             # 项目唯一入口：加载配置、初始化依赖、启动 HTTP 服务
│
├── bootstrap/                  # 项目启动初始化
│   ├── config.go               # 配置加载（viper/env）
│   ├── logger.go               # 日志初始化
│   ├── gorm.go                 # MySQL 初始化
│   ├── redis.go                # Redis 初始化
│   ├── router.go               # Gin 路由注册
│   └── app.go                  # 应用启动编排（聚合初始化）
│
├── configs/                    # 多环境配置
│   ├── dev.yaml
│   ├── test.yaml
│   └── prod.yaml
│
├── internal/                   # 核心业务代码（禁止外部引用）
│   ├── handler/                # HTTP 接口层
│   │   ├── user/
│   │   │   └── handler.go
│   │   └── goods/
│   │       └── handler.go
│   │
│   ├── service/                # 业务逻辑层（纯业务）
│   │   ├── user/
│   │   │   └── service.go
│   │   └── goods/
│   │       └── service.go
│   │
│   ├── dao/                    # 数据访问层
│   │   ├── user/
│   │   │   ├── mysql.go
│   │   │   └── redis.go
│   │   └── goods/
│   │       └── mysql.go
│   │
│   ├── entity/                 # 领域实体（DB 映射）
│   │   ├── user.go
│   │   └── goods.go
│   │
│   ├── dto/                    # 数据传输对象（Request / Response）
│   │   ├── user.go
│   │   └── goods.go
│   │
│   ├── middleware/             # 中间件
│   │   ├── cors.go
│   │   ├── jwt.go
│   │   └── logger.go
│   │
│   └── constant/               # 项目常量 & 枚举（替代 global）
│       ├── error_code.go
│       └── common.go
│
├── pkg/                        # 项目级公共组件（有业务语义）
│   ├── resp/                   # 统一响应封装
│   ├── logger/                 # 日志封装
│   ├── encrypt/                # 加密（hash / jwt）
│   └── validator/              # 数据校验
│
├── utils/                      # 通用工具函数（无业务语义）
│   ├── stringx/
│   ├── timex/
│   └── filex/
│
├── docs/                       # Swagger / 接口文档
├── logs/                       # 日志文件
├── scripts/                    # 构建 / 迁移 / 启动脚本
├── test/                       # 单元测试 / 集成测试
│   ├── service/
│   └── api/
│
├── Makefile                    # 一键操作
├── go.mod
├── go.sum
└── README.md
```


### 1.项目整体设计原则

* **高内聚、低耦合**
* **分层清晰，职责单一**
* **业务逻辑独立于框架**
* **可测试、可扩展**

---

### 2.顶层目录规范

| 目录         | 说明                 |
| ------------ | -------------------- |
| `cmd/`       | 程序入口，仅负责启动 |
| `bootstrap/` | 初始化配置、资源     |
| `internal/`  | 核心业务代码         |
| `pkg/`       | 项目级公共组件       |
| `utils/`     | 通用工具函数         |
| `configs/`   | 配置文件             |
| `scripts/`   | 构建 / 运维脚本      |

---

### 3.分层架构规范（强制）

```text
HTTP Request
    ↓
Handler（接口层）
    ↓
Service（业务层）
    ↓
DAO（数据层）
    ↓
Database / Cache
```

---

### 4.各层职责定义

#### Handler 层

* 解析 HTTP 请求
* 参数校验
* 调用 Service
* 返回统一响应

❌ 禁止：

* 写业务逻辑
* 直接操作数据库

---

#### Service 层（核心）

* 编写业务规则
* 组合多个 DAO
* 处理事务

✅ 特性：

* 不依赖 Gin / HTTP
* 可被单元测试直接调用

---

#### DAO 层

* 数据库 CRUD
* Redis 访问
* 不包含业务规则

---

#### Entity 层

* 表结构映射
* 领域模型
* 仅定义结构体

---

#### DTO 层

* 请求参数结构
* 响应结构
* 仅服务于 HTTP 层

---

### 5.internal 目录规范

* `internal` 下的包：

  * ❌ 禁止被外部项目引用
  * ✅ 仅服务于本项目
* 所有核心业务代码必须放在 `internal`

---

### 6.pkg vs utils 使用规范

#### pkg

* 项目专属
* 含业务语义
* 不可直接复用到其他项目

#### utils

* 通用工具
* 与业务无关
* 可跨项目拷贝使用

---

### 7.常量与全局状态规范

* 使用 `constant/` 存放：

  * 错误码
  * 枚举
  * Context Key
* ❌ 禁止滥用全局变量
* ❌ 禁止存放 DB / Redis 实例

---

### 8.设计合理性自检清单（面试级）

* Service 是否可脱离 HTTP 使用？
* Entity 是否未被 HTTP 污染？
* Handler 是否只做接口适配？
* 是否遵守单向依赖？

---

### 9.推荐实践（进阶）

* 使用 Makefile 管理常用命令
* 使用接口（interface）隔离 Service / DAO
* 为 Service 编写单元测试
* 通过构造函数注入依赖

---

# 托希歌的个人博客

一个从零搭建、**前后端分离**的个人博客系统。后端基于 **Go + Gin** 提供 RESTful API，前端基于 **Vue 3 + TypeScript** 构建单页应用，数据层采用 **MySQL + Elasticsearch + Redis** 的组合——关系数据存 MySQL，文章正文与统计存 Elasticsearch 以支撑全文搜索，热榜、浏览量、登录状态等走 Redis 缓存。

项目覆盖了一个个人博客所需的完整闭环：Markdown 文章发布与全文搜索、标签云、树形评论、多平台热榜聚合、用户注册登录（邮箱 / QQ）、反馈、图片与站点配置管理，并内置前台展示与后台管理两套界面。既可作个人博客直接部署，也可作为前后端分离项目、Go + Vue 技术栈的参考实现。

## 🛠 技术栈

| 层 | 技术 |
|---|---|
| **后端** | Go 1.26 · Gin · GORM · Elasticsearch 8 · Redis · JWT · Zap · Cron · 七牛云 |
| **前端** | Vue 3 · TypeScript · Vite 5 · Pinia · Vue Router · Element Plus · md-editor-v3 · Lucide |
| **数据 / 中间件** | MySQL 8 · Redis · Elasticsearch 8 |

## 🧩 架构与技术亮点

- **分层后端**：`router → api → service → model` 清晰分层，Gin 提供 RESTful API
- **Elasticsearch 全文搜索**：文章正文存储于 ES，支持关键词 + 标签搜索，Painless 脚本增量更新浏览 / 评论 / 收藏量
- **双 Token 鉴权**：access + refresh 自动续期，Redis + 本地缓存黑名单实现多点登录互踢
- **Redis 多级缓存**：热榜、浏览量、天气、JWT 状态，Cron 定时同步与预热
- **统一响应**：所有接口统一 `{code, data, msg}` 结构，前端拦截器统一处理

## ✨ 功能特性

- 📝 **文章**：Elasticsearch 全文搜索、标签云、收藏、浏览量统计（Redis 缓存 + 定时同步）
- 💬 **评论**：树形嵌套评论、回复、删除（仅本人或管理员）
- 📊 **热榜**：聚合 8 个平台热搜（百度 / B站 / 知乎 / 抖音 / HelloGitHub / IT之家 / 掘金 / CSDN），Redis 缓存 + 定时预热
- 👤 **用户**：邮箱注册 / 登录、QQ 登录、忘记密码、冻结 / 解冻、登录日志
- 📮 **反馈**：用户反馈 + 管理员回复
- 🖼 **图片**：本地上传 / 七牛云，多类别管理（插图 / 封面 / 友链等）
- 🗣 **碎碎念**：发布 / 删除
- ⚙️ **在线配置**：网站 / 系统 / 邮箱 / QQ / 七牛 / JWT / 高德 配置（动态写回 `config.yaml`）
- 🔐 **双 Token 鉴权**：access token + refresh token 自动续期，多点登录互踢
- 🌤 **天气**：高德 IP 定位 + 实时天气（Redis 缓存）

## 📁 项目结构

```
myBlog/
├── server/                 # Go 后端
│   ├── main.go             # 入口
│   ├── config.yaml         # 配置文件（含密钥，已 gitignore，需手动创建）
│   ├── api/                # 接口层（handler）
│   ├── router/             # 路由
│   ├── middleware/         # 中间件（JWT、管理员鉴权、登录记录、日志）
│   ├── service/            # 业务逻辑层
│   ├── model/              # 数据模型（database / request / response / ...）
│   ├── initialize/         # 初始化（DB / ES / Redis / 路由 / 定时任务）
│   ├── core/               # 配置加载 / 日志 / 服务器
│   ├── flag/               # CLI 工具（迁移、导入导出、建管理员）
│   ├── task/               # 定时任务
│   ├── utils/              # 工具函数
│   └── uploads/            # 本地上传图片（gitignore）
│
└── web/                    # Vue 3 前端
    ├── src/
    │   ├── api/            # API 封装
    │   ├── components/     # 公共组件
    │   ├── views/          # 页面（前台 + 后台）
    │   ├── stores/         # Pinia 状态管理
    │   ├── router/         # 路由
    │   └── assets/         # 主题样式
    ├── public/             # 静态资源（默认头像 / 兜底图标 / logo）
    └── vite.config.ts
```

## 🚀 快速开始

### 环境要求

- Go 1.26+
- Node.js 18+（推荐 20+）
- MySQL 8
- Redis
- Elasticsearch 8

### 1. 后端

```bash
cd server

# ① 创建配置文件（config.yaml 已被 gitignore，需手动创建，见下方「配置说明」）

# ② 初始化 MySQL 表结构
go run . -sql

# ③ 初始化 Elasticsearch 索引
go run . -es

# ④ 创建管理员账号（交互式输入邮箱和密码）
go run . -admin

# ⑤ 启动服务（默认 8080）
go run .
```

### 2. 前端

```bash
cd web
npm install
npm run dev
```

开发服务器默认跑在 `80` 端口（`vite.config.ts` 里配置），并将 `/api`、`/uploads` 代理到后端 `http://127.0.0.1:8080`。若 80 端口被占用，可用 `npm run dev -- --port 5173` 换端口。

## ⚙️ 配置说明

`server/config.yaml` 是全局配置，包含数据库、缓存、邮件、对象存储、第三方登录等。**因为含密钥，已被 `.gitignore` 忽略**，需要手动创建。主要字段：

| 配置块 | 说明 |
|---|---|
| `mysql` | MySQL 连接（host / port / db_name / 账号密码） |
| `redis` | Redis 连接 |
| `es` | Elasticsearch 连接 |
| `system` | 服务地址 / 端口 / 环境 / 路由前缀 / 单点登录 |
| `jwt` | access / refresh token 密钥与有效期 |
| `email` | SMTP 发信配置（验证码邮件） |
| `qiniu` | 七牛云对象存储 |
| `qq` | QQ 登录 appid / key / 回调 |
| `gaode` | 高德天气 key |
| `captcha` | 图形验证码参数 |
| `upload` | 上传大小 / 本地存储路径 |
| `website` | 网站信息（标题 / 备案 / 社交链接 / 联系方式等） |
| `hot` | 热榜平台列表（开关 / 图标） |
| `yiyan` | 一言开关与默认句 |
| `zap` | 日志配置 |

> 密钥（邮箱 secret、七牛 AK/SK、JWT secret 等）建议部署时通过环境变量或配置模板注入，避免打进镜像 / 提交到仓库。

## 🧰 CLI 工具

后端支持以下命令行参数（`go run . <flag>`）：

| 参数 | 说明 |
|---|---|
| `-sql` | 初始化 MySQL 表结构 |
| `-sql-export` | 导出表结构到 `.sql` |
| `-sql-import <file>` | 导入 SQL 数据 |
| `-es` | 初始化 ES 索引 |
| `-es-export` | 导出 ES 数据到 `.json` |
| `-es-import <file>` | 导入 ES 数据 |
| `-admin` | 交互式创建管理员 |

> 一次只能传一个 flag。不带任何 flag 则正常启动服务。

## 🔧 部署

生产环境需要以下组件：**Go 后端 + 前端静态资源（Nginx）+ MySQL + Redis + Elasticsearch**。

- 中间件（MySQL / Redis / ES）可直接用官方 Docker 镜像。
- 后端与前端建议容器化后部署；前端打包后交给 Nginx 托管，并反向代理 `/api`、`/uploads` 到后端。
- 密钥一律通过环境变量注入，不要写死在镜像里。

（容器化与 `docker-compose` 配置可按需补充。）

## 📄 License

MIT

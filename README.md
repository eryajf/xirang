<div align="center">

<h1 align="center">XiRang</h1>

[![Auth](https://img.shields.io/badge/Auth-eryajf-ff69b4)](https://github.com/eryajf)
[![Go Version](https://img.shields.io/github/go-mod/go-version/eryajf/xirang)](https://github.com/eryajf/xirang)
[![Gin Version](https://img.shields.io/badge/Gin-1.6.3-brightgreen)](https://github.com/eryajf/xirang)
[![Gorm Version](https://img.shields.io/badge/Gorm-1.24.5-brightgreen)](https://github.com/eryajf/xirang)
[![GitHub Issues](https://img.shields.io/github/issues/eryajf/xirang.svg)](https://github.com/eryajf/xirang/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/eryajf/xirang)](https://github.com/eryajf/xirang/pulls)
[![GitHub Pull Requests](https://img.shields.io/github/stars/eryajf/xirang)](https://github.com/eryajf/xirang/stargazers)
[![HitCount](https://views.whatilearened.today/views/github/eryajf/xirang.svg)](https://github.com/eryajf/xirang)
[![GitHub license](https://img.shields.io/github/license/eryajf/xirang)](https://github.com/eryajf/xirang/blob/main/LICENSE)

<p> 🐉 简单好用，不缠不绕，直接上手的go-web框架 </p>

<img src="https://cdn.jsdelivr.net/gh/eryajf/tu@main/img/image_20240420_214408.gif" width="800"  height="3">
</div><br>

<p align="center">
  <a href="" rel="noopener">
 <img src="https://cdn.jsdelivr.net/gh/eryajf/tu/img/image_20220826_101156.png" alt="Project logo"></a>
</p>

> `xirang`，亦即息壤，大概在我读小学时，父亲从外地回来，并带回来几本读物，其中一本是迅哥儿小时候最爱读的《山海经》，也是在那时，我第一次认识这个词语。
>
> 维基百科对其解释是：息壤是中国古代传说中的一种神物，所谓“息壤”就是自己可以自动生长的土壤。用“息壤”修筑的堤坝，洪水长一米，堤坝也自动长一米。
>
> 我将项目命名为`xirang`，亦是希望在自己不间断地学习积累之中，项目也能够越发成长，并有其丰富的力量。

## 🥸 项目介绍

`xirang` 是一个非常简单的 `gin+gorm` 框架的基础架构，你只需要修改简单的代码，即可开始上手编写你的接口。

只需要根据情况修改配置`config.yml`，然后配置里边的数据库配置信息，即可开始开发。

数据库支持 MySQL 与 sqlite3(无 CGO 依赖)，如果你的系统仅为运维内部一个小系统，则推荐你使用 sqlite3。数据表会自动映射并创建。

## 👨‍💻 项目地址

| 分类 |                 GitHub                  |                    Gitee                     |
| :--: | :-------------------------------------: | :------------------------------------------: |
| 后端 |  https://github.com/eryajf/xirang.git   |  https://gitee.com/eryajf-world/xirang.git   |
| 前端 | https://github.com/eryajf/xirang-ui.git | https://gitee.com/eryajf-world/xirang-ui.git |

## 📖 目录结构

```
xirang
├── config----------------配置文件读取
├── controller------------控制层
├── logic-----------------逻辑层
├── middleware------------中间件
├── model-----------------对象定义
├── public----------------一些公共组件与工具
├── routers---------------路由
├── service---------------服务层
├── test------------------一些测试
├── config.yml------------配置文件
└── main.go---------------程序入口
```

## 👀 功能概览

|  ![登录页](https://cdn.jsdelivr.net/gh/eryajf/tu/img/image_20220830_234917.png)  | ![首页](https://cdn.jsdelivr.net/gh/eryajf/tu/img/image_20220830_233946.png)     |
| :------------------------------------------------------------------------------: | -------------------------------------------------------------------------------- |
| ![用户管理](https://cdn.jsdelivr.net/gh/eryajf/tu/img/image_20220830_234015.png) | ![分组管理](https://cdn.jsdelivr.net/gh/eryajf/tu/img/image_20220830_234043.png) |
| ![角色管理](https://cdn.jsdelivr.net/gh/eryajf/tu/img/image_20220830_234122.png) | ![菜单管理](https://cdn.jsdelivr.net/gh/eryajf/tu/img/image_20220830_234153.png) |
| ![接口管理](https://cdn.jsdelivr.net/gh/eryajf/tu/img/image_20220830_234218.png) | ![操作日志](https://cdn.jsdelivr.net/gh/eryajf/tu/img/image_20220830_234245.png) |

## 🚀 快速开始

xirang 项目的基础依赖项只有 MySQL，本地准备好这个服务之后，就可以启动项目，进行调试。

`注意：xirang还支持sqlite3(无CGO依赖)，默认配置文件即指向sqlite3，你可以不准备任何依赖，直接运行项目。`

### 拉取代码

```sh
# 后端代码
$ git clone https://github.com/eryajf/xirang.git

# 前端代码
$ git clone https://github.com/eryajf/xirang-ui.git
```

### 更改配置

```sh
# 修改后端配置
$ cd xirang
# 文件路径 config.yml, 根据自己本地的情况，调整数据库等配置信息。
$ vim config.yml
```

### 启动服务

```sh
# 启动后端
$ cd xirang
$ go mod tidy
$ make run

# 启动前端
$ cd xirang-ui
$ git config --global url."https://".insteadOf git://
$ npm install --registry=http://registry.npmmirror.com
$ yarn dev
```

本地访问：http://localhost:8090，用户名/密码：admin/123456

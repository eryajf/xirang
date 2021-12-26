<h1 align="center">XiRang</h1>

<div align="center">
简单好用，不缠不绕，直接上手的go-web框架
<p align="center">
<img src="https://img.shields.io/github/go-mod/go-version/eryajf/xirang" alt="Go version"/>
<img src="https://img.shields.io/badge/Gin-1.6.3-brightgreen" alt="Gin version"/>
<img src="https://img.shields.io/badge/Gorm-1.22.3-brightgreen" alt="Gorm version"/>
<img src="https://img.shields.io/github/license/eryajf/xirang" alt="License"/>
</p>
</div>


>`xirang`，亦即息壤，大概在我读小学时，父亲从外地回来，并带回来几本读物，其中一本是迅哥儿小时候最爱读的《山海经》，也是在那时，我第一次认识这个词语。
>
>维基百科对其解释是：息壤是中国古代传说中的一种神物，所谓“息壤”就是自己可以自动生长的土壤。用“息壤”修筑的堤坝，洪水长一米，堤坝也自动长一米。
>
>我将项目命名为`xirang`，亦是希望在自己不间断地学习积累之中，项目也能够越发成长，并有其丰富的力量。

## 项目介绍

[gitee地址](https://gitee.com/eryajf/xirang): https://gitee.com/eryajf/xirang

`xirang` 是一个非常简单的 `gin+gorm` 框架的基础架构，你只需要修改简单的代码，即可开发出你想要的接口。

只需要将`.env.example`改为`.env`，然后配置里边的数据库配置信息，即可开始开发。

数据表会自动创建，也可以通过docs下的sql自行创建。

## 重构更新

将gorm更新到v2版本，优化了项目初始化逻辑，优化了整体的内容，废除掉一些不必要的内容。

## 目录结构

```
xirang
├── controller----------------控制层
├── docs----------------------提供的原料信息
├── main.go-------------------入口文件
├── middleware----------------中间件
├── model---------------------对象定义
├── public--------------------一些公共组件
├── router--------------------路由
└── service-------------------服务层
```
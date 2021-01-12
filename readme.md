# 关于CORM
CORM是一个简单的golang语言编写的ORM框架，借助golang的`database/sql`包，可以很高效的开发一款支持多个数据库的ORM工具。本项目根据[!7天用Go从零实现ORM框架](https://geektutu.com/post/geeorm.html)编写的。
## 什么是ORM框架
对象关系映射（Object Relational Mapping，简称ORM）是通过使用描述对象和数据库之间映射的元数据，将面向对象语言程序中的对象自动持久化到关系数据库中。**简单说就是我们可以直接传入结构体参数来操作数据库，而不是通过传入sql语句**。举个栗子：
```golang
type User struct {
    Name string
    Age  int
}
// 把这个User存入数据库，传统的sql会这么写
// CREATE TABLE `User` (`Name` text, `Age` integer);
// INSERT INTO `User` (`Name`, `Age`) VALUES ("Tom", 18);
// SELECT * FROM `User`;

// 用ORM框架，可以这么写
orm.CreateTable(&User{})
orm.Save(&User{"Tom", 18})
var users []User
orm.Find(&users)
```
ORM框架相当于对象和数据库中间的一个桥梁，省去了繁琐的SQL语言
## 关于CORM
实现部分ORM功能
- 表的创建、删除、迁移
- 记录的增删查改，查询条件的链式操作
- 记录的增删查改，查询条件的链式操作
- 钩子(在创建/更新/删除/查找之前或之后)
- 钩子(在创建/更新/删除/查找之前或之后)
## 目录结构
```
corm
├── cc.db // sqlite3 数据文件
├── clause // 子句实现
│   ├── clause.go 
│   ├── clause_test.go
│   └── generator.go 
├── corm.go // corm 入口
├── corm_test.go
├── dialect // 不同数据库差异兼容database/sql
│   ├── dialect.go
│   ├── sqlite3.go
│   └── sqlite3_test.go
├── go.mod
├── go.sum
├── log // 基于logger的简单log库
│   ├── log.go
│   └── log_test.go
├── main // main函数测试
│   └── main.go
├── readme.md
├── schema
│   ├── schema.go // 对象到表的映射实现
│   └── schema_test.go
├── session // 核心结构，与数据库的交互
│   ├── hooks.go // 实现钩子
│   ├── hooks_test.go
│   ├── raw.go // 数据库操作的封装
│   ├── raw_test.go
│   ├── record.go // 增删改查的实现
│   ├── record_test.go
│   ├── table.go // 对象-表的操作
│   ├── table_test.go
│   └── transaction.go // 事务实现
```
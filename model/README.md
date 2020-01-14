#### 数据库

采用 MySQL 存储好友关系、群、群成员等 IM 服务的核心数据

数据模型 protobuf 中定义，主要是为了减少数据类型的转换

#### 模块

* db 直接通过 SQL 语句操作数据库
* cache 访问磁盘的数据的性能不如访问内存数据库效率高，因此热点数据会被保存在缓存中，提高数据的访问效率。
* simple 不依赖第三方的组件（即数据库），但程序退出后数据会丢失，仅用于测试
* sql 数据表结构


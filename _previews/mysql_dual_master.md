---
layout: post
title:  "mysq dual master"
date:   2022-02-13 12:00:00 +0800
tags:   tech
---

* category
{:toc}


mysql dual master

双主配置并非master官方支持，但实践中简单场景也能用。

### 重置用户密码和权限
```sh
$ sudo grep 'temporary password' /var/log/mysqld.log
2022-01-28T04:30:28.815290Z 1 [Note] A temporary password is generated for root@localhost: passwd123456

$ mysql -uroot -p -h127.0.0.1
```

```sql
ALTER USER 'root'@'localhost' IDENTIFIED BY 'passwd123456';

create database db_quick；
CREATE USER 'repl'@'%' IDENTIFIED BY 'pwd123456*';
GRANT REPLICATION SLAVE ON *.* TO 'repl'@'%';

CREATE USER 'tmpuser'@'%' IDENTIFIED BY 'pwd123456';
ALTER USER 'tmpuser'@'%' IDENTIFIED BY 'pwd123456';
GRANT ALL PRIVILEGES ON *.* TO 'tmpuser'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;
```

### 主从两库设置不同的自增序列

```sh
host1$ tail /etc/my.cnf -n 5 # 主库奇数：
log-bin=/data/mysql/binlog
server-id=1
auto_increment_offset=1
auto_increment_increment=2
log_slave_updates=ON
host2$ tail /etc/my.cnf -n 5 # 从库偶数：
log-bin=/data/mysql/binlog
server-id=2
auto_increment_offset=2
auto_increment_increment=2
log_slave_updates=ON
```

```sql
SHOW VARIABLES LIKE 'auto_inc%'; 
```

### 初始数据
```sql
CREATE TABLE `tab_test_repl` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `app_id` varchar(64) NOT NULL DEFAULT '' COMMENT '应用标识' ,
    `key_id` varchar(64) NOT NULL DEFAULT '' COMMENT '密钥标识',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间',

        KEY (`app_id`),
        KEY (`key_id`),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='密钥表';

INSERT INTO `tab_test_repl` (`app_id`, `key_id`) VALUES ('app1', 'key1');
INSERT INTO `tab_test_repl` (`app_id`, `key_id`) VALUES ('app2', 'key2');

SELECT * FROM `tab_test_repl`;
```

在 master 备份快照，在 slave 恢复快照
```sh
$ mysqldump --master-data -h 192.168.0.1 -u root -p "db_quick" > /tmp/repl.sql
```
> mysqldump --master-data 操作生成的 sql 文件其实已经包含了CHANGE MASTER 语句，而且指定了正确的 log position

```sh
$ mysql -h 192.168.0.2 -u root -p pwd123456 db_quick < /tmp/repl.sql
```

### 分别在两个sql实例中将对方设置为 master 
```sql
STOP SLAVE;
CHANGE MASTER TO MASTER_HOST='192.168.0.1',MASTER_USER='tmpuser',MASTER_PASSWORD='pwd123456',MASTER_LOG_FILE='binlog.000004', MASTER_LOG_POS=499;
START SLAVE;
SHOW SLAVE STATUS\G
```

同步 bin log position
在 master host
```sql
SHOW MASTER STATUS;
```
slave host
```sql
flush tables with read lock;
select master_pos_wait('binlog.000003',154);
unlock tables;
```

参考：
1. 简朝阳. MySQL性能调优与架构设计 (博文视点LAMP精品书廊) (Kindle 位置 4088-4089). 电子工业出版社. Kindle 版本. 


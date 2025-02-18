

create database if not exists blc_gene;



use blc_gene;

show tables ;



# ================================  数据库表结构  ================================

# files-文件表
drop table if exists files;
create table if not exists files (
                                     file_name varchar(32) comment '文件名',
                                     file_hash varchar(64) primary key comment '文件哈希,唯一索引',
                                     file_shares int comment '文件nft份额',
                                     data mediumtext comment '文件数据,最大16MB',
                                     race int comment '种族',
                                     age int comment '年龄',
                                     blood_type int comment '血型',
                                     gender boolean comment '性别',
                                     height float comment '身高',
                                     weight float comment '体重',
                                     smk_stat int comment '烟龄',
                                     alc_stat int comment '饮酒频率',
                                     other varchar(128) comment '其他特征',
                                     description varchar(1024) comment '描述',
                                     create_time datetime comment '创建时间',
                                     update_time datetime comment '更新时间'
) engine=innodb default charset=utf8 comment='文件表';


# permission/用户对哪些基因有访问权限（就是用户买了那些数据，一对多的关系）
create table  if not exists permission(
                                          id int comment '用户ID',
                                          file_hash int comment '文件哈希',
                                          create_time datetime comment '创建时间'
) engine=innodb default charset=utf8 comment='用户对哪些数据的权限';

# users/用户列表
drop table if exists users;
create table if not exists users(
                                    id int primary key auto_increment comment '用户ID，自增、唯一索引',
                                    address varchar(32) comment '用户账号地址',
                                    name varchar(32) comment '用户名',
                                    email varchar(32) comment '邮箱',
                                    picture varchar(1024) comment '头像地址',
                                    create_time datetime comment '注册时间'
) engine=innodb default charset=utf8 comment='用户表';



describe files;

describe permission;

describe users;

show tables ;

select * from files;


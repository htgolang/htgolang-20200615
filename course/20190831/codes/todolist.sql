create database todolist2 default charset utf8mb4;


use database todolist2;

create table task (
    id int,
    name varchar(64),
    `desc` varchar(1024),
    status int,
    create_time datetime,
    complete_time datetime
) engine=innodb default charset utf8mb4;


整数: int
浮点数：float, numeric(精度, 小数点位数)
布尔: boolean
字符串: varchar/text
时间: datetime, date, time


create table user3 (
    id int,
    name varchar(64) not null default "",
    password varchar(1024),
    `desc` text not null default "",
    sex boolean not null default false,
    height float not null default 0.0,
    birthday date,
    create_time datetime
) engine=innodb default charset utf8mb4;


key: 主键 primary key
            unique

int primary key, 自动增长 auto_increment


create table user4 (
    id int primary key,
    name varchar(64) not null default "",
    password varchar(1024),
    `desc` text not null default "",
    sex boolean not null default false,
    height float not null default 0.0,
    birthday date,
    create_time datetime
) engine=innodb default charset utf8mb4;


create table user5 (
    id int primary key,
    name varchar(64) unique not null default "",
    password varchar(1024),
    `desc` text not null default "",
    sex boolean not null default false,
    height float not null default 0.0,
    birthday date,
    create_time datetime
) engine=innodb default charset utf8mb4;


create table user6 (
    id int primary key auto_increment,
    name varchar(64) unique not null default "",
    password varchar(1024),
    `desc` text not null default "",
    sex boolean not null default false,
    height float not null default 0.0,
    birthday date,
    create_time datetime
) engine=innodb default charset utf8mb4;



create table test01 (
    col1 varchar(10),
    col2 varchar(10) not null default ""
) engine=innodb default charset utf8mb4;

结构体
    {}
    {
        attr1,
        attr2
        attr3,
        ...
        attrn,
    }
    {
        attr1=xxx,
        attr2=xxx,
        attr4=xxx,
        attr3=xx
    }

插入数据
insert into `tablename` values('col1', ..., 'coln');
insert into `tablename`(col1, col2, ..., coln) values('col1', ..., 'coln')


insert into test01 values('1', '2');
insert into test01(col1) values('1');
insert into test01(col2) values('2');
insert into test01(col2, col1) values('20', '10');


create table test02(
    id int primary key,
    name varchar(32) unique not null default ''
) engine=innodb default charset utf8mb4;


insert into test02 values(1, "kk");


insert into test02(id) values(2);


create table test03(
    id int primary key auto_increment,
    name varchar(32) unique not null default ''
) engine=innodb default charset utf8mb4;


设计两张表 用户/任务
用户表
create table todolist_user (
    id int primary key auto_increment,
    name varchar(64) unique not null default '',
    password varchar(1024) not null default '',
    sex boolean not null default true,
    birthday date not null,
    tel varchar(32) not null default '',
    addr varchar(128) not null default '',
    `desc` text not null default '',
    create_time datetime not null
    ) engine=innodb default charset utf8mb4;

任务表
type Task struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Progress int    `json:"progress"`
	User     string `json:"user"`
	Desc     string `json:"desc"`
	Status   string `json:"staus"`
}

create table todolist_task(
    id int primary key auto_increment,
    name varchar(64) not null default '',
    progress float not null default 0,
    user varchar(64) not null default '',
    `desc` text not null default '',
    status int not null default 0,
    create_time datetime not null,
    complete_time datetime
)engine=innodb default charset utf8mb4;


insert into todolist_user(name, password, sex, birthday, create_time) values('kk', md5('kk'), 1, '1988-11-11', now());
insert into todolist_user(name, password, sex, birthday, create_time) values
('xiangge11', md5('kk1'), 1, '1988-11-12', now()),
('xiangge22', md5('kk2'), 1, '1988-11-13', now()),
('xiangge33', md5('kk3'), 1, '1988-11-15', now()),
('xiaolin44', md5('kk4'), 0, '1988-11-11', now()),
('xiaolin11', md5('kk1'), 0, '1988-11-12', now()),
('xiaolin22', md5('kk2'), 0, '1988-11-13', now()),
('xiaolin33', md5('kk3'), 0, '1988-11-15', now()),
('xiangge44', md5('kk4'),1, '1988-11-11', now());



where条件

colname condtion value;

condtion:
    关系运算 等于 =
            不等于 !=
            >
            <
            >=
            <=
    in: colname in (v1, v2)
    id 2 4
    between and: id between 2 and 4
    like name 以 kk结尾的所有数据    %kk
         name 包含 kk的所有数据      %kk%
         name 以 kk开头的所有数据    kk%

布尔运算
    与 and
    或 or
    非 not



更新

update tablename
    set colname = value1, colname2 = value2
    where

update todolist_user
set `desc` ='test', sex = 0
where id > 30


delete from tablename
where id = 9
;


create table accesslog(
    id int primary key auto_increment,
    ip varchar(256) not null default '',
    logtime datetime not null,
    method varchar(8) not null default '',
    url varchar(4096) not null default '',
    status_code int not null default 0,
    bytes int not null default 0
) engine=innodb default charset utf8mb4;

insert into accesslog(ip, logtime, method, url, status_code, bytes) Values
('1.1.1.1', '2018-10-11', 'GET', '/', 200, 250),
('1.1.1.2', '2018-10-11', 'POST', '/login/', 200, 250),
('1.1.1.3', '2018-10-12', 'GET', '/logout/', 200, 250),
('1.1.1.4', '2018-10-13', 'POST', '/login/', 302, 250),
('2.1.1.1', '2018-10-11', 'GET', '/logout/', 200, 250),
('2.2.1.1', '2018-10-11', 'GET', '/error/', 400, 250),
('2.31.1.1', '2018-10-11', 'GET', '/', 500, 250),
('2.3.3.3', '2018-10-11', 'GET', '/', 200, 250),
('1.1.1.1', '2018-10-11', 'GET', '/', 200, 250),
('1.1.1.2', '2018-10-12', 'GET', '/login/', 200, 250),
('1.1.1.3', '2018-10-13', 'GET', '/logout/', 200, 250),
('1.1.1.4', '2018-10-14', 'POST', '/login/', 200, 250),
('2.1.1.1', '2018-10-11', 'GET', '/logout/', 200, 250),
('2.2.1.1', '2019-10-11', 'GET', '/error/', 400, 250),
('2.31.1.1', '2019-10-11', 'POST', '/', 500, 250),
('2.3.3.3', '2018-10-11', 'GET', '/', 200, 250),
('1.1.1.1', '2018-10-11', 'GET', '/', 200, 250),
('1.1.1.2', '2018-10-11', 'POST', '/login/', 200, 250),
('1.1.1.3', '2018-10-12', 'GET', '/logout/', 200, 250),
('1.1.1.4', '2018-10-13', 'POST', '/login/', 200, 250),
('2.1.1.1', '2018-10-11', 'GET', '/logout/', 200, 250),
('2.2.1.1', '2018-10-11', 'GET', '/error/', 400, 250),
('2.31.1.1', '2018-10-11', 'GET', '/', 500, 250),
('2.3.3.3', '2018-10-11', 'GET', '/', 302, 250),
('1.1.1.1', '2018-10-11', 'GET', '/', 200, 250),
('1.2.1.2', '2018-10-12', 'GET', '/login/', 200, 250),
('1.1.1.3', '2018-10-13', 'GET', '/logout/', 200, 250),
('1.1.1.4', '2018-10-14', 'POST', '/login/', 200, 250),
('2.1.1.1', '2018-10-11', 'GET', '/logout/', 302, 250),
('2.2.1.1', '2019-10-11', 'GET', '/error/', 400, 250),
('2.31.1.1', '2019-10-11', 'POST', '/', 500, 250),
('2.3.3.3', '2018-10-11', 'GET', '/', 200, 250);


[{'ip' : '', status_code:''}, ]
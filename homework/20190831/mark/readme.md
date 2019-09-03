

```
create table todolist_user1 (
    id int primary key auto_increment,
    name varchar(64) unique not null default '',
    password varchar(1024) not null default '',
    sex varchar(32) not null default '',
    birthday date not null,
    tel varchar(32) not null default '',
    addr varchar(128) not null default '',
    `desc` text not null default '',
    create_time datetime not null,
    change_time datetime not null
    ) engine=innodb default charset utf8mb4;
```



```
insert into todolist_user1(name,password,sex,birthday,create_time) values('mark',md5('password'),'男','2000-06-06',now());
```


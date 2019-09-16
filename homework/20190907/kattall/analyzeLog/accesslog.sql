create table accesslog(
    id int primary key auto_increment,
    ip varchar(256) not null default '',
    accesstime datetime not null,
    method varchar(8) not null default '',
    url varchar(2048) not null default '',
		protocol varchar(32) not null default '',
    statuscode int not null default 0,
    length int not null default 0) engine=innodb default charset utf8mb4;
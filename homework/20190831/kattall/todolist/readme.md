用户管理：
    增删查改
    
任务管理：
    增删查改
    
任务状态显示：
statuFunc := template.FuncMap{
    "statusType": func(t int) string {
        if t == 0 {
            return "新建"
        } else if t == 1 {
            return "正在进行"
        } else if t == 2 {
            return "停止"
        } else if t == 3 {
            return "完成"
        } else {
            return "未知状态"
        }
    },
}
tpl := template.Must(template.New("task.html").Funcs(statuFunc).ParseFiles("views/task/task.html"))
tpl.Execute(w, modules.GetTasks())


create table user (
    id int primary key auto_increment,
    name varchar(64) unique not null default '',
    password varchar(128) not null default '',
    sex boolean not null default true,
    birthday date not null, 
    tel varchar(32) not null default '',
    addr varchar(128) not null default '',
    `desc` varchar(1024) not null default '',
    create_time datetime not null) engine=innodb default charset utf8mb4;
    
create table task (
    id int primary key auto_increment,
    name varchar(64) unique not null default '',
    progress float not null default 0,
    user varchar(64) not null default '',
    `desc` varchar(64) not null default '',
    status int not null default 0,
    create_time datetime not null,
    complate_time datetime) engine=innodb default charset utf8mb4;
    
create table accesslog(
    id int primary key auto_increment,
    ip varchar(256) not null default '',
    logtime datetime not null,
    method varchar(8) not null default '',
    url varchar(2048) not null default '',
    status_code int not null default 0,
    bytes int not null default 0) engine=innodb default charset utf8mb4;
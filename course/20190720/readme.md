type TypeName struct{
    Field01 FieldType01
    Field02 FieldType02
    Field03 FieldType03
    ....
}



type User struct {
    id int
    name string
}


var me User

var point *User

point = &me




var name struct {
    Field01 FielType01
    ...
    FieldN FieldTypeN
}


name := struct{
    Field01 FielType01
    ...
    FieldN FieldTypeN
}{FieldVal0, ..., FieldNValN}


type TypeName struct{
    Field01 FieldType01
    Field02 FieldType02
    structVar01 StructName01
    ....
}


TypeName{FieldVal01, FieldVal01, FieldValN, StructName01{FieldVal01, FieldVal02, ...， FieldValN}, StructName02{FieldVal01, ..., FieldValN}, ..., StructNameN{FieldVal01, ..., FieldValN}}


type Addr struct {
    region string
    street string
    no string
}

type User struct {
    name string
    Addr
}


var me User

me.addr.region


me2 := me

me2.name = 'kk'
me



读
    a. 文件位置
    b. 打开文件
    c. 读取文件内容(1, 20)
    d. 关闭文件

写
    a. 创建文件
    b. 写入内容
    c. 关闭

追加


文件路径:
    绝对路径： 程序不管在什么位置运行，操作的文件都不会变化(从根路径/盘符开始书写的路径)
    相对路径: 与程序的位置有关(程序运行的位置)_ ./xxxx ../xxx   xxxx
            程序放置的位置  c:/htgolang/main.exe, e:/htgolang/main.exe
            程序运行的位置  e:/ c:/htgolang/main.exe



type User struct {
    name string
    Addr
}


User{}

var user User

user := User{}


练习1：

md5sum
-s "abc" 计算-s参数字母的md5
-f path 计算-f后面路径文件的md5值

1. 复习整理课程内容
2. 用户管理
    用户改为struct，并且针对属性使用不同数据类型
    ID int
    Name string
    birthday time.Time
    tel string
    addr string
    desc string

    注意需要进行类型转换
3. 在添加和修改数据时对用户名进行检查，用户名不能为空，且用户名必须唯一
    a. 添加kk
    b. 再添加kk 不能添加(用户名已经存在)

    c.有kk
    d.修改一个非kk用户名 => kk (用户名已经存在)
    e.修改kk => kk (修改成功)
4. 在查询时，让用户输入排序属性，按照属性值从小到大顺序显示
    id, name, birthday, tel,

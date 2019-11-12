AKDxc4

< &lt;
> &gt;
" &quot;
' &#39;
& &amp;


outercondition = innercondition and

(name like '%xxx%' or desc like '%xxx%') and create_user = 1


total, totalFilter

管理员
    total  select count(*) from user

    totalFilter select count(*) from user where q=xxxx

普通用户
     total  select count(*) from user where create_user = xxxx

    totalFilter select count(*) from user where q=xxxx and create_user  = xxx


    先设置用户条件 求count total
    再设置查询条件 求count totalFilter


第二页 10 10


1. 排序，后端需要维护0 => name关系
    => 能不能前端 直接告诉后端用哪个列进行排序(列名)

2. 页面上有不能排序的列
    => datatable能不能针对某一列指定进行排序

3. key => search[value] order[0][dir]/order[0][column]
    自定定义
4. 参数传递了一堆 columns[7]

    => 能不能自定义提交发起ajax的参数

1. 修改datatable的布局 dom: lftip 定义一个放置button按钮div
    l: 显示分页数据
    f: 搜索
    t: 表格内容
    i: 搜索数量
    p: 翻页

    < f>
    <"className" >
    <"#id">

    <"row" <"col-2" l><"col-2" i><"col-8"p>>

    <div class="row">
        <div class="col-2"> l</div>
        <div class="col-2"> i</div>
        <div class="col-2"> p</div>
    </div>
2. datatable绘制完成后 使用jquery再button div中插入咱们的按钮
    div html方法

    <"row" <"col-5" l><"col-6" f><"#buttons.col-1">>tip

        <div class="row">
        <div class="col-5"> l</div>
        <div class="col-6"> f</div>
        <div class="col-1" id="buttons"></div>
    </div>
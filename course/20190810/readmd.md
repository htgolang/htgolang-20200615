1. 发起请求获取http页面
2. 在页面中查找相应的标签
3. 在标签中获取属性/内容


query 使用
1. 创建document对象
https://www.zhipin.com/job_detail/?query=go&city=100010000&industry=&position=

2. 查找标签 Find
    标签名          Find("a")
    ID   baidu      Find("#baidu")
    Class           Find(".x")
    属性名+属性值    Find("[name=kk]")

    组合
        标签a 具有class x 的标签 Find("a.x")

    复杂的
        子孙标签        Find("table a.x")
        子标签          Find("td form input.txt")

3. 遍历所有的标签 Each
4. 获取标签属性的值，标签的内容
    Attr， Text/Html


找信息: 公司名称，职位名称，薪资


tag: div class: job-primary
    div.job-primary
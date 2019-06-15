# htgolang-201906 #
手撕Golang 201906期

## 不要删除别人的代码!!! ##

## 目录结构 ##
 
+ homework/20190615/ 第一次作业提交的目录
    - kk kk的目录
        + table.go 作业代码文件


## Git使用方法 ##

### 视频 ###

+ git使用 [百度下载](https://pan.baidu.com/s/13aXQVZ0VkHkUZxqRiqXm1Q "08.git使用") 提取码：re96
+ git windows使用 [百度下载](https://pan.baidu.com/s/1ezf4-fox_glUy3WmeUdigQ "09.git windows使用") 提取码：tct0


### 命令行添加代码 ###

```
第一次使用
git clone https://github.com/51reboot/actual_06_homework_mage.git

更新本地代码
git pull

查看代码状态
git status

后面添加代码，只需要下面三行即可
git add .
git commit -m "first commit"
git push -u origin master
```

> 用命令行操作，要添加ssh的公钥到github里，操作方法


```

创建SSH key的方法很简单，执行如下命令就可以：
ssh-keygen
生成的SSH key文件保存在中～/.ssh/id_rsa.pub

然后用文本编辑工具打开该文件，我用的是vim,所以命令是：
vim ~/.ssh/id_rsa.pub

接着拷贝.ssh/id_rsa.pub文件内的所以内容，将它粘帖到github帐号管理中的添加SSH key界面中。
打开github帐号管理中的添加SSH key界面的步骤如下：
1. 登录github
2. 点击右上方的Accounting settings图标
3. 选择 SSH key
4. 点击 Add SSH key
在出现的界面中填写SSH key的名称，填一个你自己喜欢的名称即可
然后将上面拷贝的~/.ssh/id_rsa.pub文件内容粘贴到key一栏，在点击“add key”按钮就可以了。
添加过程github会提示你输入一次你的github密码

添加完成后再次执行git clone就可以成功克隆github上的代码库了。

```

账号没有加到reboot群组里的，请随时联系我！


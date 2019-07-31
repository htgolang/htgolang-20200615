### 用户系统管理    
1. 运行main文件，即可运行该程序     
2. 密码保存在passwd-dir/passwd.gob文件中，并同步到passwdfile.csv文件，验证密码时从passwdfile.csv读取     
3. 用户信息保存在users_dir目录下，每更新一次用户数据即生成一个新的文件，仅保留最新的三个文件    
4. 日志文件保存在user.log文件中     
5. 获取用户信息，仅从users_dir目录下的最近生成的一个文件读取数据    
6. 当users_dir目录不存在时，将创建该目录，当users_dir目录下没有任何一个用户文件时，选择(显示/查询/修改/删除)操作时，都将直接返回主菜单    
7. 当passwd-dir/passwd.gob文件不存在时，将重新设置本系统密码    
8. 支持修改密码操作    
9. json 数据持久化 
  
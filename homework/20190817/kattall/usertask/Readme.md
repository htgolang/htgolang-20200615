// 用户（包含任务列表）
type User struct {
	Id       int       `json:"id"`
	UserName string    `json:"username"`
	Tasks    []Task    `json:"tasks"`
	Birthday time.Time `json:"birthday"`
	Addr     string    `json:"addr"`
	Desc     string    `json:"desc"`
	Password string    `json:"password"`
}

// 任务
type Task struct {
	Id       int    `"json"": "id"`
	Name     string `"json"": "name"`
	Progress int    `"json"": "progress"`
	Desc     string `"json"": "desc"`
	Status   string `"json"": "status"`
}


## 用户任务系统
*  用户登陆         已完成
    * 通关表单提交，后台通过PostFormValue("username")获取数据，把用户，密码和json文件得数据做比较。
*  添加用户         已完成     
    * 不能添加相同用户，前台会报400
*  删除用户         已完成     
    * get请求前端通过get请求获取用户id, 后台拿到id, 进行删除。
*  修改用户         已完成     
    * 前端通过get请求获取用户id, 后台拿到id. 取出id得数据, 把数据返回给模板，进行渲染。
    * post请求，后台通过PostFormValue("username")获取属性，后台修改。 重定向到用户列表页面。
*  修改用户密码     已完成
    

## 用户下可添加任务，修改任务，删除任务

* 添加任务  
    功能完成。
    在添加任务得时候，必须要需要登陆。 登陆后对任务进行添加，删除，修改等操作。
    在添加得时候，会把username传过去，可以通过username查询当前用得任务，进行添加，添加后，302重定向得时候，会通过get方法，所以这边可能页面不太对。 还没学session, 要不然应该可以解决。
* 修改任务
* 删除任务
    * 修改任务和删除任务还没完成
    原因： 登陆过后， 任务列表使用一个table展现出来。 删除或者删除只能传入一个id得标签，传不了用户名，或者用户id信息。暂时不知道如何区操作。
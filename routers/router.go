package routers

import (
	"fabcar/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	////建立路由过滤器,用于登录校验   参数一是过滤匹配支持正则    参数二过滤位置    参数三过滤操作（函数） 参数是context
	//beego.InsertFilter("/article/*", beego.BeforeRouter, filter)
	//// 登录
	beego.Router("/login", &controllers.UserController{}, "get:ShowLogin;post:HandleLogin")
	//// 注册
	//beego.Router("/register", &controllers.UserController{}, "get:ShowRegister;post:HandleRegister")
	//// 展示首页
	////beego.Router("/article/index", &controllers.ArticleController{}, "get:ShowIndex;post:SelectType")
	//beego.Router("/article/index", &controllers.ArticleController{}, "get:ShowIndex")
	//// 添加文章
	//beego.Router("/article/add", &controllers.ArticleController{}, "get:ShowAdd;post:HandleAdd")
	//// 文章详情
	////beego.Router("/article/content", &controllers.ArticleController{}, "get:ShowContent")
	//// 编辑文章
	//beego.Router("/article/update", &controllers.ArticleController{}, "get:ShowUpdate;post:HandleUpdate")
	//// 删除文章
	//beego.Router("/article/delete", &controllers.ArticleController{}, "get:ShowDelete;post:HandleDelete")
	//// 文章类型
	////beego.Router("/article/addType", &controllers.ArticleController{}, "get:ShowArtType;post:AddType")
	//// 退出
	beego.Router("/logout", &controllers.UserController{}, "get:LogOut")

	beego.Router("/", &controllers.MainController{})
	beego.InsertFilter("/article/*", beego.BeforeRouter, filter)
	beego.Router("/article/index", &controllers.ArticleController{}, "get:ShowIndex")
}

//  过滤函数
func filter(ctx *context.Context) {
	// 获取session中的用户名
	userName := ctx.Input.Session("username")
	if userName == nil {
		// 跳转到登录页
		ctx.Redirect(302, "/login")
		return
	}
}

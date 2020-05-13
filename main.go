package main

import (
	"github.com/kataras/iris"
	"code/irisCms/config"
	"code/irisCms/controller"
	"code/irisCms/datasource"
	"code/irisCms/service"
	//"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"time"
)

/**
 * 程序主入口
 */
func main() {
	app := newApp()

	// 应用App设置
	configation(app)

	// 路由设置
	mvcHandle(app)

	config := config.InitConfig()
	addr := ":" + config.Port
	app.Run(
		iris.Addr(addr), //在端口9000进行监听
		iris.WithoutServerError(iris.ErrServerClosed), //无服务错误提示
		iris.WithOptimizations,  //对json数据序列化更快的配置
	)
}

// 构建APP
func newApp() *iris.Application {
	app := iris.New()

	// 设置日志级别，开发阶段为 debug
	app.Logger().SetLevel("debug")

	// 注册静态资源
	// StaticWeb
	app.HandleDir("/static", "./static")
	app.HandleDir("/manage/static", "./static")

	// 注册视图文件
	app.RegisterView(iris.HTML("./static", ".html"))
	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})

	return app
}

/**
 * MVC 架构模式处理
 */
func mvcHandle(app *iris.Application) {

	//启用session
	sessManager := sessions.New(sessions.Config{
		Cookie:  "sessioncookie",
		Expires: 24 * time.Hour,
	})

	engine := datasource.NewMysqlEngine()

	//管理员模块功能
	adminService := service.NewAdminService(engine)
	admin := mvc.New(app.Party("/admin"))
	admin.Register(
		adminService,
		sessManager.Start,
	)
	admin.Handle(new(controller.AdminController))

	//用户功能模块
	userService := service.NewUserService(engine)
	user := mvc.New(app.Party("/v1/users"))
	user.Register(
		userService,
		sessManager.Start,
	)
	user.Handle(new(controller.UserController))

	//统计功能模块
	statisService := service.NewStatisService(engine)
	statis := mvc.New(app.Party("/statis/{model}/{date}/"))
	statis.Register(
		statisService,
		sessManager.Start,
	)
	statis.Handle(new(controller.StatisController))

	//订单模块
	orderService := service.NewOrderService(engine)
	order := mvc.New(app.Party("/bos/orders/"))
	order.Register(
		orderService,
		sessManager.Start,
	)
	order.Handle(new(controller.OrderController)) //控制器
}

// 项目设置
func configation(app *iris.Application) {
	// 配置字符编码
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))

	//错误配置
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"errmsg": iris.StatusNotFound,
			"msg": " not found ",
			"data": iris.Map{},
		})
	})

	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"errmsg": iris.StatusInternalServerError,
			"msg": " internal error ",
			"data": iris.Map{},
		})
	})
}
package main

import (
	"code/irisCms/config"
	"github.com/kataras/iris"
)

/**
 * 程序主入口
 */
func main() {
	app := newApp()

	// 应用App设置
	configation(app)

	// 路由设置

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
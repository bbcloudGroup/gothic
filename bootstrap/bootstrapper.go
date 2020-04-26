package bootstrap

import (
	//"github.com/gorilla/securecookie"

	"github.com/bbcloudGroup/gothic/config"
	"github.com/kataras/iris/v12"
	"unsafe"

	//"github.com/kataras/iris/v12/middleware/logger"
	//"github.com/kataras/iris/v12/middleware/recover"
	//"github.com/kataras/iris/v12/sessions"
	//"github.com/kataras/iris/v12/websocket"
)


var app *Bootstrapper

func App() Bootstrapper {
	return *app
}

type Bootstrapper struct {
	*iris.Application
	configurators [] iris.Configurator
	//AppName      string
	//AppOwner     string
	//AppSpawnDate time.Time
	//Sessions *sessions.Sessions
}

// New returns a new Bootstrapper.
func New(cfgs ...iris.Configurator) *Bootstrapper {

	b := &Bootstrapper{
		//AppName:      appName,
		//AppOwner:     appOwner,
		//AppSpawnDate: time.Now(),
		Application:  iris.New(),
	}

	b.SetupUp(cfgs...)

	app = b

	return b
}


func (b *Bootstrapper) SetupUp(cfgs ...iris.Configurator) {
	for _, cfg := range cfgs {
		b.configurators = append(b.configurators, cfg)
	}
}



func (b *Bootstrapper) Run(withOrWithout ...iris.Configurator) error {

	for _, cfg := range withOrWithout {
		b.configurators = append(b.configurators, cfg)
	}

	b.Application.Configure(b.configurators...)

	app := config.GetApp(b)

	return b.Application.Run(iris.Addr(":" + app.AppPort))
}


func Application(app *iris.Application) *Bootstrapper {
	return (*Bootstrapper)(unsafe.Pointer(app))
}








//// SetupViews loads the templates.
//func (b *Bootstrapper) SetupViews(viewsDir string) {
//	b.RegisterView(iris.HTML(viewsDir, ".html").Layout("shared/layout.html"))
//}
//
//// SetupSessions initializes the sessions, optionally.
//func (b *Bootstrapper) SetupSessions(expires time.Duration, cookieHashKey, cookieBlockKey []byte) {
//	b.Sessions = sessions.New(sessions.Config{
//		Cookie:   "SECRET_SESS_COOKIE_" + b.AppName,
//		Expires:  expires,
//		Encoding: securecookie.New(cookieHashKey, cookieBlockKey),
//	})
//}
//
//// SetupWebsockets prepares the websocket server.
//func (b *Bootstrapper) SetupWebsockets(endpoint string, handler websocket.ConnHandler) {
//	ws := websocket.New(websocket.DefaultGorillaUpgrader, handler)
//
//	b.Get(endpoint, websocket.Handler(ws))
//}
//
//// SetupErrorHandlers prepares the http error handlers
//// `(context.StatusCodeNotSuccessful`,  which defaults to < 200 || >= 400 but you can change it).
//func (b *Bootstrapper) SetupErrorHandlers() {
//	b.OnAnyErrorCode(func(ctx iris.Context) {
//		err := iris.Map{
//			"app":     b.AppName,
//			"status":  ctx.GetStatusCode(),
//			"message": ctx.Values().GetString("message"),
//		}
//
//		if jsonOutput := ctx.URLParamExists("json"); jsonOutput {
//			ctx.JSON(err)
//			return
//		}
//
//		ctx.ViewData("Err", err)
//		ctx.ViewData("Title", "Error")
//		ctx.View("shared/error.html")
//	})

package logerr

import (
	"fmt"
	"github.com/bbcloudGroup/gothic/config"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/logger"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func logHandler(log config.Log, cfgs ...logger.Config) context.Handler {

	handler := logger.New()
	if len(cfgs) > 0 {
		handler = logger.New(cfgs[0])
	} else {
		handler = logger.New(logger.Config{
			Status: log.Status,
			IP: log.IP,
			Method: log.Method,
			Path: log.Path,
			Query: log.Query,
			Columns: log.Columns,
			MessageContextKeys: log.MessageHeaderKeys,
			MessageHeaderKeys: log.MessageHeaderKeys,
		})
	}

	return handler
}


func getRequestLogs(ctx context.Context) string {
	var status, ip, method, path string
	status = strconv.Itoa(ctx.GetStatusCode())
	path = ctx.Path()
	method = ctx.Method()
	ip = ctx.RemoteAddr()
	// the date should be logged by iris' Logger, so we skip them
	return fmt.Sprintf("%v %s %s %s", status, path, method, ip)
}


func Recover() context.Handler {
	return recoverHandler()
}

func recoverHandler() context.Handler {
	return func(ctx context.Context) {
		defer func() {
			if err := recover(); err != nil {
				if ctx.IsStopped() {
					return
				}

				var stacktrace string
				for i := 1; ; i++ {
					_, f, l, got := runtime.Caller(i)
					if !got {
						break
					}

					stacktrace += fmt.Sprintf("%s:%d\n", f, l)
				}

				// when stack finishes
				requestLog := getRequestLogs(ctx)
				logMessage := fmt.Sprintf("Recovered from a route's Handler('%s')\n", ctx.HandlerName())
				logMessage += fmt.Sprintf("At Request: %s\n", requestLog)
				logMessage += fmt.Sprintf("Trace: %s\n", err)
				logMessage += fmt.Sprintf("\n%s", stacktrace)
				ctx.Application().Logger().Error(logMessage)

				problem := NewProblem(ctx).Status(iris.StatusInternalServerError)
				if config.GetApp(ctx.Application()).AppEnv != config.Production {
					_ = problem.Title(fmt.Sprintf("%s", err)).
						Key("stacktrace", strings.Split(stacktrace, "\n")).
						Detail(requestLog)
				} else {
					_ = problem.Title("服务异常:" + time.Now().Format("2006-01-02 15:04"))
				}
				_, _ = ctx.Problem(problem)

				ctx.StopExecution()
			}
		}()

		ctx.Next()
	}
}



func NewProblem(ctx context.Context) iris.Problem {
	problem := iris.NewProblem().
		Type(ctx.FullRequestURI()).
		Status(ctx.GetStatusCode())

	return problem
}
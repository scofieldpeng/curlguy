package curlguy

import (
	"flag"
	"github.com/labstack/echo"
	"github.com/scofieldpeng/curlguy/api/request"
	"github.com/labstack/echo/engine/standard"
)

var(
	appPort *int = flag.Int("port","8899","app listen port")
	app *echo.Echo = echo.New()
)

func init() {
	flag.Parse()
}

func main() {
	app.GET(`/api/v1/request`,request.Request) // api request
	app.Run(standard.New(":" + *appPort))
}
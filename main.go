package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/urfave/cli"
	"os"
)

type H map[string]interface{}

func mainHandler(c echo.Context) error {
	ip := c.RealIP()
	lang := c.Request().Header.Get("Accept-Language")
	agent := c.Request().Header.Get("User-Agent")
	return c.JSON(200, H{"ipaddress": ip, "language": lang, "software": agent})
}

func start(c *cli.Context) error {
	port := c.Int("port")
	e := echo.New()
	e.GET("/", mainHandler)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
	return nil
}

func main() {
	app := cli.NewApp()
	app.Author = "Alain Gilbert"
	app.Email = "alain.gilbert.15@gmail.com"
	app.Name = "Request Header Parser Microservice"
	app.Usage = "Request Header Parser Microservice"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:   "port",
			Value:  3001,
			Usage:  "Webserver port",
			EnvVar: "PORT",
		},
	}
	app.Action = start
	app.Run(os.Args)
}

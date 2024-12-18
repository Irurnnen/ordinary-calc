//go:build debug
// +build debug

package main

import (
	_ "github.com/Irurnnen/ordinary-calc/docs"
	"github.com/Irurnnen/ordinary-calc/internal/application"
)

// @title		Ordinary Calc
// @version		0.0.1
// @description	This is ordinary calculator with http server).

// @license.name	MIT
// @license.url		https://mit-license.org/

// @host		127.0.0.1:8080
// @BasePath	/api/v1

func main() {
	app := application.NewDebug()
	app.Run()
}

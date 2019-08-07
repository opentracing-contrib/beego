package main

import (
	"github.com/astaxie/beego"
	apmbeego "github.com/opentracing-contrib/beego"
	"github.com/opentracing-contrib/beego/examples/tracer"
	"github.com/opentracing/opentracing-go"
)

const (
	DefaultComponentName = "beego-demo"
)

type helloController struct{ beego.Controller }

func (this *helloController) Hello() {
	span, _ := opentracing.StartSpanFromContext(this.Ctx.Request.Context(), "helloController.Hello")
	defer span.Finish()
	this.Ctx.WriteString("hello world")
}

func main() {

	// 1. init tracer
	tracer, closer := tracer.Init(DefaultComponentName)
	if closer != nil {
		defer closer.Close()
	}

	// 2. ste the global tracer
	if tracer != nil {
		opentracing.SetGlobalTracer(tracer)
	}

	beego.Router("/hello", &helloController{}, "get:Hello")

	// 3. use the middleware
	beego.RunWithMiddleWares("localhost:8080", apmbeego.Middleware(DefaultComponentName))

}

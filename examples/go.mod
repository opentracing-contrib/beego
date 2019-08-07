module github.com/opentracing-contrib/beego/examples

go 1.12

replace github.com/opentracing-contrib/beego => ../

require (
	github.com/astaxie/beego v1.12.0
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/opentracing-contrib/beego v0.0.0-00010101000000-000000000000
	github.com/opentracing/opentracing-go v1.1.0
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/uber-go/atomic v1.4.0 // indirect
	github.com/uber/jaeger-client-go v2.16.0+incompatible
	github.com/uber/jaeger-lib v2.0.0+incompatible // indirect
	go.uber.org/atomic v1.4.0 // indirect
)

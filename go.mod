module github.com/temp-go-dev/sample-swagger

go 1.12

require (
	github.com/gin-contrib/pprof v1.2.0
	github.com/gin-gonic/gin v1.4.0
	github.com/go-openapi/errors v0.19.2
	github.com/go-openapi/runtime v0.19.2
	github.com/go-openapi/strfmt v0.19.0
	github.com/go-openapi/swag v0.19.2
	github.com/go-openapi/validate v0.19.2
	github.com/mikkeloscar/gin-swagger v0.0.0-20190528202043-47b88fd7a7d1
	github.com/opentracing/opentracing-go v1.1.0
	github.com/sirupsen/logrus v1.4.2
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
)

//replaceしないといけないらしい？ https://github.com/gin-gonic/gin/issues/1673
replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43

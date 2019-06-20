# sample-swagger-test

[mikkeloscar/gin-swagger](https://github.com/mikkeloscar/gin-swagger)を利用した自動生成。  
OpenAPIではないため使えない。  
同じ機構を使えば実現できそう  

* 自動生成  
`gin-swagger -A my-api -f swagger.yaml`  
modelsとrestapiディレクトリが生成される。

* 実行  
`go run main.go`

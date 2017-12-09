# golang-reverse-proxy-dev

前端前后分离后, 搭建开发环境获取测试环境的数据

把测试环境的接口代理给到本地开发环境, 把本地开发文件也指定URL 某PATH地址.

例子:

访问除了'/spa/'开头的地址(如http://0.0.0.0:8888/login)都转发到测试服务器处理, 而'/spa/'下面地址(如 http://0.0.0.0:8888/spa/test.html )则返回本地文件(spa目录下)

```go run src/main.go```


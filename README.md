## 关于

> 是基于 [Gin](https://github.com/gin-gonic/gin) 进行模块化设计的 API工具，引入Gin最少化的配置。
>
> Gin 是一个用 Go (Golang) 编写的 Web 框架，小巧、性能优越。
>
> 主要为了 学习和测试用。
>
> 示例:
>
> 1、获取ip的地理位置信息
>
> http://xxxx.xxxx.xxxx/tools/get_geoip
>
> 无参数是获取请求端IP的位置信息



## 使用&安装

1、git clone 地址



2、需要下载geoip数据库

https://dev.maxmind.com/geoip/geolite2-free-geolocation-data

修改位置在代码： ***geoip2.Open***



3、go run main.go



4、交叉编译到服务器docker

4.1、Window下编译可执行程序

> GOARCH：编译目标平台的硬件体系架构（amd64, 386, arm, ppc64等）。
>
> GOOS：编译目标平台上的操作系统（darwin, freebsd, linux, windows）。
>
> CGO_ENABLED：代表是否开启CGO，1表示开启，0表示禁用。由于CGO不能支持交叉编译，所以需要禁用。

```c
#Mac
SET CGO_ENABLED=0 
SET GOOS=darwin 
SET GOARCH=amd64 
go build main.go
#Linux
SET CGO_ENABLED=0 
SET GOOS=linux 
SET GOARCH=amd64 
go build main.go
```

4.2、生成docker的相关命令

> 构建最小镜像：由于官方的golang镜像太大了，720多M，所以为了最小化我们的应用，我们采取了alpine镜像，大小连4MB都不到。

```dockerfile
#编写Dockerfile
DockerFile如下:

FROM alpine
MAINTAINER  Aze
WORKDIR /go/src/
COPY . .
EXPOSE 8777
ENTRYPOINT ["./app/main"]

#构建镜像
docker build -t goenv:v1 .

#运行程序
docker run --name golanggeoip -v /mydata/data/geoip:/go/src -p 8608:9090 -it --privileged -u root  --entrypoint /bin/sh goenv:v1
#进入容器
./main &
```

4.3、配置nginx反代理和proxy_pass转发。 就可以愉快的使用了



## 搭建过程

1、初始化项目，切换到项目根目录

``` go mod init "自己项目的项目名"
# 1 go mod init "tools_go"   # tools_go项目名称
# 创建main.go文件

# 2 进入tools_go/router目录，初始化模块
go mod init router
# 下载依赖包，router需要gin依赖
go get -u github.com/gin-gonic/gin

# 3 进入tools_go/service目录，初始化模块
go mod init service

# 4 初始化工作区
go work init ./router ./service
```

ps:

> 如果下载超时，换成国内可访问的地址，设置GOPROXY代理：
>
> go env -w GOPROXY=https://goproxy.cn,direct
>
> go env -w GOSUMDB=off （关闭包的有效性验证）
>
> go env -w GOSUMDB=“sum.golang.google.cn” （也可设置国内提供的sum 验证服务）

2、编写相关的业务代码

todo

补充

> go.work 文件中总共支持三个指令。
>
> - go：声明 go 版本号，主要用于后续新语义的版本控制。
> - use：声明应用程序所依赖的模块的特定文件路径。该路径可以是绝对的或相对的，并且可以在应用程序的命运目录之外。
> - replace：声明模块依赖的导入路径被替换，优先于 go.mod 中的 replace 指令。



### 附录

1、可以在https://dev.maxmind.com/geoip/geolite2-free-geolocation-data 上注册账号 申请key 自动更新，是免费的。

> 通过注册引导页面，点击“Generate a License Key”
>
> https://github.com/maxmind/geoipupdate
> 这里有详细的安安装及配置说明
>
> 运行geoipupdate命令并加入定时任务

2、python代码:

安装：

> pip install geoip2

vim web.py

> from flask import Flask
> import geoip2.database
>
> app = Flask(__name__)
> reader=geoip2.database.Reader('/home/geoipupdate_4.8.0_linux_amd64/GeoLite2-City.mmdb')
> @app.route("/")
> def index():
>
>    ip_addr = request.remote_addr
>
> ​    return "Hello World!"
>
> @app.route("/getip/<ip>")
> def getip(ip):
>     ipinfo=reader.city(ip)
>     ipinfo_json={'country':ipinfo.country.name,'city':ipinfo.city.name,'location':[ipinfo.location.longitude,ipinfo.location.latitude]}
>     return ipinfo_json
>
> if __name__ == "__main__":
>     app.run(host='0.0.0.0',port=8090)

访问链接 http://192.168.1.166:8090/getip/113.97.30.189

## 参考

1、https://gin-gonic.com/

2、 https://github.com/gin-gonic/gin

3、https://github.com/maxmind










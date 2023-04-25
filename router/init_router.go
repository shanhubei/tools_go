package router

// 同个目录中（不含子目录）的所有文件包名必须一致
// init_router.go和user_router.go的包名需要一致
// 同一个包下不能出现重复的函数名称（大小写敏感）
import (
	"github.com/gin-gonic/gin"
)

// 命名规范建议使用驼峰式
// 注意：向外暴漏的函数首字母大写；不向外暴漏的函数首字母小写

func InitRouter() {
	router := gin.Default()

	// 初始化工具路由
	tools_router_group := router.Group("/tools")
	InitTestRouter(tools_router_group)

	router.Run(":9090")
}

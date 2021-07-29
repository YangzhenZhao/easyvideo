package main

import (
	"github.com/YangzhenZhao/easyvideo/gin_back_end/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := routers.InitRouters()
	r.Run(":8000")
}

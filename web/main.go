package main

import (
	micpro "com.qmz.web/micpro"
	context2 "context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/consul"
	"net/http"
)

func microCall(context *gin.Context){
	//创建consul连接
	reg := consul.NewRegistry()
	//将consul连接接入到微服务
	service := micro.NewService(micro.Registry(reg))
	//调用微服务
	micServ := micpro.NewMicproService("go.micro.srv.micpro",service.Client())
	//获取微服务结果
	res,err := micServ.Call(context2.TODO(),&micpro.Request{})
	if err != nil{
		fmt.Println(err)
	}

	//处理微服务返回结果
	context.JSON(http.StatusOK,gin.H{"code":11111,"data":res.Msg})

}


func main(){

	route := gin.Default()
	route.GET("/", microCall)
	route.Run(":8080")
}

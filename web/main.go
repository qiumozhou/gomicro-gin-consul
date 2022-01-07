package main

import (
	micpro "com.qmz.web/micpro"
	context2 "context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	//"github.com/micro/go-micro/client"
	"github.com/micro/go-plugins/registry/consul"
	//"time"
)

func microCall(context *gin.Context){

	reg := consul.NewRegistry()
	service := micro.NewService(micro.Registry(reg))
	micServ := micpro.NewMicproService("go.micro.srv.micpro",service.Client())
	res,err := micServ.Call(context2.TODO(),nil)
	if err != nil{
		fmt.Println(err)
	}

	context.Writer.WriteString(res.Msg)
}


func main(){

	route := gin.Default()
	route.GET("/", microCall)
	route.Run(":8080")
}

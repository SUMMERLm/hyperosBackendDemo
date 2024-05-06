package main

import (
	"github.com/SUMMERLm/serverless/pkg/route"
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

func main() {
	r := route.InitRouter()
	// listen on 8001
	gin.SetMode("debug")
	err := r.Run(":8001")
	if err != nil {
		klog.Error(err)
	}
}

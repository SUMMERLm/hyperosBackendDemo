package api

import (
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

var stateDone = 200

func GetIndex(c *gin.Context) {
	c.String(stateDone, "HyperOS backend Demo: Get test")
}

func PostTest(c *gin.Context) {
	hpaMsg := c.GetHeader("Alertname")
	klog.Infof(hpaMsg)
	c.String(stateDone, "HyperOS backend Demo: Post test")
}

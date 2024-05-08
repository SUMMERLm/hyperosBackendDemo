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
	scnid := c.GetHeader("SCNID")
	klog.Infof("scnid is %s ", scnid)
	HOST := c.GetHeader("HOSTss")
	klog.Infof("host is %s", HOST)

	c.String(stateDone, "HyperOS backend Demo: Post test")
}

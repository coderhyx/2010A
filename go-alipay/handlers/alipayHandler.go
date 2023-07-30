package handlers

import (
	"fmt"
	"go-alipay/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/xlog"
)

func AliPayNotify(c *gin.Context) {
	tradeStatus := c.PostForm("trade_status")
	fmt.Println("-------------->", tradeStatus)
	if tradeStatus == "TRADE_CLOSED" {
		c.JSON(http.StatusOK, gin.H{
			"msg": "交易已关闭",
		})
	}
	if tradeStatus == "TRADE_SUCCESS" {
		//验签
		//todo 做自己的业务
		c.JSON(http.StatusOK, gin.H{
			"msg": "成功！",
		})
	}
}
func AliPayReturn(c *gin.Context) {
	notifyReq, err := alipay.ParseNotifyToBodyMap(c.Request)
	if err != nil {
		xlog.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}
	ok, err := alipay.VerifySign(config.AliPublicKey, notifyReq)
	if err != nil {
		xlog.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}
	msg := ""
	if ok {
		msg = "验签成功"
	} else {
		msg = "验签失败"
	}
	//TODO,做自己的业务
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

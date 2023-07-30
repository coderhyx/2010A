package main

import (
	"context"
	"fmt"
	"go-alipay/config"
	"go-alipay/handlers"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/xlog"
)

func main() {
	//1.获取url进行支付
	client, err := alipay.NewClient(config.AppId, config.PrivateKey, config.IsProduction)
	if err != nil {
		xlog.Error(err)
		return
	}
	client.SetCharset("utf-8").
		SetSignType(alipay.RSA2).
		SetNotifyUrl(config.NotifyURL).
		SetReturnUrl(config.ReturnURL)

	ts := time.Now().UnixMilli()
	fmt.Println("OutTradeNo", ts)
	outTradeNo := fmt.Sprintf("%d", ts)
	bm := make(gopay.BodyMap)
	bm.Set("subject", "2101A专题")
	funcName(bm, outTradeNo)
	bm.Set("total_amount", "10.00")
	bm.Set("product_code", config.ProductCode)

	payUrl, err := client.TradePagePay(context.Background(), bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	//if {if { if {}else{}}else}else{}
	xlog.Debugf("payUrl", payUrl)

	//2.支付成功后，支付宝回调我们
	engine := gin.Default()
	engine.POST("/notify", handlers.AliPayNotify)
	engine.GET("/return", handlers.AliPayReturn)
	engine.Run(":8999")
}

func funcName(bm gopay.BodyMap, outTradeNo string) gopay.BodyMap {
	return bm.Set("out_trade_no", outTradeNo)
}

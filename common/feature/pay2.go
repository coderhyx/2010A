package feature

import (
	"fmt"

	"github.com/smartwalle/alipay/v3"
)

func Pay2() {
	var privateKey = "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDf4/19hCVbrsShf/zXCU+/5OEQbEd4YLyOe8BYBB8L+xjfzQ6GuylME/fTQYEpdIYcHU6OfKhgmLlsQllT2OZ5UZnFAZPhpMGuWF7DIuy+x+XdRfJE8h/eJh3BYGb58iIOOajOcwUkOVpgN/kXPU/mN3CN0WT+FITBKmiEGRjIHu5h3IbdFEedFyOvdym1eaSR1Ek7os3Xa+M8rYrpkIj6fgOB0FOQ2FJquHJCdQm7qYSzVK/5jf/h3xBxnGq/iBP2+tdXbfPWXy7Bcys/J9E3GGJSHMH6kapK23Gt+JnDYglQgUIw2aE4crwyvJT/BK7mkGy2hx8RimYNNKrhXs3xAgMBAAECggEBAMWeoSZEGRGG/uqqZQuNnYX8yafMW31mrah6lPlbkefqWDda8UJ9O2N6kJo4zIBB6Qox2CAu6hRxWeGz4tL+tdhJ7ZeV3+kgmxB/0g0d66guG7gnQEQZD4XvUP5aUCq4zdSOknC/1770nNAnN8eKh9bmAoQ7WpBmnhM+kohe+p/P/o/D4EOH4MYCTJtCe4Ea6W8o0dVG5zfYR0SQ+AlSYgLYXi0dh72PmOuahNUvS3I53/7BW17YZ3xojmK+liP+wUuwZ6k/1lOhLrk3qH9rYY0Ud//hp6BBrF/T+fafkVa3fICjRdqVVX6JRPa8U3Ag86eudnnIFuiBFRNDOr6gRAECgYEA+U8fd1qociCdu2eYP/N6qjy15NEv3fYHckH9VhfdNJVIjDADwvmIbexJ2hnHSFx3tMkAWXaQ86JCB2w47qLRgryI66O3MFoCgWKrRb2r1GN4XNgcks2YOqgU86bWbHpzzwF8ZmKznMVAagJ6wYQ3t56CllzaDk9/mr12K6wAILECgYEA5eY6ytr6hwBF77qB9i6RjkjOUoKL+6T6ZLMDLBzrFs4bw0xZI/IIXJG2Ojvtz6sRWtUYMNvhIEiNfCy/tAVcRM6vqu6xwF5B+zL8F3t1FqNsPRgjMSCxPryUJiu79ASg8eKmLZEAQSi8FWi70Vg34/80tt7xKnnl8xlu9l6G0UECgYA0L0Gw5AMaUIVzss2FpVtpwud4C+lvFo6cdf+nQ7uDxDO5VFuVNlI+YBVdV8QE+4X7C4/NKipRNQeJMjgAi4g7S2eFm3E+57haiOK86GTNQjzxgjMI74wLyx8HmUaM0lznWbJGZCagjBFdn0M+uoRHJSDEhI8IK8/T/hB8N7aXIQKBgHudswk2e7UxgWloyM69tYhjP4WAKpLR381bsB39Iq9tfeIiYGACmVplAy4G4VVjr34+zLYg1MlOGb2mNiIvK7DXzf9EP5GnXSUcAg8CVDArCV1EaE/XO1b3gVWQ+Iw6HOxTKXWg3RksNQ3x9eOX4t2WcRrPf6+OQgXYLCEPLMDBAoGAQZxHWuzGCW7fEUa8W6/CpDClq2x3nxN/7n8Geadqg7wTpvO3BYiJDsx+9KK/rwU/gEva95FTs0yNXh9vGa6Mt+7SsHJHvj4eVC9lKgTeaKZ32ySbXYeoPkQKzyPALVk7KJBl59Nrsb7/oU0E7fC81mZX72YVW/BTLtsRnyRpd0c=" // 这个地方填写之前生成的RSA2 私钥
	var appId = "9021000123613685"                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              // 在沙箱环境第一个参数就是
	// 创建支付宝客户端，沙箱环境下第三个参数填false，生产环境要换成true
	var client, _ = alipay.New(appId, privateKey, false)

	// 如果涉及到资金的支出，必须用公钥证书模式，如果只涉及支付推荐使用普通公钥模式，也就是刚刚生成的RSA2密钥
	/*  公钥证书模式
	client.LoadAppPublicCertFromFile("appCertPublicKey_2017011104995404.crt") // 加载应用公钥证书
	client.LoadAliPayRootCertFromFile("alipayRootCert.crt") // 加载支付宝根证书
	client.LoadAliPayPublicCertFromFile("alipayCertPublicKey_RSA2.crt") // 加载支付宝公钥证书
	*/
	//普通公钥模式，加载支付宝公钥，在沙箱环境上传RSA2公钥后生成的支付宝公钥
	client.LoadAliPayPublicKey("MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwCOf6ASz1XTiswHCN+uLmwagdgoh0D6aEbztuFqRgmi5xi3S+JPWOvNJVpaNntsOGsVbH9KPMZR1pQsYZLJ52gD2+PUzVcPrCltTkBlR2QXamvSjwFCYRgOtbNr4FGtbgYE83Fh0TC1IoLGdlvpKrQDdItNi0nYC/yDow3xroRQR1Vplhz3F/DHDJ5efyZ/yuOUG7p0oc3s3RoUeH2KW9sIcSewN+0mW5NLgXSz/U1yIpSC5PDI1LDBZP3XEl3BfPQ+D5jYGoxC46KuZBq8ii3nKI2TUH4hnYtRLXw5o+CPABEVttywMYWWNpZ15K4lJtqPyQAsZfaV2aInGoCSpDwIDAQAB")

	var p = alipay.TradeAppPay{}
	p.NotifyURL = "http://zbiy92.natappfree.cc/callback" //回调地址
	p.Subject = "标题"
	p.OutTradeNo = "传递一个唯一单号" //自己生成一个唯一的交易单号
	p.TotalAmount = "10.00"   //交易金额

	//values := url.Values{} //回调参数，这里只能这样写，要进行urlEncode才能传给支付宝
	//需要回传的参数
	//values.Add("aaa", aaa)
	//values.Add("bbb", bbb)
	//p.PassbackParams = values.Encode() //支付宝会把passback_params={aaa=aaa&bbb=bbb}发送到回调函数

	//这里返回的url中会包含sign，直接返回给前端就ok

	var url, err = client.TradeAppPay(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("URL:", url)
}

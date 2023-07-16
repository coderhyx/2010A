package feature

import (
	"fmt"

	unisms "github.com/apistd/uni-go-sdk/sms"
)

type UniSMS struct {
	AccessKeyId  string
	AccessSecret string
	Signature    string
	TemplateId   string
}

func NewUniSMS(AccessKeyId, AccessSecret, Signature, TemplateId string) *UniSMS {
	return &UniSMS{
		AccessKeyId,
		AccessSecret,
		Signature,
		TemplateId,
	}
}

func (sms *UniSMS) SendCode(phone, code string) error {
	// 初始句柄
	client := unisms.NewClient(sms.AccessKeyId, sms.AccessSecret) // 若使用简易验签模式仅传入第一个参数即可
	// 构建信息
	message := unisms.BuildMessage()
	message.SetTo(phone)
	message.SetSignature(sms.Signature)
	message.SetTemplateId(sms.TemplateId)
	message.SetTemplateData(map[string]string{"code": code}) // 设置自定义参数 (变量短信)
	// 发送短信
	res, err := client.Send(message)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("====================>1:", res)
	return nil
}

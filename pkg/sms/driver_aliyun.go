package sms

import (
	"encoding/json"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"server/pkg/logger"
)

// Aliyun 实现 sms.Driver interface
type Aliyun struct{}

// CreateClient /**
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Dysmsapi
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

// Send 实现 sms.Driver interface 的 Send 方法
func (s *Aliyun) Send(phone string, message Message, config map[string]string) bool {
	// 请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID 和 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
	// 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例使用环境变量获取 AccessKey 的方式进行调用，仅供参考，建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html
	client, err := CreateClient(tea.String(config["access_key_id"]), tea.String(config["access_key_secret"]))
	if err != nil {
		logger.ErrorString("短信[阿里云]", "解析绑定错误", err.Error())
		return false
	}
	logger.DebugJSON("短信[阿里云]", "配置信息", config)

	param, err := json.Marshal(message.Data)
	if err != nil {
		logger.ErrorString("短信[阿里云]", "短信模板参数解析错误", err.Error())
		return false
	}

	// 发送参数
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String(config["sign_name"]),
		TemplateCode:  tea.String(message.Template),
		PhoneNumbers:  tea.String(phone),
		TemplateParam: tea.String(string(param)),
	}

	// 其他运行参数
	runtime := &util.RuntimeOptions{}

	_result, err := client.SendSmsWithOptions(sendSmsRequest, runtime)

	logger.DebugJSON("短信[阿里云]", "请求内容", _result.Body.RequestId)
	logger.DebugJSON("短信[阿里云]", "接口响应", _result.Body)

	if err != nil {
		var errs = &tea.SDKError{}
		if _t, ok := err.(*tea.SDKError); ok {
			errs = _t
		} else {
			errs.Message = tea.String(err.Error())
		}

		logger.ErrorString("短信[阿里云]", "发信失败", errs.Error())

		var r dysmsapi20170525.SendSmsResponseBody
		err = json.Unmarshal([]byte(*errs.Data), &r)
		if err != nil {
			logger.ErrorString("短信[阿里云]", "解析响应 JSON 错误", err.Error())
		}

		logger.ErrorString("短信[阿里云]", "服务商返回错误", *r.Message)
		return false
	}

	logger.DebugString("短信[阿里云]", "发信成功", "")
	return true
}

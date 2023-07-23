package mail

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/spf13/cast"
	"os"
	"server/pkg/logger"
)

// SES 实现 email.Driver interface
type SES struct{}

// Send 实现 email.Driver interface 的 Send 方法
func (s *SES) Send(email Email, config map[string]string) bool {
	// 设置访问密钥和密钥 ID
	accessKeyID := config["key"]
	secretAccessKey := config["secret"]

	// 创建 AWS 会话
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(config["region"]),
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
	})
	// 创建会话错误
	if err != nil {
		fmt.Println("Error creating session:", err)
		os.Exit(1)
	}

	// 群发
	var emails []*string
	for _, email := range email.To {
		emails = append(emails, aws.String(email))
	}

	// 组装电子邮件
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: emails,
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(cast.ToString(email.HTML)),
				},
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(cast.ToString(email.Text)),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(email.Subject),
			},
		},
		Source: aws.String(email.From.Address),
	}

	svc := ses.New(sess)
	// 尝试发送电子邮件。
	_, err = svc.SendEmail(input)

	// 如果出现错误，则显示错误消息。
	if err != nil {
		// 打印错误的消息。
		if err, ok := err.(awserr.Error); ok {
			switch err.Code() {
			case ses.ErrCodeMessageRejected:
				logger.ErrorString("发送邮件", "发件出错："+ses.ErrCodeMessageRejected, err.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				logger.ErrorString("发送邮件", "发件出错："+ses.ErrCodeMailFromDomainNotVerifiedException, err.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				logger.ErrorString("发送邮件", "发件出错："+ses.ErrCodeConfigurationSetDoesNotExistException, err.Error())
			default:
				logger.ErrorString("发送邮件", "发件出错", err.Error())
			}
		}
		return false
	}

	logger.DebugString("发送邮件", "发件成功", "")
	return true
}

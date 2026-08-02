package service

import (
	"fmt"
	"server/global"
	"server/utils"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BaseService struct{}

func (baseservice *BaseService) SendEmailVerificationCode(c *gin.Context, to string) error {
	verificationCode := utils.GenerateVerificationCode(6)
	expireTime := time.Now().Add(5 * time.Minute).Unix()

	session := sessions.Default(c)
	session.Set("varification_code", verificationCode)
	session.Set("email", to)
	session.Set("expire_time", expireTime)
	_ = session.Save()
	subject := "您的邮箱验证码"
	body := fmt.Sprintf(`亲爱的用户[%s]，<br/>
<br/>
感谢您注册%s的个人博客！为了确保您的邮箱安全，请使用以下验证码进行验证：<br/>
<br/>
验证码：[<font color="blue"><u>%s</u></font>]<br/>
该验证码在 5 分钟内有效，请尽快使用。<br/>
<br/>
如果您没有请求此验证码，请忽略此邮件。
<br/>
如有任何疑问，请联系我们的支持团队：<br/>
邮箱：%s<br/>
<br/>
祝好，<br/>%s<br/>
<br/>`, to, global.Config.Website.Name, verificationCode, global.Config.Email.From, global.Config.Website.Title)
	err := utils.Email(to, subject, body)
	if err != nil {
		global.Log.Error("发送邮件验证码失败", zap.Error(err))
	}
	return nil
}

package appTypes

import "encoding/json"

// Register 用户注册来源
type Register int

const (
	Email Register = iota // 邮箱验证码注册
	QQ                    // QQ登录注册
)

func (r Register) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r *Register) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*r = ToRegister(s)
	return nil

}

func (r Register) String() string {
	var str string
	switch r {
	case Email:
		str = "邮箱"
	case QQ:
		str = "QQ"
	default:
		return "未知"
	}
	return str
}

func ToRegister(s string) Register {
	switch s {
	case "邮箱":
		return Email
	case "QQ":
		return QQ
	default:
		return -1
	}
}

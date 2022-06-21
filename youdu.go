package youdu

type Youdu struct {
	config *Config

	dept          *Dept
	messageSender *MessageSender
	media         *Media
	user          *User
	session       *Session
	group         *Group
}

// New 创建一个 Youdu 实例
func New(config *Config) *Youdu {
	return &Youdu{
		config: config,
	}
}

// Message 创建消息相关的实例
func (y *Youdu) Message() *MessageSender {
	if y.messageSender == nil {
		y.messageSender = NewMessageSender(y.config)
	}

	return y.messageSender
}

// Media 创建媒体相关的实例
func (y *Youdu) Media() *Media {
	if y.media == nil {
		y.media = NewMedia(y.config)
	}

	return y.media
}

// Dept 创建部门相关的实例
func (y *Youdu) Dept() *Dept {
	if y.dept == nil {
		y.dept = NewDept(y.config)
	}

	return y.dept
}

// User 创建用户相关的实例
func (y *Youdu) User() *User {
	if y.user == nil {
		y.user = NewUser(y.config)
	}

	return y.user
}

// Session 创建会话相关的实例
func (y *Youdu) Session() *Session {
	if y.session == nil {
		y.session = NewSession(y.config)
	}

	return y.session
}

// Group 创建群相关的实例
func (y *Youdu) Group() *Group {
	if y.group == nil {
		y.group = NewGroup(y.config)
	}

	return y.group
}

// AccessToken 返回 accessToken
func (y *Youdu) AccessToken() (string, error) {
	return y.config.GetAccessTokenProvider().GetAccessToken()
}

// Encryptor 返回加密器
func (y *Youdu) Encryptor() *encryptor {
	return y.config.GetEncryptor()
}

// Config 获取配置
func (y *Youdu) Config() *Config {
	return y.config
}

package youdu

type Youdu struct {
	config *Config

	dept          *Dept
	messageSender *MessageSender
	media         *Media
	user          *User
	session       *Session
}

func New(config *Config) *Youdu {
	return &Youdu{
		config: config,
	}
}

func (y *Youdu) Message() *MessageSender {
	if y.messageSender == nil {
		y.messageSender = NewMessageSender(y.config)
	}

	return y.messageSender
}

func (y *Youdu) Media() *Media {
	if y.media == nil {
		y.media = NewMedia(y.config)
	}

	return y.media
}

func (y *Youdu) Dept() *Dept {
	if y.dept == nil {
		y.dept = NewDept(y.config)
	}

	return y.dept
}

func (y *Youdu) User() *User {
	if y.user == nil {
		y.user = NewUser(y.config)
	}

	return y.user
}

func (y *Youdu) Session() *Session {
	if y.session == nil {
		y.session = NewSession(y.config)
	}

	return y.session
}

func (y *Youdu) GetAccessToken() (string, error) {
	return y.config.GetAccessTokenProvider().GetAccessToken()
}

func (y *Youdu) Encryptor() *encryptor {
	return y.config.GetEncryptor()
}

func (y *Youdu) Config() *Config {
	return y.config
}

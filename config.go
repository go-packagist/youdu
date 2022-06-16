package youdu

type Config struct {
	Api    string
	Buin   int
	AppId  string
	AesKey string

	encryptor           *encryptor
	http                *Http
	accessTokenProvider *accessTokenProvider
}

func (c *Config) GetEncryptor() *encryptor {
	if c.encryptor == nil {
		c.encryptor = NewEncryptor(c)
	}

	return c.encryptor
}

func (c *Config) GetHttp() *Http {
	if c.http == nil {
		c.http = NewHttp(c)
	}

	return c.http
}

func (c *Config) GetAccessTokenProvider() *accessTokenProvider {
	if c.accessTokenProvider == nil {
		c.accessTokenProvider = NewAccessTokenProvider(c)
	}

	return c.accessTokenProvider
}

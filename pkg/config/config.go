package config

import (
	"flag"
	"net/url"
	"os"
	"strings"
)

type Config struct {
	apiKey           string
	apiKeyHeaderName string
	serverUrl        string
	username         string
	password         string
	authType         string
}

func Get() *Config {
	conf := &Config{}
	flag.StringVar(&conf.username, "username", os.Getenv("USERNAME"), "Basic Auth username")
	flag.StringVar(&conf.password, "password", os.Getenv("PASSWORD"), "Basic Auth password")
	flag.StringVar(&conf.authType, "authType", os.Getenv("AUTH_TYPE"), "Auth Type")
	flag.StringVar(&conf.apiKeyHeaderName, "apiKeyHeaderName", os.Getenv("API_KEY_HEADER_NAME"), "API Key Header Name")
	flag.StringVar(&conf.apiKey, "apiKey", os.Getenv("API_KEY"), "API Key")
	flag.StringVar(&conf.serverUrl, "serverUrl", os.Getenv("SERVER_URL"), "Server Url")

	flag.Parse()

	return conf
}

func (c *Config) GetApiKey() string {
	return c.apiKey
}

func (c *Config) GetApiKeyHeaderName() string {
	return c.apiKeyHeaderName
}

// GetAuthType returns the auth type accepted by the server
// Possible values include: API_KEY, BASIC_AUTH
func (c *Config) GetAuthType() string {
	// convert all characters to upper case
	authType := strings.ToUpper(c.authType)
	// replace space, hyphen with underscore
	authType = strings.ReplaceAll(authType, " ", "_")
	authType = strings.ReplaceAll(authType, "%20", "_")
	authType = strings.ReplaceAll(authType, "-", "_")
	return authType
}

func (c *Config) GetUsernameAndPassword() (string, string) {
	return c.username, c.password
}

func (c *Config) GetServerURL() string {
	c.serverUrl = strings.TrimSuffix(c.serverUrl, "/")
	u, _ := url.Parse(c.serverUrl)
	if u.Scheme == "" {
		return "https://" + c.serverUrl
	} else {
		return c.serverUrl
	}
}

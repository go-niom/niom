package config

const SMTPConfig = `package config

// SMTP holds the SMTP configuration
type SMTP struct {
	Host     string
	Port     string
	User     string
	Password string
	From     string
}

var smtp = &SMTP{}

// SMTPCfg returns the default SMTP configuration
func SMTPCfg() *SMTP {
	return smtp
}

// LoadSMTPCfg loads SMTP configuration
func LoadSMTPCfg() {
	smtp.Host = getEnv("SMTP_HOST", "smtp-relay.sendinblue.com")
	smtp.Port = getEnv("SMTP_PORT", "587")
	smtp.User = getEnv("SMTP_USER", "")
	smtp.Password = getEnv("SMTP_PASSWORD", "")
	smtp.From = getEnv("MAIL_FROM", "")
}`

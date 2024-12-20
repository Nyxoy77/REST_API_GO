package email

import (
	"fmt"
	"net/smtp"
	"sync"
	"time"

	"github.com/spf13/viper"
)

var emailCache = make(map[string]time.Time)
var cacheMutex sync.Mutex

func SendEmailWithCooldown(to []string, subject string, body string, cooldown time.Duration) error {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	// Check cooldown for the first recipient
	if lastSent, exists := emailCache[to[0]]; exists {
		if time.Since(lastSent) < cooldown {
			return fmt.Errorf("cooldown period not over, please wait before sending another email")
		}
	}

	// Set up SMTP authentication
	auth := smtp.PlainAuth(
		"",
		viper.GetString("FROM_EMAIL"),
		viper.GetString("APP_PASSWORD"),
		viper.GetString("SMTP_SERVER"),
	)

	// Create the email message
	message := "Subject: " + subject + "\n\n" + body

	// Send the email
	err := smtp.SendMail(
		viper.GetString("SMTP_SERVER_WITH_PORT"),
		auth,
		viper.GetString("FROM_EMAIL"),
		to,
		[]byte(message),
	)
	if err != nil {
		return err
	}

	// Update cache
	emailCache[to[0]] = time.Now()
	return nil
}

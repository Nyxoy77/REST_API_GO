package services

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Nyxoy/restAPI/email"
	"github.com/Nyxoy/restAPI/utils"
)

type EmailBody struct {
	Toaddr  string `json:"to_addr"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func HandleEmail(w http.ResponseWriter, r *http.Request, mail EmailBody) {

	to := strings.Split(mail.Toaddr, ",")
	coolTime := 120 * time.Second
	err := email.SendEmailWithCooldown(to, mail.Subject, mail.Body, coolTime)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Wait for the cooldown period")
		log.Println("Email could not be sent")
		return
	}
	utils.WriteError(w, http.StatusOK, "If you have a valid registered mail id an email will be sent shortly")

}

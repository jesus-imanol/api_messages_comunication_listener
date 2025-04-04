package adapters

import (
	"apimessages/src/messages/domain/entities"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"gopkg.in/gomail.v2"
	"github.com/joho/godotenv"
)

type SMTP struct {
	historial map[string]map[string]time.Time 
	mu        sync.Mutex
}

func NewSMTP() *SMTP {
	return &SMTP{historial: make(map[string]map[string]time.Time)}
}

func (uc *SMTP) CaseError(errorMessage entities.Message, gmail string) error {
	now := time.Now()

	uc.mu.Lock()
	defer uc.mu.Unlock()

	if usuarios, existe := uc.historial[errorMessage.Text]; existe {
		if ultimaHora, enviado := usuarios[errorMessage.User]; enviado {
			if now.Sub(ultimaHora) < time.Hour {
				fmt.Printf("El mensaje '%s' ya fue enviado al usuario '%s' recientemente.\n", errorMessage.Text, errorMessage.User)
				return nil
			}
		}
	} else {
		uc.historial[errorMessage.Text] = make(map[string]time.Time) 
	}

	err := SendGmail(errorMessage.Text, gmail)
	if err != nil {
		return err
	}

	uc.historial[errorMessage.Text][errorMessage.User] = now
	return nil
}

func SendGmail(errorMessage string, gmail string) error {
	errg := godotenv.Load()
	if errg != nil {
		fmt.Println("Error al cargar el archivo.env: ", errg)
		return errors.New("fallo al cargar el archivo.env")
	}

	from := os.Getenv("GMAIL")
	password := os.Getenv("GMAIL_PASS")
	host := "smtp.gmail.com"
	port := 587

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", gmail)
	m.SetHeader("Subject", "Nuevo mensaje de alerta")

	timestamp := time.Now().Format("02/01/2006 15:04:05")

	htmlBody := `<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <title>Alerta del Sistema</title>
    <style>
        body { font-family: Arial, sans-serif; background-color: #f7f7f7; }
        .container { max-width: 600px; margin: 0 auto; background-color: #ffffff; padding: 20px; border-radius: 5px; }
        .header { font-size: 18px; font-weight: bold; }
        .message { background-color: #ffdddd; padding: 10px; border-left: 4px solid #ff0000; }
        .footer { text-align: center; font-size: 12px; color: #666; margin-top: 20px; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">⚠️ Alerta del Sistema</div>
        <p>Se ha detectado un evento crítico que requiere atención:</p>
        <div class="message">{{errorMessage}}</div>
        <p>Fecha y hora: {{timestamp}}</p>
        <div class="footer">Este es un mensaje automático, por favor no responder.</div>
    </div>
</body>
</html>`

	htmlBody = strings.ReplaceAll(htmlBody, "{{errorMessage}}", errorMessage)
	htmlBody = strings.ReplaceAll(htmlBody, "{{timestamp}}", timestamp)

	m.SetBody("text/plain", "Evento crítico detectado: "+errorMessage)
	m.AddAlternative("text/html", htmlBody)

	d := gomail.NewDialer(host, port, from, password)

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error al enviar el correo:", err)
		return errors.New("fallo al enviar el correo")
	}

	fmt.Println("Correo enviado exitosamente:", errorMessage)
	return nil
}
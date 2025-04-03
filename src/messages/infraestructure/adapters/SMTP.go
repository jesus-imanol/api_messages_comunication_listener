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
	historial map[entities.Message]time.Time
	mu        sync.Mutex
}

func NewSMTP() *SMTP {
	return &SMTP{historial: make(map[entities.Message]time.Time)}
}

func (uc *SMTP) CaseError(errorMessage entities.Message, gmail string) error {
	now := time.Now()

	uc.mu.Lock()
	defer uc.mu.Unlock()

	if ultimaHora, existe := uc.historial[errorMessage]; existe {
		if now.Sub(ultimaHora) < time.Hour {
			fmt.Println("El error ya fue enviado recientemente:", errorMessage.Text)
			return nil
		}
	}

	err := SendGmail(errorMessage.Text, gmail)
	if err != nil {
		return err
	}

	uc.historial[errorMessage] = now
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
	m.SetHeader("Subject", "1 mensaje nuevo para ti")

	timestamp := time.Now().Format("02/01/2006 15:04:05")


	htmlBody := `<!DOCTYPE html>
<html lang="es">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Alerta de Sistema</title>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
            line-height: 1.6;
            color: #222;
            background-color: #f7f7f7;
            margin: 0;
            padding: 20px;
        }
        .container {
            max-width: 600px;
            margin: 0 auto;
            background-color: #ffffff;
            border-radius: 4px;
            overflow: hidden;
            box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
        }
        .logo {
            text-align: center;
            padding: 20px 0;
            border-bottom: 1px solid #eaeaea;
        }
        .logo img {
            height: 40px;
        }
        .header {
            text-align: center;
            padding: 30px 30px 15px;
            font-size: 24px;
            font-weight: 600;
        }
        .divider {
            height: 1px;
            background-color: #eaeaea;
            margin: 20px 0;
        }
        .content {
            padding: 20px 30px;
        }
        .message-container {
            margin-bottom: 15px;
        }
        .error-box {
            background-color: #f9f9f9;
            border-left: 3px solid #ff7a59;
            padding: 15px;
            margin: 15px 0;
            color: #333;
        }
        .button-container {
            text-align: center;
            padding: 20px 0 30px;
        }
        .button {
            display: inline-block;
            background-color: #1dbf73;
            color: white;
            font-weight: 600;
            padding: 12px 24px;
            text-decoration: none;
            border-radius: 4px;
            text-align: center;
        }
        .footer {
            padding: 20px 30px;
            text-align: center;
            font-size: 12px;
            color: #999;
            border-top: 1px solid #eaeaea;
        }
        .timestamp {
            color: #999;
            font-size: 12px;
            margin-top: 5px;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="logo">
            <svg width="90" height="27" viewBox="0 0 90 27" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M18.5 0H22.5V7.7H26.5V11.7H22.5V27H18.5V11.7H15.5V7.7H18.5V0Z" fill="#222222"/>
                <path d="M31.6 7.7H35.6V9.7C36.1 8.9 36.8 8.3 37.5 7.9C38.2 7.5 39.1 7.3 39.9 7.3C41.7 7.3 43.1 7.9 44.1 9.1C45.1 10.3 45.6 12 45.6 14.1V27H41.6V14.9C41.6 13.6 41.3 12.7 40.8 12.1C40.3 11.5 39.5 11.2 38.5 11.2C37.4 11.2 36.6 11.6 36 12.3C35.4 13 35.1 14.2 35.1 15.8V27H31.1V7.7H31.6Z" fill="#222222"/>
                <path d="M54.1 11.6V20.4C54.1 21.2 54.3 21.7 54.6 22.1C54.9 22.5 55.5 22.6 56.3 22.6H58.1V27H55.3C53.6 27 52.3 26.6 51.5 25.7C50.7 24.8 50.2 23.5 50.2 21.7V11.7H48.2V7.7H50.2V3.3H54.2V7.7H58.2V11.7H54.1V11.6Z" fill="#222222"/>
                <path d="M66 0H70V7.7H74V11.7H70V27H66V11.7H63V7.7H66V0Z" fill="#222222"/>
                <path d="M83.5 7.3C85.2 7.3 86.7 7.8 87.8 8.9C88.9 10 89.5 11.5 89.5 13.3V27H85.7V25C84.8 26.6 83.1 27.4 80.8 27.4C79.8 27.4 78.9 27.2 78.1 26.9C77.3 26.6 76.7 26.1 76.2 25.5C75.7 24.9 75.5 24.2 75.5 23.4C75.5 22 76 20.9 77.1 20.1C78.2 19.3 79.8 18.9 82 18.9H85.5C85.5 17.8 85.2 17 84.6 16.4C84 15.8 83.1 15.5 81.9 15.5C81.1 15.5 80.3 15.6 79.5 15.9C78.7 16.2 78.1 16.5 77.5 17L75.9 13.8C76.7 13.2 77.7 12.7 78.8 12.3C79.9 11.9 81.1 11.7 82.4 11.7L83.5 7.3ZM82.1 24.1C82.9 24.1 83.6 23.9 84.2 23.4C84.8 22.9 85.2 22.3 85.4 21.5V21.2H82.3C80.5 21.2 79.5 21.8 79.5 23C79.5 23.5 79.7 24 80.2 24.3C80.6 24.7 81.3 24.9 82.1 24.9V24.1Z" fill="#222222"/>
                <circle cx="87" cy="7" r="3" fill="#1DBF73"/>
            </svg>
        </div>
        <div class="header">
            1 mensaje nuevo para ti
        </div>
        <div class="divider"></div>
        <div class="content">
            <div class="message-container">
                <p>Hola,</p>
                <p>Se ha detectado un error crítico en el sistema que requiere atención inmediata:</p>
                <div class="error-box">
                    {{errorMessage}}
                    <div class="timestamp">{{timestamp}}</div>
                </div>
            </div>
        </div>
        <div class="footer">
            <p>Este es un mensaje automático. Por favor, no responda a este correo.</p>
            <p>© 2025 - Sistema de Monitoreo API Messages</p>
        </div>
    </div>
</body>
</html>`

	htmlBody = strings.ReplaceAll(htmlBody, "{{errorMessage}}", errorMessage)
	htmlBody = strings.ReplaceAll(htmlBody, "{{timestamp}}", timestamp)

	m.SetBody("text/plain", "Error crítico en el sistema: "+errorMessage)
	m.AddAlternative("text/html", htmlBody)

	d := gomail.NewDialer(host, port, from, password)

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error al enviar el correo:", err)
		return errors.New("fallo al enviar el correo")
	}

	fmt.Println("Correo enviado exitosamente:", errorMessage)
	return nil
}
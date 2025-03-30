package adapters

import (
	"apimessages/src/messages/domain/entities"
	"errors"
	"fmt"
	"os"
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
	m.SetHeader("Subject", "Alerta de Error Crítico")
	m.SetBody("text/plain", "Se detectó un error crítico en el sistema:\n\n"+errorMessage)

	d := gomail.NewDialer(host, port, from, password)

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error al enviar el correo:", err)
		return errors.New("fallo al enviar el correo")
	}

	fmt.Println("Correo enviado exitosamente:", errorMessage)
	return nil
}
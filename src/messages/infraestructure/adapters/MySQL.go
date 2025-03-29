package adapters

import (
	"apimessages/src/core"
	"apimessages/src/messages/domain/entities"
	"fmt"
	"log"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() (*MySQL, error) {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}, nil
}

func (mysql *MySQL) CreateMessage(humidity entities.Message) (*entities.Message, error) {
	query := `INSERT INTO messages (type, quantity, text) VALUES (?, ?, ?)`
	result, err := mysql.conn.ExecutePreparedQuery(query, humidity.Type, humidity.Quantity, humidity.Text)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if result != nil {
		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 1 {
			log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
			lastInsertID, err := result.LastInsertId()
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
		humidity.ID = lastInsertID
		} else {
			log.Printf("[MySQL] - Ninguna fila fue afectada.")
		}
	} else {
		log.Printf("[MySQL] - Resultado de la consulta es nil.")
	}
	return &humidity, nil 
}

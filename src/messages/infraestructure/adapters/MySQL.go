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
	query := `INSERT INTO messages (type, quantity, text, created_at, username) VALUES (?, ?, ?, ?, ?)`
	result, err := mysql.conn.ExecutePreparedQuery(query, humidity.Type, humidity.Quantity, humidity.Text, humidity.CreatedAt, humidity.User)
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

func (mysql *MySQL) GetGmailByUserName(userName string) (string, error){
	query := `SELECT gmail FROM user WHERE username =?`
    rows, err := mysql.conn.FetchRows(query, userName)
    if err != nil {
        return "", fmt.Errorf("error al ejecutar la consulta SELECT: %v", err)
    }

    var gmail string
    for rows.Next() {
        err := rows.Scan(&gmail)
        if err != nil {
            return "", fmt.Errorf("error al leer el resultado de la consulta: %v", err)
        }
    }

    if gmail == "" {
        return "", fmt.Errorf("usuario %s no encontrado", userName)
    }

    return gmail, nil
}
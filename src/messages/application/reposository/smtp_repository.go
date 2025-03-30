package repository

import "apimessages/src/messages/domain/entities"

type ISmtp interface {
	CaseError(errorMessage entities.Message, gmail string) error
}

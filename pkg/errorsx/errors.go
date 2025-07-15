package errorsx

import (
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"
)

var (
	ErrNotFound     = errors.New("record not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")
	ErrConflict     = errors.New("conflict")
	ErrBadRequest   = errors.New("bad request")
	ErrInternal     = errors.New("internal server error")
)

func IsUniqueConstraintError(err error) (bool, string) {
	if err == nil {
		return false, ""
	}

	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		if mysqlErr.Number == 1062 {
			message := mysqlErr.Message
			if strings.Contains(message, "email") {
				return true, "email"
			}
			return true, "unknown"
		}
	}

	return false, ""
}

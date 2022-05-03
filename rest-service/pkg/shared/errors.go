package shared

import (
	"errors"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

var (
	ErrDatabase       = errors.New("database error")
	ErrPersonNotFound = errors.New("person not found")
	ErrGetAll         = errors.New("get all people failed")
	ErrFindPerson     = errors.New("could not find person")
	ErrInvalidUUID    = errors.New("invalid uuid")
	ErrJSONInvalid    = errors.New("invalid json request")
	ErrInvalidRequest = errors.New("invalid request")
)

type InternalCode int

const (
	BadRequest = iota + 1
	Unauthorized
	Forbidden
	NotFound
	Gone
	Internal
	BadGateway
	NoContent
)

type CodeMap map[int]int

var (
	CodesMap = CodeMap{
		BadRequest:   http.StatusBadRequest,
		Unauthorized: http.StatusUnauthorized,
		Forbidden:    http.StatusForbidden,
		NotFound:     http.StatusNotFound,
		Gone:         http.StatusGone,
		Internal:     http.StatusInternalServerError,
		BadGateway:   http.StatusBadGateway,
		NoContent:    http.StatusNoContent,
	}
	DefaultMSError = NewMSError("", 0, "", "", nil)
	DefaultDBError = NewDBError("", nil)
)

type AppLevel string

const (
	ServiceLevel   = "Service"
	TransportLevel = "Transport"
)

func GetMSError(err error) *MSError {
	if err != nil {
		msError := DefaultMSError
		if errors.Is(err, msError) && errors.As(err, &msError) {
			return msError.(*MSError)
		}
	}

	return nil
}

func NewMSError(text string, code InternalCode, op string, level AppLevel, err error) error {
	return &MSError{
		internalCode: code,
		Message:      text,
		Op:           op,
		Level:        level,
		err:          err,
		ErrStack:     "",
	}
}

type MSError struct {
	internalCode InternalCode
	Message      string
	Op           string
	Level        AppLevel
	err          error
	ErrStack     string
}

func (ms *MSError) Error() string {
	return ms.Message
}

func (ms *MSError) Code() int {
	return int(ms.internalCode)
}

func (ms *MSError) Unwrap() error {
	return ms.err
}

func (ms *MSError) Is(e error) bool {
	_, ok := e.(*MSError)
	return ok
}

func (ms *MSError) As(e error) bool {
	_, ok := e.(*MSError)
	return ok
}

func (ms *MSError) GetError() error {
	return ms.err
}

func (ms *MSError) Stack() error {
	err := ms.GetError()
	stack := ""
	count := 0
	for err != nil {
		count++
		stack += fmt.Sprintf("%d-: %s ", count, err.Error())
		err = errors.Unwrap(err)
	}
	ms.ErrStack = stack
	return ms
}

func NewDBError(message string, err error) error {
	return &DBError{
		Message: message,
		Err:     err,
	}
}

type DBError struct {
	Message string
	Err     error
}

func (db *DBError) Error() string {
	return db.Message
}

func (db *DBError) Unwrap() error {
	return db.Err
}

func (db *DBError) Is(e error) bool {
	_, ok := e.(*DBError)
	return ok
}

func HandleDbError(err error, message string, op string) error {
	if !errors.Is(err, DefaultDBError) {
		err = NewMSError(message, NotFound, op, ServiceLevel, err)
	} else {
		err = NewMSError(err.Error(), Internal, op, ServiceLevel, err)
	}
	return err
}

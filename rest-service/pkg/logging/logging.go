package logging

import (
	"io"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/wclaro123/stackpath/backend-developer-tests/rest-service/pkg/shared"
)

var (
	defaultLogLevel = []level.Option{level.AllowInfo()}
)

func NewLogger() log.Logger {
	logLevel := defaultLogLevel

	var w io.Writer = os.Stdout

	logger := log.NewJSONLogger(log.NewSyncWriter(w))
	logger = level.NewFilter(logger, logLevel...)
	logger = log.With(logger,
		"ts", log.DefaultTimestampUTC,
		"caller", log.DefaultCaller,
	)
	return logger
}

type LogError struct {
	Code    string `json:"code"`
	Source  string `json:"source"`
	Message string `json:"message"`
	Stack   string `json:"err_stack"`
}

func HandleError(err error) string {
	errorStack := ""
	thr := shared.GetMSError(err)
	if thr != nil {
		err = thr.Stack()
		errorStack = thr.ErrStack
	}

	return errorStack
}

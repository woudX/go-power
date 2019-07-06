// Package errors implements functions to manipulate errors.
package powerr

import (
	"fmt"
	"runtime/debug"
	"strings"
)

const (
	DefaultErrorSep = "||"
)

// More powerful error tool, it contains multi functions:
// - Record err params in map
// - Easy analysis with sep
// - Record backtrace
type PowError struct {
	msg   string
	data  map[string]interface{}
	stack []byte
}

//	New returns a PowErr with err message
func New(msg string) *PowError {
	return &PowError{
		msg:  msg,
		data: make(map[string]interface{}),
	}
}

//	NewE returns a PowErr with error, if e is PowError func just return error ref, else return with new error with msg
func NewE(e error) *PowError {
	if powErr, ok := e.(*PowError); ok {
		return powErr
	}

	return New(e.Error())
}

//	Print all error infos in PowError, user can select sep they like
func (e *PowError) ErrorWithSep(sep string) string {
	msgList := make([]string, 0, len(e.data)+1)

	msgList = append(msgList, e.msg)
	for key, val := range e.data {
		msgList = append(msgList, fmt.Sprintf("%v=%v", key, val))
	}

	if len(e.stack) > 0 {
		msgList = append(msgList, fmt.Sprintf("stack=%v", e.stack))
	}

	return strings.Join(msgList, sep)
}

//	Print all error info in PowError
func (e *PowError) Error() string {
	return e.ErrorWithSep(DefaultErrorSep)
}

//	Store kv params to error
func (e *PowError) StoreKV(key string, val interface{}) *PowError {
	e.data[key] = val
	return e
}

//	Store stack to error, only one stack can be stored
func (e *PowError) StoreStack() *PowError {
	e.stack = debug.Stack()
	return e
}

//	Get stack store in error
func (e *PowError) Stack() []byte {
	return e.stack
}
package utils

type ErrorType uint64

var (
	ErrorTypeNoUser ErrorType = 1 << 63
	ErrorTypePsd    ErrorType = 1 << 62
	ErrorTypeAny    ErrorType = 1<<64 - 1
)

type ServiceError struct {
	Err  error
	Type ErrorType
}

func NewServiceError(err error, t ErrorType) *ServiceError {
	return &ServiceError{
		Err:  err,
		Type: t,
	}
}

func (e *ServiceError) Error() string {
	return e.Err.Error()
}

func (e *ServiceError) SetType(t ErrorType) {
	e.Type = t
}

func (e *ServiceError) IsType(t ErrorType) bool {
	return (e.Type & t) > 0
}

func (e *ServiceError) Unwrap() error {
	return e.Err
}

// func NewError(t ErrorType, format string, arg ...any) *ServiceError {
// 	e := fmt.Errorf(format, arg...)
// 	return &ServiceError{
// 		err: e,
// 		t:   t,
// 	}
// }

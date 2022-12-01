package cryptom

//仿golang-jwt的errors处理，非常规整

//错误码 参考google.golang.org/grpc/codes

import (
	"errors"
)

// 简单型，本身就是底层，供Is对比用
var (
	ErrContentEncryptionKeyNotFoundInContext = errors.New("未能从context中获得CEK（ContentEncryptionKey），请确保中间已正确配置")
	ErrCEKNotFoundInContext                  = ErrContentEncryptionKeyNotFoundInContext
	ErrECEKNotFoundInHeader                  = errors.New("header中ECEK字段不存在")
)

//
// 复合型的Type枚举，用以组装CryptomError类型错误（带cause）
type ErrType uint

const (
	ErrTypeECEKDecryptFailure ErrType = 1 << iota
	ErrTypeContentDecryptFailure
	ErrTypeContentEncryptFailure
)

func NewCryptomError(errType ErrType, cause error, message string) *CryptomError {
	return &CryptomError{errType: errType, cause: cause, message: message}
}

// 复合型，会带cause，判断时要用AS并比对ErrType（e.g.见standardDemo的manualDemoLogic）
type CryptomError struct {
	errType ErrType
	cause   error
	message string
}

func (e CryptomError) Error() string {
	if e.cause == nil {
		return e.message
	} else {
		return e.message + ": " + e.cause.Error() //与github.com/pkg/errors  的withMessage 统一风格
	}
}
func (e *CryptomError) Unwrap() error {
	return e.cause
}

func (e *CryptomError) ErrType() ErrType {
	return e.errType
}

//
//// No errors
//func (e *CryptomError) valid() bool {
//	return e.Errors == 0
//}
//
//// Is checks if this CryptomError is of the supplied error. We are first checking for the exact error message
//// by comparing the inner error message. If that fails, we compare using the error flags. This way we can use
//// custom error messages (mainly for backwards compatability) and still leverage errors.Is using the global error variables.
//func (e *CryptomError) Is(err error) bool {
//	// Check, if our inner error is a direct match
//	if errors.Is(errors.Unwrap(e), err) {
//		return true
//	}
//
//	// Otherwise, we need to match using our error flags
//	switch err {
//	case ErrTokenMalformed:
//		return e.Errors&ValidationErrorMalformed != 0
//	case ErrTokenUnverifiable:
//		return e.Errors&ValidationErrorUnverifiable != 0
//	case ErrTokenSignatureInvalid:
//		return e.Errors&ValidationErrorSignatureInvalid != 0
//	case ErrTokenInvalidAudience:
//		return e.Errors&ValidationErrorAudience != 0
//	case ErrTokenExpired:
//		return e.Errors&ValidationErrorExpired != 0
//	case ErrTokenUsedBeforeIssued:
//		return e.Errors&ValidationErrorIssuedAt != 0
//	case ErrTokenInvalidIssuer:
//		return e.Errors&ValidationErrorIssuer != 0
//	case ErrTokenNotValidYet:
//		return e.Errors&ValidationErrorNotValidYet != 0
//	case ErrTokenInvalidId:
//		return e.Errors&ValidationErrorId != 0
//	case ErrTokenInvalidClaims:
//		return e.Errors&ValidationErrorClaimsInvalid != 0
//	}
//
//	return false
//}

package error_wrapper

import "fmt"

func (e *ErrorWrapper) Set(data ...interface{}) *ErrorWrapper {
	for _, row := range data {
		e.Data = append(e.Data, row)
	}
	return e

}

func (e *ErrorWrapper) Get() string {
	var temp string
	if len(e.Data) > 0 {
		for _, row := range e.Data {
			temp += fmt.Sprintf(`%v. `, row)
		}
		e.tempMessage = fmt.Sprintf(`%v. data: %v`, e.tempMessage, e.Data)
	}

	if e.IsMasked {
		e.Message = [2]string{e.MaskedMessage, e.tempMessage}
	} else {
		e.Message = [2]string{e.tempMessage, e.tempMessage}
	}

	if e.IsMasked {
		return fmt.Sprintf(`%s (%d)`, e.MaskedMessage, e.ErrorCode)
	} else {
		return fmt.Sprintf(`%s (%d)`, e.Message[0], e.ErrorCode)
	}
}

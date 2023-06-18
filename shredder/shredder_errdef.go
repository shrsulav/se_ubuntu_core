package shredder

import "fmt"

type ShredErrCode int

const (
	ShredErrProcessing 		ShredErrCode = -5
	ShredErrFileNotExist 	ShredErrCode = -4
	ShredErrNoExecutePerm 	ShredErrCode = -3
	ShredErrNotAFile 		ShredErrCode = -2
	ShredErrFileOpen 		ShredErrCode = -1
	ShredErrSuccess 		ShredErrCode = 0
	ShredErrFileWrite		ShredErrCode = 1
	ShredErrFileDelete 		ShredErrCode = 2
)

func (errCode ShredErrCode) ShredErrString() string {
	switch errCode {
	case ShredErrFileNotExist:
		return "file does not exist"
	case ShredErrNoExecutePerm:
		return "no execute permission in parent directory"
	case ShredErrNotAFile:
		return "given path is not a file, but a directory"
	case ShredErrFileOpen:
		return "file open error"
	case ShredErrSuccess:
		return "shredding successful"
	case ShredErrFileWrite:
		return "error while writing to file"
	case ShredErrFileDelete:
		return "error while deleting the file"
	}
	return "Unknown ShredErrorCode"
}

type ShredderError struct {
	ErrMessage 	string
	ErrCode  	ShredErrCode
}

func (w *ShredderError) Error() string {
	return fmt.Sprintf("%s: %v", w.ErrMessage, w.ErrCode)
}

func ReturnInfo(code ShredErrCode, msg string) *ShredderError {
	return &ShredderError{
		ErrMessage: msg,
		ErrCode:    code,
	}
}

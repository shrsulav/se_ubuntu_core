package shredder

import "fmt"

type ShredErrCode int

const (
	ShredErrProcessing 		ShredErrCode = -7
	ShredErrFileNotExist 	ShredErrCode = -6
	ShredErrNoExecutePerm 	ShredErrCode = -5
	ShredErrNotAFile 		ShredErrCode = -4
	ShredErrFileOpen 		ShredErrCode = -3
	ShredErrFileWrite		ShredErrCode = -2
	ShredErrFileDelete 		ShredErrCode = -1
	ShredErrSuccess 		ShredErrCode = 0
)

func (errCode ShredErrCode) ShredErrString() string {
	switch errCode {
	case ShredErrProcessing:
		return "error while processing the file"
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

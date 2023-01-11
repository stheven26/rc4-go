package helpers

import (
	"errors"
	"fmt"
	"mime/multipart"
	"strings"
)

var (
	fileTypes = []string{"pdf", "doc"}
)

func CheckFileType(filename string, filetypes []string) bool {
	filenameSplit := strings.Split(filename, ".")
	fileType := filenameSplit[len(filenameSplit)-1]
	for _, allowedType := range filetypes {
		if fileType == allowedType {
			return true
		}
	}
	return false
}

func ValidateFileHeader(fileHeader *multipart.FileHeader) (err error) {
	filenameSplit := strings.Split(fileHeader.Filename, ".")
	fileType := filenameSplit[len(filenameSplit)-1]
	for _, allowedType := range fileTypes {
		if fileType == allowedType {
			if (fileType == "pdf" || fileType == "doc") && fileHeader.Size > 10<<20 { // 10mb for pdf & doc
				err = errors.New(fmt.Sprintf("%s is too large", fileHeader.Filename))
			}
			return
		}
	}
	err = errors.New(fmt.Sprintf("invalid filetype for %s", fileHeader.Filename))
	return
}

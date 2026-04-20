package usecases

import (
	"errors"
	"slices"
	"strings"

	"github.com/google/uuid"
)

type FileMeta struct{
	Filename 	string
	Size 		int64
	ContentType string
}

type UploadResult struct{
	Filename 	string
	Url 		string
}

func UploadCase(file FileMeta) (UploadResult, error) {

	if file.Size > 5*1024*1024 {
		return UploadResult{}, errors.New("file size exceeds 5mb")
	}

	allowedTypes := []string{"image/jpeg", "image/png", "image/jpg"}
	isAllowed := slices.Contains(allowedTypes, file.ContentType)

	if !isAllowed{
		return UploadResult{}, errors.New("file type not allowed")
	}

	extSplit := strings.Split(file.Filename, ".")
	ext := extSplit[len(extSplit)-1]

	newFileName := uuid.New().String() + "." + ext

	return UploadResult{
		Filename: newFileName,
		Url: "/uploads/" + newFileName,
	}, nil
}
package utils

import (
	"io"
	"mime/multipart"
	_ "mime/multipart"
	"os"
	"path/filepath"
)

//多文件上传
func UploadFiles(files []*multipart.FileHeader, dest string) (dests string) {
	failures := 0
	for _, file := range files {
		resPath, err := saveUploadedFile(file, dest)
		if err != nil {
			failures++
			return ""
		} else {
			dests = dests + resPath + ";"
		}
	}
	return
}

//单文件上传
func UploadFile(file multipart.File, info *multipart.FileHeader, dest string) (res string) {
	defer file.Close()
	uploadRes, err := saveUploadedFile(info, dest)
	if err != nil {
		return ""
	} else {
		res = uploadRes
	}
	return
}

//多文件存储
func saveUploadedFile(fh *multipart.FileHeader, destDirectory string) (string, error) {
	src, err := fh.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	resPath := filepath.Join(destDirectory, fh.Filename)
	out, err := os.OpenFile(resPath,
		os.O_WRONLY|os.O_CREATE, os.FileMode(0666))
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, resErr := io.Copy(out, src)
	if resErr == nil {
		return resPath, nil
	} else {
		return "", nil
	}

}


package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

var (
	filePath, filetype string
)

func main() {
	fmt.Printf("请输入要上出传文件的文件绝对路径: ")
	fmt.Scanln(&filePath) //Scanln 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行。
	if Exists(filePath) {
		f, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		//获取类型
		contentType, err := GetFileContentType(f)
		if err != nil {
			panic(err)
		}
		if contentType == "image/gif" || contentType == "image/jpg" || contentType == "image/png" || contentType == "image/jpeg" {
			// fileType = "img"

			fmt.Printf("支持服务商有:Ali,Sogou,Vimcn,Niupic,Juejin,UploadLiu,Catbox,NetEasy,Prnt,Qihoo,Souhu,Toutiao,Xiaomi,ImgTg\n多个服务商输入:Ali,Sogou\n请输入云服务商:")
			fmt.Scanln(&filetype)
			extraParams := map[string]string{
				"type": filetype,
			}
			request, err := newfileUploadRequest("http://v1.alapi.cn/api/image", extraParams, "image", filePath)
			if err != nil {
				log.Fatal(err)
			}
			client := &http.Client{}
			resp, err := client.Do(request)
			if err != nil {
				log.Fatal(err)
			} else {
				body := &bytes.Buffer{}
				_, err := body.ReadFrom(resp.Body)
				if err != nil {
					log.Fatal(err)
				}
				resp.Body.Close()
				fmt.Println(resp.StatusCode)
				fmt.Println(resp.Header)
				fmt.Println(body)
			}
		} else {
			fmt.Printf("%s 不是图片类型", filePath)

		}
	} else {
		fmt.Printf("%s 文件不存在", filePath)
	}
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 获取类型
func GetFileContentType(out *os.File) (string, error) {
	// 只需要前 512 个字节就可以了
	buffer := make([]byte, 512)
	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}
	contentType := http.DetectContentType(buffer)
	return contentType, nil
}

// Creates a new file upload http request with optional extra params
func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}

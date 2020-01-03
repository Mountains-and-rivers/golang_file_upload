# golang 2种文件上传方式


# octet-stream方式上传！文件：golang_file_upload\octet-stream_upload.go   
请求头添加 binary/octet-stream
```
	file, err := os.Open("./filename")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	res, err := http.Post("http://127.0.0.1:5050/upload", "binary/octet-stream", file)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	message, _ := ioutil.ReadAll(res.Body)
	fmt.Printf(string(message))
}  
```


# multipart【表单方式，携带key value】方式上传！文件：golang_file_upload\octet-stream_upload.go   
借助了 开源库： mime/multipart  
```
func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}
	file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, fi.Name())
	if err != nil {
		return nil, err
	}
	part.Write(fileContents)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	return http.NewRequest("POST", uri, body)
}

func main() {
	// path, _ := os.Getwd()
	// path += "/test.test"
	extraParams := map[string]string{
		"keykeykeykeykeykeykey":       "valuevaluevaluevaluevaluevaluevaluevaluevalue",
	}
	request, err := newfileUploadRequest("http://localhost:5050/upload", extraParams, "file", "./filename")
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		var bodyContent []byte
		fmt.Println(resp.StatusCode)
		fmt.Println(resp.Header)
		resp.Body.Read(bodyContent)
		resp.Body.Close()
		fmt.Println(bodyContent)
	}
}  
```

# 接收文件服务器 receiveFileServer.go

package wwabago

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path"
)

const (
	GraphBaseAPI = "https://graph.facebook.com/v22.0/%s/%s"
)
const (
	logcredentialserror = "whatsapp credentials cannot be empty"
)


func GetFileContentType(output *os.File) (string, error) {
    // to sniff the content type only the first 512 bytes are used
    buf := make([]byte, 512)
    _, err := output.Read(buf)
    if err != nil {
        return "", err
    }

    // the function that actually does the trick
    contentType := http.DetectContentType(buf)
    return contentType, nil
}

func CreateFormData(form map[string]string, fileFieldName, filePath, mimeType string) (string, io.Reader, error) {
	body := new(bytes.Buffer)
	mp := multipart.NewWriter(body)

	// Add text fields
	for key, val := range form {
		if key != fileFieldName {
			_ = mp.WriteField(key, val)
		}
	}

	// Add file with custom content-type
	file, err := os.Open(filePath)
	if err != nil {
		return "", nil, err
	}
	defer file.Close()

	fileWriter, err := mp.CreatePart(
		textproto.MIMEHeader{
			"Content-Disposition": []string{`form-data; name="` + fileFieldName + `"; filename="` + path.Base(filePath) + `"`},
			"Content-Type":        []string{mimeType},
		})
	if err != nil {
		return "", nil, err
	}

	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return "", nil, err
	}

	mp.Close()

	return mp.FormDataContentType(), body, nil
}

func ReadResponseBody(response *http.Response) ([]byte, error) {

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func UploadWabaFile(w *Wwaba, contentType string, body io.Reader) *MediaResponse{


	graphApi := fmt.Sprintf(GraphBaseAPI, w.PhoneID, "media")
	authorization := fmt.Sprintf("Bearer %s", w.Authorization)


	// Create a new HTTP request
	req, err := http.NewRequest("POST", graphApi, body)
	req.Header.Add("Authorization", authorization)
	
	if err != nil {
		fmt.Printf("%v",err)
	}
		
	// Set the content type and authorization header
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	  
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("%v",err)
	}


	defer response.Body.Close()

	
	responseBody , err := ReadResponseBody(response)
	if err != nil {
		fmt.Printf("%v",err)
	}
   
	

   mediaResponse := MediaResponse{}
   err = json.Unmarshal(responseBody, &mediaResponse)
   if err != nil {
	fmt.Printf("%v",err)
   }


   return &mediaResponse
}

package wwabago

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)


func SendTextMessage(wwaba *Wwaba, msg interface{}) (*WhatsappMessageSuccess, error) {
	
	message, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	print(string(message))

	graphApi := fmt.Sprintf(GraphBaseAPI, wwaba.PhoneID, "messages")
	authorization := fmt.Sprintf("Bearer %s", wwaba.Authorization)

	req, err := http.NewRequest(http.MethodPost, graphApi, bytes.NewBuffer(message))
	if err != nil {
		fmt.Print("25")
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authorization)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("failed to read response body: %v", err)
			return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}

		var wppMessageError WhatsappMessageError
		if err := json.Unmarshal(body, &wppMessageError); err != nil {
			return nil, fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, string(body))
		}

		return nil, fmt.Errorf("error from API: %s", wppMessageError.Error.Message)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}


	var responseMessage WhatsappMessageSuccess

	err = json.Unmarshal(body, &responseMessage)

	if err != nil {
		log.Println("faild to unmarshal image reponse.")
	}


	return &responseMessage, nil
}

func SendImageMessage(wwaba *Wwaba, imageMsg *ImageMessage) (*WhatsappMessageSuccess, error){
	

	_, err := os.Stat(imageMsg.File.Path)

	if  err != nil {
		return nil, fmt.Errorf("caminho do arquivo: %s não encontrado", err)
	}
	
	file, err := os.Open(imageMsg.File.Path)
	if err != nil {
		return nil, fmt.Errorf("image file not found: %v", err)
	}

	mimeType, err  := GetFileContentType(file)
	if err != nil {
		return nil, fmt.Errorf("erro ao detectar conteúdo da imagem: %v", err)
	}

		
	formFields  := map[string]string{
		"messaging_product": imageMsg.MessagingProduct,
		"type":              "image", // ou o typeMessage que estiver usando
	}


	contentType, body, err := CreateFormData(formFields, "file",imageMsg.File.Path,mimeType)
	if err != nil {
		return nil, fmt.Errorf("erro ao montar cabeçalho %v", err)
	}

	imageResponse := UploadWabaFile(wwaba, contentType, body)

	imageMsg.Image = &ImageInfo{
		Id: imageResponse.ID,
		Caption: imageMsg.File.Caption,
	}
	
	imageMsg.File = nil

	jsonData, err := json.MarshalIndent(imageMsg, "", "  ")
	if err != nil {
		return nil, err
	}


	graphApi := fmt.Sprintf(GraphBaseAPI, wwaba.PhoneID, "messages")

	authorization := fmt.Sprintf("Bearer %s", wwaba.Authorization)

	req, err := http.NewRequest("POST", graphApi, bytes.NewBuffer(jsonData))


	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authorization) // Substitua pelo seu token de acesso
	
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to send message: %s", response.Status)
	}

	responseBody, err := ReadResponseBody(response)
	if err != nil {
		return nil, err
	}

	var responseMessage WhatsappMessageSuccess

	err = json.Unmarshal(responseBody, &responseMessage)

	if err != nil {
		log.Println("faild to unmarshal image reponse.")
	}


	return &responseMessage, nil
}
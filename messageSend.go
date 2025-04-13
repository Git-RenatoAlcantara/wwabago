package wwabago

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SendTextMessage(wwaba *Wwaba, msg *MessageConfig) (string, error) {
	message, err := json.MarshalIndent(msg, "", " ")

	if err != nil {
		return "", fmt.Errorf("failed to marshal message: %w", err)
	}

	buff := bytes.NewBuffer(message)

	graphApi := fmt.Sprintf(GraphBaseAPI,  wwaba.PhoneID, "messages")
	
	authorization := fmt.Sprintf("Bearer %s", wwaba.Authorization)

	// implement the logic to send the message
	resp, err := http.NewRequest(http.MethodPost, graphApi, buff)
	if err != nil {
		return "", err
	}

	resp.Header.Add("Content-Type", "application/json")
	resp.Header.Add("Authorization", authorization)


	response, err := http.DefaultClient.Do(resp)

	
	//Handle Error
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}
	defer response.Body.Close()
	//Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}
	sb := string(body)
	log.Println(sb)
	log.Printf("Status: %d", response.StatusCode)

	return "", nil
}
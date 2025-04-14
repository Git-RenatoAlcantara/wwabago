package wwabago

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


func SendTextMessage(wwaba *Wwaba, msg *MessageConfig) (string, error) {
	
	message, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}

	graphApi := fmt.Sprintf(GraphBaseAPI, wwaba.PhoneID, "messages")
	authorization := fmt.Sprintf("Bearer %s", wwaba.Authorization)

	req, err := http.NewRequest(http.MethodPost, graphApi, bytes.NewBuffer(message))
	if err != nil {
		fmt.Print("25")
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authorization)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("failed to read response body: %v", err)
			return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}

		var wppMessageError WhatsappMessageError
		if err := json.Unmarshal(body, &wppMessageError); err != nil {
			return "", fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, string(body))
		}

		return "", fmt.Errorf("error from API: %s", wppMessageError.Error.Message)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}


package wwabago

import (
	"bytes"
	"encoding/json"
	"errors"
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
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authorization)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return "", errors.New(string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
package wwabago

import (
	"context"
	"errors"
)


type Wwaba struct {
	Authorization string
	PhoneID string
}


func  NewWwaba(authorization, phoneID string) (*Wwaba, error){
	return CreateWwaba(authorization, phoneID)
}


func CreateWwaba(authorization, phoneID string) (*Wwaba, error){
	if authorization == "" || phoneID == ""{
		return nil, errors.New(logcredentialserror)
	}
	wwaba := &Wwaba{
		Authorization: authorization,
		PhoneID: phoneID,
	}

	return wwaba, nil
}


func (wwaba *Wwaba) Send(ctx context.Context, messageBase interface{}) (*WhatsappMessageSuccess, error) {

	switch msg := messageBase.(type) {
	case TextMessage:
		return SendTextMessage(ctx, wwaba, &msg)
	case ImageMessage:
		return SendImageMessage(wwaba, &msg)
	case VideoMessage:
		return SendVideoMessage(wwaba, &msg)
	default:
		return nil, errors.New("unsupported message type")
	}
}
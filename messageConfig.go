package wwabago

const (
    MessagingProductWhatsapp = "whatsapp"
    RecipientTypeIndividual = "individual"
    TypeText = "text"
    TypeImage = "image"
)
// chatID is where to send it, text is the message text.
func NewMessage(phone_number string, text string) TextMessage {
	return 	TextMessage{
        MessagingProduct: MessagingProductWhatsapp,
        RecipientType:    RecipientTypeIndividual,
        To:               phone_number,
        Type:             TypeText,
        Text: &TextInfo{
            PreviewURL: false,
            Body:       text,
        },
    }
}

func NewImageMessage(phone_number, path,caption string) ImageMessage{
    return ImageMessage{
        MessagingProduct: MessagingProductWhatsapp,
        RecipientType:    RecipientTypeIndividual,
        To:               phone_number,
        Type:             TypeImage,
        File: &FileConfig{
            Path:    path,
            Caption: caption,
        },
    }
}
package wwabago

// chatID is where to send it, text is the message text.
func NewMessage(phone_number string, text string) MessageConfig {
	return MessageConfig{
	    MessagingProduct: "whatsapp",
        RecipientType:    "individual",
        To:               phone_number, // certifique-se que o número está registrado no WABA
        Type:             "text",
        BaseChat: BaseConfig{
            PreviewURL: true,
            Body:       text,
        },
	}
}
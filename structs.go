package wwabago

// content represents the content of a message.
type BaseConfig struct {
	PreviewURL bool   `json:"preview_url"`
	Body       string `json:"body"`
}

// message represents the structure of a message to be sent.
type MessageConfig struct {
	MessagingProduct string   `json:"messaging_product"`
	RecipientType    string   `json:"recipient_type"`
	To               string   `json:"to"`
	Type             string   `json:"type"`
	BaseChat         BaseConfig `json:"text,omitempty"`
}
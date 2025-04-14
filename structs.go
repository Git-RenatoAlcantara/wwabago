package wwabago

// waba represents the WhatsApp Business API credentials.
type waba struct {
	Authorization string
	PhoneID       string
}

// MessageInfo holds the recipient's information.
type MessageContact struct {
	Number string `json:"to"`
}

// TextMessage represents a text message to be sent.
type TextConfig struct {
	Info      MessageContact
	PreviewURL bool `json:"preview_url"`
	Body      string `json:"body"`
}

// content represents the content of a message.
type TextInfo struct {
	PreviewURL bool   `json:"preview_url"`
	Body       string `json:"body"`
}

// message represents the structure of a message to be sent.
type TextMessage struct {
	MessagingProduct string   `json:"messaging_product"`
	RecipientType    string   `json:"recipient_type"`
	To               string   `json:"to"`
	Type             string   `json:"type"`
	Text             *TextInfo `json:"text,omitempty"`
}

// message represents the structure of a message to be sent.
type ImageMessage struct {
	MessagingProduct string   `json:"messaging_product"`
	RecipientType    string   `json:"recipient_type"`
	To               string   `json:"to"`
	Type             string   `json:"type"`
	File             *FileConfig    `json:"file,omitempty"`
	Image            *ImageInfo   `json:"image,omitempty"`
}



// image represents an image to be sent.
type ImageInfo struct {
	Id      string `json:"id"`
	Caption string `json:"caption"`
}

// ImageResponse represents the response structure for an image.
type ImageResponse struct {
	ID string `json:"id"`
}

// file represents a file to be sent.
type FileConfig struct {
	Path    string `json:"path"`
	Caption string `json:"caption"`
}

// ImageMessage represents an image message to be sent.
type ImageComplete struct {
	Info    MessageContact
	Path    string
	Caption string
}


type WhatsappMessageError struct {
	Error struct {
		Message      string `json:"message"`
		Type         string `json:"type"`
		Code         int    `json:"code"`
		ErrorSubcode int    `json:"error_subcode"`
		FbtraceID    string `json:"fbtrace_id"`
	} `json:"error"`
}

type WhatsappMessageSuccess struct {
	MessagingProduct string `json:"messaging_product"`
	Contacts         []struct {
		Input string `json:"input"`
		WaID  string `json:"wa_id"`
	} `json:"contacts"`
	Messages []struct {
		ID string `json:"id"`
	} `json:"messages"`
}
package wwabago

import (
	"testing"
)

const (
	TestAuthentication = "EAAH36KWr86sBO3SFL2LNebrFRrhAiavvNkCbmZBrJqR1hvjnKTPo4IFUYex15uUsgUeqjUsJi37lTEwg2yGSexFAIIDfrT9JHc0cvpDJcZAGXnsuHIV9ZAN1ofd4yM2QSJmitfXpZCLYBWSsaGE53PTEP8oYA46ZAk0xAioPFQbgI6paUvjIbF2gbJlZBquoJLcVRVlJ0Kk8EKiTawjmoipE3jiNpoowZDZD"
	TestePhoneID = "495746173617776"
)

type testLogger struct {
	t *testing.T
}

func (t testLogger) Println(v ...interface{}) {
	t.t.Log(v...)
}

func (t testLogger) Printf(format string, v ...interface{}) {
	t.t.Logf(format, v...)
}

func getClient(t *testing.T) (*Wwaba, error){

	wwaba, err :=	CreateWwaba(TestAuthentication, TestePhoneID)
	
	logger := testLogger{t}
 	SetLogger(logger)

	if err != nil {
		t.Error(t)
	}

	return wwaba, err

}


func TestCreateWwaba_noCredentials(t *testing.T){

	_, err := CreateWwaba("", "")

	if err == nil {
		t.Error(err)
	}

}

func TestSendWithMessage(t *testing.T) {
	wwaba, _ := getClient(t)

	msg := NewMessage("5521967056425", "Mensagem de teste")
	_, err := wwaba.Send(msg)
	if err != nil {
		t.Error(err)
	}

}



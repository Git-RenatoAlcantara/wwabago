package wwabago

import (
	"context"
	"testing"
)

const (
	TestAuthentication = ""
	TestePhoneID = ""
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

	ctx := context.Background()

	msg := NewMessage("+00(00)00000-0000", "Mensagem de teste")
	_, err := wwaba.Send(ctx, msg)
	if err != nil {
		t.Error(err)
	}

}

func TestSendWithCancelledContext(t *testing.T) {
    wwaba, _ := getClient(t)

    ctx, cancel := context.WithCancel(context.Background())
    cancel() // Cancela o contexto imediatamente

    msg := NewMessage("+00(00)00000-0000", "Mensagem de teste")
    _, err := wwaba.Send(ctx, msg)
    if err == nil {
        t.Error("expected error due to cancelled context, got nil")
    }
}


func TestSendMessageWithImageFromPath(t *testing.T){
	wwaba, _ := getClient(t)

	ctx := context.Background()

	msg := NewImageMessage(
		"+00(00)00000-0000",
		"./image/node-js.png",
		"Image demo",
	)

	_, err := wwaba.Send(ctx, msg)
	if err != nil {
		t.Error(err)
	}
}

func TestSendMessageWithVideoFromPath(t *testing.T){
	wwaba, _ := getClient(t)

	ctx := context.Background()

	msg := NewVideoMessage(
		"+00(00)00000-0000", 
		"./video/demo.mp4",
		"Video demo",
	)
	
	_, err := wwaba.Send(ctx, msg)
	if err != nil {
		t.Error(err)
	}
}




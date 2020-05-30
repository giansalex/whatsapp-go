package cl

import (
	"encoding/gob"
	"fmt"
	"os"
	"os/signal"
	"time"

	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/Rhymen/go-whatsapp"
)

// WhatsappClient connect
type WhatsappClient struct {
	wac *whatsapp.Conn
}

// NewClient create whatsapp client
func NewClient() *WhatsappClient {
	wac, err := whatsapp.NewConn(8 * time.Second)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating connection: %v\n", err)
		os.Exit(1)
		return nil
	}

	wac.SetClientName("Giansalex Whatsapp Web API", "Whatsapp API")
	wac.SetClientVersion(0, 4, 1307)
	err = login(wac)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error logging in: %v\n", err)
		os.Exit(1)
		return nil
	}

	fmt.Println("Whatsapp Connected!")
	return &WhatsappClient{wac}
}

// Listen text messages
func (wp *WhatsappClient) Listen(f messageListener) {
	wp.wac.AddHandler(&messageHandler{f, time.Now().Unix()})

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)

	fmt.Println("Press ctrl+c to exit.")

	<-sigs
	fmt.Println("Shutdown.")
	os.Exit(0)
}

// SendText send text message
func (wp *WhatsappClient) SendText(to string, text string) {
	reply := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: to,
		},
		Text: text,
	}

	msgID, err := wp.wac.Send(reply)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error sending message: %v", err)
	} else {
		fmt.Println("Message Sent -> ID : " + msgID)
	}
}

// GetConnection return whatsapp connection
func (wp *WhatsappClient) GetConnection() *whatsapp.Conn {
	return wp.wac
}

func login(wac *whatsapp.Conn) error {
	//load saved session
	session, err := readSession()
	if err == nil {
		//restore session
		session, err = wac.RestoreWithSession(session)
		if err != nil {
			return fmt.Errorf("restoring failed: %v", err)
		}
	} else {
		//no saved session -> regular login
		qr := make(chan string)
		go func() {
			terminal := qrcodeTerminal.New()
			terminal.Get(<-qr).Print()
		}()
		session, err = wac.Login(qr)
		if err != nil {
			return fmt.Errorf("error during login: %v", err)
		}
	}

	//save session
	err = writeSession(session)
	if err != nil {
		return fmt.Errorf("error saving session: %v", err)
	}
	return nil
}

func readSession() (whatsapp.Session, error) {
	session := whatsapp.Session{}
	file, err := os.Open(os.TempDir() + "/whatsappSession.gob")
	if err != nil {
		return session, err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&session)
	if err != nil {
		return session, err
	}
	return session, nil
}

func writeSession(session whatsapp.Session) error {
	file, err := os.Create(os.TempDir() + "/whatsappSession.gob")
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(session)
	if err != nil {
		return err
	}
	return nil
}

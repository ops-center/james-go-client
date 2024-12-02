package inbox

import (
	"os"
	"testing"
)

func TestSendEmail(t *testing.T) {
	os.Setenv("INBOX_SERVER_TOKEN", "<jwt token here>")

	jmapClient, err := NewJMAPClient(
		&JMAPConf{
			JMAPSessionEndpoint: "http://192.168.0.62:80/jmap/session",
		},
	)
	if err != nil {
		t.Fatalf("failed to create jmap client : %v", err)
	}

	t.Log("Created jmap client")

	emailOptions := []Option{
		WithSubject("hehe2"),
		WithTextBody(" "),
		WithRecipients([]string{"recipient.acc@cloud.appscode.com"}),
		WithReferences([]string{"hehe2"}),
		WithInReplyTo([]string{"hehe"}),
	}

	newEmail, err := jmapClient.NewEmail(emailOptions...)
	if err != nil {
		t.Fatalf("failed to create new email : %v", err)
	}

	if err = jmapClient.SendEmail(newEmail); err != nil {
		t.Fatalf("failed to send email : %v", err)
	}
}

package mail

import (
	"testing"

	"github.com/mikcatta/simple_bank/util"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	config, err := util.LoadConfig("../")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "A test email"
	content := `
	<h1>Hello World</h1>
	<p>
	  this is a test message from <a href="http://techschool.guru">Tech School</a>
	</p>`

	to := []string{"techschool.guru@gmail.com"}
	cc := []string{"mike_cat_taylor@hotmail.com"}
	attachFiles := []string{"../README.md"}

	err = sender.SendEmail(subject, content, to, cc, nil, attachFiles)
	require.NoError(t, err)

}

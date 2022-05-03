package main

/*
	https://zhuanlan.zhihu.com/p/446124221
	https://zhuanlan.zhihu.com/p/357556162
*/

import (
	"github.com/emersion/go-imap"
	id "github.com/emersion/go-imap-id"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/charset"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	SimpleUsage()
}

func CustomerImapClient() (*client.Client, error) {
	return NewImapClient("ycliu912@163.com", "XWOLXSWHQYVQTXEN")
}

func NewImapClient(username, password string) (*client.Client, error) {
	imap.CharsetReader = charset.Reader

	log.Println("Connecting to server...")

	//c, err := client.DialTLS("smtp.exmail.qq.com:993", nil)
	c, err := client.DialTLS("imap.163.com:993", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	idClient := id.NewClient(c)
	idClient.ID(
		id.ID{id.FieldName: "IMAPClient", id.FieldVersion: "1.2.0"},
	)

	if err := c.Login(username, password); err != nil {
		return nil, err
	}

	log.Println("Logged in")

	return c, nil
}

func SimpleUsage() {
	c, err := CustomerImapClient()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Logout()

	//获取文件夹列表
	mailboxs := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.List("", "*", mailboxs)
	}()
	log.Println("邮件文件夹：")
	for m := range mailboxs {
		log.Println("* " + m.Name)
	}
	if err := <-done; err != nil {
		log.Fatal(err)
	}

	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}

	from := uint32(5)
	to := mbox.Messages
	if mbox.Messages > 5 {
		from = mbox.Messages - 5
	}
	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)

	messages := make(chan *imap.Message, 10)
	//done := make(chan error, 1)
	go func() {
		done <- c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope}, messages)
	}()

	for msg := range messages {
		log.Println("* " + msg.Envelope.Subject)
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}
}

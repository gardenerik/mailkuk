package main

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/emersion/go-smtp"
	"io"
	"time"
)

type Backend struct{}

func (bkd *Backend) NewSession(_ *smtp.Conn) (smtp.Session, error) {
	return &Session{}, nil
}

type Session struct {
	rcptTo string
}

func (s *Session) AuthPlain(_, _ string) error {
	return fmt.Errorf("auth is not supported")
}

func (s *Session) Mail(from string, opts *smtp.MailOptions) error {
	log.Debugf("Receiving mail from %s", from)
	return nil
}

func (s *Session) Rcpt(to string) error {
	log.Debugf("Addressed to %s", to)
	s.rcptTo = to
	return nil
}

func (s *Session) Data(r io.Reader) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	sender, exists := Router[s.rcptTo]
	if !exists {
		log.Debug("Ignoring mail to unknown address", "to", s.rcptTo)
		return nil
	}

	go func() {
		log.Infof("Received mail for %s.", s.rcptTo)
		err := sender.AcceptMail(data)
		if err != nil {
			log.Error("Error while accepting mail", "err", err)
		}
	}()

	return nil
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {
	return nil
}

var SMTPServer *smtp.Server

func startServer(cfg Server) error {
	be := &Backend{}
	SMTPServer = smtp.NewServer(be)
	SMTPServer.Addr = cfg.ListenAddr
	SMTPServer.Domain = cfg.Domain
	SMTPServer.ReadTimeout = 100 * time.Second
	SMTPServer.WriteTimeout = 10 * time.Second
	SMTPServer.MaxMessageBytes = 5242880 // 5 MB
	SMTPServer.AuthDisabled = true
	return SMTPServer.ListenAndServe()
}

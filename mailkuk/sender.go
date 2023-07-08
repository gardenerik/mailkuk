package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

type Sender struct {
	url     string
	headers map[string]string
}

func (s Sender) AcceptMail(data []byte) error {
	buf := bytes.NewReader(data)
	req, err := http.NewRequest("POST", s.url, buf)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "message/rfc822")
	for h, v := range s.headers {
		req.Header.Add(h, v)
	}

	client := http.Client{
		Timeout: 15 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("http server responded with status %d", resp.StatusCode)
	}

	return nil
}

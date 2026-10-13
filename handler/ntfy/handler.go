package ntfy

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/ushis/gesundheit/handler"
	"github.com/ushis/gesundheit/result"
)

type Handler struct {
	Url      string
	Token    string
	Topic    string
	Priority int
}

func init() {
	handler.RegisterSimple("ntfy", New)
}

func New(configure func(interface{}) error) (handler.Simple, error) {
	h := Handler{Priority: 3}

	if err := configure(&h); err != nil {
		return nil, err
	}
	if h.Topic == "" {
		return nil, errors.New("missing topic")
	}
	u, err := url.Parse(h.Url)

	if err != nil {
		return nil, err
	}
	u.Path = strings.TrimSuffix(u.Path, "/") + "/" + h.Topic
	h.Url = u.String()
	return h, nil
}

func (h Handler) Handle(e result.Event) error {
	req, err := http.NewRequest("POST", h.Url, strings.NewReader(e.Message))

	if err != nil {
		return err
	}
	req.Header.Set("Title", fmt.Sprintf("(%s) %s %s", e.NodeName, e.CheckDescription, e.Status))
	req.Header.Set("Priority", fmt.Sprintf("%d", h.Priority))
	req.Header.Set("Content-Type", "text/plain")

	if e.Status == result.StatusOK {
		req.Header.Set("Tags", "heavy_check_mark")
	} else {
		req.Header.Set("Tags", "warning")
	}

	if h.Token != "" {
		req.Header.Set("Authorization", "Bearer "+h.Token)
	}
	resp, err := (&http.Client{}).Do(req)

	if err != nil {
		return err
	}
	if resp.StatusCode/100 != 2 {
		return errors.New("unexpected response status: " + resp.Status)
	}
	return nil
}

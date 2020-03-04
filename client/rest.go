package client

import (
	"net/http"
	"time"
)

type InvalidResponse struct{}

func (InvalidResponse) Error() string {
	return "Invalid Response Status Code!"
}

type RESTConn struct {
	Conn *http.Client
}

func (c RESTConn) Close() error {
	c.Conn.CloseIdleConnections()
	return nil
}

func (c RESTConn) fetch(input string) (float64, error) {
	t := time.Now()
	response, err := c.Conn.Get(input)
	elapsed := time.Since(t)
	if err != nil{
		return 0, err
	} else if response.StatusCode >= 300{
		return 0, InvalidResponse{}
	}
	return float64(elapsed), nil
}

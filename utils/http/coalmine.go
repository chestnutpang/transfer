package http

import (
	"bytes"
	"cts/plugin"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CoalmineClient struct {
	header   map[string]string
	ip       string
	port     int
	protocol string
}

type CameraStatus struct {
	Company     string `json:"company"`
	Biz         string `json:"biz"`
	BizType     string `json:"biz_type"`
	Freq        int    `json:"freq"`
	FreqChannel string `json:"freq_channel"`
	Status      int    `json:"status"`
}

func (c *CoalmineClient) SetHeader(k, v string) {
	c.header[k] = v
}

func (c *CoalmineClient) Get(uri string, args map[string]string) {
	url := fmt.Sprintf("%s://%s:%d%s", c.protocol, c.ip, c.port, uri)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
}

func (c *CoalmineClient) Post(uri string, params io.Reader) {
	url := fmt.Sprintf("%s://%s:%d%s", c.protocol, c.ip, c.port, uri)
	req, err := http.NewRequest("POST", url, params)
	if err != nil {

	}
	for k, v := range c.header {
		req.Header.Set(k, v)
	}

}

func (c *CoalmineClient) CameraStatus(company, biz, bizType string, status int) {
	uri := ""
	body := CameraStatus{
		Company:     company,
		Biz:         biz,
		BizType:     bizType,
		Freq:        0,
		FreqChannel: "",
		Status:      status,
	}
	b, err := json.Marshal(body)
	if err != nil {

	}
	c.Post(uri, bytes.NewBuffer(b))
}

func (c *CoalmineClient) Step1(p plugin.IPluginParams) {
	fmt.Println("step1")
}

func (c *CoalmineClient) Step2(p plugin.IPluginParams) {
	fmt.Println("step2")
}

func (c *CoalmineClient) Step3(p plugin.IPluginParams) {
	fmt.Println("step3")
}

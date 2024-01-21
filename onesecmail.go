package onesecmail

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const baseUrl = "https://www.1secmail.com/api/v1/"

type EmailAddresses []string
type ActiveDomains []string
type Messages []Message

type Client struct {
	httpClient http.Client
}

type Message struct {
	Id      int    `json:"id"`
	From    string `json:"from"`
	Subject string `json:"subject"`
	Date    string `json:"date"`
}
type Attachment struct {
	Filename    string `json:"filename"`
	ContentType string `json:"contentType"`
	Size        int    `json:"size"`
}

type MessageDetailed struct {
	Id          int          `json:"id"`
	From        string       `json:"from"`
	Subject     string       `json:"subject"`
	Date        string       `json:"date"`
	Attachments []Attachment `json:"attachments"`
	Body        string       `json:"body"`
	TextBody    string       `json:"textBody"`
	HtmlBody    string       `json:"htmlBody"`
}

func NewClient() *Client {
	return &Client{
		httpClient: http.Client{},
	}
}

func (c *Client) GenerateRandomEmailAddresses(count int) (EmailAddresses, error) {
	var emails = EmailAddresses{}

	req, err := http.NewRequest("GET", baseUrl+"?action=genRandomMailbox&count="+strconv.Itoa(count), nil)
	if err != nil {
		return emails, errors.New("error creating new http request")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return emails, errors.New("error doing http request")
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return emails, errors.New("error reading response body")
	}
	res.Body.Close()

	err = json.Unmarshal(data, &emails)
	if err != nil {
		return emails, errors.New("error decoding json")
	}

	return emails, nil
}

func (c *Client) GetAllActiveDomains() (ActiveDomains, error) {
	var domains = ActiveDomains{}

	req, err := http.NewRequest("GET", baseUrl+"?action=getDomainList", nil)
	if err != nil {
		return domains, errors.New("error creating new http request")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return domains, errors.New("error doing http request")
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return domains, errors.New("error reading response body")
	}
	res.Body.Close()

	err = json.Unmarshal(data, &domains)
	if err != nil {
		return domains, errors.New("error decoding json")
	}

	return domains, nil
}

func (c *Client) CheckMailbox(address string) (Messages, error) {
	var messages = Messages{}
	parts := strings.Split(address, "@")
	if len(parts) != 2 {
		return messages, errors.New("error parsing address; it should be name@domain; use GenerateRandomEmailAddresses func")
	}
	req, err := http.NewRequest("GET", baseUrl+"?action=getMessages&login="+parts[0]+"&domain="+parts[1], nil)
	if err != nil {
		return messages, errors.New("error creating new http request")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return messages, errors.New("error doing http request")
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return messages, errors.New("error reading response body")
	}
	res.Body.Close()

	err = json.Unmarshal(data, &messages)
	if err != nil {
		return messages, errors.New("error decoding json")
	}

	return messages, nil
}

func (c *Client) ReadEmail(address string, id int) (MessageDetailed, error) {
	var message = MessageDetailed{}
	parts := strings.Split(address, "@")
	if len(parts) != 2 {
		return message, errors.New("error parsing address; it should be name@domain; use GenerateRandomEmailAddresses func")
	}
	req, err := http.NewRequest("GET", baseUrl+"?action=readMessage&login="+parts[0]+"&domain="+parts[1]+"&id="+strconv.Itoa(id), nil)
	if err != nil {
		return message, errors.New("error creating new http request")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return message, errors.New("error doing http request")
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return message, errors.New("error reading response body")
	}
	res.Body.Close()

	err = json.Unmarshal(data, &message)
	if err != nil {
		return message, errors.New("error decoding json")
	}

	return message, nil
}

func (c *Client) DownloadAttachmentUrl(address string, messageID int, filename string) (string, error) {
	var url = ""

	parts := strings.Split(address, "@")
	if len(parts) != 2 {
		return url, errors.New("error parsing address; it should be name@domain; use GenerateRandomEmailAddresses func")
	}

	url = baseUrl + "?action=download&login=" + parts[0] + "&domain=" + parts[1] + "&id=" + strconv.Itoa(messageID) + "&file=" + filename

	return url, nil
}

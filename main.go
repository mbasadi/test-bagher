package notification

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

var __client_id,__client_secret string;  
type EmailAttachments struct {
	Filename string `json:"filename"`
	Url string `json:"url"`
}
type SendRequestEmailOptions struct {
		ReplyToAddresses  []string `json:"replyToAddresses"`
		CcAddresses []string `json:"ccAddresses"`
		BccAddresses []string `json:"bccAddresses"`
		Attachments []EmailAttachments `json:"attachments"`
}
type SendRequestOptions struct {
	Email SendRequestEmailOptions `json:"email"`
}
type User struct{
	Id string `json:"id"`
	Email string `json:"email"`
	Number string `json:"number"`
  }

type SendRequest struct {
	NotificationId  string `json:"notificationId"`
	User  User `json:"user"`
	MergeTags  map[string]string `json:"mergeTags"`
	Replace map[string]string `json:"replace"`
	ForceChannels  []string `json:"forceChannels"`
	TemplateId  string `json:"templateId"`
	SubNotificationId string `json:"subNotificationId"`
	Options SendRequestOptions `json:"options"`
}
type RetractRequest struct {
	NotificationId  string `json:"notificationId"`
	UserId  string `json:"userId"`
}
type CreateSubNotificationRequest struct {
	NotificationId  string 
	Title  string 
	SubNotificationId  string 
}
type DeleteSubNotificationRequest struct {
	NotificationId  string 
	SubNotificationId  string 
}
type SetUserPreferencesRequest struct {
	NotificationId  string 
	Channel  string 
	State bool 
	SubNotificationId string 
}
func Init(client_id,client_secret string) error {
	if client_id == "" {
        return  errors.New("Bad client_id")
    }
	if client_secret == "" {
        return  errors.New("Bad client_secret")
    }
	__client_id = client_id
	__client_secret= client_secret
return nil

}
func basicAuth(client_id, client_secret string) string {
	auth := client_id + ":" + client_secret
	return base64.StdEncoding.EncodeToString([]byte(auth))
  }
func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}
func request(client *http.Client, method,uri string, data *bytes.Buffer) error{
	endpoint := "https://api.notificationapi.com/" + __client_id + "/" + uri
	req, err := http.NewRequest(method, endpoint, data)
	if err != nil {
		log.Fatalf("Error Occurred. %+v", err)
	}
	req.Header.Add("Authorization","Basic " + basicAuth(__client_id,__client_secret))
	response, err := client.Do(req)

	if err != nil {
		log.Fatalf("Error sending request to API endpoint. %+v", err)
	}

	// Close the connection to reuse it
	defer response.Body.Close()

	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
	}
	if response.StatusCode==202 {
		fmt.Printf("NotificationAPI request ignored.")
	}

	if response.StatusCode==500 {
		return errors.New("NotificationAPI request failed.")
	}

	return nil

}

func Send(params SendRequest) error{
	c := httpClient()
	sendRequest, err := json.Marshal(params)
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
	}
	return request(c, http.MethodPost, "sender",bytes.NewBuffer(sendRequest))
  }
  func Retract(params RetractRequest) error{
	c := httpClient()
	retractRequest, err := json.Marshal(params)
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
	}
	return request(c, http.MethodPost, "sender/retract",bytes.NewBuffer(retractRequest))
  }
  func CreateSubNotification(params CreateSubNotificationRequest) error{
	c := httpClient()
	createSubNotificationRequest, err := json.Marshal(map[string]string{ "title": params.Title })
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
	}
	return request(c, http.MethodPut,  "notifications/"+params.NotificationId+"/subNotifications/"+params.SubNotificationId,bytes.NewBuffer(createSubNotificationRequest))
  }
  



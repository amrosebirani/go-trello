package trello

import (
	"net/http"
	"log"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/url"
	"github.com/mrjones/oauth"
	"bytes"
)

type Webhook struct {
	ID          string   `json:"id"`
	Description string   `json:"description"`
	IDModel     string   `json:"idModel"`
	CallbackURL string   `json:"callbackURL"`
	Active      bool     `json:"active"`
}

type WebhookFactory struct {
	c *oauth.Consumer
	a *oauth.AccessToken
}

func NewWebhookFactory(c *oauth.Consumer, a *oauth.AccessToken) * WebhookFactory{
	return &WebhookFactory{
		c: c,
		a: a,
	}
}

func (w *WebhookFactory) CreateTrelloWebhook(callback_url string, id_model string, description string, webhookData *Webhook) error{
	form := url.Values{}
	form.Add("description", description)
	form.Add("callbackURL", callback_url)
	form.Add("idModel", id_model)
	queryData := form.Encode()
	req, err := http.NewRequest("POST", "https://trello.com/1/webhooks", bytes.NewBuffer([]byte(queryData)))
	if err != nil {
		log.Print(err)
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client, err := w.c.MakeHttpClient(w.a)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	//fmt.Println(resp.Status)
	defer resp.Body.Close()
	htmlData, err := ioutil.ReadAll(resp.Body) //<--- here!
	data := string(htmlData)
	fmt.Println(data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// print out

	err = json.Unmarshal(htmlData, webhookData)
	return err
}

func (w *WebhookFactory) DeleteTrelloWebhook () {

}

func (w *WebhookFactory) UpdateTrelloWebhook() {

}

func (w *WebhookFactory) GetAllTrelloWebhooks () {

}

func (w *WebhookFactory) GetTrelloWebhook () {

}
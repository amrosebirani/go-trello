package trello

import (
	"net/http"
	"log"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/url"
)

type TrelloWebhook struct {
	ID          string   `json:"id"`
	Description string   `json:"description"`
	IDModel     string   `json:"idModel"`
	CallbackURL string   `json:"callbackURL"`
	Active      bool     `json:"active"`
}

func CreateTrelloWebhook(callback_url string, id_model string, description string, webhookData *TrelloWebhook) error{
	req, err := http.NewRequest("POST", "https://trello.com/1/webhooks", nil)
	if err != nil {
		log.Print(err)
		return err
	}
	form := url.Values{}
	form.Add("description", description)
	form.Add("callbackURL", callback_url)
	form.Add("idModel", id_model)
	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
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

func DeleteTrelloWebhook () {

}

func UpdateTrelloWebhook() {

}

func GetAllTrelloWebhooks () {

}

func GetTrelloWebhook () {

}
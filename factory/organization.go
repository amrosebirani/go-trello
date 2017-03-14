package factory

import (
	"github.com/mrjones/oauth"
	"net/url"
	"net/http"
	"bytes"
	"log"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

type Organization struct {
	Id          string         `json:"id"`
	Name        string         `json:"name"`
	DisplayName string         `json:"displayName"`
	Desc        string         `json:"desc"`
	DescData    interface{}    `json:"descData"`
	Url         string         `json:"url"`
	Website     string         `json:"website"`
	LogoHash    string         `json:"logoHash"`
	Products    []string       `json:"products"`
	PowerUps    []string       `json:"powerUps"`
}

type OrganizationFactory struct {
	c *oauth.Consumer
	a *oauth.AccessToken
}

func NewOrganizationFactory(c *oauth.Consumer, a *oauth.AccessToken) * OrganizationFactory{
	return &OrganizationFactory{
		c: c,
		a: a,
	}
}

func (o *OrganizationFactory) CreateTrelloOrganization (name string, description string, orgData *Organization) error{
	form := url.Values{}
	form.Add("desc", description)
	form.Add("name", name)
	form.Add("displayName", name)
	queryData := form.Encode()
	req, err := http.NewRequest("POST", "https://trello.com/1/organizations", bytes.NewBuffer([]byte(queryData)))
	if err != nil {
		log.Print(err)
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client, err := o.c.MakeHttpClient(o.a)
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

	err = json.Unmarshal(htmlData, orgData)
	return err

}
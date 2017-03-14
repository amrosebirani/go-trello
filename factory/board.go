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

type Board struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	DescData struct {
		Emoji struct{} `json:"emoji"`
	} `json:"descData"`
	Closed         bool   `json:"closed"`
	IdOrganization string `json:"idOrganization"`
	Pinned         bool   `json:"pinned"`
	Url            string `json:"url"`
	ShortUrl       string `json:"shortUrl"`
	Prefs          struct {
		PermissionLevel       string            `json:"permissionLevel"`
		Voting                string            `json:"voting"`
		Comments              string            `json:"comments"`
		Invitations           string            `json:"invitations"`
		SelfJoin              bool              `json:"selfjoin"`
		CardCovers            bool              `json:"cardCovers"`
		CardAging             string            `json:"cardAging"`
		CalendarFeedEnabled   bool              `json:"calendarFeedEnabled"`
		Background            string            `json:"background"`
		BackgroundColor       string            `json:"backgroundColor"`
		BackgroundImage       string            `json:"backgroundImage"`
		BackgroundImageScaled []BoardBackground `json:"backgroundImageScaled"`
		BackgroundTile        bool              `json:"backgroundTile"`
		BackgroundBrightness  string            `json:"backgroundBrightness"`
		CanBePublic           bool              `json:"canBePublic"`
		CanBeOrg              bool              `json:"canBeOrg"`
		CanBePrivate          bool              `json:"canBePrivate"`
		CanInvite             bool              `json:"canInvite"`
	} `json:"prefs"`
	LabelNames struct {
		Red    string `json:"red"`
		Orange string `json:"orange"`
		Yellow string `json:"yellow"`
		Green  string `json:"green"`
		Blue   string `json:"blue"`
		Purple string `json:"purple"`
	} `json:"labelNames"`
}

type BoardBackground struct {
	width  int    `json:"width"`
	height int    `json:"height"`
	url    string `json:"url"`
}


type BoardFactory struct {
	c *oauth.Consumer
	a *oauth.AccessToken
}

func NewBoardFactory(c *oauth.Consumer, a *oauth.AccessToken) * BoardFactory{
	return &BoardFactory{
		c: c,
		a: a,
	}
}

func (b *BoardFactory) CreateTrelloBoard (name string, description string, trello_org_id string, boardData *Board) error{
	form := url.Values{}
	form.Add("desc", description)
	form.Add("name", name)
	form.Add("defaultLists", "false")
	form.Add("idOrganization", trello_org_id)
	queryData := form.Encode()
	req, err := http.NewRequest("POST", "https://trello.com/1/organizations", bytes.NewBuffer([]byte(queryData)))
	if err != nil {
		log.Print(err)
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client, err := b.c.MakeHttpClient(b.a)
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

	err = json.Unmarshal(htmlData, boardData)
	return err

}
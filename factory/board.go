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
	IdOrganization string `json:"idOrganization" bson:"idOrganization"`
	Pinned         bool   `json:"pinned"`
	Url            string `json:"url"`
	ShortUrl       string `json:"shortUrl" bson:"shortUrl"`
	Prefs          struct {
		PermissionLevel       string            `json:"permissionLevel" bson:"permissionLevel"`
		Voting                string            `json:"voting"`
		Comments              string            `json:"comments"`
		Invitations           string            `json:"invitations"`
		SelfJoin              bool              `json:"selfjoin"`
		CardCovers            bool              `json:"cardCovers" bson:"cardCovers"`
		CardAging             string            `json:"cardAging" bson:"cardAGing"`
		CalendarFeedEnabled   bool              `json:"calendarFeedEnabled" bson:"calendarFeedEnabled"`
		Background            string            `json:"background"`
		BackgroundColor       string            `json:"backgroundColor" bson:"backgroundColor"`
		BackgroundImage       string            `json:"backgroundImage" bson:"backgroundImage"`
		BackgroundImageScaled []BoardBackground `json:"backgroundImageScaled" bson:"backgroundImageScaled"`
		BackgroundTile        bool              `json:"backgroundTile" bson:"backgroundTile"`
		BackgroundBrightness  string            `json:"backgroundBrightness" bson:"backgroundBrightness"`
		CanBePublic           bool              `json:"canBePublic" bson:"canBePublic"`
		CanBeOrg              bool              `json:"canBeOrg" bson:"canBeOrg"`
		CanBePrivate          bool              `json:"canBePrivate" bson:"canBePrivate"`
		CanInvite             bool              `json:"canInvite" bson:"canInvite"`
	} `json:"prefs"`
	LabelNames struct {
		Red    string `json:"red"`
		Orange string `json:"orange"`
		Yellow string `json:"yellow"`
		Green  string `json:"green"`
		Blue   string `json:"blue"`
		Purple string `json:"purple"`
	} `json:"labelNames" bson:"labelNames"`
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
	req, err := http.NewRequest("POST", "https://trello.com/1/boards", bytes.NewBuffer([]byte(queryData)))
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
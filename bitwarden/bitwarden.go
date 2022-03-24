package bitwarden

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type BitwardenClient interface {
	GetProperty(property string) string
	SetProperty(property string, value string) string
	GetItems(folder string) (Items, error)
	GetItem(itemID string) (Item, error)
	FindItem(name string) (string, error)
}

type bitwardenClient struct {
	BaseUrl string
	Port    string
	Client  *resty.Client
}

// New generate a new gitlab client
func New(baseUrl string, port string) BitwardenClient {

	// TODO: Add TLS Insecure && Pass in CA CRT for authentication
	restClient := resty.New()

	if port == "" {
		port = "7787"
	}

	return &bitwardenClient{
		BaseUrl: baseUrl,
		Port:    port,
		Client:  restClient,
	}
}

func (r *bitwardenClient) GetProperty(property string) string {
	switch property {
	case "BaseUrl":
		return r.BaseUrl
	case "Port":
		return r.Port
	}
	return ""
}

func (r *bitwardenClient) SetProperty(property string, value string) string {
	switch property {
	case "BaseUrl":
		r.BaseUrl = value
		return r.BaseUrl
	case "Port":
		r.Port = value
		return r.Port
	}
	return ""
}

func (r *bitwardenClient) GetItems(folder string) (Items, error) {

	// TODO: detect if there are no options passed in, ? verus & for page option
	fetchUri := fmt.Sprintf("http://localhost:%s/list/object/items", r.Port)
	// logrus.Warn(fetchUri)
	resp, resperr := r.Client.R().
		Get(fetchUri)

	if resperr != nil {
		logrus.WithError(resperr).Error("Oops")
		return Items{}, resperr
	}
	var items Items
	marshErr := json.Unmarshal(resp.Body(), &items)
	if marshErr != nil {
		logrus.Fatal("Cannot marshall Pipeline", marshErr)
		return Items{}, resperr
	}

	return items, nil

}

func (r *bitwardenClient) GetItem(itemID string) (Item, error) {

	// TODO: detect if there are no options passed in, ? verus & for page option
	fetchUri := fmt.Sprintf("http://localhost:%s/object/item/%s", r.Port, itemID)
	// logrus.Warn(fetchUri)
	resp, resperr := r.Client.R().
		Get(fetchUri)

	if resperr != nil {
		logrus.WithError(resperr).Error("Oops")
		return Item{}, resperr
	}

	var items Item
	marshErr := json.Unmarshal(resp.Body(), &items)
	if marshErr != nil {
		logrus.Fatal("Cannot marshall Pipeline", marshErr)
		return Item{}, resperr
	}

	return items, nil

}

func (r *bitwardenClient) FindItem(name string) (string, error) {

	fetchUri := fmt.Sprintf("http://localhost:%s/list/object/items", r.Port)
	// logrus.Warn(fetchUri)
	resp, resperr := r.Client.R().
		Get(fetchUri)

	if resperr != nil {
		logrus.WithError(resperr).Error("Oops")
		return "", resperr
	}
	var items Items
	marshErr := json.Unmarshal(resp.Body(), &items)
	if marshErr != nil {
		logrus.Fatal("Cannot marshall Pipeline", marshErr)
		return "", resperr
	}

	for _, v := range items.Data.Data {
		if v.Name == name {
			return v.ID, nil
		}
	}
	return "", nil
}

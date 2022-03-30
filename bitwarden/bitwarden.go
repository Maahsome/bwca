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
	GetUsername(itemID string) (string, error)
	GetPassword(itemID string) (string, error)
	GetTOTP(itemID string) (string, error)
	NewItem(newlogin Newlogin) (ReturnStatus, error)
	UpdateItem(id string, updatelogin ItemData) (ReturnStatus, error)
	DeleteItem(itemID string) (bool, error)
	GetFolders() (Folders, error)
	GetFolder(folderID string) (Folder, error)
	NewFolder(newfolder Newfolder) (ReturnStatus, error)
	FindFolder(name string) (string, error)
	Sync() (SyncReturn, error)
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

	var item Item
	marshErr := json.Unmarshal(resp.Body(), &item)
	if marshErr != nil {
		logrus.Fatal("Cannot marshall Pipeline", marshErr)
		return Item{}, resperr
	}

	return item, nil

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

func (r *bitwardenClient) GetUsername(itemID string) (string, error) {

	// TODO: detect if there are no options passed in, ? verus & for page option
	fetchUri := fmt.Sprintf("http://localhost:%s/object/username/%s", r.Port, itemID)
	// logrus.Warn(fetchUri)
	resp, resperr := r.Client.R().
		Get(fetchUri)

	if resperr != nil {
		logrus.WithError(resperr).Error("Oops")
		return "", resperr
	}

	var pw Password
	marshErr := json.Unmarshal(resp.Body(), &pw)
	if marshErr != nil {
		logrus.Fatal("Cannot marshall Pipeline", marshErr)
		return "", resperr
	}

	return pw.Data.Data, nil
}

func (r *bitwardenClient) GetPassword(itemID string) (string, error) {

	// TODO: detect if there are no options passed in, ? verus & for page option
	fetchUri := fmt.Sprintf("http://localhost:%s/object/password/%s", r.Port, itemID)
	// logrus.Warn(fetchUri)
	resp, resperr := r.Client.R().
		Get(fetchUri)

	if resperr != nil {
		logrus.WithError(resperr).Error("Oops")
		return "", resperr
	}

	var pw Password
	marshErr := json.Unmarshal(resp.Body(), &pw)
	if marshErr != nil {
		logrus.Fatal("Cannot marshall Pipeline", marshErr)
		return "", resperr
	}

	return pw.Data.Data, nil
}

func (r *bitwardenClient) GetTOTP(itemID string) (string, error) {

	// TODO: detect if there are no options passed in, ? verus & for page option
	fetchUri := fmt.Sprintf("http://localhost:%s/object/totp/%s", r.Port, itemID)
	// logrus.Warn(fetchUri)
	resp, resperr := r.Client.R().
		Get(fetchUri)

	if resperr != nil {
		logrus.WithError(resperr).Error("Oops")
		return "", resperr
	}

	var totp Totp
	marshErr := json.Unmarshal(resp.Body(), &totp)
	if marshErr != nil {
		logrus.Fatal("Cannot marshall Pipeline", marshErr)
		return "", resperr
	}

	return totp.Data.Data, nil
}

func (r *bitwardenClient) NewItem(newlogin Newlogin) (ReturnStatus, error) {

	newLoginJSON, merr := json.Marshal(newlogin)
	if merr != nil {
		logrus.WithError(merr).Error("Oops")
		return ReturnStatus{}, merr
	}

	// fmt.Println(string(newLoginJSON[:]))

	fetchUri := fmt.Sprintf("http://localhost:%s/object/item", r.Port)
	// logrus.Warn(fetchUri)
	resp, resperr := r.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(string(newLoginJSON[:])).
		Post(fetchUri)

	if resperr != nil {
		logrus.WithError(resperr).Error("Oops")
		return ReturnStatus{}, resperr
	}

	var status ReturnStatus
	marshErr := json.Unmarshal(resp.Body(), &status)
	if marshErr != nil {
		logrus.Fatal("Cannot marshall Pipeline", marshErr)
		return ReturnStatus{}, resperr
	}

	return status, nil
}

func (r *bitwardenClient) UpdateItem(id string, updatelogin ItemData) (ReturnStatus, error) {

	newLoginJSON, merr := json.Marshal(updatelogin)
	if merr != nil {
		logrus.WithError(merr).Error("Oops")
		return ReturnStatus{}, merr
	}

	fetchUri := fmt.Sprintf("http://localhost:%s/object/item/%s", r.Port, id)
	// logrus.Warn(fetchUri)
	resp, resperr := r.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(string(newLoginJSON[:])).
		Put(fetchUri)

	if resperr != nil {
		logrus.WithError(resperr).Error("Oops")
		return ReturnStatus{}, resperr
	}

	var status ReturnStatus
	marshErr := json.Unmarshal(resp.Body(), &status)
	if marshErr != nil {
		logrus.Fatal("Cannot marshall Pipeline", marshErr)
		return ReturnStatus{}, resperr
	}

	return status, nil
}

func (r *bitwardenClient) DeleteItem(itemID string) (bool, error) {

	fetchUri := fmt.Sprintf("http://localhost:%s/object/item/%s", r.Port, itemID)
	resp, resperr := r.Client.R().
		Delete(fetchUri)

	if resperr != nil {
		logrus.WithError(resperr).Error("Oops")
		return false, resperr
	}

	var success Success
	marshErr := json.Unmarshal(resp.Body(), &success)
	if marshErr != nil {
		logrus.Fatal("Cannot marshall Pipeline", marshErr)
		return false, resperr
	}

	return success.Success, nil
}

func (r *bitwardenClient) GetFolders() (Folders, error) {

	// TODO: detect if there are no options passed in, ? verus & for page option
	fetchUri := fmt.Sprintf("http://localhost:%s/list/object/folders", r.Port)
	// logrus.Warn(fetchUri)
	resp, resperr := r.Client.R().
		Get(fetchUri)

	if resperr != nil {
		logrus.WithError(resperr).Error("Oops")
		return Folders{}, resperr
	}
	var folders Folders
	marshErr := json.Unmarshal(resp.Body(), &folders)
	if marshErr != nil {
		logrus.Fatal("Cannot marshall Pipeline", marshErr)
		return Folders{}, resperr
	}

	return folders, nil

}

func (r *bitwardenClient) GetFolder(folderID string) (Folder, error) {

	// TODO: detect if there are no options passed in, ? verus & for page option
	fetchUri := fmt.Sprintf("http://localhost:%s/object/folder/%s", r.Port, folderID)
	// logrus.Warn(fetchUri)
	resp, resperr := r.Client.R().
		Get(fetchUri)

	if resperr != nil {
		logrus.WithError(resperr).Error("Oops")
		return Folder{}, resperr
	}

	var folder Folder
	marshErr := json.Unmarshal(resp.Body(), &folder)
	if marshErr != nil {
		logrus.Fatal("Cannot marshall Pipeline", marshErr)
		return Folder{}, resperr
	}

	return folder, nil

}

func (r *bitwardenClient) NewFolder(newfolder Newfolder) (ReturnStatus, error) {

	newFolderJSON, merr := json.Marshal(newfolder)
	if merr != nil {
		logrus.WithError(merr).Error("Oops")
		return ReturnStatus{}, merr
	}

	// fmt.Println(string(newFolderJSON[:]))

	fetchUri := fmt.Sprintf("http://localhost:%s/object/folder", r.Port)
	// logrus.Warn(fetchUri)
	resp, resperr := r.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(string(newFolderJSON[:])).
		Post(fetchUri)

	if resperr != nil {
		logrus.WithError(resperr).Error("Oops")
		return ReturnStatus{}, resperr
	}

	var status ReturnStatus
	marshErr := json.Unmarshal(resp.Body(), &status)
	if marshErr != nil {
		logrus.Fatal("Cannot marshall Pipeline", marshErr)
		return ReturnStatus{}, resperr
	}

	return status, nil
}

func (r *bitwardenClient) FindFolder(name string) (string, error) {

	fetchUri := fmt.Sprintf("http://localhost:%s/list/object/folders", r.Port)
	// logrus.Warn(fetchUri)
	resp, resperr := r.Client.R().
		Get(fetchUri)

	if resperr != nil {
		logrus.WithError(resperr).Error("Oops")
		return "", resperr
	}
	var folders Folders
	marshErr := json.Unmarshal(resp.Body(), &folders)
	if marshErr != nil {
		logrus.Fatal("Cannot marshall Pipeline", marshErr)
		return "", resperr
	}

	for _, v := range folders.Data.Data {
		if v.Name == name {
			return v.ID, nil
		}
	}
	return "", nil
}

func (r *bitwardenClient) Sync() (SyncReturn, error) {

	fetchUri := fmt.Sprintf("http://localhost:%s/sync", r.Port)
	resp, resperr := r.Client.R().
		SetHeader("Content-Type", "application/json").
		Post(fetchUri)

	if resperr != nil {
		logrus.WithError(resperr).Error("Oops")
		return SyncReturn{}, resperr
	}

	var status SyncReturn
	marshErr := json.Unmarshal(resp.Body(), &status)
	if marshErr != nil {
		logrus.Fatal("Cannot marshall Pipeline", marshErr)
		return SyncReturn{}, resperr
	}

	return status, nil
}

package bitwarden

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/maahsome/gron"
	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Items struct {
	Data    ItemsData `json:"data"`
	Success bool      `json:"success"`
}

type ItemsData struct {
	Data   []ItemsDataDatum `json:"data"`
	Object string           `json:"object"`
}

type ItemsDataDatum struct {
	Card            ItemsDataDatumCard              `json:"card"`
	CollectionIds   []interface{}                   `json:"collectionIds"`
	DeletedDate     interface{}                     `json:"deletedDate"`
	Favorite        bool                            `json:"favorite"`
	Fields          []ItemsDataDatumField           `json:"fields"`
	FolderID        string                          `json:"folderId"`
	ID              string                          `json:"id"`
	Identity        ItemsDataDatumIdentity          `json:"identity"`
	Login           ItemsDataDatumLogin             `json:"login"`
	Name            string                          `json:"name"`
	Notes           interface{}                     `json:"notes"`
	Object          string                          `json:"object"`
	OrganizationID  interface{}                     `json:"organizationId"`
	PasswordHistory []ItemsDataDatumPasswordHistory `json:"passwordHistory"`
	Reprompt        int                             `json:"reprompt"`
	RevisionDate    string                          `json:"revisionDate"`
	SecureNote      ItemsDataDatumSecureNote        `json:"secureNote"`
	Type            int                             `json:"type"`
}

type ItemsDataDatumCard struct {
	Brand          string `json:"brand"`
	CardholderName string `json:"cardholderName"`
	Code           string `json:"code"`
	ExpMonth       string `json:"expMonth"`
	ExpYear        string `json:"expYear"`
	Number         string `json:"number"`
}

type ItemsDataDatumField struct {
	LinkedID interface{} `json:"linkedId"`
	Name     string      `json:"name"`
	Type     int         `json:"type"`
	Value    string      `json:"value"`
}

type ItemsDataDatumIdentity struct {
	Address1       string      `json:"address1"`
	Address2       interface{} `json:"address2"`
	Address3       interface{} `json:"address3"`
	City           string      `json:"city"`
	Company        string      `json:"company"`
	Country        string      `json:"country"`
	Email          string      `json:"email"`
	FirstName      string      `json:"firstName"`
	LastName       string      `json:"lastName"`
	LicenseNumber  interface{} `json:"licenseNumber"`
	MiddleName     string      `json:"middleName"`
	PassportNumber interface{} `json:"passportNumber"`
	Phone          string      `json:"phone"`
	PostalCode     string      `json:"postalCode"`
	Ssn            string      `json:"ssn"`
	State          string      `json:"state"`
	Title          string      `json:"title"`
	Username       string      `json:"username"`
}

type ItemsDataDatumLogin struct {
	Password             interface{}              `json:"password"`
	PasswordRevisionDate interface{}              `json:"passwordRevisionDate"`
	Totp                 interface{}              `json:"totp"`
	Uris                 []ItemsDataDatumLoginURI `json:"uris"`
	Username             interface{}              `json:"username"`
}

type ItemsDataDatumLoginURI struct {
	Match interface{} `json:"match"`
	URI   string      `json:"uri"`
}

type ItemsDataDatumPasswordHistory struct {
	LastUsedDate string `json:"lastUsedDate"`
	Password     string `json:"password"`
}

type ItemsDataDatumSecureNote struct {
	Type int `json:"type"`
}

type Item struct {
	Data    ItemData `json:"data"`
	Success bool     `json:"success"`
}

type ItemData struct {
	CollectionIds   []interface{}             `json:"collectionIds"`
	DeletedDate     interface{}               `json:"deletedDate"`
	Favorite        bool                      `json:"favorite"`
	FolderID        string                    `json:"folderId"`
	ID              string                    `json:"id"`
	Login           ItemDataLogin             `json:"login"`
	Name            string                    `json:"name"`
	Notes           string                    `json:"notes"`
	Object          string                    `json:"object"`
	OrganizationID  interface{}               `json:"organizationId"`
	PasswordHistory []ItemDataPasswordHistory `json:"passwordHistory"`
	Reprompt        int                       `json:"reprompt"`
	RevisionDate    string                    `json:"revisionDate"`
	Type            int                       `json:"type"`
}

type ItemDataLogin struct {
	Password             string `json:"password"`
	PasswordRevisionDate string `json:"passwordRevisionDate"`
	Totp                 string `json:"totp"`
	Username             string `json:"username"`
}

type ItemDataPasswordHistory struct {
	LastUsedDate string `json:"lastUsedDate"`
	Password     string `json:"password"`
}

type Totp struct {
	Data    TotpData `json:"data"`
	Success bool     `json:"success"`
}

type TotpData struct {
	Data   string `json:"data"`
	Object string `json:"object"`
}

type Password struct {
	Data    PasswordData `json:"data"`
	Success bool         `json:"success"`
}

type PasswordData struct {
	Data   string `json:"data"`
	Object string `json:"object"`
}

type Newlogin struct {
	CollectionID   string            `json:"collectionId,omitempty"`
	Favorite       bool              `json:"favorite"`
	Fields         map[string]string `json:"fields,omitempty"`
	FolderID       string            `json:"folderId,omitempty"`
	Login          NewloginLogin     `json:"login"`
	Name           string            `json:"name"`
	Notes          string            `json:"notes"`
	OrganizationID string            `json:"organizationId,omitempty"`
	Reprompt       int               `json:"reprompt"`
	Type           int               `json:"type"`
}

type NewloginLogin struct {
	Password string             `json:"password"`
	Totp     string             `json:"totp,omitempty"`
	Uris     []NewloginLoginURI `json:"uris"`
	Username string             `json:"username"`
}

type NewloginLoginURI struct {
	Match int    `json:"match,omitempty"`
	URI   string `json:"uri,omitempty"`
}

type ReturnStatus struct {
	Data         interface{} `json:"data"`
	DeleteDate   string      `json:"deleteDate"`
	RevisionDate string      `json:"revisionDate"`
	Success      bool        `json:"success"`
}

type Success struct {
	Success bool `json:"success"`
}

type Folders struct {
	Data    FoldersData `json:"data"`
	Success bool        `json:"success"`
}

type FoldersData struct {
	Data   []FoldersDataDatum `json:"data"`
	Object string             `json:"object"`
}

type FoldersDataDatum struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Object string `json:"object"`
}

type Folder struct {
	Data    FolderData `json:"data"`
	Success bool       `json:"success"`
}

type FolderData struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Object string `json:"object"`
}

type Newfolder struct {
	Name string `json:"name"`
}

// ToJSON - Write the output as JSON
func (it *Items) ToJSON() string {
	itJSON, err := json.MarshalIndent(it, "", "  ")
	if err != nil {
		logrus.WithError(err).Error("Error extracting JSON")
		return ""
	}
	return string(itJSON[:])
}

func (it *Items) ToGRON() string {
	itJSON, err := json.MarshalIndent(it, "", "  ")
	if err != nil {
		logrus.WithError(err).Error("Error extracting JSON for GRON")
	}
	subReader := strings.NewReader(string(itJSON[:]))
	subValues := &bytes.Buffer{}
	ges := gron.NewGron(subReader, subValues)
	ges.SetMonochrome(false)
	if serr := ges.ToGron(); serr != nil {
		logrus.WithError(serr).Error("Problem generating GRON syntax")
		return ""
	}
	return string(subValues.Bytes())
}

func (it *Items) ToYAML() string {
	itYAML, err := yaml.Marshal(it)
	if err != nil {
		logrus.WithError(err).Error("Error extracting YAML")
		return ""
	}
	return string(itYAML[:])
}

func (it *Items) ToTEXT(noHeaders bool) string {
	buf, row := new(bytes.Buffer), make([]string, 0)

	// ************************** TableWriter ******************************
	table := tablewriter.NewWriter(buf)
	if !noHeaders {
		table.SetHeader([]string{"ID", "NAME", "FOLDERID"})
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	}

	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)

	for _, v := range *&it.Data.Data {
		row = []string{
			v.ID,
			v.Name,
			v.FolderID,
		}
		table.Append(row)
	}

	table.Render()

	return buf.String()

}

// ToJSON - Write the output as JSON
func (it *Item) ToJSON() string {
	itJSON, err := json.MarshalIndent(it, "", "  ")
	if err != nil {
		logrus.WithError(err).Error("Error extracting JSON")
		return ""
	}
	return string(itJSON[:])
}

func (it *Item) ToGRON() string {
	itJSON, err := json.MarshalIndent(it, "", "  ")
	if err != nil {
		logrus.WithError(err).Error("Error extracting JSON for GRON")
	}
	subReader := strings.NewReader(string(itJSON[:]))
	subValues := &bytes.Buffer{}
	ges := gron.NewGron(subReader, subValues)
	ges.SetMonochrome(false)
	if serr := ges.ToGron(); serr != nil {
		logrus.WithError(serr).Error("Problem generating GRON syntax")
		return ""
	}
	return string(subValues.Bytes())
}

func (it *Item) ToYAML() string {
	itYAML, err := yaml.Marshal(it)
	if err != nil {
		logrus.WithError(err).Error("Error extracting YAML")
		return ""
	}
	return string(itYAML[:])
}

func (it *Item) ToTEXT(noHeaders bool) string {
	buf, row := new(bytes.Buffer), make([]string, 0)

	// ************************** TableWriter ******************************
	table := tablewriter.NewWriter(buf)
	if !noHeaders {
		table.SetHeader([]string{"ID", "NAME", "FOLDERID"})
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	}

	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)

	// for _, v := range *&it.Data.Data {
	row = []string{
		it.Data.ID,
		it.Data.Name,
		it.Data.FolderID,
	}
	table.Append(row)
	// }

	table.Render()

	return buf.String()

}

// ToJSON - Write the output as JSON
func (fl *Folders) ToJSON() string {
	flJSON, err := json.MarshalIndent(fl, "", "  ")
	if err != nil {
		logrus.WithError(err).Error("Error extracting JSON")
		return ""
	}
	return string(flJSON[:])
}

func (fl *Folders) ToGRON() string {
	itJSON, err := json.MarshalIndent(fl, "", "  ")
	if err != nil {
		logrus.WithError(err).Error("Error extracting JSON for GRON")
	}
	subReader := strings.NewReader(string(itJSON[:]))
	subValues := &bytes.Buffer{}
	ges := gron.NewGron(subReader, subValues)
	ges.SetMonochrome(false)
	if serr := ges.ToGron(); serr != nil {
		logrus.WithError(serr).Error("Problem generating GRON syntax")
		return ""
	}
	return string(subValues.Bytes())
}

func (fl *Folders) ToYAML() string {
	itYAML, err := yaml.Marshal(fl)
	if err != nil {
		logrus.WithError(err).Error("Error extracting YAML")
		return ""
	}
	return string(itYAML[:])
}

func (fl *Folders) ToTEXT(noHeaders bool) string {
	buf, row := new(bytes.Buffer), make([]string, 0)

	// ************************** TableWriter ******************************
	table := tablewriter.NewWriter(buf)
	if !noHeaders {
		table.SetHeader([]string{"ID", "NAME"})
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	}

	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)

	for _, v := range *&fl.Data.Data {
		row = []string{
			v.ID,
			v.Name,
		}
		table.Append(row)
	}

	table.Render()

	return buf.String()

}

// ToJSON - Write the output as JSON
func (fl *Folder) ToJSON() string {
	flJSON, err := json.MarshalIndent(fl, "", "  ")
	if err != nil {
		logrus.WithError(err).Error("Error extracting JSON")
		return ""
	}
	return string(flJSON[:])
}

func (fl *Folder) ToGRON() string {
	flJSON, err := json.MarshalIndent(fl, "", "  ")
	if err != nil {
		logrus.WithError(err).Error("Error extracting JSON for GRON")
	}
	subReader := strings.NewReader(string(flJSON[:]))
	subValues := &bytes.Buffer{}
	ges := gron.NewGron(subReader, subValues)
	ges.SetMonochrome(false)
	if serr := ges.ToGron(); serr != nil {
		logrus.WithError(serr).Error("Problem generating GRON syntax")
		return ""
	}
	return string(subValues.Bytes())
}

func (fl *Folder) ToYAML() string {
	flYAML, err := yaml.Marshal(fl)
	if err != nil {
		logrus.WithError(err).Error("Error extracting YAML")
		return ""
	}
	return string(flYAML[:])
}

func (fl *Folder) ToTEXT(noHeaders bool) string {
	buf, row := new(bytes.Buffer), make([]string, 0)

	// ************************** TableWriter ******************************
	table := tablewriter.NewWriter(buf)
	if !noHeaders {
		table.SetHeader([]string{"ID", "NAME"})
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	}

	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)

	// for _, v := range *&fl.Data.Data {
	row = []string{
		fl.Data.ID,
		fl.Data.Name,
	}
	table.Append(row)
	// }

	table.Render()

	return buf.String()

}

package mgostore

import (
	"testing"

	"github.com/robjsliwa/osin"
	"gopkg.in/mgo.v2/bson"
)

func TestCopyToAccessData(t *testing.T) {
	client := &osin.DefaultClient{
		Id:          "1234",
		Secret:      "aabbccdd",
		RedirectUri: "http://localhost:14000/appauth",
	}

	accessData := &DefaultAccessData{
		Client:        *client,
		AuthorizeData: nil,
		AccessData:    nil,
		AccessToken:   "kksdjskdj",
		RefreshToken:  "kjdskdjskdjks",
		ExpiresIn:     3600,
		Scope:         "skdjksjd",
		RedirectUri:   "http://localhost:14000/appauth",
		CreatedAt:     bson.Now(),
	}

	toAccessData := new(osin.AccessData)

	accessData.CopyToAccessData(toAccessData)

	if !compareAccessData(accessData, toAccessData) {
		t.Errorf("TestCopyToAccessData failed, expected: '%+v', got: '%+v'", accessData, toAccessData)
	}
}

func compareAccessData(from *DefaultAccessData, to *osin.AccessData) bool {
	if from.Client.GetId() == to.Client.GetId() &&
		from.Client.GetSecret() == to.Client.GetSecret() &&
		from.Client.GetRedirectUri() == to.Client.GetRedirectUri() &&
		from.Client.GetUserData() == to.Client.GetUserData() &&
		from.AccessToken == to.AccessToken &&
		from.RefreshToken == to.RefreshToken &&
		from.ExpiresIn == to.ExpiresIn &&
		from.Scope == to.Scope &&
		from.RedirectUri == to.RedirectUri &&
		from.CreatedAt == to.CreatedAt {
		return true
	}

	return false
}

package mgostore

import (
	"testing"

	"github.com/robjsliwa/osin"
	"gopkg.in/mgo.v2/bson"
)

func TestCopyToAuthorizeData(t *testing.T) {
	client := &osin.DefaultClient{
		Id:          "1234",
		Secret:      "aabbccdd",
		RedirectUri: "http://localhost:14000/appauth",
	}

	authorizeData := &DefaultAuthorizeData{
		Client:      *client,
		Code:        "9999",
		ExpiresIn:   3600,
		CreatedAt:   bson.Now(),
		RedirectUri: "http://localhost:14000/appauth",
	}

	toAuthorizeData := new(osin.AuthorizeData)

	authorizeData.CopyToAuthorizeData(toAuthorizeData)

	if !compareAuthorizeData(authorizeData, toAuthorizeData) {
		t.Errorf("TestCopyToAuthorizeData failed, expected: '%+v', got: '%+v'", authorizeData, toAuthorizeData)
	}
}

func compareAuthorizeData(from *DefaultAuthorizeData, to *osin.AuthorizeData) bool {
	if from.Client.GetId() == to.Client.GetId() &&
		from.Client.GetSecret() == to.Client.GetSecret() &&
		from.Client.GetRedirectUri() == to.Client.GetRedirectUri() &&
		from.Client.GetUserData() == to.Client.GetUserData() &&
		from.Code == to.Code &&
		from.ExpiresIn == to.ExpiresIn &&
		from.CreatedAt == to.CreatedAt &&
		from.RedirectUri == to.RedirectUri {
		return true
	}

	return false
}

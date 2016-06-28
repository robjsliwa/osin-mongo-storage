package mgostore

import (
	"time"

	"github.com/robjsliwa/osin"
)

// AccessData represents an access grant (tokens, expiration, client, etc)
type DefaultAccessData struct {
	// Client information
	Client osin.DefaultClient `bson:'client'`

	// Authorize data, for authorization code
	AuthorizeData *DefaultAuthorizeData `bson:'authorizedata'`

	// Previous access data, for refresh token
	AccessData *DefaultAccessData `bson:'accessdata'`

	// Access token
	AccessToken string `bson:'accesstoken'`

	// Refresh Token. Can be blank
	RefreshToken string `bson:'refreshtoken'`

	// Token expiration in seconds
	ExpiresIn int32 `bson:'expiresin'`

	// Requested scope
	Scope string `bson:'scope'`

	// Redirect Uri from request
	RedirectUri string `bson:'redirecturi'`

	// Date created
	CreatedAt time.Time `bson:'createdat'`

	// Data to be passed to storage. Not used by the library.
	UserData interface{} `bson:'userdata'`
}

func (defautlAccessData *DefaultAccessData) CopyToAccessData(accessData *osin.AccessData) {
	if defautlAccessData == nil || accessData == nil {
		return
	}

	client := new(osin.Client)
	*client = &defautlAccessData.Client
	accessData.Client = *client
	if defautlAccessData.AuthorizeData != nil {
		if accessData.AuthorizeData == nil {
			accessData.AuthorizeData = new(osin.AuthorizeData)
		}
		defautlAccessData.AuthorizeData.CopyToAuthorizeData(accessData.AuthorizeData)
	} else {
		accessData.AuthorizeData = nil
	}
	if defautlAccessData.AccessData != nil {
		if accessData.AccessData == nil {
			accessData.AccessData = new(osin.AccessData)
		}
		defautlAccessData.AccessData.CopyToAccessData(accessData.AccessData)
	} else {
		accessData.AccessData = nil
	}
	accessData.AccessToken = defautlAccessData.AccessToken
	accessData.RefreshToken = defautlAccessData.RefreshToken
	accessData.ExpiresIn = defautlAccessData.ExpiresIn
	accessData.Scope = defautlAccessData.Scope
	accessData.RedirectUri = defautlAccessData.RedirectUri
	accessData.CreatedAt = defautlAccessData.CreatedAt
	accessData.UserData = defautlAccessData.UserData
}

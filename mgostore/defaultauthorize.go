package mgostore

// Authorization data
import (
	"time"

	"github.com/robjsliwa/osin"
)

type DefaultAuthorizeData struct {
	// Client information
	Client osin.DefaultClient `bson:'client'`

	// Authorization code
	Code string `bson:'code'`

	// Token expiration in seconds
	ExpiresIn int32 `bson:'expiresin'`

	// Requested scope
	Scope string `bson:'scope'`

	// Redirect Uri from request
	RedirectUri string `bson:'redirecturi'`

	// State data from request
	State string `bson:'state'`

	// Date created
	CreatedAt time.Time `bson:'createdat'`

	// Data to be passed to storage. Not used by the library.
	UserData interface{} `bson:'userdata'`
}

func (defaultAuthorizeData *DefaultAuthorizeData) CopyToAuthorizeData(authorizeData *osin.AuthorizeData) {
	if authorizeData == nil {
		return
	}

	client := new(osin.Client)
	*client = &defaultAuthorizeData.Client
	authorizeData.Client = *client
	authorizeData.Code = defaultAuthorizeData.Code
	authorizeData.ExpiresIn = defaultAuthorizeData.ExpiresIn
	authorizeData.Scope = defaultAuthorizeData.Scope
	authorizeData.RedirectUri = defaultAuthorizeData.RedirectUri
	authorizeData.State = defaultAuthorizeData.State
	authorizeData.CreatedAt = defaultAuthorizeData.CreatedAt
	authorizeData.UserData = defaultAuthorizeData.UserData
}

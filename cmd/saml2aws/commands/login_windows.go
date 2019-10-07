package commands

import (
	"github.com/tinder-edwardowens/saml2aws/helper/credentials"
	"github.com/tinder-edwardowens/saml2aws/helper/wincred"
)

func init() {
	credentials.CurrentHelper = &wincred.Wincred{}
}

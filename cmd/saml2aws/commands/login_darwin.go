package commands

import (
	"github.com/tinder-edwardowens/saml2aws/helper/credentials"
	"github.com/tinder-edwardowens/saml2aws/helper/osxkeychain"
)

func init() {
	credentials.CurrentHelper = &osxkeychain.Osxkeychain{}
}

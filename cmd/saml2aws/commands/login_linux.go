package commands

import (
	"github.com/tinder-edwardowens/saml2aws/helper/credentials"
	"github.com/tinder-edwardowens/saml2aws/helper/linuxkeyring"
)

func init() {
	if keyringHelper, err := linuxkeyring.NewKeyringHelper(); err == nil {
		credentials.CurrentHelper = keyringHelper
	}
}

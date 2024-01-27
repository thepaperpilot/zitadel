package domain

import (
	"golang.org/x/text/language"

	es_models "github.com/zitadel/zitadel/internal/eventstore/v1/models"
	"github.com/zitadel/zitadel/internal/zerrors"
)

type Profile struct {
	es_models.ObjectRoot

	FirstName          string
	LastName           string
	NickName           string
	DisplayName        string
	PreferredLanguage  language.Tag
	Gender             Gender
	PreferredLoginName string
	LoginNames         []string
}

func (p *Profile) Validate() error {
	if p == nil {
		return zerrors.ThrowInvalidArgument(nil, "PROFILE-GPY3p", "Errors.User.Profile.Empty")
	}
	return nil
}

func AvatarURL(prefix, resourceOwner, key string) string {
	if prefix == "" || resourceOwner == "" || key == "" {
		return ""
	}
	return prefix + "/" + resourceOwner + "/" + key
}

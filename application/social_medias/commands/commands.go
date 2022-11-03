package social_media_commands

import (
	contracts "rocky.my.id/git/mygram/application/social_medias/contracts"
)

type SocialMediaCommands struct {
	Repository contracts.SocialMediaRepositoryContract
}

func NewSocialMediaCommands(repository contracts.SocialMediaRepositoryContract) *SocialMediaCommands {
	return &SocialMediaCommands{Repository: repository}
}

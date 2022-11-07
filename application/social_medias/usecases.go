package social_media_use_cases

import contracts "rocky.my.id/git/mygram/application/social_medias/contracts"

type SocialMediaUseCases struct {
	Commands contracts.SocialMediaCommandsContract
	Queries  contracts.SocialMediaQueriesContract
}

func NewSocialMediaUseCases(commands contracts.SocialMediaCommandsContract, queries contracts.SocialMediaQueriesContract) *SocialMediaUseCases {
	return &SocialMediaUseCases{Commands: commands, Queries: queries}
}

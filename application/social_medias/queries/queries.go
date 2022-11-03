package social_media_queries

import (
	contracts "rocky.my.id/git/mygram/application/social_medias/contracts"
)

type SocialMediaQueries struct {
	Repository contracts.SocialMediaRepositoryContract
}

func NewSocialMediaQueries(repository contracts.SocialMediaRepositoryContract) *SocialMediaQueries {
	return &SocialMediaQueries{Repository: repository}
}

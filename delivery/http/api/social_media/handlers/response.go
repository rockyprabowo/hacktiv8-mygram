package social_media_handlers

import dto "rocky.my.id/git/mygram/application/social_medias/dto"

type SocialMediaCollectionResponse struct {
	SocialMedias []dto.SocialMediaWithUserDTO `json:"social_medias"`
}

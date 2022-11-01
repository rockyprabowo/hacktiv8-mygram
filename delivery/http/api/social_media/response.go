package social_media_http_delivery

import dto "rocky.my.id/git/mygram/application/social_medias/dto"

const DeleteSuccessMessage = "Your social media has been successfully deleted."

type SocialMediaCollectionResponse struct {
	SocialMedias []dto.SocialMediaWithUserDTO `json:"social_medias"`
}

package responses

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jtonynet/cine-catalogo/models"
)

var MoviePosterContentType = "image/png"

type basePoster struct {
	UUID            uuid.UUID `gorm:"type:uuid;unique;not null"`
	Name            string
	ContentType     string
	AlternativeText string
	Path            string
}

type Poster struct {
	basePoster

	Templates interface{} `json:"_templates,omitempty"`
}

func NewPoster(
	model models.Poster,
	movieUUID uuid.UUID,
	baseURL string,
	templates interface{},
) Poster {
	poster := Poster{
		basePoster{
			UUID:            model.UUID,
			Name:            model.Name,
			ContentType:     model.ContentType,
			AlternativeText: model.AlternativeText,
		},

		//NewPosterLinks(model.Movie.UUID, model.UUID, baseURL, model.Path),
		NewPosterLinks(movieUUID, model.UUID, baseURL, model.Path),

		//templates,
	}

	return poster
}

type HATEOASPosterItemLinks struct {
	Links *HATEOASPosterLinks `json:"_links,omitempty"`
}

type HATEOASPosterLinks struct {
	Self         HATEOASLink `json:"self"`
	Image        HATEOASLink `json:"image"`
	UpdatePoster HATEOASLink `json:"update-poster"`
	DeletePoster HATEOASLink `json:"delete-poster"`
}

func NewPosterLinks(
	movieUUID,
	posterUUID uuid.UUID,
	baseURL,
	posterPath string) *HATEOASPosterLinks {
	return &HATEOASPosterLinks{
		Self:         HATEOASLink{HREF: fmt.Sprintf("%s/movies/%s/posters/%s", baseURL, movieUUID, posterUUID)},
		UpdatePoster: HATEOASLink{HREF: fmt.Sprintf("%s/movies/%s/posters/%s", baseURL, movieUUID, posterUUID)},
		DeletePoster: HATEOASLink{HREF: fmt.Sprintf("%s/movies/%s/posters/%s", baseURL, movieUUID, posterUUID)},
		Image:        HATEOASLink{HREF: fmt.Sprintf("%s/%s", baseURL, posterPath)},
	}
}

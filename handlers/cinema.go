package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jtonynet/cine-catalogo/handlers/requests"
	"github.com/jtonynet/cine-catalogo/handlers/responses"
	"github.com/jtonynet/cine-catalogo/internal/database"
	"github.com/jtonynet/cine-catalogo/models"
)

func CreateCinemas(ctx *gin.Context) {
	requestList := []requests.Cinema{}
	err := ctx.ShouldBindJSON(&requestList)
	if err != nil {
		//TODO: Implements in future
		return
	}

	cinemaList := []models.Cinema{}
	for _, request := range requestList {
		cinema, err := models.NewCinema(
			uuid.New(),
			request.Name,
			request.Description,
			request.Capacity,
		)
		if err != nil {
			//TODO: Implements in future
			return
		}

		cinemaList = append(cinemaList, cinema)
	}

	if err := database.DB.Create(&cinemaList).Error; err != nil {
		//TODO: Implements in future
		return
	}

	responseList := []responses.Cinema{}
	for _, cinema := range cinemaList {
		responseList = append(responseList,
			responses.Cinema{
				UUID:        cinema.UUID,
				Name:        cinema.Name,
				Description: cinema.Description,
				Capacity:    cinema.Capacity,
			},
		)
	}

	responses.SendSuccess(ctx, http.StatusOK, "CreateCinemas", responseList)
}

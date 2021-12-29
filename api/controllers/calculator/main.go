package calculator

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/msfidelis/nutrition/pkg/bmr"
	"github.com/msfidelis/nutrition/pkg/imc"
	"github.com/msfidelis/nutrition/pkg/logger"
)

type Request struct {
	Age    int     `json:"age"`
	Weight float64 `json:"weight"`
	Height float64 `json:"height"`
	Gender string  `json:"gender"`
}

type Response struct {
	Status int `json:"status" binding:"required"`
	Imc    struct {
		Result float64 `json:"result"`
		Class  string  `json:"class"`
	} `json:"imc"`
	Basal struct {
		BMR float64 `json:"bmr"`
	} `json:"basal"`
	HealthInfo struct {
		Age    int     `json:"age"`
		Weight float64 `json:"weight"`
		Height float64 `json:"height"`
		Gender string  `json:"gender"`
	} `json:"health_info"`
	Recomendations struct {
		Protein struct {
			Value int    `json:"value"`
			Unit  string `json:"unit"`
		} `json:"protein"`
		Water struct {
			Value float64 `json:"value"`
			Unit  string  `json:"unit"`
		} `json:"water"`
	} `json:"recomendations"`
}

// Ok godoc
// @Summary Return IAM and Calc
// @Tags HealthCalculator
// @Produce json
// @Success 200 {object} Response
// @Router /calculator [post]
func Post(c *gin.Context) {
	var response Response
	var request Request

	log := logger.Instance()

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// IMC
	imcValue, class := imc.Calc(request.Weight, request.Height)

	log.Info().
		Str("Gender", request.Gender).
		Int("Age", request.Age).
		Float64("Weight", request.Weight).
		Float64("Height", request.Height).
		Float64("Body mass index", imcValue).
		Msg("Body mass index")

	// Calcular o Metabolismo Basal

	basal := bmr.Calc(request.Gender, request.Weight, request.Height, request.Age)

	log.Info().
		Str("Gender", request.Gender).
		Int("Age", request.Age).
		Float64("Weight", request.Weight).
		Float64("Height", request.Height).
		Float64("Basal", basal).
		Msg("Basal Metabolic Rate")

	response.Imc.Result = imcValue
	response.Imc.Class = class
	response.Basal.BMR = basal

	// Recomendations
	response.Recomendations.Protein.Value = int(request.Weight) * 2
	response.Recomendations.Protein.Unit = "g"
	response.Recomendations.Water.Value = request.Weight * float64(35)
	response.Recomendations.Water.Unit = "ml"

	// Health Info
	response.HealthInfo.Age = request.Age
	response.HealthInfo.Gender = request.Gender
	response.HealthInfo.Weight = request.Weight
	response.HealthInfo.Height = request.Height

	response.Status = http.StatusOK
	c.JSON(http.StatusOK, response)
}

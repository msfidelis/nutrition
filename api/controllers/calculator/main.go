package calculator

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
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
		Protein int     `json:"protein"`
		Water   float64 `json:"water"`
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

	log.Info().
		Str("Gender", request.Gender).
		Int("Age", request.Age).
		Float64("Weight", request.Weight).
		Float64("Height", request.Height).
		Msg("Initial Health Calc")

	// Calcular IMC
	var imc float64
	var class string
	imc = (request.Weight / (request.Height * request.Height)) * 1

	if imc <= 18.5 {
		class = "underweight"
	} else if imc > 18.5 && imc < 24.9 {
		class = "healthy"
	} else if imc >= 25 && imc <= 29.9 {
		class = "overweight"
	} else if imc > 30 {
		class = "obese"
	}

	log.Info().
		Str("Gender", request.Gender).
		Int("Age", request.Age).
		Float64("Weight", request.Weight).
		Float64("Height", request.Height).
		Float64("Body mass index", imc).
		Msg("Body mass index")

	// Calcular o Metabolismo Basal
	log.Info().
		Str("Gender", request.Gender).
		Int("Age", request.Age).
		Float64("Weight", request.Weight).
		Float64("Height", request.Height).
		Msg("Initial Basal Metabolic Rate Calc")

	var basal float64

	if strings.ToUpper(request.Gender) == "M" {
		basal = 66 + (13.75 * request.Weight) + (5.0 * (request.Height * 100)) - (6.8 * float64(request.Age))
	} else {
		basal = 665 + (9.56 * request.Weight) + (1.8 * (request.Height * 100)) - (4.7 * float64(request.Age))
	}

	log.Info().
		Str("Gender", request.Gender).
		Int("Age", request.Age).
		Float64("Weight", request.Weight).
		Float64("Height", request.Height).
		Float64("Basal", basal).
		Msg("Basal Metabolic Rate")

	response.Imc.Result = imc
	response.Imc.Class = class
	response.Basal.BMR = basal

	// Recomendations
	response.Recomendations.Protein = int(request.Weight) * 2
	response.Recomendations.Water = request.Weight * float64(35)

	// Health Info
	response.HealthInfo.Age = request.Age
	response.HealthInfo.Gender = request.Gender
	response.HealthInfo.Weight = request.Weight
	response.HealthInfo.Height = request.Height

	response.Status = http.StatusOK
	c.JSON(http.StatusOK, response)
}

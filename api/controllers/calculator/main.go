package calculator

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/msfidelis/nutrition/pkg/bmr"
	"github.com/msfidelis/nutrition/pkg/calories"
	"github.com/msfidelis/nutrition/pkg/imc"
	"github.com/msfidelis/nutrition/pkg/logger"
	"github.com/msfidelis/nutrition/pkg/proteins"
	"github.com/msfidelis/nutrition/pkg/water"
)

type Request struct {
	Age               int     `json:"age" binding:"required"`
	Weight            float64 `json:"weight" binding:"required"`
	Height            float64 `json:"height" binding:"required"`
	Gender            string  `json:"gender" binding:"required,oneof=M F"`
	ActivityIntensity string  `json:"activity_intensity" binding:"required,oneof=sedentary lightly_active moderately_active very_active extra_active"`
}

type Response struct {
	Status int `json:"status" binding:"required"`
	Imc    struct {
		Result float64 `json:"result"`
		Class  string  `json:"class"`
	} `json:"imc"`
	Basal struct {
		BMR struct {
			Value float64 `json:"value"`
			Unit  string  `json:"unit"`
		} `json:"bmr"`
		Necessity struct {
			Value float64 `json:"value"`
			Unit  string  `json:"unit"`
		} `json:"necessity"`
	} `json:"basal"`
	HealthInfo struct {
		Age               int     `json:"age"`
		Weight            float64 `json:"weight"`
		Height            float64 `json:"height"`
		Gender            string  `json:"gender"`
		ActivityIntensity string  `json:"activity_intensity"`
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
		Calories struct {
			Maintain struct {
				Value float64 `json:"value"`
				Unit  string  `json:"unit"`
			} `json:"maintain_weight"`
			Loss struct {
				Value float64 `json:"value"`
				Unit  string  `json:"unit"`
			} `json:"loss_weight"`
			Gain struct {
				Value float64 `json:"value"`
				Unit  string  `json:"unit"`
			} `json:"gain_weight"`
		} `json:"calories"`
	} `json:"recomendations"`
}

// Ok godoc
// @Summary Return IAM and Calc
// @Tags HealthCalculator
// @Produce json
// @Success 200 {object} Response
// @Param message body Request true "Health Information"
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
	basal, necessity := bmr.Calc(request.Gender, request.Weight, request.Height, request.Age, request.ActivityIntensity)

	log.Info().
		Str("Gender", request.Gender).
		Int("Age", request.Age).
		Float64("Weight", request.Weight).
		Float64("Height", request.Height).
		Float64("Basal", basal).
		Msg("Basal Metabolic Rate")

	// IMC Calc
	response.Imc.Result = imcValue
	response.Imc.Class = class

	// Basal Calc
	response.Basal.BMR.Value = basal
	response.Basal.BMR.Unit = "kcal"
	response.Basal.Necessity.Value = necessity
	response.Basal.Necessity.Unit = "kcal"

	// Consume recomendations
	proteinsToConsume := proteins.Calc(int(request.Weight))
	waterToConsume := water.Calc(request.Weight)

	response.Recomendations.Protein.Value = proteinsToConsume
	response.Recomendations.Protein.Unit = "g"
	response.Recomendations.Water.Value = waterToConsume
	response.Recomendations.Water.Unit = "ml"

	// @TODO - Ajustar calculos
	caloriesMaintain := calories.Maintain(necessity)
	caloriesLoss := calories.Loss(necessity)
	caloriesGain := calories.Gain(necessity)

	// Maintain Weight
	response.Recomendations.Calories.Maintain.Value = caloriesMaintain
	response.Recomendations.Calories.Maintain.Unit = "kcal"

	// Loss Weight
	response.Recomendations.Calories.Loss.Value = caloriesLoss
	response.Recomendations.Calories.Loss.Unit = "kcal"

	// Gain Weight
	response.Recomendations.Calories.Gain.Value = caloriesGain
	response.Recomendations.Calories.Gain.Unit = "kcal"

	// Health Info
	response.HealthInfo.Age = request.Age
	response.HealthInfo.Gender = request.Gender
	response.HealthInfo.Weight = request.Weight
	response.HealthInfo.Height = request.Height
	response.HealthInfo.ActivityIntensity = request.ActivityIntensity

	response.Status = http.StatusOK

	c.JSON(http.StatusOK, response)
}

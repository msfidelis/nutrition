package main

import (
	"testing"

	"github.com/msfidelis/nutrition/pkg/bmr"
	"github.com/msfidelis/nutrition/pkg/calories"
	"github.com/msfidelis/nutrition/pkg/imc"
	"github.com/msfidelis/nutrition/pkg/proteins"
	"github.com/msfidelis/nutrition/pkg/water"
)

func TestBMRSedentary(t *testing.T) {

	t.Run("BMR Calc - Male Sedentary", func(t *testing.T) {

		basal, necessity := bmr.Calc("M", 89.0, 1.77, 26, "sedentary")
		want := 1997.95
		wantNecessity := 2397.54
		if basal != want {
			t.Errorf("got %v want %v", basal, want)
		}
		if necessity != wantNecessity {
			t.Errorf("Necessity calc error, got %v want %v", necessity, wantNecessity)
		}

	})

	t.Run("BMR Calc - Female Sedentary", func(t *testing.T) {

		basal, necessity := bmr.Calc("F", 89.0, 1.77, 26, "sedentary")
		want := 1712.24
		wantNecessity := 2054.688
		if basal != want {
			t.Errorf("Basal calc error, got %v want %v", basal, want)
		}
		if necessity != wantNecessity {
			t.Errorf("Necessity calc error, got %v want %v", necessity, wantNecessity)
		}

	})

}

func TestBMRLightlyActive(t *testing.T) {

	t.Run("BMR Calc - Male Lightly Active", func(t *testing.T) {

		basal, necessity := bmr.Calc("M", 89.0, 1.77, 26, "lightly_active")
		want := 1997.95
		wantNecessity := 2747.18125
		if basal != want {
			t.Errorf("got %v want %v", basal, want)
		}
		if necessity != wantNecessity {
			t.Errorf("Necessity calc error, got %v want %v", necessity, wantNecessity)
		}

	})

	t.Run("BMR Calc - Female Lightly Active", func(t *testing.T) {

		basal, necessity := bmr.Calc("F", 89.0, 1.77, 26, "lightly_active")
		want := 1712.24
		wantNecessity := 2354.33
		if basal != want {
			t.Errorf("Basal calc error, got %v want %v", basal, want)
		}
		if necessity != wantNecessity {
			t.Errorf("Necessity calc error, got %v want %v", necessity, wantNecessity)
		}

	})

}

func TestBMRModeratelyActive(t *testing.T) {

	t.Run("BMR Calc - Male Moderately Active", func(t *testing.T) {

		basal, necessity := bmr.Calc("M", 89.0, 1.77, 26, "moderately_active")
		want := 1997.95
		wantNecessity := 3096.8225
		if basal != want {
			t.Errorf("got %v want %v", basal, want)
		}
		if necessity != wantNecessity {
			t.Errorf("Necessity calc error, got %v want %v", necessity, wantNecessity)
		}

	})

	t.Run("BMR Calc - Female Moderately Active", func(t *testing.T) {

		basal, necessity := bmr.Calc("F", 89.0, 1.77, 26, "moderately_active")
		want := 1712.24
		wantNecessity := 2653.972
		if basal != want {
			t.Errorf("Basal calc error, got %v want %v", basal, want)
		}
		if necessity != wantNecessity {
			t.Errorf("Necessity calc error, got %v want %v", necessity, wantNecessity)
		}

	})

}

func TestBMRVeryActive(t *testing.T) {

	t.Run("BMR Calc - Male Very Active", func(t *testing.T) {

		basal, necessity := bmr.Calc("M", 89.0, 1.77, 26, "very_active")
		want := 1997.95
		wantNecessity := 3446.4637500000003
		if basal != want {
			t.Errorf("got %v want %v", basal, want)
		}
		if necessity != wantNecessity {
			t.Errorf("Necessity calc error, got %v want %v", necessity, wantNecessity)
		}

	})

	t.Run("BMR Calc - Female Very Active", func(t *testing.T) {

		basal, necessity := bmr.Calc("F", 89.0, 1.77, 26, "very_active")
		want := 1712.24
		wantNecessity := 2953.614
		if basal != want {
			t.Errorf("Basal calc error, got %v want %v", basal, want)
		}
		if necessity != wantNecessity {
			t.Errorf("Necessity calc error, got %v want %v", necessity, wantNecessity)
		}

	})

}

func TestBMRExtraActive(t *testing.T) {

	t.Run("BMR Calc - Male Extra Active", func(t *testing.T) {

		basal, necessity := bmr.Calc("M", 89.0, 1.77, 26, "extra_active")
		want := 1997.95
		wantNecessity := 3796.105
		if basal != want {
			t.Errorf("got %v want %v", basal, want)
		}
		if necessity != wantNecessity {
			t.Errorf("Necessity calc error, got %v want %v", necessity, wantNecessity)
		}

	})

	t.Run("BMR Calc - Female Extra Active", func(t *testing.T) {

		basal, necessity := bmr.Calc("F", 89.0, 1.77, 26, "extra_active")
		want := 1712.24
		wantNecessity := 3253.256
		if basal != want {
			t.Errorf("Basal calc error, got %v want %v", basal, want)
		}
		if necessity != wantNecessity {
			t.Errorf("Necessity calc error, got %v want %v", necessity, wantNecessity)
		}

	})

}

func TestIMC(t *testing.T) {

	t.Run("IMC Calc - Underweight", func(t *testing.T) {

		imcValue, class := imc.Calc(50, 1.77)
		wantImc := 15.959653994701394
		if imcValue != wantImc {
			t.Errorf("got %v want %v", imcValue, wantImc)
		}

		if class != "underweight" {
			t.Errorf("got %v want %v", class, "underweight")
		}

	})

	t.Run("IMC Calc - Overweight", func(t *testing.T) {

		imcValue, class := imc.Calc(89, 1.77)
		wantImc := 28.40818411056848
		if imcValue != wantImc {
			t.Errorf("got %v want %v", imcValue, wantImc)
		}

		if class != "overweight" {
			t.Errorf("got %v want %v", class, "overweight")
		}

	})

	t.Run("IMC Calc - Healthy", func(t *testing.T) {

		imcValue, class := imc.Calc(77, 1.77)
		wantImc := 24.577867151840145
		if imcValue != wantImc {
			t.Errorf("got %v want %v", imcValue, wantImc)
		}

		if class != "healthy" {
			t.Errorf("got %v want %v", class, "healthy")
		}

	})

	t.Run("IMC Calc - Obese", func(t *testing.T) {

		imcValue, class := imc.Calc(100, 1.77)
		wantImc := 31.91930798940279
		if imcValue != wantImc {
			t.Errorf("got %v want %v", imcValue, wantImc)
		}

		if class != "obese" {
			t.Errorf("got %v want %v", class, "obese")
		}

	})

}

func TestWaterRecomendations(t *testing.T) {

	t.Run("Water Recomendations", func(t *testing.T) {
		recomendations := water.Calc(90)
		want := 3150.0
		if recomendations != want {
			t.Errorf("got %v want %v", recomendations, want)
		}
	})

}

func TestProteinRecomendations(t *testing.T) {

	t.Run("Proteins Recomendations", func(t *testing.T) {
		recomendations := proteins.Calc(90)
		want := 180
		if recomendations != want {
			t.Errorf("got %v want %v", recomendations, want)
		}
	})

}

func TestCaloriesRecomendations(t *testing.T) {

	t.Run("Calories Recomendations to Maintain", func(t *testing.T) {
		recomendations := calories.Maintain(1000.00)
		want := 1000.00
		if recomendations != want {
			t.Errorf("got %v want %v", recomendations, want)
		}
	})

	t.Run("Calories Recomendations to Loss", func(t *testing.T) {
		recomendations := calories.Loss(1000.00)
		want := 900.00
		if recomendations != want {
			t.Errorf("got %v want %v", recomendations, want)
		}
	})

	t.Run("Calories Recomendations to Gain", func(t *testing.T) {
		recomendations := calories.Gain(1000.00)
		want := 1300.00
		if recomendations != want {
			t.Errorf("got %v want %v", recomendations, want)
		}
	})

}

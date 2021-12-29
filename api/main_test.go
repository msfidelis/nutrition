package main

import (
	"testing"

	"github.com/msfidelis/nutrition/pkg/bmr"
	"github.com/msfidelis/nutrition/pkg/imc"
)

func TestBMR(t *testing.T) {

	t.Run("BMR Calc - Male", func(t *testing.T) {

		basal := bmr.Calc("M", 89.0, 1.77, 26)
		want := 1997.95
		if basal != want {
			t.Errorf("got %v want %v", basal, want)
		}

	})

	t.Run("BMR Calc - Female", func(t *testing.T) {

		basal := bmr.Calc("F", 89.0, 1.77, 26)
		want := 1712.24
		if basal != want {
			t.Errorf("got %v want %v", basal, want)
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

# Nutritional Calculator Application
![Version](https://img.shields.io/badge/version-v1-blue.svg?cacheSeconds=2592000)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Twitter: fidelissauro](https://img.shields.io/twitter/follow/fidelissauro.svg?style=social)](https://twitter.com/fidelissauro)

> Simple project do calculate healthcare metrics and nutrition suggestions

## Academic References

[Como calcular metabolismo basal e gasto calótico - PTBR](https://www.brdanutricao.com.br/calculadora-taxa-metabolica-basal-tmb/#Variacoes_da_Calculadora_Taxa_Metabolica_Basal_TMB)

[How many calories you need everyday(https://www.healthline.com/health/what-is-basal-metabolic-rate#How-many-calories-you-need-everyday-) 

[Calorie Calculator](https://www.calculator.net/calorie-calculator.html)


## Usage 

```bash
curl --location --request POST '0.0.0.0:8080/calculator' \
--header 'Content-Type: application/json' \
--data-raw '{ 
   "age": 26,
   "weight": 90.0,
   "height": 1.77,
   "gender": "M", 
   "activity_intensity": "very_active"
} ' --silent | jq .
```

```json
{
    "status": 200,
    "imc": {
        "result": 28.72737719046251,
        "class": "overweight"
    },
    "basal": {
        "bmr": {
            "value": 2011.7,
            "unit": "kcal"
        },
        "necessity": {
            "value": 3470.1825000000003,
            "unit": "kcal"
        }
    },
    "health_info": {
        "age": 26,
        "weight": 90,
        "height": 1.77,
        "gender": "M",
        "activity_intensity": "very_active"
    },
    "recomendations": {
        "protein": {
            "value": 180,
            "unit": "g"
        },
        "water": {
            "value": 3150,
            "unit": "ml"
        },
        "calories": {
            "maintain_weight": {
                "value": 3470.1825000000003,
                "unit": "kcal"
            },
            "loss_weight": {
                "value": 3123.1642500000003,
                "unit": "kcal"
            },
            "gain_weight": {
                "value": 5205.27375,
                "unit": "kcal"
            }
        }
    }
}
```


### Swagger

Access Swagger in `http://0.0.0.0/swagger/index.html`

Swagger runs on production build on `Dockerfile`. 

## Author

👤 **Matheus Fidelis**

* Website: https://raj.ninja
* Twitter: [@fidelissauro](https://twitter.com/fidelissauro)
* Github: [@msfidelis](https://github.com/msfidelis)
* LinkedIn: [@msfidelis](https://linkedin.com/in/msfidelis)

## 🤝 Contributing

Contributions, issues and feature requests are welcome!

Feel free to check [issues page](/issues). 

## Show your support

Give a ⭐️ if this project helped you!


## 📝 License

Copyright © 2020 [Matheus Fidelis](https://github.com/msfidelis).

This project is [MIT](LICENSE) licensed.

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_
# Nutritional Calculator Application
![Version](https://img.shields.io/badge/version-v1-blue.svg?cacheSeconds=2592000)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Twitter: fidelissauro](https://img.shields.io/twitter/follow/fidelissauro.svg?style=social)](https://twitter.com/fidelissauro)

> Simple project do calculate healthcare metrics and nutrition suggestions

## Academic References


## Usage 

```bash
curl --location --request POST '0.0.0.0:8080/calculator' \
--header 'Content-Type: application/json' \
--data-raw '{
   "age": 26,
   "weight": 89.0,
   "height": 1.77,
   "gender": "M"
} ' --silent | jq .
```

```json
{
  "status": 200,
  "imc": {
    "result": 28.40818411056848,
    "class": "overweight"
  },
  "basal": {
    "bmr": 1997.95
  },
  "health_info": {
    "age": 26,
    "weight": 89,
    "height": 1.77,
    "gender": "M"
  },
  "recomendations": {
    "protein": 178,
    "water": 3115
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
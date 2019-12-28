package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"../handler"
	"../models"
)

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(
			"Content-Type",
			"application/json",
		)

		var (
			res models.ResponseResult
		)

		tokenString := r.Header.Get("Authorization")
		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method")
				}

				return []byte(secretKey), nil
			},
		)
		if err != nil {
			res.Error = err.Error()
			_ = json.NewEncoder(w).Encode(res)
			return
		}

		if !token.Valid {
			res.Error = "Token is invalid!"
			_ = json.NewEncoder(w).Encode(res)
		}

		h.ServeHTTP(w, r)
	})
}

func GetWeather(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	const url string = `https://api.openweathermap.org/data/2.5/weather?id=%d&appid=b725faec3be149931cdf9b6773e4f321&lang=ru`
	requestToOWM(env, w, r, url)
	return nil
}

func requestToOWM(env *handler.Env, w http.ResponseWriter, r *http.Request, url string) error {
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	var (
		res         models.ResponseResult
		requestArgs models.OWMRequestArgs
	)

	requestBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(requestBody, &requestArgs)
	if err != nil {
		res.Error = err.Error()
		_ = json.NewEncoder(w).Encode(res)
		return nil
	}

	var cityId = requestArgs.CityId
	var currentWeatherURL = fmt.Sprintf(
		url,
		cityId,
	)
	remoteApiResponse, err := http.Get(currentWeatherURL)
	if err != nil {
		res.Error = err.Error()
		_ = json.NewEncoder(w).Encode(res)
		return nil
	}

	var responseBody interface{}
	body, _ := ioutil.ReadAll(remoteApiResponse.Body)
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		res.Error = err.Error()
		_ = json.NewEncoder(w).Encode(res)
		return nil
	}

	_ = json.NewEncoder(w).Encode(responseBody)
	return nil
}

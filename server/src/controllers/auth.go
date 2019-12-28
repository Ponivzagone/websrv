package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"../handler"
	"../models"
	"golang.org/x/crypto/bcrypt"
)

const secretKey = "asd"

func RegisterHandler(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	var (
		res  models.ResponseResult
		user models.User
	)

	log.Printf("Start")
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		res.Error = err.Error()
		_ = json.NewEncoder(w).Encode(res)
		return nil
	}
	log.Printf("response")
	log.Printf("%s", user.Username)

	rows, err := env.DB.Query("SELECT 1 FROM APP.USERS WHERE USERNAME LIKE '$1' AND ACTIVE = B'1'", user.Username)
	if err != nil {
		res.Error = err.Error()
		_ = json.NewEncoder(w).Encode(res)
		return nil
	}

	if !rows.Next() {
		res.Result = "Not Uniq Username"
		_ = json.NewEncoder(w).Encode(res)
		return nil
	}

	txn, err := env.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}

	var id int
	err = env.DB.QueryRow("INSERT INTO APP.USERS(USERNAME,MODIFATE) VALUES($1,CURRENT_TIMESTAMP) RETURNING USERID", user.Username).Scan(&id)
	if err != nil {
		res.Error = err.Error()
		_ = json.NewEncoder(w).Encode(res)
		return nil
	}

	err = env.DB.QueryRow("INSERT INTO APP.PASSWORD(USERID,PASSWORD,MODIFATE) VALUES($1,$2,CURRENT_TIMESTAMP) RETURNING PASSWORDID", id, user.Password).Scan(&id)
	if err != nil {
		res.Error = err.Error()
		_ = json.NewEncoder(w).Encode(res)
		return nil
	}
	defer txn.Commit()

	res.Result = "Registration Successful"
	_ = json.NewEncoder(w).Encode(res)
	log.Printf("Stop")
	return nil
}

func LoginHandler(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	var (
		res  models.ResponseResult
		user models.User
	)

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		res.Error = err.Error()
		_ = json.NewEncoder(w).Encode(res)
		return nil
	}

	rows, err := env.DB.Query("SELECT 1 FROM APP.USERS AS S INNER JOIN APP.PASSWORD AS P ON P.USERID = S.USERID AND P.ACTIVE = S.ACTIVE WHERE S.USERNAME LIKE $1 AND P.PASSWORD LIKE $2 AND S.ACTIVE = B'1'", user.Username, user.Password)
	if err != nil {
		res.Error = err.Error()
		_ = json.NewEncoder(w).Encode(res)
		return nil
	}

	succes := 0
	if !rows.Next() {
		succes = 1
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		5,
	)
	if err != nil {
		res.Error = "Error create token, Try Again"
		_ = json.NewEncoder(w).Encode(res)
		return nil
	}

	var id int
	err = env.DB.QueryRow("INSERT INTO APP.LOGINATTEMPT(USERNAME,PASSWORD,IPNUMBER,BROWSERTYPE,SUCCESS,TOCKEN) VALUES($1,$2,'0','0',$3,$4) RETURNING loginattemptid", user.Username, user.Password, succes, hash).Scan(&id)
	if err != nil {
		res.Error = err.Error()
		_ = json.NewEncoder(w).Encode(res)
		return nil
	}

	user.Token = string(hash)
	user.Password = ""
	_ = json.NewEncoder(w).Encode(user)
	return nil
}

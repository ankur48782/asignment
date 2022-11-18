package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Token struct {
	Token string
}

func CallApi() {

	// var input UserInput
	// if err := ctx.ShouldBindJSON(&input); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"Error": err.Error(),
	// 	})
	// 	return
	// }
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"msg": input,
	// })
	// fmt.Println(input.custId)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failure in loading env file")
	}
	ApiKey := os.Getenv("API_KEY")
	// fmt.Printf(ApiKey)
	url := os.Getenv("GET_TOKEN_URL")
	method := "GET"
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("api-key", ApiKey)
	req.Header.Add("Cookie", "ci_session=fj4v10njc1upd36uj7armtnjjnaabn17")

	// fmt.Println("req is 1: ", req)

	resp, getErr := client.Do(req)

	if getErr != nil {
		log.Fatal(getErr)
	}
	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var token Token
	json.Unmarshal([]byte(body), &token)
	getToken := token.Token
	statusUrl := os.Getenv("GET_MANDATE_STATUS")
	// fmt.Println(statusUrl)
	finalUrl := statusUrl + "c70eb939df40155d44"

	statuReq, err := http.NewRequest(method, finalUrl, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Printf(ApiKey)

	statuReq.Header.Add("Token", getToken)
	//statuReq.Header.Add("api-key", ApiKey)
	statuReq.Header.Add("Cookie", "ci_session=0pbg4mr3cf10bvht2o66g85t9jc4lrc4")
	statuReq.Header["api-key"] = []string{ApiKey}
	StatusRes, err := client.Do(statuReq)
	// fmt.Println("req is: ", statuReq)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer StatusRes.Body.Close()

	reSbody, err := ioutil.ReadAll(StatusRes.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(reSbody))

	// fmt.Println(getToken)
	// fmt.Println(finalUrl)
	// fmt.Println(statusUrl)

}

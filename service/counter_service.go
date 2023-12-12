package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"io"
	"net/http"
	"os"
)

// JsonResult 返回结构
type JsonResult struct {
	Code     int         `json:"code"`
	ErrorMsg string      `json:"errorMsg,omitempty"`
	Data     interface{} `json:"data"`
}

func Code2Id(context iris.Context) {
}

func Login(context iris.Context) {
	handler := context.Request().Header
	loginUrl := os.Getenv("LOGIN_URL")
	if loginUrl == "" {
		fmt.Printf("login url is empty")
		context.WriteString("login url is empty")
		return
	}
	body, _ := json.Marshal(handler)
	fmt.Printf("login url: %s, body: %s", loginUrl, string(body))
	resp, err := http.Post(loginUrl, "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("login error: %s", err.Error())
		context.WriteString(err.Error())
		return
	}
	defer resp.Body.Close()
	resBody, _ := io.ReadAll(resp.Body)
	context.Write(resBody)
}

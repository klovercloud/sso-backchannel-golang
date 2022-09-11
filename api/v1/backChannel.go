package v1

import (
	"backChannel/config"
	"backChannel/dto"
	"backChannel/helper"
	"backChannel/model"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type BackChannelInf interface {
	Create(e echo.Context) error
	Get(e echo.Context) error
}

type BackChannelInstance struct {
}

func BackChannelController() BackChannelInf {
	return new(BackChannelInstance)
}

func (c BackChannelInstance) Create(e echo.Context) error {
	//Business Logic
	var formData dto.BackChannelAuthDto
	err := e.Bind(&formData)
	if err != nil {
		log.Println("[ERROR] JSON bind error: ", err.Error())
		return e.JSON(http.StatusBadRequest, nil)
	}
	serverModel := model.AuthServerRequestModel{
		GrantType:    config.GRANT_AUTH_CODE,
		Code:         formData.Code,
		ClientID:     config.ClientId,
		ClientSecret: config.ClientSecret,
		RedirectUri:  config.RedirectURI,
	}
	toMap, err := helper.ToMap(serverModel, "json")
	if err != nil {
		log.Println("[ERROR] mapConvert error")
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	data := url.Values{}
	for key, value := range toMap {
		strVal := fmt.Sprint(value)
		data.Add(key, strVal)
	}
	encodedData := data.Encode()
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodPost, config.AuthorizeURI+"/"+config.TOKEN, strings.NewReader(encodedData))
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	response, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("[ERROR] ", err.Error())
		}
	}(response.Body)
	var responseModel interface{}
	err = json.NewDecoder(response.Body).Decode(&responseModel)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(response.StatusCode, responseModel)
}
func (c BackChannelInstance) Get(e echo.Context) error {
	infoDto := dto.FrontChannelInfoDto{
		AuthorizeURI: config.AuthorizeURI + "/" + config.AUTH,
		RedirectURI:  config.RedirectURI,
		ClientID:     config.ClientId,
		Scope:        config.SCOPE_OPENID,
	}
	return e.JSON(http.StatusOK, infoDto)

}

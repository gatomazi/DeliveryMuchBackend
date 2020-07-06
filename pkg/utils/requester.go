package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

var messages = map[string]string{
	"not-found":           "Usuário não encontrado",
	"unexpected":          "Algo inesperado aconteceu.",
	"unexpected-read-all": "Não foi possível acessar este recurso",
}

//RequestGenerator -
func RequestGenerator(urlTarget, typeRequest string, payload interface{}) (body interface{}, err error) {
	var req *http.Request

	cli := &http.Client{}
	//GIPHY tava dando erro ao passar espaço na URL
	urlTarget = strings.ReplaceAll(urlTarget, " ", "%20")
	if typeRequest == "POST" {
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(payload)
		req, err = http.NewRequest(typeRequest, urlTarget, reqBodyBytes)
	} else {
		req, err = http.NewRequest(typeRequest, urlTarget, nil)
	}

	if err != nil {
		return nil, err
	}

	res, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	json.NewDecoder(res.Body).Decode(&body)

	return body, nil
}

package dodjo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"io"
	"log"
	"mime/multipart"
	"os"
)

type Products struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Product `json:"results"`
}

type Engagements struct {
	Count    int          `json:"count"`
	Next     interface{}  `json:"next,omitempty"`
	Previous interface{}  `json:"previous,omitempty"`
	Results  []Engagement `json:"results"`
}

type Client struct {
	apiUrl   string
	apiToken string
}

func (client *Client) makeRequest(sMethod, sEndpoint string, baRequestBody []byte, bIsFile bool, mFormDataVars map[string]string, pResponseOutForm interface{}) error {
	logrus.Debug("Making request")
	pRequest := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(pRequest)
	pRequest.SetRequestURI(client.apiUrl + sEndpoint)
	pRequest.Header.Set("Authorization", fmt.Sprintf("Token %s", client.apiToken))
	pRequest.Header.SetMethod(sMethod)
	if !bIsFile {
		pRequest.Header.Set("Content-Type", "application/json")
	} else {
		pMultiPartWriter := multipart.NewWriter(pRequest.BodyWriter())
		pRequest.Header.Set("Content-Type", pMultiPartWriter.FormDataContentType())
		for k, v := range mFormDataVars {
			if k != "file" {
				err := pMultiPartWriter.WriteField(k, v)
				if err != nil {
					logrus.Fatal(err)
				}

			} else {
				w, err := pMultiPartWriter.CreateFormFile(k, v)
				if err != nil {
					logrus.Fatal(err)
				}
				f, err := os.Open(v)
				if err != nil {
					logrus.Fatal(err)
				}
				_, err = io.Copy(w, f)
				if err != nil {
					logrus.Fatal(err)
				}
			}
		}
		err := pMultiPartWriter.Close()
		if err != nil {
			logrus.Panic(err)
		}
	}
	if (sMethod == "POST" || sMethod == "PATCH") && !bIsFile {
		if len(baRequestBody) > 0 {
			pRequest.SetBody(baRequestBody)
		} else {
			return errors.New("POST/PATCH without a body")
		}
	}
	pClient := &fasthttp.Client{}
	pResponse := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(pResponse)
	err := pClient.Do(pRequest, pResponse)
	if err != nil {
		panic(err)
	}
	iStatusCode := pResponse.StatusCode()
	if iStatusCode != fasthttp.StatusOK && iStatusCode != fasthttp.StatusCreated && iStatusCode != 404 {
		var mBodyJson map[string]interface{}
		err = json.Unmarshal(pResponse.Body(), &mBodyJson)
		if err != nil {
			log.Fatal(err)
		}
		sDetailResp := ""
		_, ok := mBodyJson["detail"]
		if ok {
			sDetailResp = mBodyJson["detail"].(string)
		}
		logrus.Debugf("Cannot successfully process response.\nError:\n%s", pResponse.Body())
		return errors.New(fmt.Sprintf("cannot parse response. Status code = %d. Detail: %s", iStatusCode, sDetailResp))
	}
	err = json.Unmarshal(pResponse.Body(), pResponseOutForm)
	if err != nil {
		return err
	}
	return nil
}

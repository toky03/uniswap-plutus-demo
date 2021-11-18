package adapter

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/toky03/oracle-swap-demo/model"
)

const (
	restEndpoint = "/api/contract/instance"
)

type oracleAdapter struct {
	baseUrl string
}


func CrateAdapter() *oracleAdapter {
	url := os.Getenv("ORACLE_URL")
	port := os.Getenv("ORACLE_PORT")
	if url == "" {
		url = "http://127.0.0.1"
	}
	if port == "" {
		port = "9080"
	}
	return &oracleAdapter{
		baseUrl: url + ":" + port + restEndpoint,
	}
}

func (u *oracleAdapter) AddFunds(uuid string, input model.AddPayload) error {
	payload, err := json.Marshal(input)
	if err != nil {
		return err
	}
	return u.createPostRequest("endpoint/add", uuid, strings.NewReader(string(payload)))
}

func (u *oracleAdapter) ClosePool(uuid string, input model.ClosePayload) error {
	payload, err := json.Marshal(input)
	if err != nil {
		return err
	}
	return u.createPostRequest("endpoint/close", uuid, strings.NewReader(string(payload)))
}

func (u *oracleAdapter) CreatePool(uuid string, input model.CreatePayload) error {
	payload, err := json.Marshal(input)
	if err != nil {
		return err
	}
	return u.createPostRequest("endpoint/create", uuid, bytes.NewBuffer(payload))
}

func (u *oracleAdapter) ReadFunds(uuid string) error {
	return u.createPostRequest("endpoint/funds", uuid, strings.NewReader("[]"))
}

func (u *oracleAdapter) ReadPools(uuid string) error {
	return u.createPostRequest("endpoint/pools", uuid, strings.NewReader("[]"))
}

func (u *oracleAdapter) RemoveFunds(uuid string, input model.RemovePayload) error {
	payload, err := json.Marshal(input)
	if err != nil {
		return err
	}
	return u.createPostRequest("endpoint/remove", uuid, bytes.NewBuffer(payload))
}

func (u *oracleAdapter) Swap(uuid string, input model.SwapPayload) error {
	payload, err := json.Marshal(input)
	if err != nil {
		return err
	}
	return u.createPostRequest("endpoint/swap", uuid, bytes.NewBuffer(payload))
}

func (u *oracleAdapter) ReadStatusPool(uuid string) (model.WalletStatusPool, error) {
	url := fmt.Sprintf("%s/%s/status", u.baseUrl, uuid)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return model.WalletStatusPool{}, err
	}
	req.Header.Add("accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return model.WalletStatusPool{}, err
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return model.WalletStatusPool{}, err
	}
	var walletStatus model.WalletStatusPool
	err = json.Unmarshal(body, &walletStatus)
	res.Body.Close()
	return walletStatus, nil
}


func (u *oracleAdapter) ReadStatus(uuid string) (model.WalletStatus, error) {

	url := fmt.Sprintf("%s/%s/status", u.baseUrl, uuid)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return model.WalletStatus{}, err
	}
	req.Header.Add("accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return model.WalletStatus{}, err
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return model.WalletStatus{}, err
	}
	var walletStatus model.WalletStatus
	err = json.Unmarshal(body, &walletStatus)
	res.Body.Close()
	return walletStatus, nil

}

func (u *oracleAdapter) createPostRequest(endpoint, uuid string, payload io.Reader) error {
	url := fmt.Sprintf("%s/%s/%s", u.baseUrl, uuid, endpoint)
	retries := 5
	var response *http.Response
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return err
	}
	req.Header.Add("content-type", "application/json")

	for retries > 0 {
		response, err = http.DefaultClient.Do(req)
		if err != nil {
			log.Println(err)
			retries -= 1
		} else if response.StatusCode != http.StatusOK {
			body, _ := ioutil.ReadAll(response.Body)
			err = errors.New("Status does not match expectation of 200 actual status is: " + response.Status + " content " + string(body))
			log.Println(err)
			retries -= 1
		} else {
			break
		}
		time.Sleep(3 * time.Second)
	}
	return err
}

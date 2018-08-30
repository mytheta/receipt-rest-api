package main

import (
	"bytes"
	"encoding/json"
	"github.com/hikaru7719/receipt-rest-api/domain/model"
	"github.com/hikaru7719/receipt-rest-api/infrastructure/datastore"
	"github.com/hikaru7719/receipt-rest-api/interface/server/form"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestAPI(t *testing.T) {
	datastore.CreateConnection(datastore.GetDBEnv())
	r := router()
	testServer := httptest.NewServer(r)
	client := new(http.Client)
	testGetReceipt(t, client, testServer)
	testPostReceipt(t, client, testServer)
	testDeleteReceipt(t, client, testServer)
	testPostCredit(t, client, testServer)
}

func testGetReceipt(t *testing.T, client *http.Client, testServer *httptest.Server) {
	testData, _ := model.NewReceipt(1, "test", 1000, "日用品", "2018-08-08", "memo", 1)
	datastore.DB.Create(testData)
	id := strconv.Itoa(testData.ID)
	req, _ := http.NewRequest("GET", testServer.URL+"/v1/receipt/"+id, nil)
	res, _ := client.Do(req)

	if res.StatusCode != 200 {
		t.Error("Get Request /v1/receipt/:id Not Working")
	}
}

func testPostReceipt(t *testing.T, client *http.Client, testServer *httptest.Server) {
	testData := &form.ReceiptForm{UserID: 2, Name: "test", Price: 1000, Kind: "日用品", Date: "2018-08-08", Memo: "test", CreditID: 2}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(testData)
	req, _ := http.NewRequest("POST", testServer.URL+"/v1/receipt", buf)
	req.Header.Set("Content-Type", "application/json")
	res, _ := client.Do(req)

	if res.StatusCode != 201 {
		t.Error("Post Request /v1/receipt Not Working")
	}
}

func testDeleteReceipt(t *testing.T, client *http.Client, testServer *httptest.Server) {
	testData, _ := model.NewReceipt(3, "test", 1000, "日用品", "2018-08-08", "memo", 1)
	datastore.DB.Create(testData)
	id := strconv.Itoa(testData.ID)
	req, _ := http.NewRequest("DELETE", testServer.URL+"/v1/receipt/"+id, nil)
	res, _ := client.Do(req)

	if res.StatusCode != 202 {
		t.Error("Delete Request /v1/receipt/:id Not Working")
	}

}

func testPostCredit(t *testing.T, client *http.Client, testServer *httptest.Server) {
	testData := &form.CreditForm{UserID: 1, CardName: "アメリカンエクスプレス", FinishDate: 10, WithdrawalDate: 4, LaterPaymentMonth: 1}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(testData)
	req, _ := http.NewRequest("POST", testServer.URL+"/v1/credit", buf)
	req.Header.Set("Content-Type", "application/json")
	res, _ := client.Do(req)

	if res.StatusCode != 201 {
		t.Error("Post Request /v1/credit Not Working")
	}
}

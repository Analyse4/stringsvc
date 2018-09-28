package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestStringservice_Uppercase(t *testing.T) {

	//statusCode := 200
	var url = "http://localhost:9001/uppercase"
	data := map[string]string{"s": "slimshady"}
	jsondata, _ := json.Marshal(data)
	correctdata := "SLIMSHADY"
	var res uppercaseResponse

	t.Log("Given the need to uppercase API.")
	{
		t.Logf("\tWhen checking \"%s\" for the same string", url)
		{
			resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsondata))
			if err != nil {
				t.Fatal("\t\tShould be able to make the Post call.", ballotX, err)
			}
			t.Log("\t\tShould be able to make the Post call.", checkMark)
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatal("\t\tShould be able to make the read call.", ballotX, err)
			}
			t.Log("\t\tShould be able to make the read call.", checkMark)
			err = json.Unmarshal(body, &res)
			if err != nil {
				t.Fatal("\t\tjson unmarshal failed.", ballotX, err)
			}

			if res.Err != "" {
				t.Errorf("\t\tShould receive a nil err. %v", ballotX)
			}
			if res.V == correctdata {
				t.Logf("\t\tShould receive string SLIMSHADY, actually receive %s. %v", res.V, checkMark)
			} else {
				t.Errorf("\t\tShould receive string SLIMSHADY, actually receive %s. %v", res.V, ballotX)
			}
		}
	}
}

func TestStringservice_Count(t *testing.T) {
	var url = "http://localhost:9001/count"
	data := map[string]string{"s": "slimshady"}
	jsondata, err := json.Marshal(data)
	if err != nil {
		t.Fatal("Marshal test data is failed", ballotX, err)
	}
	correctdata := "9"
	var res countResponse

	t.Logf("\tWhen checking \"%s\" for the same string", url)
	{

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsondata))
		if err != nil {
			t.Fatal("\t\tShould be able to make the Post call.", ballotX, err)
		}
		t.Log("\t\tShould be able to make the Post call.", checkMark)
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("\t\tShould be able to make the read call.", ballotX, err)
		}
		err = json.Unmarshal(body, &res)
		if err != nil {
			t.Fatal("\t\tjson unmarshal failed.", ballotX, err)
		}
		if res.V != correctdata {
			t.Errorf("\t\tShould receive int %s, actually receive %s. %v", correctdata, res.V, ballotX)
		} else {
			t.Logf("\t\tShould receive int %s, actually receive %s. %v", correctdata, res.V, checkMark)
		}
	}
}

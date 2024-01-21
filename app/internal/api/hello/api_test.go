package hello_api

import (
	"testing"

	cmp "github.com/google/go-cmp/cmp"

	"encoding/json"
	"io"
	"net/http"
)

// http://localhost:3000/ にGETでアクセスし、戻り値を検証する
func TestHello(t *testing.T) {
	res, err := http.Get("http://localhost:3000/")
	if err != nil {
			t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
			t.Errorf("want %d, got %d", http.StatusOK, res.StatusCode)
	}

	resBodyByte, _ := io.ReadAll(res.Body)
	var actual = &Hello{}
	json.Unmarshal(resBodyByte, &actual)

	var expect = &Hello{
		Message: "Hello, World!",
	}

	if !cmp.Equal(actual, expect) {
		t.Errorf("expected %s, actual %s", expect, actual)		
	}
}
package api2captcha

import (
	"context"
	"log"
	"testing"
)

func TestClient_GetBalance(t *testing.T) {

}

func TestClient_Funcaptcha(t *testing.T) {

}

func TestClient_ImageToText(t *testing.T) {

	cli := NewClient("APIKEY")
	rs, err := cli.ImageToText(context.TODO(), "Image Base64")
	if err != nil {
		panic(err)
	}

	log.Println(rs.Text)
}

package whois_test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/linzeyan/whois"
)

const domain string = "google.com"

func TestWhoisXMLApi(t *testing.T) {
	var data whois.WhoisXML
	resp, err := data.Request(domain)
	if err != nil {
		log.Println(err)
		return
	}
	resp.Json()
}

func TestIp2Whois(t *testing.T) {
	var data whois.Ip2Whois
	resp, err := data.Request(domain)
	if err != nil {
		log.Println(err)
		return
	}
	resp.Json()
}

func TestWhoApi(t *testing.T) {
	var data whois.WhoApi
	resp, err := data.Request(domain)
	if err != nil {
		log.Println(err)
		return
	}
	resp.Yaml()
}

func TestApiNinjas(t *testing.T) {
	var data whois.ApiNinjas
	resp, err := data.Request(domain)
	if err != nil {
		log.Println(err)
		return
	}
	resp.Json()
}

func TestVerisign(t *testing.T) {
	var data whois.Verisign
	resp, err := data.Request(domain)
	if err != nil {
		log.Println(err)
		return
	}
	resp.String()
	resp.Json()
	resp.Yaml()
}

func TestIana(t *testing.T) {
	result, err := whois.RequestIana(domain)
	if err != nil {
		log.Println(err)
		return
	}
	out, err := json.MarshalIndent(whois.ParseIana(result), "", "  ")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(out))
}

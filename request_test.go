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
	resp := whois.Request(data, domain)
	resp.Json()
}

func TestIp2Whois(t *testing.T) {
	var data whois.Ip2Whois
	resp := whois.Request(data, domain)
	resp.Json()
}

func TestWhoApi(t *testing.T) {
	var data whois.WhoApi
	resp := whois.Request(data, domain)
	resp.Yaml()
}

func TestApiNinjas(t *testing.T) {
	var data whois.ApiNinjas
	resp := whois.Request(data, domain)
	resp.Json()
}

func TestVerisign(t *testing.T) {
	var data whois.Verisign
	resp := whois.Request(data, domain)
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

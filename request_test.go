package whois_test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/linzeyan/whois"
)

const domain string = "google.com"

var data whois.Servers

func TestWhoisXMLApi(t *testing.T) {
	data = &whois.WhoisXML{}
	resp, err := whois.Request(data, domain)
	if err != nil {
		log.Println(err)
	}
	resp.Json()
}

func TestIp2Whois(t *testing.T) {
	data = &whois.Ip2Whois{}
	resp, err := whois.Request(data, domain)
	if err != nil {
		log.Println(err)
	}
	resp.Json()
}

func TestWhoApi(t *testing.T) {
	data = &whois.WhoApi{}
	resp, err := whois.Request(data, domain)
	if err != nil {
		log.Println(err)
	}
	resp.Yaml()
}

func TestApiNinjas(t *testing.T) {
	data = &whois.ApiNinjas{}
	resp, err := whois.Request(data, domain)
	if err != nil {
		log.Println(err)
	}
	resp.Json()
}

func TestVerisign(t *testing.T) {
	data = &whois.Verisign{}
	resp, err := whois.Request(data, domain)
	if err != nil {
		log.Println(err)
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

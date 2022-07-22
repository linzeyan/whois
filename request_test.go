package whois_test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/linzeyan/whois"
	"gopkg.in/yaml.v3"
)

const domain string = "google.com"

func TestWhoisXMLApi(t *testing.T) {
	var data whois.WhoisXML
	resp, err := data.Request(domain)
	if err != nil {
		log.Println(err)
		return
	}
	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(out))
}

func TestIp2Whois(t *testing.T) {
	var data whois.Ip2Whois
	resp, err := data.Request(domain)
	if err != nil {
		log.Println(err)
		return
	}
	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(out))
}

func TestWhoApi(t *testing.T) {
	var data whois.WhoApi
	resp, err := data.Request(domain)
	if err != nil {
		log.Println(err)
		return
	}
	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(out))
}

func TestApiNinjas(t *testing.T) {
	var data whois.ApiNinjas
	resp, err := data.Request(domain)
	if err != nil {
		log.Println(err)
		return
	}
	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(out))
}

func TestVerisign(t *testing.T) {
	result, err := whois.RequestVerisign(domain)
	if err != nil {
		log.Println(err)
		return
	}
	out, err := json.MarshalIndent(whois.ParseVerisign(result), "", "  ")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(out))
	a, err := yaml.Marshal(whois.ParseVerisign(result))
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(a))
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

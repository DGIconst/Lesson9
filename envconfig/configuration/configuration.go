package configuration

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"
)

type AllUrl url.URL

type Config struct {
	Port        int      `json:"port"`
	DbURL       AllUrl   `json:"db_url"`
	JaegerURL   AllUrl   `json:"jaeger_url"`
	SentryURL   AllUrl   `json:"sentry_url"`
	SomeAppID   string   `json:"some_app_id"`
	SomeAppKey  string   `json:"some_app_key"`
	KafkaBroker []string `json:"kafka_broker"`
}

func (u *AllUrl) UnmarshalJSON(value []byte) error {
	val, err := url.Parse(strings.Replace(string(value), "\"", "", -1))
	if err != nil {
		return err
	}

	*u = AllUrl(*val)
	return nil
}

func GetConfig() Config {
	date, err := os.ReadFile("configuration/data.json")
	if err != nil {
		panic(err)
	}

	var g Config

	if err = json.Unmarshal(date, &g); err != nil {
		panic(err)
	}

	fmt.Println(g)
	return Config{}
}

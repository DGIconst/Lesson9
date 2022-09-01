package configuration

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/kelseyhightower/envconfig"
)

type UrlMap []string

func (u *UrlMap) Decode(value string) error {
	for _, u := range *u {
		_, err := url.Parse(u)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	*u = strings.Split(value, ",")
	return nil
}

type KafkaMap map[string]int

type Config struct {
	Port         int
	Id           string
	Key          string
	Url          UrlMap
	Kafka_broker KafkaMap
}

func GetConfig() Config {
	var g Config
	err := envconfig.Process("myapp", &g)
	if err != nil {
		log.Fatal(err.Error())
	}
	format := "Port: %d\nId: %s\nKey: %s\n"
	_, err = fmt.Printf(format, g.Port, g.Id, g.Key)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Url:")
	for _, u := range g.Url {
		fmt.Printf("%s\n", u)
	}
	fmt.Printf("Kafka_broker:")
	for k, v := range g.Kafka_broker {
		fmt.Printf("  %s: %d\n", k, v)
	}
	return Config{}
}

//  export MYAPP_PORT=8080
//  export MYAPP_URL="db_url: postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable,jaeger_url: http://jaeger:16686,sentry_url: http://sentry:9000"
//  export MYAPP_KAFKA_BROKER="kafka:9092"
//  export MYAPP_ID=testid
//  export MYAPP_KEY=testkey

// сломал всю голову перечитал пересмотрел и только больше запутался c
// проверкой валидности Url мне кажеться что я перемудрил (())
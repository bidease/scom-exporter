package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"

	yml "gopkg.in/yaml.v2"
)

type config struct {
	WebMetricsAddr  string `yaml:"web_metrics_addr"`
	MetricsEndpoint string `yaml:"metrics_endpoint"`
	ScrapeInterval  int64  `yaml:"scrape_interval"`
	Auth            struct {
		Email string
		Token string
	}
}

var conf config

func (conf *config) Read(f string) {
	if !path.IsAbs(f) && f[:1] == "~" {
		f = path.Join(os.Getenv("HOME"), f[1:])
	}

	bytes, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatalf("Read file %s is failed: %s", f, err)
	}

	err = yml.Unmarshal(bytes, conf)
	if err != nil {
		log.Fatalf("Read config is failed: %s", err)
	}

	if conf.WebMetricsAddr == "" {
		conf.WebMetricsAddr = "localhost:9999"
		log.Println("Web metrics address:", conf.WebMetricsAddr)
	}

	if conf.MetricsEndpoint == "" {
		conf.MetricsEndpoint = "/metrics"
		log.Println("Metrics endpoint:", conf.MetricsEndpoint)
	}

	if conf.ScrapeInterval == 0 {
		conf.ScrapeInterval = 300
		log.Println("Scrape interval:", conf.ScrapeInterval)
	}
}

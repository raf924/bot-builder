{{ if false }}
package main
{{ end }}

import (
    "flag"
    "gopkg.in/yaml.v2"
    "os"
    cnf "github.com/raf924/bot/v2/pkg/config/connector"
    "github.com/raf924/bot/v2/pkg/connector"
	"log"
	"context"
)

func main() {
    configFile := flag.String("config", os.Getenv("HCCONNECTOR_CONFIG"), "Path to yaml config file: ex /opt/hcbot/config.yaml")
    flag.Parse()
    if *configFile == "" {
        log.Println("Connector quit with error: %v", "No config file")
        os.Exit(1)
    }
    f, err := os.Open(*configFile)
    if err != nil {
        log.Println("Connector quit with error: %v", err)
        os.Exit(1)
    }
    var config cnf.Config
    if err = yaml.NewDecoder(f).Decode(&config); err != nil {
        log.Println("Connector quit with error: %v", err)
        os.Exit(1)
    }
    hcConnector := connector.NewConnector(config)
    if hcConnector == nil {
        log.Println("Connector quit with error: no connector created")
        os.Exit(1)
    }
    if err = hcConnector.Start(context.Background()); err != nil {
        log.Println("Connector quit with error: %v", err)
        os.Exit(1)
    }
    _ = <-hcConnector.Done()
    if hcConnector.Err() != nil {
        log.Println("Connector quit with error: %v", hcConnector.Err())
        os.Exit(1)
    }
}
{{ if false }}
package main
{{ end }}

import (
    "flag"
    "gopkg.in/yaml.v2"
    "os"
    cnf "github.com/raf924/bot/pkg/config/connector"
    "github.com/raf924/bot/pkg/connector"
)

func main() {
configFile := flag.String("config", os.Getenv("HCCONNECTOR_CONFIG"), "Path to yaml config file: ex /opt/hcbot/config.yaml")
flag.Parse()
if *configFile == "" {
panic("No config file")
}
f, err := os.Open(*configFile)
if err != nil {
panic(err)
}
var config cnf.Config
if err = yaml.NewDecoder(f).Decode(&config); err != nil {
panic(err)
}
hcConnector := connector.NewConnector(config)
if hcConnector == nil {
return
}
if err = hcConnector.Start(); err != nil {
panic(err)
}
_ = <-hcConnector.Done()
}
{{- /*gotype: github.com/raf924/bot-builder/cmd/bot-builder.botConfig*/ -}}
//go:generate go get github.com/raf924/bot
//go:generate go get {{ .BotRelay }}
{{ range .CmdModules}}
//go:generate go get {{ . }}
{{ end }}

package main

import (
"flag"
_ "{{ .BotRelay }}"
{{ range .CmdModules}}
_ "{{.}}"
{{end}}
"github.com/raf924/bot/pkg/bot"
botConfig "github.com/raf924/bot/pkg/config/bot"
"gopkg.in/yaml.v2"
"log"
"os"
"time"
)

func main() {
configFile := flag.String("config", os.Getenv("HCBOT_CONFIG"), "Path to yaml config file: ex /opt/hcbot/config.yaml")
flag.Parse()
if *configFile == "" {
panic("No config file")
}
f, err := os.Open(*configFile)
if err != nil {
panic(err)
}
var config botConfig.Config
if err = yaml.NewDecoder(f).Decode(&config); err != nil {
panic(err)
}
hcBot := bot.NewBot(config)
for {
if err = hcBot.Start(); err != nil {
panic(err)
}
<-hcBot.Done()
log.Println("Reconnecting")
time.Sleep(5 * time.Second)
}
}
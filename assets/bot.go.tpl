{{ if false }}
package main
{{ end }}

import (
    "flag"
    "github.com/raf924/bot/pkg/bot"
    botConfig "github.com/raf924/bot/pkg/config/bot"
    "gopkg.in/yaml.v2"
    "log"
    "os"
    "time"
	"context"
)

func main() {
    configFile := flag.String("config", os.Getenv("HCBOT_CONFIG"), "Path to yaml config file: ex /opt/hcbot/config.yaml")
    flag.Parse()
    if *configFile == "" {
        log.Println("Bot quit with error: %v", "No config file")
        os.Exit(1)
    }
    f, err := os.Open(*configFile)
    if err != nil {
        log.Println("Bot quit with error: %v", err)
        os.Exit(1)
    }
    var config botConfig.Config
    if err = yaml.NewDecoder(f).Decode(&config); err != nil {
        panic(err)
    }
    hcBot := bot.NewBot(config)
    for {
        if err = hcBot.Start(context.Background()); err != nil {
            panic(err)
        }
        <-hcBot.Done()
        if hcBot.Err() != nil {
            log.Println("Bot quit with error: %v", hcBot.Err())
        }
        log.Println("Reconnecting")
        time.Sleep(5 * time.Second)
    }
}
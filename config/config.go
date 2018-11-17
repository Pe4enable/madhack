package config

import (
	"log"
	"os"
	"strconv"
	"github.com/spf13/viper"
)

type Config struct {
	Ticker string // BTC
	Port   string //8000
}

func LoadConfig(filePath string) (conf *Config, err error) {
	viper.SetConfigFile(filePath)

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	log.Printf("Read config")

	conf = new(Config)
	conf.Ticker = GetStringEnvValue("TICKER", viper.GetString("ticker"))
	conf.Port = GetStringEnvValue("PORT", ":" + viper.GetString("port"))
	//conf.StartBlockHeight = GetIntEnvValue("START_BLOCK_HEIGHT", viper.GetInt64("block"))
	//conf.Confirmations = GetIntEnvValue("CONFIRM_COUNT", viper.GetInt64("confirmations"))
	//
	//conf.NodeType = GetStringEnvValue("NODE_TYPE", viper.GetString("node.type"))
	//conf.NodeConn = new(rpcclient.ConnConfig)
	//conf.NodeConn.Host = GetStringEnvValue("NODE_HOST_URL", viper.GetString("node.host"))
	//conf.NodeConn.User = GetStringEnvValue("NODE_USER", viper.GetString("node.user"))
	//conf.NodeConn.Pass = GetStringEnvValue("NODE_PASS", viper.GetString("node.pass"))
	//conf.NodeConn.HTTPPostMode = true
	//conf.NodeConn.DisableTLS = true

	log.Printf("Read env config")
	return
}

func GetIntEnvValue(name string, defValue int64) (result int64) {
	if startBlockHeight, ok := os.LookupEnv(name); ok {
		s, err := strconv.ParseInt(startBlockHeight, 10, 64)
		if err != nil {
			log.Printf("ENV-%s not found, used default value %s", name, defValue)
			return defValue
		}
		return s
	} else {
		log.Printf("ENV-%s not found, used default value %s", name, defValue)
		return defValue
	}
}

func GetStringEnvValue(name string, defValue string) (result string) {
	if u, ok := os.LookupEnv(name); ok {
		return u
	} else {
		log.Printf("ENV-%s not found, used default value '%s'", name, defValue)
		return defValue
	}
}

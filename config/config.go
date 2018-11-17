package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Mongo           Mongo
	TimeOffset      TimeOffset
	Port            string
	Interval        int    // Seconds
}

// For debug purposes
type TimeOffset struct {
	Years, Months, Days int
}

type Mongo struct {
	Url        string
	DB         string
	Collection string
}

var config Config

func Load(configFile string) (conf *Config, err error) {
	viper.SetConfigFile(configFile)
	err = viper.ReadInConfig()
	if err != nil {
		log.Panic("Fatal error config file:", err)
	}

	conf = new(Config)

	conf.Mongo.Url = GetStringEnvValue("MONGO_URL", viper.GetString("mongo.url"))
	conf.Mongo.DB = GetStringEnvValue("MONGO_DB", viper.GetString("mongo.db"))
	conf.Mongo.Collection = GetStringEnvValue("MONGO_COLLECTION", viper.GetString("mongo.collection"))

	conf.Interval = GetIntEnvValue("INTERVAL", viper.GetInt("interval"))
	// For debug purposes
	conf.TimeOffset.Years = GetIntEnvValue("TIMEOFFSET_YEARS", viper.GetInt("timeoffset.years"))
	conf.TimeOffset.Months = GetIntEnvValue("TIMEOFFSET_MONTHS", viper.GetInt("timeoffset.months"))
	conf.TimeOffset.Days = GetIntEnvValue("TIMEOFFSET_DAYS", viper.GetInt("timeoffset.days"))

	conf.Port = ":" + GetStringEnvValue("PORT", viper.GetString("port"))
	log.Printf("Loaded config: %#v", conf)
	return
}

func GetIntEnvValue(name string, defValue int) (result int) {
	if startBlockHeight, ok := os.LookupEnv(name); ok {
		s, err := strconv.Atoi(startBlockHeight)
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

func GetStringArrayEncValue(name string, defValue []string) (result []string) {
	if u, ok := os.LookupEnv(name); ok {
		return strings.Fields(u)
	} else {
		log.Printf("ENV-%s not found, used default value '%s'", name, defValue)
		return defValue
	}
}

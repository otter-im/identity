package config

import (
	"fmt"
	"github.com/golang/glog"
	"net"
	"os"
	"strconv"
	"strings"
)

func Environment() string {
	return envString("APP_ENV", "dev")
}

func ServiceHost() string {
	return envString("IDENTITY_SERVICE_HOST", "0.0.0.0")
}

func ServicePort() string {
	return fmt.Sprintf("%v", envInt("IDENTITY_SERVICE_PORT", 50050))
}

func PasswordMinLength() int {
	return envInt("PASSWORD_MIN_LENGTH", 12)
}

func PasswordBcryptCost() int {
	return envInt("PASSWORD_BCRYPT_COST", 12)
}

func PostgresAddress() string {
	return envString("POSTGRES_ADDRESS", "localhost:5432")
}

func PostgresUser() string {
	return envString("POSTGRES_USER", "")
}

func PostgresPassword() string {
	return envString("POSTGRES_PASSWORD", "")
}

func PostgresDatabase() string {
	return envString("POSTGRES_DATABASE", "otter_identity")
}

func RedisNodes() map[string]string {
	nodes := envString("REDIS_NODES", "localhost:6379")
	result := make(map[string]string)
	for _, node := range strings.Split(nodes, ",") {
		host, _, err := net.SplitHostPort(node)
		if err != nil {
			glog.Warningln(err)
			continue
		}
		result[host] = node
	}
	return result
}

func RedisPassword() string {
	return envString("REDIS_PASSWORD", "")
}

func RedisDB() int {
	return envInt("REDIS_DB", 0)
}

func envString(key string, v string) string {
	result, ok := os.LookupEnv(key)
	if !ok {
		return v
	}
	return result
}

func envInt(key string, v int) int {
	resultString := envString(key, fmt.Sprintf("%v", v))
	if resultString == "" {
		return v
	}

	result, err := strconv.Atoi(resultString)
	if err != nil {
		glog.Warningf("invalid environment variable \"%s\", using default of %v\n", key, v)
		return v
	}
	return result
}

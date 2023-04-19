package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"golang-program-structure/common/logging"
	_ "golang-program-structure/docs"
	"golang-program-structure/service"
)

/*
@title Golang-Program-Structure
@version 1.0
@description Golang-Program-Structure
@contact.name Golang-Program-Structure Support
@contact.url http://www.mindinventory.com
@license.name Apache 2.0
@license.url http://www.apache.org/licenses/LICENSE-2.0.html
@BasePath /
@schemes http
*/
func main() {
	// Load the environment vars from a .env file if present
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Output the environment, sorted by var, for debug purposes
	var envVars map[string]string
	envVars, _ = godotenv.Read()
	keys := make([]string, 0, len(envVars))
	for key := range envVars {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Load the config struct with values from the environment without any prefix (i.e. "")
	var config service.Config
	err = envconfig.Process("", &config)
	if err != nil {
		log.Fatal(err)
	}

	// Set up the logger and attach to the config struct
	logging.SetupLogger(config.Env)
	logger := logging.GetLogger()
	config.Logger = logger

	// Output the config for debugging
	logger.Infof("%+v\n", config)

	// Set the router/handler environment level and initialize
	mode := gin.ReleaseMode
	if config.Env == "local" {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)

	// Service HTTP/2 requests using unencrypted method, H2C
	address := fmt.Sprintf(":%v", config.ServicePort)
	h := NewHandler(&config)

	server := &http.Server{
		Addr:    address,
		Handler: h2c.NewHandler(h, &http2.Server{}),
	}

	logger.Infof("Listening on %s", address)
	logger.Fatal(server.ListenAndServe())
}

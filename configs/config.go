package configs

import (
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"github.com/nhannvt/resume/pkg/util"
)

// envPrefix is prefix of system environment sforum uses.
// All system environments have to be named with SFORUM prefix like SFORUM_DYNAMODB_ENDPOINT
const envPrefix = "SFORUM"

var (

	// Stage Name
	stage = "local"

	// API endpoint
	defaultAPIEndpoint = "defaultAPIEndpoint"

	// API Endpoint for Test
	testAPIEndpoint = "testAPIEndpoint"

	// DynamoDB Endpoint
	dynamodbEndpoint = ""

	// ElasticSearch Endpoint
	elasticsearchEndpoint = ""

	// Default context when tag search is done without context parameter.
	// You can define multiple contexts by using comma delimited. test,resume
	defaultSearchContext = "test"

	// Default search hits when tag search is done without hits parameter
	defaultSearchHits = "20"

	// AWS Region
	awsRegion = "ap-northeast-1"

	// AWS AccessKeyId
	awsAccessKeyID = ""

	// AWS SecretAccessKey
	awsSecretAccessKey = ""

	// Default server port
	defaultServerPort = "8080"

	// Default period(sec) to stop its service when server recieves SIGTERM
	defaultTerminationPeriod = "10"
)

// sharedInstance is only struct which is refered from anywhere.
var sharedInstance *config
var once sync.Once

type Config interface {
	Get(key string) string
	GetInt(key string) int
	GetInt64(key string) int64
}

// config struct represents configurations regarding sforum api.
// By default, struct value is set with local variables.
type config struct {
	stage                    string
	defaultAPIEndpoint       string
	dynamodbEndpoint         string
	elasticsearchEndpoint    string
	defaultSearchContext     string
	defaultSearchHits        string
	awsRegion                string
	awsAccessKeyID           string
	awsSecretAccessKey       string
	defaultServerPort        string
	defaultTerminationPeriod string
}

type configOptions func(*config)

// TestAPIEndpoint changes endpoint name to the one for testing.
func TestAPIEndpoint() configOptions {
	return func(c *config) {
		c.defaultAPIEndpoint = testAPIEndpoint
	}
}

// GetConfig returns singleton config struct.
// With configOptions, you can change configurations.
func GetConfig(configOptions ...configOptions) *config {
	once.Do(func() {
		c := config{
			stage:                    stage,
			defaultAPIEndpoint:       defaultAPIEndpoint,
			dynamodbEndpoint:         dynamodbEndpoint,
			elasticsearchEndpoint:    elasticsearchEndpoint,
			defaultSearchContext:     defaultSearchContext,
			defaultSearchHits:        defaultSearchHits,
			awsRegion:                awsRegion,
			awsAccessKeyID:           awsAccessKeyID,
			awsSecretAccessKey:       awsSecretAccessKey,
			defaultServerPort:        defaultServerPort,
			defaultTerminationPeriod: defaultTerminationPeriod,
		}
		for _, opt := range configOptions {
			opt(&c)
		}
		sharedInstance = &c
	})
	return sharedInstance
}

// Get returns string type value of configuraiton with a specified key.
func (c *config) Get(key string) string {
	result := c.Getenv(key)
	if result == "" {
		value := reflect.ValueOf(c).Elem()
		if value.FieldByName(key).IsValid() {
			result = value.FieldByName(key).String()
		}
	}

	return result
}

// GetInt returns int type value of configuraiton with a specified key.
func (c *config) GetInt(key string) int {
	v, _ := strconv.Atoi(c.Get(key))
	return v
}

// GetInt64 returns int64 type value of configuraiton with a specified key.
func (c *config) GetInt64(key string) int64 {
	v, _ := strconv.ParseInt(c.Get(key), 10, 64)
	return v
}

// Getenv returns specified system environment.
// System enviroment key is "SFORUM" + given key which will be converted to snakecase.
func (c *config) Getenv(key string) string {
	return os.Getenv(c.envKey(key))
}

// envKey make key string to get a certain system environment for sforum.
// The given key should be camel case or snake case.
func (c *config) envKey(key string) string {
	return envPrefix + "_" + strings.ToUpper(util.CamelCaseToSnakeCase(key))
}

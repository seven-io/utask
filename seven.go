package seven_utask

import (
	"errors"
	"fmt"
	"github.com/ovh/utask/pkg/plugins/taskplugin"
	"github.com/seven-io/go-client/sms77api"
	"os"
)

//goland:noinspection GoUnusedGlobalVariable
var (
	Plugin = taskplugin.New("seven_utask", "v0.1", exec,
		taskplugin.WithConfig(validConfig, Config{}))
)

type Config struct {
	ApiKey    string
	SmsParams sms77api.SmsBaseParams
}

func validConfig(config interface{}) (err error) {
	cfg := config.(*Config)

	if len(cfg.ApiKey) < 1 {
		cfg.ApiKey = os.Getenv("SEVEN_API_KEY")
	}

	if cfg.ApiKey == "" {
		return errors.New("ApiKey is missing")
	}

	return nil
}

func exec(stepName string, config interface{}, ctx interface{}) (output interface{}, metadata interface{}, err error) {
	cfg := config.(*Config)

	var client = sms77api.New(sms77api.Options{
		ApiKey: cfg.ApiKey,
	})

	var json, smsError = client.Sms.Json(cfg.SmsParams)
	if smsError == nil {
		output = json
		fmt.Println(fmt.Sprintf("%s", json.Success))
	} else {
		err = smsError
		fmt.Println(smsError.Error())
	}

	return output, metadata, err
}

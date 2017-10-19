package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kelseyhightower/envconfig"
)

type MattermostPoster struct {
	MattermostURL   string `envconfig:"MATTERMOST_URL" required:"true"`
	IncomingHookKey string `envconfig:"INCOMING_HOOK_KEY" required:"true"`
	HTTPClient      http.Client
}

func ParseMattermostConfigFromEnv() (MattermostPoster, error) {
	var mattermostPoster MattermostPoster
	err := envconfig.Process("", &mattermostPoster)
	if err != nil {
		return MattermostPoster{}, err
	}
	return mattermostPoster, nil
}

func (mc MattermostPoster) postURL() string {
	return mc.MattermostURL + "/hooks/" + mc.IncomingHookKey
}

func (mc MattermostPoster) PostMessage(msg []byte) error {
	body := bytes.NewReader(msg)
	resp, err := mc.HTTPClient.Post(mc.postURL(), "", body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Status code %d: %s", resp.StatusCode, respBody)
	}
	return nil
}

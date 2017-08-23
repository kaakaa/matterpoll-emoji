package poll

import (
	"fmt"
	"github.com/mattermost/platform/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCommandCorrect(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	message := "What do you gys wanna grab for lunch?"
	emojis := ":pizza: :sushi:"
	c, err := getConfig("sample_conf.json")
	require.Nil(err)
	ps := PollServer{c}

	payload := fmt.Sprintf("token=%s&channel_id=%s&text=\"%s\"%s", c.Token, model.NewId(), message, emojis)
	response := sendHttpRequest(require, &ps, payload)

	assert.Equal(ResponseUsername, response.Username)
	assert.Equal(ResponseIconUrl, response.IconURL)
	assert.Equal(model.COMMAND_RESPONSE_TYPE_IN_CHANNEL, response.ResponseType)
	assert.Equal(message+" #poll", response.Text)
}

func TestCommandWronMessageFormat(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	message := model.NewRandomString(20)
	emojis := ""
	c, err := getConfig("sample_conf.json")
	require.Nil(err)
	ps := PollServer{c}

	payload := fmt.Sprintf("token=%s&channel_id=%s&text=\"%s\"%s", model.NewId(), model.NewId(), message, emojis)
	response := sendHttpRequest(require, &ps, payload)

	assert.Equal(ResponseUsername, response.Username)
	assert.Equal(ResponseIconUrl, response.IconURL)
	assert.Equal(model.COMMAND_RESPONSE_TYPE_EPHEMERAL, response.ResponseType)
	assert.Equal(ErrorTextWrongFormat, response.Text)
}

func TestCommandTokenMissmatch(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	message := "What do you gys wanna grab for lunch?"
	emojis := ":pizza: :sushi:"
	c, err := getConfig("sample_conf.json")
	require.Nil(err)
	ps := PollServer{c}

	payload := fmt.Sprintf("token=%s&channel_id=%s&text=\"%s\"%s", model.NewId(), model.NewId(), message, emojis)
	response := sendHttpRequest(require, &ps, payload)

	assert.Equal(ResponseUsername, response.Username)
	assert.Equal(ResponseIconUrl, response.IconURL)
	assert.Equal(model.COMMAND_RESPONSE_TYPE_EPHEMERAL, response.ResponseType)
	assert.Equal(ErrorTokenMissmatch, response.Text)
}

func TestHeaderMediaTypeWrong(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	message := "What do you gys wanna grab for lunch?"
	emojis := ":pizza: :sushi:"
	c, err := getConfig("sample_conf.json")
	require.Nil(err)
	ps := PollServer{c}

	payload := fmt.Sprintf("token=%s&channel_id=%s&text=\"%s\"%s", c.Token, model.NewId(), message, emojis)
	reader := strings.NewReader(payload)
	r, err := http.NewRequest("POST", "localhost:8505/poll", reader)
	require.Nil(err)
	require.NotNil(r)

	recorder := httptest.NewRecorder()
	ps.PollCmd(recorder, r)
	assert.Equal(http.StatusUnsupportedMediaType, recorder.Code)
}

func TestURLFormat(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	c, err := getConfig("sample_conf.json")
	require.Nil(err)
	ps := PollServer{c}

	payload := "%"
	reader := strings.NewReader(payload)
	r, err := http.NewRequest("POST", "localhost:8505/poll", reader)
	require.Nil(err)
	require.NotNil(r)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()
	ps.PollCmd(recorder, r)
	assert.Equal(http.StatusBadRequest, recorder.Code)
}

func getConfig(path string) (*PollConf, error) {
	p, err := getTestFilePath(path)
	if err != nil {
		return nil, err
	}
	conf, err := LoadConf(p)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func sendHttpRequest(require *require.Assertions, ps *PollServer, payload string) (response *model.CommandResponse) {
	reader := strings.NewReader(payload)

	r, err := http.NewRequest(http.MethodPost, "localhost:8505/poll", reader)
	require.Nil(err)
	require.NotNil(r)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()
	ps.PollCmd(recorder, r)
	response = model.CommandResponseFromJson(recorder.Result().Body)
	require.NotNil(response)
	return
}

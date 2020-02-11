package Feature

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/RobyFerro/go-web/app/config"
	"github.com/RobyFerro/go-web/database"
	"github.com/RobyFerro/go-web/database/model"
	"io/ioutil"
	"net/http"
	"testing"
)

const (
	URL = "http://localhost:8005"
)

var (
	ActiveUser = newUser{
		Name:           "Marco",
		Surname:        "Aurelio",
		Username:       "mAurelio",
		Password:       "tester",
		RepeatPassword: "tester",
	}
)

type newUser struct {
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat-password"`
}

// Execute request to insert new user.
func TestInsertNewUser(t *testing.T) {
	serializedUser, err := json.Marshal(&ActiveUser)
	if err != nil {
		t.Error("Cannot serialize user data")
	}

	nS, err := convertStructInRequestFormat(&serializedUser)
	if err != nil {
		t.Errorf("Failed to parse payload: %s", err.Error())
	}

	resp, err := http.Post(fmt.Sprintf("%s/users", URL), "application/json", nS)
	if err != nil {
		t.Errorf("Failed to execute the request: %s", err.Error())
	}

	if resp != nil {
		defer closeRequest(resp)
		if resp.StatusCode != 200 {
			t.Errorf("Wrong status code: %d", resp.StatusCode)
		}

		assertUserExist(&ActiveUser, t)
	}
}

// Try to execute user login
// TODO: missing user auth test
func TestUserLogin(t *testing.T) {
	user := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		Username: ActiveUser.Username,
		Password: ActiveUser.Password,
	}

	serializedUser, err := json.Marshal(&user)
	if err != nil {
		t.Error("Cannot serialize user data")
	}

	nS, err := convertStructInRequestFormat(&serializedUser)
	if err != nil {
		t.Errorf("Failed to parse payload: %s", err.Error())
	}

	resp, err := http.Post(fmt.Sprintf("%s/login", URL), "application/json", nS)
	if err != nil {
		t.Errorf("Failed to execute the request: %s", err.Error())
	}

	if resp != nil {
		defer closeRequest(resp)

		if resp.StatusCode != 200 {
			t.Errorf("Wrong status code: %d", resp.StatusCode)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("Cannot parse request response: %s", err.Error())
		}

		type AccessToken struct {
			Token string `json:"token"`
		}

		var tk AccessToken
		_ = json.Unmarshal(body, &tk)

	}
}

// Convert structure into buffer.
// Operation required by HTTP request
func convertStructInRequestFormat(n *[]byte) (*bytes.Buffer, error) {
	return bytes.NewBuffer(*n), nil
}

// Check that user exists in Database.
func assertUserExist(u *newUser, t *testing.T) {
	var user model.User
	db := database.ConnectDB(config.Configuration())
	if err := db.Where("username = ?", &u.Username).Find(&user).Error; err != nil {
		t.Error("Cannot find user")
	}

	if user.Username != u.Username {
		t.Error("Cannot find user")
	}
}

// Simple wrapper to close request in on line
func closeRequest(r *http.Response) {
	_ = r.Body.Close()
}

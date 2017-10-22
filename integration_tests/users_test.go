package integration_tests

import (
	"testing"
	"github.com/golang_social_auth/models"
	"encoding/json"
	"net/http"
	"github.com/golang_social_auth/controllers"
	"github.com/stretchr/testify/assert"
	"strings"
)

func TestSignUp(t *testing.T) {
	in := struct {
		Type     string `json:"type"`
		Value    string `json:"value"`
		Password string `json:"password"`
	}{
		Type:     models.ProviderTypeEmail,
		Value:    "siemensptomaster@gmail.com",
		Password: "xyz",
	}
	body, err := json.Marshal(in)
	if err != nil {
		t.Error(err)
	}

	var u map[string]controllers.User
	resp := newRequest(t, http.StatusOK, "POST", "/signup", body)
	json.Unmarshal(resp.Body.Bytes(), &u)

	assert.Equal(t, in.Type, u["user"].Type)
	assert.Equal(t, in.Value, u["user"].Value)
	assert.True(t, len(u["user"].Token) > 0)
}

func TestSignUpMulti(t *testing.T) {
	in := struct {
		Type     string `json:"type"`
		Value    string `json:"value"`
		Password string `json:"password"`
	}{
		Type:     models.ProviderTypeEmail,
		Value:    "siemensptomaster@gmail.com",
		Password: "xyz",
	}
	body, err := json.Marshal(in)
	if err != nil {
		t.Error(err)
	}

	var e map[string]string
	resp := newRequest(t, http.StatusBadRequest, "POST", "/signup", body)
	json.Unmarshal(resp.Body.Bytes(), &e)

	assert.Equal(t, "user is in use", e["message"])
}

func TestSignUpSocial(t *testing.T) {
	in := struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	}{
		Type:  models.ProvideTypeLinkedIn,
		Value: "i_esVlB5FI",
	}
	body, err := json.Marshal(in)
	if err != nil {
		t.Error(err)
	}

	var u map[string]controllers.User
	resp := newRequest(t, http.StatusOK, "POST", "/signup", body)
	json.Unmarshal(resp.Body.Bytes(), &u)

	assert.Equal(t, in.Type, u["user"].Type)
	assert.Equal(t, strings.ToLower(in.Value), u["user"].Value)
	assert.True(t, len(u["user"].Token) > 0)
}

func TestSignUpMultiSocial(t *testing.T) {
	in := struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	}{
		Type:  models.ProvideTypeLinkedIn,
		Value: "i_esVlB5FI",
	}
	body, err := json.Marshal(in)
	if err != nil {
		t.Error(err)
	}

	var e map[string]string
	resp := newRequest(t, http.StatusBadRequest, "POST", "/signup", body)
	json.Unmarshal(resp.Body.Bytes(), &e)

	assert.Equal(t, "user is in use", e["message"])
}

func TestLogin(t *testing.T) {
	in := struct {
		Type     string `json:"type"`
		Value    string `json:"value"`
		Password string `json:"password"`
	}{
		Type:     models.ProviderTypeEmail,
		Value:    "siemensptomaster@gmail.com",
		Password: "xyz",
	}
	body, err := json.Marshal(in)
	if err != nil {
		t.Error(err)
	}

	var u map[string]controllers.User
	resp := newRequest(t, http.StatusOK, "POST", "/login", body)
	json.Unmarshal(resp.Body.Bytes(), &u)

	assert.Equal(t, in.Type, u["user"].Type)
	assert.Equal(t, in.Value, u["user"].Value)
	assert.True(t, len(u["user"].Token) > 0)
}

func TestLoginSocial(t *testing.T) {
	in := struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	}{
		Type:  models.ProvideTypeLinkedIn,
		Value: "i_esVlB5FI",
	}
	body, err := json.Marshal(in)
	if err != nil {
		t.Error(err)
	}

	var u map[string]controllers.User
	resp := newRequest(t, http.StatusOK, "POST", "/login", body)
	json.Unmarshal(resp.Body.Bytes(), &u)

	assert.Equal(t, in.Type, u["user"].Type)
	assert.Equal(t, strings.ToLower(in.Value), u["user"].Value)
	assert.True(t, len(u["user"].Token) > 0)
}

func TestLoginByUnknownType(t *testing.T) {
	in := struct {
		Type     string `json:"type"`
		Value    string `json:"value"`
		Password string `json:"password"`
	}{
		Type:  "unknown",
		Value: "siemensptomaster@gmail.com",
	}
	body, err := json.Marshal(in)
	if err != nil {
		t.Error(err)
	}

	var e map[string]string
	resp := newRequest(t, http.StatusBadRequest, "POST", "/login", body)
	json.Unmarshal(resp.Body.Bytes(), &e)

	assert.Equal(t, "type is invalid", e["message"])
}

func TestLoginByUnknownValue(t *testing.T) {
	in := struct {
		Type     string `json:"type"`
		Value    string `json:"value"`
		Password string `json:"password"`
	}{
		Type:  models.ProviderTypeEmail,
		Value: "unknown",
	}
	body, err := json.Marshal(in)
	if err != nil {
		t.Error(err)
	}

	var e map[string]string
	resp := newRequest(t, http.StatusBadRequest, "POST", "/login", body)
	json.Unmarshal(resp.Body.Bytes(), &e)

	assert.Equal(t, "record not found", e["message"])
}

func TestLoginByEmptyPassword(t *testing.T) {
	in := struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	}{
		Type:  models.ProviderTypeEmail,
		Value: "siemensptomaster@gmail.com",
	}
	body, err := json.Marshal(in)
	if err != nil {
		t.Error(err)
	}

	var e map[string]string
	resp := newRequest(t, http.StatusBadRequest, "POST", "/login", body)
	json.Unmarshal(resp.Body.Bytes(), &e)

	assert.Equal(t, "password invalid", e["message"])
}

func TestLoginByWrongPassword(t *testing.T) {
	in := struct {
		Type     string `json:"type"`
		Value    string `json:"value"`
		Password string `json:"password"`
	}{
		Type:     models.ProviderTypeEmail,
		Value:    "siemensptomaster@gmail.com",
		Password: "wrong",
	}
	body, err := json.Marshal(in)
	if err != nil {
		t.Error(err)
	}

	var e map[string]string
	resp := newRequest(t, http.StatusBadRequest, "POST", "/login", body)
	json.Unmarshal(resp.Body.Bytes(), &e)

	assert.Equal(t, "password invalid", e["message"])
}

func TestUpdateMe(t *testing.T) {
	in := struct {
		Type     string `json:"type"`
		Value    string `json:"value"`
		Password string `json:"password"`
	}{
		Type:     models.ProviderTypeEmail,
		Value:    "siemensptomaster@gmail.com",
		Password: "xyz",
	}
	body, err := json.Marshal(in)
	if err != nil {
		t.Error(err)
	}

	var u map[string]controllers.User
	resp := newRequest(t, http.StatusOK, "POST", "/login", body)
	json.Unmarshal(resp.Body.Bytes(), &u)

	assert.Equal(t, in.Type, u["user"].Type)
	assert.Equal(t, in.Value, u["user"].Value)
	assert.True(t, len(u["user"].Token) > 0)

	user := struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Username  string `json:"username"`
	}{
		FirstName: "ievgen",
		LastName:  "maksymenko",
		Username:  "skews",
	}
	body, err = json.Marshal(user)
	if err != nil {
		t.Error(err)
	}

	resp = newRequestAuth(t, u["user"].Token, http.StatusOK, "PUT", "/user/me", body)
	json.Unmarshal(resp.Body.Bytes(), &u)

	assert.Equal(t, user.FirstName, u["user"].FirstName)
	assert.Equal(t, user.LastName, u["user"].LastName)
	assert.Equal(t, user.Username, u["user"].Username)
}

func TestUpdateMeByUnknownToken(t *testing.T) {
	user := struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Username  string `json:"username"`
	}{
		FirstName: "ievgen",
		LastName:  "maksymenko",
		Username:  "skews",
	}
	body, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}

	var e map[string]string
	resp := newRequestAuth(t, "unknown", http.StatusUnauthorized, "PUT", "/user/me", body)
	json.Unmarshal(resp.Body.Bytes(), &e)

	assert.Equal(t, "Login required", e["message"])

}

func TestPasswordResetByEmail(t *testing.T) {
	in := struct {
		Email string `json:"email"`
	}{
		Email: "siemensptomaster@gmail.com",
	}
	body, err := json.Marshal(in)
	if err != nil {
		t.Error(err)
	}

	newRequest(t, http.StatusOK, "POST", "/password_reset", body)
}

func TestPasswordResetByUnknownEmail(t *testing.T) {
	in := struct {
		Email string `json:"email"`
	}{
		Email: "unknown",
	}
	body, err := json.Marshal(in)
	if err != nil {
		t.Error(err)
	}

	var e map[string]string
	resp := newRequest(t, http.StatusBadRequest, "POST", "/password_reset", body)
	json.Unmarshal(resp.Body.Bytes(), &e)

	assert.Equal(t, "record not found", e["message"])
}

func TestPasswordReset(t *testing.T) {
	in := struct {
		Type     string `json:"type"`
		Value    string `json:"value"`
		Password string `json:"password"`
	}{
		Type:     models.ProviderTypeEmail,
		Value:    "siemensptomaster@gmail.com",
		Password: "xyz",
	}
	body, err := json.Marshal(in)
	if err != nil {
		t.Error(err)
	}

	var u map[string]controllers.User
	resp := newRequest(t, http.StatusOK, "POST", "/login", body)
	json.Unmarshal(resp.Body.Bytes(), &u)

	assert.Equal(t, in.Type, u["user"].Type)
	assert.Equal(t, in.Value, u["user"].Value)

	tokenPreFlight := u["user"].Token
	assert.True(t, len(tokenPreFlight) > 0)

	pass := struct {
		Password string `json:"password"`
	}{
		Password: "xyz1",
	}

	passBody, err := json.Marshal(pass)
	if err != nil {
		t.Error(err)
	}

	newRequest(t, http.StatusOK, "POST", "/"+tokenPreFlight+"/password_reset", passBody)

	in.Password = pass.Password
	body, err = json.Marshal(in)
	if err != nil {
		t.Error(err)
	}
	resp = newRequest(t, http.StatusOK, "POST", "/login", body)
	json.Unmarshal(resp.Body.Bytes(), &u)

	assert.NotEqual(t, tokenPreFlight, u["user"].Token)
}

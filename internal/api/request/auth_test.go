package request

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ew0s/trade-bot/pkg/testhelpers"
	"github.com/stretchr/testify/require"
)

func TestSignIn_Bind_Success(t *testing.T) {
	tests := []struct {
		name           string
		reqQueryParams map[string]string
		reqBody        string
		expected       SignIn
	}{
		{
			name: "successfully bind sign in request",
			reqQueryParams: map[string]string{
				"uid": "123e4567-e89b-12d3-a456-426655440000",
			},
			reqBody: `{"password": "password"}`,
			expected: SignIn{
				UID:      "123e4567-e89b-12d3-a456-426655440000",
				Password: "password",
			},
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			req := testhelpers.MakeTestHTTPRequest(http.MethodPost, "/auth/sign-in", tc.reqBody, tc.reqQueryParams, nil)

			var actual SignIn
			err := actual.Bind(req)

			require.NoError(t, err)
			require.Equal(t, tc.expected, actual)
		})
	}
}

func TestSignIn_Bind_Error(t *testing.T) {
	tests := []struct {
		name               string
		reqQueryParams     map[string]string
		reqBody            string
		expectedErrMessage string
	}{
		{
			name:               "uid param not presented",
			reqBody:            `{"password": "password"}`,
			expectedErrMessage: "getting uid from query: does not contain uid key",
		},
		{
			name: "uid param is invalid",
			reqQueryParams: map[string]string{
				"uid": "",
			},
			reqBody:            `{"password": "password"}`,
			expectedErrMessage: "getting uid from query: validating uid: uuid: incorrect UUID length 0 in string \"\"",
		},
		{
			name: "invalid body passed",
			reqQueryParams: map[string]string{
				"uid": "123e4567-e89b-12d3-a456-426655440000",
			},
			reqBody:            `{"password": }`,
			expectedErrMessage: "decoding body: invalid character '}' looking for beginning of value",
		},
		{
			name: "password is empty",
			reqQueryParams: map[string]string{
				"uid": "123e4567-e89b-12d3-a456-426655440000",
			},
			reqBody:            `{"password": ""}`,
			expectedErrMessage: "validating SignIn: validating struct: password: cannot be blank.",
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			httpReq := testhelpers.MakeTestHTTPRequest(http.MethodPost, "/auth/sign-in", tc.reqBody, tc.reqQueryParams, nil)

			var req SignIn
			actualErr := req.Bind(httpReq)

			require.EqualError(t, actualErr, tc.expectedErrMessage)
		})
	}
}

func TestSignUp_Bind_Success(t *testing.T) {
	tests := []struct {
		name     string
		reqBody  string
		expected SignUp
	}{
		{
			name: "successfully bind sign up request",
			reqBody: `
				{
					"name":     "Joe",
					"username": "Jokovi32",
					"password": "qwerty"
				}
			`,
			expected: SignUp{
				Name:     "Joe",
				Username: "Jokovi32",
				Password: "qwerty",
			},
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			req := testhelpers.MakeTestHTTPRequest(http.MethodPost, "/auth/sign-up", tc.reqBody, nil, nil)

			var actual SignUp
			err := actual.Bind(req)

			require.NoError(t, err)
			require.Equal(t, tc.expected, actual)
		})
	}
}

func TestSignUp_Bind_Error(t *testing.T) {
	tests := []struct {
		name               string
		reqBody            string
		expectedErrMessage string
	}{
		{
			name: "body is invalid",
			reqBody: `
				{
					"name":
				}
			`,
			expectedErrMessage: "decoding body: invalid character '}' looking for beginning of value",
		},
		{
			name: "name field not presented",
			reqBody: `
				{
					"username": "Jokovi32",
					"password": "qwerty"
				}
			`,
			expectedErrMessage: "validating SignUp: validating struct: name: cannot be blank.",
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			httpReq := testhelpers.MakeTestHTTPRequest(http.MethodPost, "/auth/sign-up", tc.reqBody, nil, nil)

			var req SignUp
			actualErr := req.Bind(httpReq)

			require.EqualError(t, actualErr, tc.expectedErrMessage)
		})
	}
}

func TestLogout_Bind_Success(t *testing.T) {
	tests := []struct {
		name                 string
		accessTokenHeaderVal string
		expected             Logout
	}{
		{
			name:                 "successfully bind logout request",
			accessTokenHeaderVal: "123e4567-e89b-12d3-a456-426655440000",
			expected: Logout{
				AccessToken: "123e4567-e89b-12d3-a456-426655440000",
			},
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, "/auth/logout", nil)
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tc.accessTokenHeaderVal))

			var actual Logout
			err := actual.Bind(req)

			require.NoError(t, err)
			require.Equal(t, tc.expected, actual)
		})
	}
}

func TestLogout_Bind_Error(t *testing.T) {
	tests := []struct {
		name                 string
		accessTokenHeaderVal string
		expectedErrMessage   string
	}{
		{
			name:                 "successfully bind logout request",
			accessTokenHeaderVal: "",
			expectedErrMessage:   "getting bearer token: empty bearer token in auth header",
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, "/auth/logout", nil)
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tc.accessTokenHeaderVal))

			var actual Logout
			err := actual.Bind(req)

			require.EqualError(t, err, tc.expectedErrMessage)
		})
	}
}

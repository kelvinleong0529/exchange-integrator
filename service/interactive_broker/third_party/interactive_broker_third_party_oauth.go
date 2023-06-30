package interactive_broker_third_party

import "github.com/go-playground/validator/v10"

// Obtain a request token. See section 6.1 of the OAuth v1.0a specification for more information. http://oauth.net/core/1.0a/#auth_step1
// Note we do not return an oauth_token_secret in the response as we are using RSA signatures rather than PLAINTEXT authentication.
func (i *InteractiveBrokerThirdPartyManager) ObtainRequestToken(queryParam *ObtainRequestTokenReq) (*ObtainRequestTokenRes, error) {
	if err := validator.New().Struct(queryParam); err != nil {
		return nil, err
	}
	result := new(ObtainRequestTokenRes)
	return result, nil
}

// Obtain an access token using the request token and the verification code you received after the user provided authorization
// See section 6.3 of the OAuth v1.0a specification for more details.
func (i *InteractiveBrokerThirdPartyManager) ObtainAccessToken(queryParam *ObtainAccessTokenReq) (*ObtainAccessTokenRes, error) {
	if err := validator.New().Struct(queryParam); err != nil {
		return nil, err
	}
	result := new(ObtainAccessTokenRes)
	return result, nil
}

package provider

import (
	"context"
	"fmt"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/option"
	siteverification "google.golang.org/api/siteverification/v1"
	"io/ioutil"
)

type client struct {
	jsonCredentialsFile string
	impersonateEmail    string
	customerNumber      string
}

func (cl client) jsonCredentialsContent() (content []byte, err error) {
	content, err = ioutil.ReadFile(cl.jsonCredentialsFile)
	if err != nil {
		err = fmt.Errorf("unable to read jsonCredentialsFile %s: %s", cl.jsonCredentialsFile, err)
	}

	return
}

func (cl client) jwtConfig() (jwtConfig *jwt.Config, err error) {
	jsonCredentialsContent, err := cl.jsonCredentialsContent()
	if err != nil {
		err = fmt.Errorf("unable to get jsonCredentialsContent: %s", err)
		return
	}

	jwtConfig, err = google.JWTConfigFromJSON(jsonCredentialsContent)
	if err != nil {
		err = fmt.Errorf("unable to create jwtConfig: %s", err)
		return
	}

	if cl.impersonateEmail != "" {
		jwtConfig.Subject = cl.impersonateEmail
	}

	return
}

func (cl client) jwtConfigWithScopes(scopes ...string) (jwtConfig *jwt.Config, err error) {
	jwtConfig, err = cl.jwtConfig()
	if err != nil {
		err = fmt.Errorf("unable to create jwtConfigWithScopes(%v): %s", scopes, err)
	}

	jwtConfig.Scopes = scopes

	return
}

func (cl client) newAdminServiceWithScopes(ctx context.Context, scopes ...string) (service *admin.Service, err error) {
	jwtConfig, err := cl.jwtConfigWithScopes(scopes...)
	if err != nil {
		err = fmt.Errorf("unable to get jwtConfigWithScopes(%v): %s", scopes, err)
		return
	}

	jwtClient := jwtConfig.Client(ctx)

	service, err = admin.NewService(ctx, option.WithHTTPClient(jwtClient))
	if err != nil {
		err = fmt.Errorf("unable to create admin.NewService: %s", err)
		// this return is unnecessary, but included to be unambiguous that if there were other
		// logic below (other than the main function return), that it should return immediately
		return
	}

	return
}

func (cl client) newSiteVerificationServiceWithScopes(ctx context.Context, scopes ...string) (service *siteverification.Service, err error) {
	jwtConfig, err := cl.jwtConfigWithScopes(scopes...)
	if err != nil {
		err = fmt.Errorf("unable to get jwtConfigWithScopes(%v): %s", scopes, err)
		return
	}

	jwtClient := jwtConfig.Client(ctx)

	service, err = siteverification.NewService(ctx, option.WithHTTPClient(jwtClient))
	if err != nil {
		err = fmt.Errorf("unable to create siteverification.NewService: %s", err)
		// this return is unnecessary, but included to be unambiguous that if there were other
		// logic below (other than the main function return), that it should return immediately
		return
	}

	return
}

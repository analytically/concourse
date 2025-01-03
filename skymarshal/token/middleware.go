package token

import (
	"net/http"
	"time"
)

//counterfeiter:generate . Middleware
type Middleware interface {
	SetAuthToken(http.ResponseWriter, string, time.Time) error
	UnsetAuthToken(http.ResponseWriter)
	GetAuthToken(*http.Request) string

	SetCSRFToken(http.ResponseWriter, string, time.Time) error
	UnsetCSRFToken(http.ResponseWriter)
	GetCSRFToken(*http.Request) string

	SetStateToken(http.ResponseWriter, string, time.Time) error
	UnsetStateToken(http.ResponseWriter)
	GetStateToken(*http.Request) string
}

type middleware struct {
	sameSite      http.SameSite
	secureCookies bool
}

func NewMiddleware(sameSite http.SameSite, secureCookies bool) Middleware {
	return &middleware{sameSite: sameSite, secureCookies: secureCookies}
}

const stateCookieName = "skymarshal_state"
const authCookieName = "skymarshal_auth"
const csrfCookieName = "skymarshal_csrf"

func (m *middleware) UnsetAuthToken(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     authCookieName,
		Path:     "/",
		MaxAge:   -1,
		Secure:   m.secureCookies,
		SameSite: m.sameSite,
		HttpOnly: true,
	})
}

func (m *middleware) SetAuthToken(w http.ResponseWriter, tokenStr string, expiry time.Time) error {
	http.SetCookie(w, &http.Cookie{
		Name:     authCookieName,
		Value:    tokenStr,
		Path:     "/",
		Expires:  expiry,
		HttpOnly: true,
		SameSite: m.sameSite,
		Secure:   m.secureCookies,
	})

	return nil
}

func (m *middleware) GetAuthToken(r *http.Request) string {
	cookie, err := r.Cookie(authCookieName)
	if err != nil {
		return ""
	}
	return cookie.Value
}

func (m *middleware) UnsetCSRFToken(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     csrfCookieName,
		Path:     "/",
		MaxAge:   -1,
		Secure:   m.secureCookies,
		HttpOnly: true,
	})
}

func (m *middleware) SetCSRFToken(w http.ResponseWriter, csrfToken string, expiry time.Time) error {
	http.SetCookie(w, &http.Cookie{
		Name:     csrfCookieName,
		Value:    csrfToken,
		Path:     "/",
		Expires:  expiry,
		Secure:   m.secureCookies,
		HttpOnly: true,
	})

	return nil
}

func (m *middleware) GetCSRFToken(r *http.Request) string {
	cookie, err := r.Cookie(csrfCookieName)
	if err != nil {
		return ""
	}
	return cookie.Value
}

func (m *middleware) UnsetStateToken(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     stateCookieName,
		Path:     "/",
		MaxAge:   -1,
		Secure:   m.secureCookies,
		HttpOnly: true,
	})
}

func (m *middleware) SetStateToken(w http.ResponseWriter, stateToken string, expiry time.Time) error {
	http.SetCookie(w, &http.Cookie{
		Name:     stateCookieName,
		Value:    stateToken,
		Path:     "/",
		Expires:  expiry,
		Secure:   m.secureCookies,
		HttpOnly: true,
	})

	return nil
}

func (m *middleware) GetStateToken(r *http.Request) string {
	cookie, err := r.Cookie(stateCookieName)
	if err != nil {
		return ""
	}
	return cookie.Value
}

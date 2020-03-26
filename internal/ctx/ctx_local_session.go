package ctx

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/henderjon/jwt"
	"github.com/henderjon/logger"
	uuid "github.com/satori/go.uuid"
)

// contexts want you to use custom types ... smells odd #thanksgoogle

type LocalSessionGetter func(ctx context.Context) (*LocalSession, bool)

type KeyLocalSession int

// LocalSession is a set of data for the vasili JWT. LastAccess, Tally
// and PrevTally are all int64 because time.Unix() is int64
type LocalSession struct {
	jwt.RegisteredClaims
}

var LocalSessionKey KeyLocalSession

func SetLocalSession(ctx context.Context, b *LocalSession) context.Context {
	return context.WithValue(ctx, LocalSessionKey, b)
}

func GetLocalSession(ctx context.Context) (*LocalSession, bool) {
	b, ok := ctx.Value(LocalSessionKey).(*LocalSession)
	return b, ok
}

// newLocalSession creates a new local session JWT
func NewLocalSession(exp time.Duration) *LocalSession {
	return &LocalSession{
		RegisteredClaims: jwt.RegisteredClaims{
			JWTID:          uuid.NewV4().String(),
			IssuedAt:       time.Now().UTC().Unix(),
			ExpirationTime: time.Now().UTC().Add(exp).Unix(),
		},
	}
}

// LocalSessionParser defines a func for parsing a local session
type LocalSessionParser func(r *http.Request) (*LocalSession, error)

// NewLocalSessionParser returns a new LocalSessionParser
func NewLocalSessionParser(localCookieName, localCookieSalt string, exp time.Duration) LocalSessionParser {
	signer := jwt.NewHMACSigner(jwt.HS256, []byte(localCookieSalt))
	return LocalSessionParser(func(r *http.Request) (*LocalSession, error) {
		var err error
		// re/generate a new session to use for new cookies
		session := NewLocalSession(exp)

		cookie, err := r.Cookie(localCookieName)
		if err != nil {
			return session, nil // no cookie was found, give us the new one
		}

		// if the cookie exists, set it's value
		err = jwt.Unserialize(cookie.Value, signer, session)
		if err != nil {
			// if the token is invalid, tell us why, but give us a fresh one
			return session, err
		}

		if !session.Valid() {
			// if the token is invalid, tell us why, but give us a fresh one
			return session, errors.New("token expired or not active")
		}

		if err != nil {
			// if the token is invalid, tell us why, but give us a fresh one
			return session, err // if the token is invalid, tell us why, but give us a fresh one
		}

		return session, nil
	})
}

// NewSessionHandler returns a cookie handler
func NewLocalSessionHandler(opts CookieParams, salt string, logger logger.Logger) http.Handler {
	localSessionParser := NewLocalSessionParser(opts.Name, salt, time.Duration(opts.TTL))
	signer := jwt.NewHMACSigner(jwt.HS256, []byte(salt))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// re/generate a new session to use for new cookies
		localSess, err := localSessionParser(r)
		if err == nil {
			// update the cookie's expiration
			localSess.ExpirationTime = time.Now().UTC().Add(opts.TTL).Unix()
		} else {
			// if the token is invalid, create a new LocalSession
			logger.Log(err, true)
		}
		// TODO fail on JWT parse error

		token, err := jwt.Serialize(localSess, signer)
		if err != nil {
			logger.Log(err, true)
		}

		opts.Value = token
		http.SetCookie(w, NewCookie(opts))

		ctxTmp := r.Context()
		ctxTmp = SetLocalSession(ctxTmp, localSess)
		*r = *r.WithContext(ctxTmp)

	})
}

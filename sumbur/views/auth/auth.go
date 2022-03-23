package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"sumbur/views"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"
)

type Config struct {
	Cookie   string
	Password string
	TTL      time.Duration

	Redis redis.Options
}

var (
	config *Config
	rdb    *redis.Client
)

func Init(cfg *Config) {
	config = cfg
	rdb = redis.NewClient(&config.Redis)
}

func Check(ctx *atreugo.RequestCtx) error {
	var err error

	views.AuthState = views.AuthDefault

	session_b := ctx.Request.Header.Cookie(config.Cookie)
	if session_b == nil {
		return ctx.Next()
	}

	session := string(session_b)

	views.AuthState, err = rdb.Get(ctx, session).Int()
	if err != nil {
		if err != redis.Nil {
			panic(err)
		}

		views.AuthState = views.AuthDefault
	}

	if views.AuthState == views.AuthFail {
		rdb.Del(ctx, session)
	}

	return ctx.Next()
}

// Login

type Login struct {
	*views.BasePage
}

func AuthPOST(ctx *atreugo.RequestCtx) error {
	// Session

	var session string

	session_b := ctx.Request.Header.Cookie(config.Cookie)

	if session_b == nil {
		random := make([]byte, 16)
		rand.Read(random)
		session = hex.EncodeToString(random)
	} else {
		session = string(session_b)
	}

	// Cookie

	cookie := fasthttp.AcquireCookie()
	defer fasthttp.ReleaseCookie(cookie)

	cookie.SetKey(config.Cookie)
	cookie.SetValue(session)

	hash := sha256.Sum256(ctx.FormValue("password"))

	if hex.EncodeToString(hash[:]) == config.Password {
		cookie.SetExpire(time.Now().Add(time.Hour * config.TTL))
		rdb.SetEX(ctx, session, views.AuthOK, time.Hour*config.TTL)
	} else {
		cookie.SetExpire(time.Now().Add(time.Second))
		rdb.SetEX(ctx, session, views.AuthFail, time.Hour*config.TTL)
	}

	ctx.Response.Header.SetCookie(cookie)

	// Redirect

	return ctx.RedirectResponse(string(ctx.Referer()), 302)
}

// Restriction

type Forbidden struct {
	*views.BasePage
}

func Restrict(ctx *atreugo.RequestCtx) error {
	if views.AuthState == views.AuthOK {
		return ctx.Next()
	}

	ctx.SetStatusCode(403)
	views.WritePage(ctx, -1, &Forbidden{})

	return nil
}

type Restricted struct {
	*views.BasePage
}

func RestrictedGET(ctx *atreugo.RequestCtx) error {
	views.WritePage(ctx, -1, &Restricted{})
	return nil
}

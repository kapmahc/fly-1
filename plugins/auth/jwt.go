package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/jwt"
	"github.com/google/uuid"
	"github.com/kapmahc/fly/web"
	"github.com/kapmahc/fly/web/i18n"
)

const (
	// TOKEN token session key
	TOKEN = "token"
	// UID uid key
	UID = "uid"
	// CurrentUser current-user key
	CurrentUser = "currentUser"
	// IsAdmin is-admin key
	IsAdmin = "isAdmin"
)

//Jwt jwt helper
type Jwt struct {
	Key    []byte               `inject:"jwt.key"`
	Method crypto.SigningMethod `inject:"jwt.method"`
	Dao    *Dao                 `inject:""`
	I18n   *i18n.I18n           `inject:""`
}

//Validate check jwt
func (p *Jwt) Validate(buf []byte) (jwt.Claims, error) {
	tk, err := jws.ParseJWT(buf)
	if err != nil {
		return nil, err
	}
	if err = tk.Validate(p.Key, p.Method); err != nil {
		return nil, err
	}
	return tk.Claims(), nil
}

func (p *Jwt) parse(r *http.Request) (jwt.Claims, error) {
	tk, err := jws.ParseJWTFromRequest(r)
	if err != nil {
		return nil, err
	}
	if err = tk.Validate(p.Key, p.Method); err != nil {
		return nil, err
	}
	return tk.Claims(), nil
}

//Sum create jwt token
func (p *Jwt) Sum(cm jws.Claims, exp time.Duration) ([]byte, error) {
	kid := uuid.New().String()
	now := time.Now()
	cm.SetNotBefore(now)
	cm.SetExpiration(now.Add(exp))
	cm.Set("kid", kid)
	//TODO using kid

	jt := jws.NewJWT(cm, p.Method)
	return jt.Serialize(p.Key)
}

func (p *Jwt) getUserFromRequest(r *http.Request) (*User, error) {
	lng := r.Context().Value(web.K(i18n.LOCALE)).(string)
	cm, err := p.parse(r)
	if err != nil {
		return nil, err
	}
	user, err := p.Dao.GetUserByUID(cm.Get(UID).(string))
	if err != nil {
		return nil, err
	}
	if !user.IsConfirm() {
		return nil, p.I18n.E(http.StatusForbidden, lng, "auth.errors.user-not-confirm")
	}
	if user.IsLock() {
		return nil, p.I18n.E(http.StatusForbidden, lng, "auth.errors.user-is-lock")
	}
	return user, nil
}

// CurrentUserMiddleware current-user middleware
func (p *Jwt) CurrentUserMiddleware(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	if user, err := p.getUserFromRequest(r); err == nil {
		ctx := context.WithValue(r.Context(), web.K(CurrentUser), user)
		ctx = context.WithValue(ctx, web.K(IsAdmin), p.Dao.Is(user.ID, RoleAdmin))
		n(w, r.WithContext(ctx))
	} else {
		n(w, r)
	}
}

// MustSignInMiddleware must-sign-in middleware
func (p *Jwt) MustSignInMiddleware(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	if _, ok := r.Context().Value(web.K(CurrentUser)).(*User); ok {
		n(w, r)
	} else {
		lng := r.Context().Value(web.K(i18n.LOCALE)).(string)
		http.Error(w, p.I18n.T(lng, "auth.errors.user-must-sign-in"), http.StatusForbidden)
	}
}

// MustAdminMiddleware must-admin middleware
func (p *Jwt) MustAdminMiddleware(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	if is, ok := r.Context().Value(web.K(IsAdmin)).(bool); ok && is {
		n(w, r)
	} else {
		lng := r.Context().Value(web.K(i18n.LOCALE)).(string)
		http.Error(w, p.I18n.T(lng, "auth.errors.user-must-is-admin"), http.StatusForbidden)
	}
}

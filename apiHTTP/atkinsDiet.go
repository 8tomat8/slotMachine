package apiHTTP

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/8tomat8/slotMachine/machines/atkinsDiet"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

//go:generate easyjson -all .

type UserData struct {
	UID   string `json:"uid"`   // user id
	Chips uint   `json:"chips"` // chips balance
	Bet   uint   `json:"bet"`   // bet size
}

func (u UserData) Valid() error {
	if u.UID == "" || u.Chips == 0 || u.Bet == 0 {
		return ErrInvalidUserData
	}
	return nil
}

func (h Handler) Spins(w http.ResponseWriter, r *http.Request) {
	token, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.Log.Debug(errors.Wrap(err, "failed to read request"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO: implement pool
	u := &UserData{}
	t, err := jwt.ParseWithClaims(string(token), u, func(token *jwt.Token) (interface{}, error) {
		// TODO: Check supported sign methods
		if token == nil {
			return nil, ErrInvalidJWT
		}

		return h.JWTSecret, nil
	})
	if err != nil {
		h.Log.Debug(errors.Wrap(err, "failed to parse JWT"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !t.Valid {
		h.Log.Debug("failed to validate JWT")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err = u.Valid(); err != nil {
		h.Log.Debug(errors.Wrap(err, "failed to validate Payload"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u.Chips -= u.Bet
	if u.Chips < 0 {
		h.Log.Debug("not enough chips")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if u.Bet < 20 {
		h.Log.Debug("bet is too small")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rsp := resp{
		// 20 lines is hardcoded due to API limitations
		Spins: atkinsDiet.Run(u.Bet, 20),
	}

	for _, spin := range rsp.Spins {
		rsp.Total += spin.Total
	}

	u.Chips += rsp.Total
	t = jwt.NewWithClaims(jwt.SigningMethodHS512, u)
	rsp.JWT, err = t.SignedString(h.JWTSecret)
	if err != nil {
		h.Log.Warn(errors.Wrap(err, "failed to create new signed token"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(rsp)
	if err != nil {
		h.Log.Info(errors.Wrap(err, "failed to write response"))
	}
}

type resp struct {
	Total uint              `json:"total"`
	Spins []atkinsDiet.Spin `json:"spins"`
	JWT   string            `json:"jwt"`
}

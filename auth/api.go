package auth

import (
	"encoding/base64"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/macaroon.v1"
	"log"
	"net/http"
)

var (
	Key       = []byte("auth-super-secret-key")
	ServiceID = "auth-service-id"
	Location  = "http://localhost:8080/auth"
)

func Register(router *httprouter.Router) {
	router.POST("/auth/tokens", func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		requestData := struct {
			ClientID string `json:"id"`
		}{}

		if err := json.NewDecoder(req.Body).Decode(&requestData); err != nil {
			log.Print(err)
			w.WriteHeader(500)
			w.Write([]byte("Something went wrong when issuing a new token..."))
			return
		}
		log.Println("#### LOLOLOLOLOLOL!")

		// TODO In a real world context the auth service would first verify that the provided ClientID is valid
		// failing the request otherwise...

		m, err := macaroon.New(Key, ServiceID, Location)
		if err != nil {
			log.Print(err)
			w.WriteHeader(500)
			w.Write([]byte("Something went wrong when issuing a new token..."))
			return
		}

		if err := m.AddFirstPartyCaveat("Client:" + requestData.ClientID); err != nil {
			log.Print(err)
			w.WriteHeader(500)
			w.Write([]byte("Something went wrong when issuing a new token..."))
			return
		}

		token, err := m.MarshalBinary()
		if err != nil {
			log.Print(err)
			w.WriteHeader(500)
			w.Write([]byte("Something went wrong when issuing a new token..."))
			return
		}

		j, _ := m.MarshalJSON()

		response := struct {
			Token string `json:"token"`
			LOL   string
		}{
			Token: base64.URLEncoding.EncodeToString(token),
			LOL:   string(j),
		}

		log.Println("#### Token:", response.Token)

		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Print(err)
			w.WriteHeader(500)
			w.Write([]byte("Something went wrong when issuing a new token..."))
			return
		}
	})
}

package storage

import (
	"encoding/base64"
	"encoding/json"
	"github.com/drborges/macaroons-spike/auth"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/macaroon.v1"
	"log"
	"net/http"
)

var (
	Key      = []byte("storage-super-secret-key")
	Location = "http://localhost:8080/storage/files"
)

func noCaveatsChecker(caveat string) error {
	return nil
}

func Register(router *httprouter.Router) {
	router.GET("/storage/files", func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		encodedToken := req.Header.Get("X-Auth-Token")
		//		authMacaroon, err := macaroon.New(auth.Key, auth.ServiceID, auth.Location)
		//		if err != nil {
		//			log.Print(err)
		//			w.WriteHeader(500)
		//			w.Write([]byte("Oops, something went wrong..."))
		//			return
		//		}

		token, err := base64.URLEncoding.DecodeString(encodedToken)
		if err != nil {
			log.Print(err)
			w.WriteHeader(500)
			w.Write([]byte("Oops, something went wrong..."))
			return
		}

		userMacaroon := &macaroon.Macaroon{}
		if err := userMacaroon.UnmarshalBinary(token); err != nil {
			log.Print(err)
			w.WriteHeader(500)
			w.Write([]byte("Oops, something went wrong..."))
			return
		}

		log.Printf("#### Macaroon caveats: %+v\n", userMacaroon.Caveats())
		log.Printf("#### Macaroon signature: %+v\n", userMacaroon.Signature())
		log.Printf("#### Macaroon id: %+v\n", userMacaroon.Id())
		log.Printf("#### Macaroon location: %+v\n", userMacaroon.Location())

		if err := userMacaroon.Verify(auth.Key, noCaveatsChecker, nil); err != nil {
			log.Print(err)
			w.WriteHeader(401)
			w.Write([]byte(err.Error()))
			return
		}

		response := struct {
			Files []string `json:"files"`
		}{
			Files: []string{
				"http://localhost:6061/storage/files/1",
				"http://localhost:6061/storage/files/2",
				"http://localhost:6061/storage/files/3",
			},
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Print(err)
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
		}
	})
}

package main

import (
	"gopkg.in/macaroon.v1"
	"log"
	"time"
	"strconv"
)

func checkCaveats(caveat string) error {
	log.Println("Checking:", caveat)
	return nil
}

func main() {
	// In this example the third party key is shared with the storage service though it is not required to be a shared key.
	// One could use pub/priv keys to generate the same secrete in both parties for instance
	thirdPartyKey := []byte("auth key")
	thirdPartyCaveatId := "auth id"
	thirdPartyLocation := "auth location"

	storageKey := []byte("storage secret key")
	storageId := "storage id"
	storageLocation := "storage location"

	now := strconv.Itoa(time.Now().Add(time.Hour).Nanosecond())

	storageMacaroon, _ := macaroon.New(storageKey, storageId, storageLocation)
	storageMacaroon.AddFirstPartyCaveat("file:cat.png")
	storageMacaroon.AddFirstPartyCaveat("time < " + now)
	storageMacaroon.AddThirdPartyCaveat(thirdPartyKey, thirdPartyCaveatId, thirdPartyLocation)

	// Discharge macaroons are created by third party services upon a client's request, for example an authentication request
	//
	// It is a good practice to add an expiration time to the discharge macaroon so it may be valid for a finite amount of time
	//
	// In this context, in order to access the storage service, a user in addition to having to comply with all caveats she must provide
	// a proof that she is logged into the auth service, by sending along with the storage token the discharge token retrieved from
	// the auth service upon authentication.
	//
	// Discharge macaroons must be bound to the original/requesting macaroon before being sent back to clients. This ensures the given
	// discharge is only valid for that particular macaroon furthermore the discharge macaroon cannot be further updated.
	dischargeAuthMacaroon, _ := macaroon.New(thirdPartyKey, thirdPartyCaveatId, thirdPartyLocation)
	dischargeAuthMacaroon.AddFirstPartyCaveat("time < " + now)
	dischargeAuthMacaroon.Bind(storageMacaroon.Signature())

	if err := storageMacaroon.Verify(storageKey, checkCaveats, []*macaroon.Macaroon{dischargeAuthMacaroon}); err != nil {
		log.Fatal(err)
	}
}

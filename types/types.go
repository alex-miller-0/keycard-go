package types

import "github.com/alex-miller-0/keycard-go/apdu"
import "github.com/ethereum/go-ethereum/log"

var logger = log.New("package", "types")
// Channel is an interface with a Send method to send apdu commands and receive apdu responses.
type Channel interface {
	Send(*apdu.Command) (*apdu.Response, error)
}

type PairingInfo struct {
	Key   []byte
	Index int
}

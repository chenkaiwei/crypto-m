package client_demo

import (
	"encoding/hex"
	"math/rand"
	"time"
)

const PUB_KEY = "MIGJAoGBAL0btTrul8Md8B9S8yRna4E/IRu+piuWvnESvbO9ygrFd91+Hj6f//JwPbBMTgSTDAAmJUCppPezUqKWGQx/hGcHmwTUWPQCruf7y1iyZN0r3kEgR5Ia47apkieLlLsRoBlVJCK/wcZt52W14C/YlrVdFmycih7QT2pAE4ICGkHpAgMBAAE="

func newRandomCek(uLen uint) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, uLen)

	rand.Read(b)

	rand_str := hex.EncodeToString(b)[0:uLen]

	return rand_str
}

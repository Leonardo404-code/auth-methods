package cookies

import (
	"encoding/gob"
	"encoding/hex"
	"log"
)

type User struct {
	Name string
	Age  int
}

var secretKey []byte

func init() {
	hex, err := hex.DecodeString("13d6b4dff8f84a10851021ec8608f814570d562c92fe6b5ec4c9f595bcb3234b")
	if err != nil {
		log.Fatal(err)
	}

	secretKey = hex

	gob.Register(&User{})
}

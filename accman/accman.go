// hamlet/accman

package accman

import (
    "golang.org/x/crypto/bcrypt"
    "fmt"
)

type Account struct {
    ID int
    Name string
    PW_hash []byte
}
func NewPW (new_pw []byte) []byte {
    pw_hash,err := bcrypt.GenerateFromPassword(new_pw, 14)
    if err != nil {
        fmt.Println(err)
    }
    return pw_hash
}

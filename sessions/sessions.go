// sessions/sessions.go

package sessions

import (
    "hamlet/accman"
    "time"
    "crypto/md5"
    "log"
    "bufio"
    "encoding/json"
)

type Session struct {
    ID [16]byte
    Conn_rw *bufio.ReadWriter
    Account *accman.Account
}

func New(rw *bufio.ReadWriter) Session {
    bin_time,err := time.Now().MarshalBinary()
    if err != nil {
        log.Println("Time Error: ",err)
    }
    id := md5.Sum(bin_time)
    return Session{
        ID: id,
        Conn_rw: rw,
    }
}

func (session *Session) Login() {
}

func (session *Session) UpdateClient() {
    for {
        message, err := json.Marshal(session)
        if err != nil {
            log.Println("ERROR JSON marshaling error: ",err)
        }
        session.Conn_rw.Write(message)
        session.Conn_rw.WriteByte('\n')
        session.Conn_rw.Flush()
        time.Sleep(500 * time.Millisecond)
    }
}

// hamlet.go


package main

import (
    "hamlet/accman"
    "hamlet/sessions"
    "encoding/json"
    "io/ioutil"
    "fmt"
    "bufio"
    "strings"
    "os"
    "net"
    "log"
)

func newAcc(id int) accman.Account {
    var name,pw string
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("enter name:")
    name,_ = reader.ReadString('\n')
    name = strings.Replace(name, "\n", "", -1)
    fmt.Println("enter password:")
    pw,_ = reader.ReadString('\n')
    pw = strings.Replace(pw, "\n", "", -1)
    return accman.Account{
        ID: id,
        Name: name,
        PW_hash: accman.NewPW([]byte(pw)),
    }
}

var G_accounts []accman.Account
var G_sessions []*sessions.Session
var logger *log.Logger

func handle_connection(conn net.Conn) {
    conn_rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
    my_session := sessions.New(conn_rw)
    logger.Println("new session: ", my_session)
    log.Println("new session: ", my_session)
    G_sessions = append(G_sessions, &my_session)
    go my_session.UpdateClient()
    for {
        if my_session.Account == nil {
            message,_ := json.Marshal([...]string{"login", "new", "quit"})
            nn,err := my_session.Conn_rw.Write(message)
            if err != nil {
                logger.Println("ERROR writing message: ",err)
                log.Println("ERROR writing message: ",err)
                break
            }

            err = my_session.Conn_rw.WriteByte('\n')
            if err != nil {
                logger.Println("ERROR writing end byte")
                log.Println("ERROR writing end byte")
                break
            }

            my_session.Conn_rw.Flush()
            if err != nil {
                logger.Println("ERROR flushing buffer: ",err)
                log.Println("ERROR flushing buffer: ",err)
                break
            }

            log.Println("sent ", nn, " bytes: ", string(message))
        }
        message,err := my_session.Conn_rw.ReadBytes('\n')
        if err != nil {
            logger.Println("ERROR reading message: ",err)
            log.Println("ERROR reading message: ",err)
            break
        }
        logger.Println("message: ",message)
        log.Println("message: ",message)
    }
}

func main() {
    //init

    // initializing log
    error_log,err := os.OpenFile("error.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755 )
    if err != nil {
        log.Println("ERROR opening error.log: ",err)
    }
    logger = log.New(error_log, "hamlet ", log.Lshortfile)

    //initializing accounts
    acc_data,err := ioutil.ReadFile("accounts.json")
    if err != nil {
        logger.Println("ERROR opening accounts.json: ",err)
        log.Println("ERROR opening accounts.json: ",err)
    }

    err = json.Unmarshal(acc_data, &G_accounts)
    if err != nil {
        logger.Println("ERROR parsing sessions.json: ",err)
        log.Println("ERROR parsing sessions.json: ",err)
    }

    //initializing net
    listener,err := net.Listen("tcp", ":6666")
    if err != nil {
        logger.Println("ERROR opening tcp socket on port 6666: ",err)
        log.Println("ERROR opening tcp socket on port 6666: ",err)
    }
    for {
        conn,err := listener.Accept()
        if err != nil {
            logger.Println("ERROR accepting connection: ",err)
            log.Println("ERROR accepting connection: ",err)
        }
        logger.Println("new connection: ",conn)
        log.Println("new connection: ",conn)
        go handle_connection(conn)
    }




    acc_data,err = json.Marshal(G_accounts)
    if err != nil {
        logger.Println("ERROR parsing acc file to json: ", err)
        log.Println("ERROR parsing acc file to json: ", err)
    }
    log.Println("accd data", acc_data)
    ioutil.WriteFile("sessions.json", acc_data, 0644)
}

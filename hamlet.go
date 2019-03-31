// hamlet.go

package main

import (
    "strings"
    "fmt"
    "net"
    "bytes"
    "hamlet/dudes"
    "strconv"
)

// Variables
func PlayerCreation(connection net.Conn) dudes.Character {
    var PlayerName string
    var PlayerHP int
    var PlayerAP int
    connection.Write([]byte("Character Creation\n"))
    connection.Write([]byte("Enter Name: "))
    PlayerName = readCommand(connection)
    connection.Write([]byte("\n"))

    connection.Write([]byte("Enter HP: "))
    PlayerHP,_ = strconv.Atoi(readCommand(connection))
    connection.Write([]byte("\n"))

    connection.Write([]byte("Enter AP: "))
    PlayerAP,_ = strconv.Atoi(readCommand(connection))
    connection.Write([]byte("\n"))
    fmt.Println(PlayerHP)
    connection.Write([]byte(strconv.Itoa(PlayerHP)))
    statistics := dudes.BaseStats{HP: PlayerHP, AP: PlayerAP}
    connection.Write([]byte(strconv.Itoa(statistics.HP)))

    return dudes.NewCharacter(PlayerName, dudes.BaseStats{HP: PlayerHP, AP: PlayerAP})
}

func readCommand(connection net.Conn) string {
    for {
        var message []byte
        buff := make([]byte, 1)
        for {
            connection.Read(buff)
            fmt.Print("buff: ")
            fmt.Println(buff)
            message = append(message, buff...)
            if l := len(message); l >= 2 && bytes.Equal(message[(l-2):(l)], []byte("\r\n")) {
                return strings.TrimSpace(string(message))
            }
        }
    }
}

func serve(connection net.Conn) {
    player := PlayerCreation(connection)
    for {
        toad := dudes.NewCharacter("toad", dudes.BaseStats{
            HP: 10,
            AP: 1,})
        connection.Write([]byte(fmt.Sprintf("A wild %s attacks\n", toad.Name)))
        connection.Write([]byte("fight!\n"))
        battle := dudes.NewBattle([]*dudes.Character{&toad}, []*dudes.Character{&player})
        battle.Do()
        connection.Write([]byte(fmt.Sprintf("%s's hitpoints: %d\n", toad.Name, toad.Stats.HP)))
        connection.Write([]byte(fmt.Sprintf("%s's hitpoints: %d\n", player.Name, player.Stats.HP)))
        connection.Write([]byte("fight!\n"))
        for {
            connection.Write([]byte("command: "))
            readCommand(connection)
            connection.Write([]byte(fmt.Sprintf("\n %s attacks %s\n", player.Name, toad.Name)))
            (&player).Attacks(&toad)
            connection.Write([]byte(fmt.Sprintf("%s's hitpoints: %d\n", toad.Name, toad.Stats.HP)))
            connection.Write([]byte(fmt.Sprintf("%s's hitpoints: %d\n", player.Name, player.Stats.HP)))
            connection.Write([]byte(fmt.Sprintf("\n %s attacks %s\n", toad.Name, player.Name)))
            (&toad).Attacks(&player)
            connection.Write([]byte(fmt.Sprintf("%s's hitpoints: %d\n", toad.Name, toad.Stats.HP)))
            connection.Write([]byte(fmt.Sprintf("%s's hitpoints: %d\n", player.Name, player.Stats.HP)))
            if (toad.Stats.HP <= 0) {
            connection.Write([]byte(fmt.Sprintf("%s died\n", toad.Name)))
            break
            }
        }
    }
}

func main() {
    telnet_listen, _ := net.Listen("tcp", ":6666")
    for {
        connection, _ := telnet_listen.Accept()
        go serve(connection)
    }
}

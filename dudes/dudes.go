// hamlet/dudes

package dudes

import (
    "time"
    "fmt"
)

type BaseStats struct {
    HP int
    AP int
}

// Object Character
type Character struct {
    Name string
    Stats BaseStats
    Action string
}
func (attacker *Character) Attacks(target *Character) {
    target.Stats.HP -= attacker.Stats.AP
    if (target.Stats.HP <= 0) {
        attacker.Stats.HP += 1
    }
}
func NewCharacter(name string, stats BaseStats) Character {
    return Character {Name: name, Stats: stats}
}

// Object Battle
type Battle struct {
    team1,team2 []*Character
    Log []string
}
func NewBattle(Team1, Team2 []*Character) Battle {
    return Battle{team1: Team1, team2: Team2}
}
func (battle Battle) Do(){
    for _,fighter := range(battle.team1) {
        go fighter.fightBattle(battle.team1, battle.team2, battle.Log)
    }
    for _,fighter := range(battle.team2) {
        go fighter.fightBattle(battle.team2, battle.team1, battle.Log)
    }
    for index,fighter := range(battle.team1) {
        if (fighter.Stats.HP > 0) {
            break
        } else if (index == len(battle.team1)) {
            return
        }
    }
    for index,fighter := range(battle.team2) {
        if (fighter.Stats.HP > 0) {
            break
        } else if (index == len(battle.team2)) {
            return
        }
    }
}


func (fighter *Character) fightBattle(friends, foes []*Character, log []string){
    if (fighter.Stats.HP <= 0) {
        return
    }
    for {
        if (fighter.Action != "") {
            fighter.Attacks(foes[1])
            log = append(log, fmt.Sprintf("%s attacks %s\n", fighter.Name, foes[1].Name))
            time.Sleep(10 * time.Second)
            break
        }
    }
}

func Gattle(team1, team2 []*Character) {
    for _,fighter:= range(team1) {
        fighter.Stats.HP -= 1
    }
    for _,fighter:= range(team2) {
        fighter.Stats.HP -= 1
    }
}


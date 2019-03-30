// hamlet/dudes.go

package dudes

import (
)

type BaseStats struct {
    HP int
    AP int
}

// Object Character
type Character struct {
    Name string
    Stats BaseStats
}
func (attacker *Character) Attacks(target *Character) {
    target.Stats.HP -= attacker.Stats.AP
    if (target.Stats.HP <= 0) {
        attacker.Stats.HP += 1
    }
}
func NewCharacter(name string, stats BaseStats) Character {
    return Character {name, stats}
}


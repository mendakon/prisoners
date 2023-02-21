package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strconv"
)

const ROUND_NUM = 1
const PRISONER_ACTION_NUM = 3
const PRISONER_NUM = 8

const SELLOUT_POINT = 0
const SELLOUTED_POINT = 10
const ACCOMPLICE_POINT = 5
const REMAIN_SILENT_POINT = 2

type Prisoner struct {
	actions       [PRISONER_ACTION_NUM]bool
	actionsString string
	point         int
}

func main() {
	var prisoners [PRISONER_NUM]Prisoner
	prisonersString := makeAllpattern(PRISONER_ACTION_NUM)

	for i := 0; i < PRISONER_NUM; i++ {
		actionsTmp := binStrToBoolAry(prisonersString[i])
		copy(prisoners[i].actions[:], actionsTmp)
		prisoners[i].actionsString = prisonersString[i]
	}
	for round := 0; round < ROUND_NUM; round++ {
		for i := 0; i < PRISONER_NUM; i++ {
			for j := i + 1; j < PRISONER_NUM; j++ {
				prisoners[i], prisoners[j] = matchPrisoners(prisoners[i], prisoners[j])
			}
		}
	}

	sort.Slice(prisoners[:], func(i int, j int) bool { return prisoners[i].point < prisoners[j].point })

	for i := 0; i < PRISONER_NUM; i++ {
		fmt.Printf("第%d位: ", i+1)
		prisoners[i].print()
	}

}

func (p Prisoner) print() {
	fmt.Printf("懲役%d年. actions %s\n", p.point, p.actionsString)
}

func makeAllpattern(actionNum int) []string {
	pNum := int(math.Pow(2, float64(actionNum)))

	b := make([]string, pNum)

	for j := 0; j < pNum; j++ {
		bit2 := fmt.Sprintf("%0"+strconv.Itoa(actionNum)+"b", j)
		b[j] = bit2
	}

	return b
}

func binStrToBoolAry(binStr string) []bool {
	boolAry := []bool{}
	for _, bin := range binStr {
		if bin == '0' {
			boolAry = append(boolAry, false)
		} else if bin == '1' {
			boolAry = append(boolAry, true)
		}
	}
	return boolAry
}

func matchPrisoners(p1 Prisoner, p2 Prisoner) (Prisoner, Prisoner) {
	trials := rand.Intn(PRISONER_ACTION_NUM)
	for i := 0; i < trials; i++ {
		p1Point, p2Point := oneBattle(p1.actions[i], p2.actions[i])
		p1.point += p1Point
		p2.point += p2Point
	}
	return p1, p2
}

func oneBattle(p1Action bool, p2Action bool) (int, int) {
	p1point := 0
	p2Point := 0

	switch {
	//両方裏切る
	case p1Action && p2Action:
		p1point = ACCOMPLICE_POINT
		p2Point = ACCOMPLICE_POINT
	//1だけ裏切る
	case p1Action && !p2Action:
		p1point = SELLOUT_POINT
		p2Point = SELLOUTED_POINT
	//2だけ裏切る
	case !p1Action && p2Action:
		p1point = SELLOUTED_POINT
		p2Point = SELLOUT_POINT
	//どちらも裏切らない
	case !p1Action && !p2Action:
		p1point = REMAIN_SILENT_POINT
		p2Point = REMAIN_SILENT_POINT
	}

	return p1point, p2Point
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bot struct {
	values  []int
	lowBot  bool
	idLow   int
	highBot bool
	idHigh  int
}

func main() {
	fmt.Println("Day 10 of Advent of Code 2016")

	data := readLines(os.Stdin)
	bots := make(map[int]*bot)

	var output = make(map[int]*[]int)
	for _, instr := range data {
		fields := strings.Fields(instr)
		var botID int
		if fields[0] == "value" {
			botID, _ = strconv.Atoi(fields[5])
			refBot := bots[botID]
			if refBot == nil {
				bots[botID] = new(bot)
				refBot = bots[botID]
			}
			val, _ := strconv.Atoi(fields[1])
			refBot.values = append(refBot.values, val)
		} else {
			botID, _ = strconv.Atoi(fields[1])
			refBot := bots[botID]
			if refBot == nil {
				bots[botID] = new(bot)
				refBot = bots[botID]
			}
			idLow, _ := strconv.Atoi(fields[6])
			botLow := fields[5] == "bot"
			idHigh, _ := strconv.Atoi(fields[11])
			botHigh := fields[10] == "bot"

			refBot.idLow = idLow
			refBot.lowBot = botLow
			refBot.idHigh = idHigh
			refBot.highBot = botHigh
		}
		//fmt.Printf("Instr: %s\nBot %d, vals %v, %v low -> %d; %v high -> %d\n", instr, bot_id, bots[bot_id].values, bots[bot_id].low_bot, bots[bot_id].id_low, bots[bot_id].high_bot, bots[bot_id].id_high)
	}

	for len(bots) > 0 {
		executed := []int{}
		for k, b := range bots {
			if len(b.values) == 2 {
				executeBot(bots, output, k)
				executed = append(executed, k)
			}
		}
		for _, k := range executed {
			delete(bots, k)
		}
	}

	/*
		for k, v := range output {
			fmt.Printf("K: %d\tV: %v\n", k, v)
		}
	*/

	result := (*output[0])[0] * (*output[1])[0] * (*output[2])[0]
	fmt.Printf("Part 2: %d\n", result)
}

func executeBot(bots map[int]*bot, out map[int]*[]int, botID int) {
	refBot := bots[botID]
	//fmt.Printf("Bot %d, vals %v, %v low -> %d; %v high -> %d\n", bot_id, bots[bot_id].values, bots[bot_id].low_bot, bots[bot_id].id_low, bots[bot_id].high_bot, bots[bot_id].id_high)
	if len(refBot.values) == 2 {
		if (refBot.values[0] == 61 && refBot.values[1] == 17) || (refBot.values[0] == 17 && refBot.values[1] == 61) {
			fmt.Printf("Part 1: id - %d vs %v\n", botID, refBot.values)
		}
		var lowDestArr *[]int
		var highDestArr *[]int

		if refBot.lowBot {
			if bots[refBot.idLow] == nil {
				bots[refBot.idLow] = new(bot)
			}
			lowDestArr = &(bots[refBot.idLow].values)
		} else {
			lowDestArr = out[refBot.idLow]
			if lowDestArr == nil {
				out[refBot.idLow] = new([]int)
				lowDestArr = out[refBot.idLow]
			}
		}

		if refBot.highBot {
			if bots[refBot.idHigh] == nil {
				bots[refBot.idHigh] = new(bot)
			}
			highDestArr = &(bots[refBot.idHigh].values)
		} else {
			highDestArr = out[refBot.idHigh]
			if highDestArr == nil {
				out[refBot.idHigh] = new([]int)
				highDestArr = out[refBot.idHigh]
			}
		}

		//fmt.Println(bot_id)
		*lowDestArr = append(*lowDestArr, min(refBot.values))
		*highDestArr = append(*highDestArr, max(refBot.values))
		refBot.values = []int{}
	}
}

func min(v []int) int {
	if v[0] < v[1] {
		return v[0]
	}
	return v[1]

}

func max(v []int) int {
	if v[0] > v[1] {
		return v[0]
	}
	return v[1]

}

func readLines(f *os.File) []string {
	var data []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}

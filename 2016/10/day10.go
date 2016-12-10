package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bot struct {
	values   []int
	low_bot  bool
	id_low   int
	high_bot bool
	id_high  int
}

func main() {
	fmt.Println("Day 10 of Advent of Code 2016")

	data := read_lines(os.Stdin)
	bots := make(map[int]*Bot)

	var output = make(map[int]*[]int)
	for _, instr := range data {
		fields := strings.Fields(instr)
		var bot_id int
		if fields[0] == "value" {
			bot_id, _ = strconv.Atoi(fields[5])
			bot := bots[bot_id]
			if bot == nil {
				bots[bot_id] = new(Bot)
				bot = bots[bot_id]
			}
			val, _ := strconv.Atoi(fields[1])
			bot.values = append(bot.values, val)
		} else {
			bot_id, _ = strconv.Atoi(fields[1])
			bot := bots[bot_id]
			if bot == nil {
				bots[bot_id] = new(Bot)
				bot = bots[bot_id]
			}
			id_low, _ := strconv.Atoi(fields[6])
			bot_low := fields[5] == "bot"
			id_high, _ := strconv.Atoi(fields[11])
			bot_high := fields[10] == "bot"

			bot.id_low = id_low
			bot.low_bot = bot_low
			bot.id_high = id_high
			bot.high_bot = bot_high
		}
		//fmt.Printf("Instr: %s\nBot %d, vals %v, %v low -> %d; %v high -> %d\n", instr, bot_id, bots[bot_id].values, bots[bot_id].low_bot, bots[bot_id].id_low, bots[bot_id].high_bot, bots[bot_id].id_high)
	}

	for len(bots) > 0 {
		executed := []int{}
		for k, b := range bots {
			if len(b.values) == 2 {
				execute_bot(bots, output, k)
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

func execute_bot(bots map[int]*Bot, out map[int]*[]int, bot_id int) {
	bot := bots[bot_id]
	//fmt.Printf("Bot %d, vals %v, %v low -> %d; %v high -> %d\n", bot_id, bots[bot_id].values, bots[bot_id].low_bot, bots[bot_id].id_low, bots[bot_id].high_bot, bots[bot_id].id_high)
	if len(bot.values) == 2 {
		if (bot.values[0] == 61 && bot.values[1] == 17) || (bot.values[0] == 17 && bot.values[1] == 61) {
			fmt.Printf("Part 1: id - %d vs %v\n", bot_id, bot.values)
		}
		var low_dest_arr *[]int
		var high_dest_arr *[]int

		if bot.low_bot {
			if bots[bot.id_low] == nil {
				bots[bot.id_low] = new(Bot)
			}
			low_dest_arr = &(bots[bot.id_low].values)
		} else {
			low_dest_arr = out[bot.id_low]
			if low_dest_arr == nil {
				out[bot.id_low] = new([]int)
				low_dest_arr = out[bot.id_low]
			}
		}

		if bot.high_bot {
			if bots[bot.id_high] == nil {
				bots[bot.id_high] = new(Bot)
			}
			high_dest_arr = &(bots[bot.id_high].values)
		} else {
			high_dest_arr = out[bot.id_high]
			if high_dest_arr == nil {
				out[bot.id_high] = new([]int)
				high_dest_arr = out[bot.id_high]
			}
		}

		//fmt.Println(bot_id)
		*low_dest_arr = append(*low_dest_arr, min(bot.values))
		*high_dest_arr = append(*high_dest_arr, max(bot.values))
		bot.values = []int{}
	}
}

func min(v []int) int {
	if v[0] < v[1] {
		return v[0]
	} else {
		return v[1]
	}
}

func max(v []int) int {
	if v[0] > v[1] {
		return v[0]
	} else {
		return v[1]
	}
}

func read_lines(f *os.File) []string {
	var data []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}

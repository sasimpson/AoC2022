/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// day02Cmd represents the day02 command
var day02Cmd = &cobra.Command{
	Use:   "day02",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: day02,
}

func init() {
	rootCmd.AddCommand(day02Cmd)
}

var day02 = func(cmd *cobra.Command, args []string) error {
	file, err := os.Open("./data/day02-a.txt")
	if err != nil {
		return err
	}

	s := bufio.NewScanner(file)
	var resultASum, resultBSum int
	for s.Scan() {
		line := s.Text()

		match := strings.Split(line, " ")
		resultA := matchCompare(match[0], match[1])
		resultB := matchCompute(match[0], match[1])
		resultASum += resultA
		resultBSum += resultB
	}
	if err := s.Err(); err != nil {
		return err
	}

	fmt.Println("part A resultSum:", resultASum)
	fmt.Println("part B resultSum:", resultBSum)

	return nil
}

// A/X Rock(1) B/Y Paper(2) C/Z Scissors(3)
func matchCompare(a, b string) int {

	grid := map[string]map[string]int{
		"A": {"X": 4, "Y": 8, "Z": 3},
		"B": {"X": 1, "Y": 5, "Z": 9},
		"C": {"X": 7, "Y": 2, "Z": 6},
	}

	return grid[a][b]
}

// A Rock(1) B Paper(2) C Scissors(3)
// X Lose(0) Y Draw(3)  Z Win(6)
func matchCompute(a, b string) int {
	grid := map[string]map[string]int{
		"A": {"X": 3, "Y": 4, "Z": 8},
		"B": {"X": 1, "Y": 5, "Z": 9},
		"C": {"X": 2, "Y": 6, "Z": 7},
	}
	return grid[a][b]
}

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"sort"
	"strconv"
)

// day01Cmd represents the day01 command
var day01Cmd = &cobra.Command{
	Use:   "day01",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: day01,
}

func init() {
	rootCmd.AddCommand(day01Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day01Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day01Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var day01 = func(cmd *cobra.Command, args []string) error {
	fmt.Println("day 1 part 1")

	var foodList []int

	file, err := os.Open("./data/day01-a.txt")
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	s := bufio.NewScanner(file)
	var foodSum, maxFoodSum int
	for s.Scan() {
		line := s.Text()
		if line == "" {
			foodList = append(foodList, foodSum)
			if foodSum > maxFoodSum {
				maxFoodSum = foodSum
			}
			foodSum = 0
			continue
		}

		foodCount, err := strconv.Atoi(line)
		if err != nil {
			return err
		}

		foodSum += foodCount
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("max calories:", maxFoodSum)

	sort.Sort(sort.Reverse(sort.IntSlice(foodList)))
	var maxSum int
	for _, i := range foodList[:3] {
		maxSum += i
	}
	fmt.Println("top three max:", maxSum)

	return nil
}

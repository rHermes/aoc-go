package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/rhermes/aoc-go/years/y2020"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "aoc-go",
		Usage: "Advent of code in go",
		Commands: []*cli.Command{
			RunCommand,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

type TaskFunc func([]byte) (string, error)

type Task struct {
	Year int
	Day  int
	Part int
}

var ydp = map[Task]TaskFunc{
	Task{2020, 1, 1}:  y2020.Day01Part01,
	Task{2020, 1, 2}:  y2020.Day01Part02,
	Task{2020, 2, 1}:  y2020.Day02Part01,
	Task{2020, 2, 2}:  y2020.Day02Part02,
	Task{2020, 3, 1}:  y2020.Day03Part01,
	Task{2020, 3, 2}:  y2020.Day03Part02,
	Task{2020, 4, 1}:  y2020.Day04Part01,
	Task{2020, 4, 2}:  y2020.Day04Part02,
	Task{2020, 5, 1}:  y2020.Day05Part01,
	Task{2020, 5, 2}:  y2020.Day05Part02,
	Task{2020, 6, 1}:  y2020.Day06Part01,
	Task{2020, 6, 2}:  y2020.Day06Part02,
	Task{2020, 7, 1}:  y2020.Day07Part01,
	Task{2020, 7, 2}:  y2020.Day07Part02,
	Task{2020, 8, 1}:  y2020.Day08Part01,
	Task{2020, 8, 2}:  y2020.Day08Part02,
	Task{2020, 9, 1}:  y2020.Day09Part01,
	Task{2020, 9, 2}:  y2020.Day09Part02,
	Task{2020, 10, 1}: y2020.Day10Part01,
	Task{2020, 10, 2}: y2020.Day10Part02,
	Task{2020, 11, 1}: y2020.Day11Part01,
	Task{2020, 11, 2}: y2020.Day11Part02,
	Task{2020, 12, 1}: y2020.Day12Part01,
	Task{2020, 12, 2}: y2020.Day12Part02,
	Task{2020, 13, 1}: y2020.Day13Part01,
	Task{2020, 13, 2}: y2020.Day13Part02,
	Task{2020, 14, 1}: y2020.Day14Part01,
	Task{2020, 14, 2}: y2020.Day14Part02,
	Task{2020, 15, 1}: y2020.Day15Part01,
	Task{2020, 15, 2}: y2020.Day15Part02,
	Task{2020, 16, 1}: y2020.Day16Part01,
	Task{2020, 16, 2}: y2020.Day16Part02,
	Task{2020, 17, 1}: y2020.Day17Part01,
	Task{2020, 17, 2}: y2020.Day17Part02,
	Task{2020, 18, 1}: y2020.Day18Part01,
	Task{2020, 18, 2}: y2020.Day18Part02,
	Task{2020, 19, 1}: y2020.Day19Part01,
	Task{2020, 19, 2}: y2020.Day19Part02,
	Task{2020, 20, 1}: y2020.Day20Part01,
	Task{2020, 20, 2}: y2020.Day20Part02,
	Task{2020, 21, 1}: y2020.Day21Part01,
	Task{2020, 21, 2}: y2020.Day21Part02,
	Task{2020, 22, 1}: y2020.Day22Part01,
	Task{2020, 22, 2}: y2020.Day22Part02,
	Task{2020, 23, 1}: y2020.Day23Part01,
	Task{2020, 23, 2}: y2020.Day23Part02,
	Task{2020, 24, 1}: y2020.Day24Part01,
	Task{2020, 24, 2}: y2020.Day24Part02,
	Task{2020, 25, 1}: y2020.Day25Part01,
	Task{2020, 25, 2}: y2020.Day25Part02,
}

var RunCommand = &cli.Command{
	Name:  "run",
	Usage: "run a task",
	Flags: []cli.Flag{
		&cli.PathFlag{
			Name:  "input",
			Usage: "file to use as input for test",
		},
		&cli.IntFlag{
			Name:     "year",
			Required: true,
		},
		&cli.IntFlag{
			Name:     "day",
			Required: true,
		},
		&cli.IntSliceFlag{
			Name: "parts",
		},
	},
	Action: func(c *cli.Context) error {
		year := c.Int("year")
		day := c.Int("day")

		fd, err := os.Open(c.Path("input"))
		if err != nil {
			return err
		}
		defer fd.Close()

		var buf bytes.Buffer

		if _, err := buf.ReadFrom(fd); err != nil {
			return err
		}

		for _, part := range c.IntSlice("parts") {
			f, ok := ydp[Task{Year: year, Day: day, Part: part}]
			if !ok {
				return errors.New("Part not found!")
			}

			ans, err := f(buf.Bytes())
			if err != nil {
				return err
			}

			fmt.Printf("Answer for %d day %d part %d: %s\n", year, day, part, ans)
		}

		return nil
	},
}

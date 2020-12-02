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
	Task{2020, 1, 1}: y2020.Day01Part01,
	Task{2020, 1, 2}: y2020.Day01Part02,
	Task{2020, 2, 1}: y2020.Day02Part01,
	Task{2020, 2, 2}: y2020.Day02Part02,
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

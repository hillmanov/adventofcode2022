package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"plugin"
	"strconv"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

type RunResult struct {
	Day       int
	Part      int
	Solution  any
	StartTime time.Time
	EndTime   time.Time
}

// get the total time taken to run the solution
func (r RunResult) elapsedTime() float64 {
	return float64(r.EndTime.Sub(r.StartTime).Microseconds()) / 1000
}

func (r RunResult) tableData() []string {
	return []string{
		strconv.Itoa(r.Day),
		strconv.Itoa(r.Part),
		fmt.Sprintf("%+v", r.Solution),
		fmt.Sprintf("%.2f", r.elapsedTime()),
	}
}

func main() {
	log.SetFlags(0)

	command := flag.String("command", "runCurrent", "Command to run")
	flag.Parse()

	switch *command {
	case "new":
		new()
	case "runCurrent":
		runCurrent()
	case "runAll":
		runAll()
	case "buildCurrent":
		buildCurrent()
	case "buildAll":
		buildAll()
	}
}

func new() {
	currentDay := getCurrentDay()
	if len(currentDay) == 0 {
		currentDay = "day0"
	}
	currentDayNum := getDayNum(currentDay)
	nextDayNum := currentDayNum + 1
	nextDay := fmt.Sprintf("%02d", nextDayNum)

	cmd := exec.Command("make", "day="+nextDay)
	if _, err := cmd.Output(); err != nil {
		log.Fatal("Error creating new day")
	}
	log.Printf("Created new day: day%s", nextDay)
}

func buildCurrent() error {
	currentDay := getCurrentDay()
	log.Printf("Building %s...", currentDay)

	return buildDay(currentDay)
}

func runCurrent() {
	if err := buildCurrent(); err != nil {
		log.Fatal(err)
	}

	currentDay := getCurrentDay()
	results, _ := runDay(currentDay)
	renderResults(results)
}

func buildAll() error {
	days := getDays()
	for _, day := range days {
		if err := buildDay(day); err != nil {
			log.Fatalf("Error building %s: %s", day, err)
			return err
		}
	}
	return nil
}

func runAll() {
	if err := buildAll(); err != nil {
		log.Fatal(err)
	}

	days := getDays()
	runResults := []RunResult{}

	for _, day := range days {
		results, _ := runDay(day)
		runResults = append(runResults, results...)
	}

	renderResults(runResults)
}

func buildDay(day string) error {
	cmd := exec.Command("go", "build", "-buildmode", "plugin")
	cmd.Dir = "./" + day
	if _, err := cmd.Output(); err != nil {
		log.Printf("Error building %s: %s", day, err)
		return err
	}
	return nil
}

func runDay(day string) ([]RunResult, error) {
	dayPluginPath := filepath.Join("./", day, day+".so")

	dayPlugin, _ := plugin.Open(dayPluginPath)

	part1Symbol, _ := dayPlugin.Lookup("Part1")
	part2Symbol, _ := dayPlugin.Lookup("Part2")

	part1 := part1Symbol.(func() any)
	part2 := part2Symbol.(func() any)

	log.Printf("Running %s...", day)

	var part1Result RunResult
	part1Result = RunResult{
		Day:       getDayNum(day),
		Part:      1,
		StartTime: time.Now(),
	}
	part1Result.Solution = part1()
	part1Result.EndTime = time.Now()

	var part2Result RunResult
	part2Result = RunResult{
		Day:       getDayNum(day),
		Part:      2,
		StartTime: time.Now(),
	}
	part2Result.Solution = part2()
	part2Result.EndTime = time.Now()

	return []RunResult{part1Result, part2Result}, nil
}

func getCurrentDay() string {
	days := getDays()
	if len(days) == 0 {
		return ""
	}
	return days[len(days)-1]
}

func getDayNum(day string) int {
	numString := strings.TrimPrefix(day, "day")
	num, _ := strconv.Atoi(numString)
	return num
}

func getDays() []string {
	days := []string{}
	fsEntries, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range fsEntries {
		if strings.HasPrefix(f.Name(), "day") {
			days = append(days, f.Name())
		}
	}
	return days
}

func renderResults(rs []RunResult) {
	totalRunTime := float64(0)
	for _, r := range rs {
		totalRunTime += r.elapsedTime()
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Day", "Part", "Solution", "Time (milliseconds)"})
	table.SetAutoMergeCellsByColumnIndex([]int{0})
	table.SetRowLine(true)
	table.SetFooter([]string{"", "", "Total milliseconds", fmt.Sprintf("%.2f", totalRunTime)})
	for _, v := range rs {
		table.Append(v.tableData())
	}
	table.Render()
}

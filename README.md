# Advent of Code 2022

### Quick hits:
* Written in `go`
* Orchestrated with a `Makefile`

### Useage
* `day=XX make init` creates a folder for the given day with the folliwing:
    * `dayXX.go`
    * `input.txt`
    * `go.mod` which includes a module rewrite to the `utils` folder at the root of this project
    * `README.md` which I use to include the problem text including my solutions
* `make run-all` runs and times all days
* `make run-current` runs and times only the latest day
  * Or run individually  with `go run .` inside the day's folder (you just won't get timing stats)

### FYI
* The `Makefile` runner compiles every day as a plugin, and rolls them up into one binary.
* Each day needs to fulfill the folliwing requirements:
  * Have a `Part1` function that takes no arguments and returns a single `any` value.
  * Have a `Part2` function that takes no arguments and returns a single `any` value.
* Each part of each day must be able to be run independently from the other part.
* The `runner`, since it rolls all the days up into a single binary, prevents easy relative file reading for each individual day. This is why the following is found in each day: 
```
//go:embed input.txt
var f embed.FS
```
* This embeds the input file into the binary itself, which almost definitely speeds up execution time as we don't need to read the input directly from disk. 

## Current Results 
*produced with `make run-all`*

```
+-----+------+------------------------------------------+---------------------+
| DAY | PART |                 SOLUTION                 | TIME (MILLISECONDS) |
+-----+------+------------------------------------------+---------------------+
|   1 |    1 |                                    66719 |                0.15 |
+     +------+------------------------------------------+---------------------+
|     |    2 |                                   198551 |                0.15 |
+-----+------+------------------------------------------+---------------------+
|   2 |    1 |                                     8890 |                0.36 |
+     +------+------------------------------------------+---------------------+
|     |    2 |                                    10238 |                0.19 |
+-----+------+------------------------------------------+---------------------+
|   3 |    1 |                                     7967 |                0.43 |
+     +------+------------------------------------------+---------------------+
|     |    2 |                                     2716 |                0.52 |
+-----+------+------------------------------------------+---------------------+
|   4 |    1 |                                      475 |                1.74 |
+     +------+------------------------------------------+---------------------+
|     |    2 |                                      825 |                0.92 |
+-----+------+------------------------------------------+---------------------+
|   5 |    1 | HNSNMTLHQ                                |                0.64 |
+     +------+------------------------------------------+---------------------+
|     |    2 | RNLFDJMCT                                |                0.58 |
+-----+------+------------------------------------------+---------------------+
|   6 |    1 |                                     1175 |                0.02 |
+     +------+------------------------------------------+---------------------+
|     |    2 |                                     3217 |                0.03 |
+-----+------+------------------------------------------+---------------------+
|   7 |    1 |                                  1391690 |                0.17 |
+     +------+------------------------------------------+---------------------+
|     |    2 |                                  5469168 |                0.17 |
+-----+------+------------------------------------------+---------------------+
|   8 |    1 |                                     1690 |                0.38 |
+     +------+------------------------------------------+---------------------+
|     |    2 |                                   535680 |                0.62 |
+-----+------+------------------------------------------+---------------------+
|   9 |    1 |                                     6354 |                1.89 |
+     +------+------------------------------------------+---------------------+
|     |    2 |                                     2651 |                8.74 |
+-----+------+------------------------------------------+---------------------+
|  10 |    1 |                                    15020 |                0.05 |
+     +------+------------------------------------------+---------------------+
|     |    2 | ####.####.#..#..##..#....###...##..###.. |                0.02 |
|     |      | #....#....#..#.#..#.#....#..#.#..#.#..#. |                     |
|     |      | ###..###..#..#.#....#....#..#.#..#.#..#. |                     |
|     |      | #....#....#..#.#.##.#....###..####.###.. |                     |
|     |      | #....#....#..#.#..#.#....#....#..#.#.... |                     |
|     |      | ####.#.....##...###.####.#....#..#.#.... |                     |
|     |      |                                          |                     |
+-----+------+------------------------------------------+---------------------+
|  11 |    1 |                                   316888 |                0.16 |
+     +------+------------------------------------------+---------------------+
|     |    2 |                              35270398814 |               16.15 |
+-----+------+------------------------------------------+---------------------+
|  12 |    1 |                                      456 |                5.24 |
+     +------+------------------------------------------+---------------------+
|     |    2 |                                      454 |               67.65 |
+-----+------+------------------------------------------+---------------------+
|  13 |    1 |                                     5208 |                1.37 |
+     +------+------------------------------------------+---------------------+
|     |    2 |                                    25792 |                3.68 |
+-----+------+------------------------------------------+---------------------+
|                         TOTAL MILLISECONDS            |       112.01        |
+-----+------+------------------------------------------+---------------------+
```

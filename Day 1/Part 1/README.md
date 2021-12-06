# [AoC 2021 - Day 1: Sonar Sweep - Part 1](https://adventofcode.com/2021/day/1)

Run
---

```
❯ go run solve.go
```

Build
-----

```
❯ go build solve.go
```

Benchmark (Uncompiled)
----------------------

```
❯ perf stat -r 10 -d go run solve.go
...

 Performance counter stats for 'go run solve.go' (10 runs):

            652.53 msec task-clock                #    1.756 CPUs utilized            ( +-  4.23% )
              2836      context-switches          #    0.004 M/sec                    ( +-  5.87% )
                56      cpu-migrations            #    0.086 K/sec                    ( +-  5.51% )
             19794      page-faults               #    0.030 M/sec                    ( +-  1.46% )
         872519896      cycles                    #    1.337 GHz                      ( +-  2.54% )  (77.95%)
         131388652      stalled-cycles-frontend   #   15.06% frontend cycles idle     ( +-  3.69% )  (78.16%)
         102723474      stalled-cycles-backend    #   11.77% backend cycles idle      ( +-  3.33% )  (77.61%)
         806440340      instructions              #    0.92  insn per cycle
                                                  #    0.16  stalled cycles per insn  ( +-  2.87% )  (75.33%)
         147952968      branches                  #  226.736 M/sec                    ( +-  2.12% )  (74.97%)
           3935972      branch-misses             #    2.66% of all branches          ( +-  2.00% )  (76.73%)
         316971840      L1-dcache-loads           #  485.755 M/sec                    ( +-  1.51% )  (78.43%)
          10125624      L1-dcache-load-misses     #    3.19% of all L1-dcache accesses  ( +-  2.39% )  (79.42%)
   <not supported>      LLC-loads
   <not supported>      LLC-load-misses

            0.3717 +- 0.0140 seconds time elapsed  ( +-  3.78% )
```

Benchmark (Compiled)
--------------------

```
❯ perf stat -r 10 -d ./solve
...

 Performance counter stats for './solve' (10 runs):

              4.23 msec task-clock                #    1.411 CPUs utilized            ( +-  5.51% )
                17      context-switches          #    0.004 M/sec                    ( +-  8.86% )
                 0      cpu-migrations            #    0.094 K/sec                    ( +- 55.28% )
               194      page-faults               #    0.046 M/sec                    ( +-  0.69% )
           4855451      cycles                    #    1.147 GHz                      ( +- 11.28% )  (91.93%)
           1447420      stalled-cycles-frontend   #   29.81% frontend cycles idle     ( +-  8.77% )
            651511      stalled-cycles-backend    #   13.42% backend cycles idle      ( +-  9.67% )
           3358331      instructions              #    0.69  insn per cycle
                                                  #    0.43  stalled cycles per insn  ( +-  0.75% )
            697257      branches                  #  164.654 M/sec                    ( +-  0.72% )
              4654      branch-misses             #    0.67% of all branches          ( +- 32.14% )
                 0      L1-dcache-loads           #    0.000 K/sec                    (30.97%)
     <not counted>      L1-dcache-load-misses                                         (0.00%)
   <not supported>      LLC-loads
   <not supported>      LLC-load-misses

          0.003002 +- 0.000141 seconds time elapsed  ( +-  4.68% )
```

Leaderboard
-----------

Time: `23:37:51`

Rank: `105980`

Score: `0`

Bonus
-----

A teammate said they had done it in Perl for nostalgia sake, and pondered if it could be done in under 40 characters.
Challenge accepted:

```
$ perl -pe'$\+=$l&&$_>$l;$l=$_}{'<input
```

1) perl -p to wrap the statement in `while (<>) { ... ; print }`, by using `}{` at the end it will only print once after execution since we isolate that out of the while loop
2) `$\` is the output separator, but is uninitialized here, so `$\+=` will increment it if the `$l&&$_>$l;` statement is `true` (`1`)
3) `$l` will be uninitialized on the first run
4) `$_` is the input from the current line from the file being passed in
5) so, if `$l` has been set, and the current line is greater than `$l`, then it returns `true` (`1`), so `$\` will increment
6) finish by setting `$l` to the current line `$_` and continue the while loop
7) our default return now uses the output record separator as part of the print statement, so it will print out our incremented value stored in `$\`

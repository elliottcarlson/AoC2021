# [AoC 2021 - Day 1: Sonar Sweep - Part 2](https://adventofcode.com/2021/day/1#part2)

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

            629.06 msec task-clock                #    1.741 CPUs utilized            ( +-  2.80% )
              2599      context-switches          #    0.004 M/sec                    ( +-  4.62% )
                56      cpu-migrations            #    0.090 K/sec                    ( +-  4.94% )
             19493      page-faults               #    0.031 M/sec                    ( +-  1.18% )
         852534131      cycles                    #    1.355 GHz                      ( +-  1.29% )  (78.05%)
         124780669      stalled-cycles-frontend   #   14.64% frontend cycles idle     ( +-  3.17% )  (76.79%)
         104973289      stalled-cycles-backend    #   12.31% backend cycles idle      ( +-  4.14% )  (75.12%)
         808888568      instructions              #    0.95  insn per cycle
                                                  #    0.15  stalled cycles per insn  ( +-  1.45% )  (75.05%)
         143457944      branches                  #  228.050 M/sec                    ( +-  2.33% )  (77.10%)
           3738226      branch-misses             #    2.61% of all branches          ( +-  1.56% )  (79.96%)
         311197005      L1-dcache-loads           #  494.699 M/sec                    ( +-  1.55% )  (79.19%)
          10118108      L1-dcache-load-misses     #    3.25% of all L1-dcache accesses  ( +-  1.76% )  (78.17%)
   <not supported>      LLC-loads
   <not supported>      LLC-load-misses

           0.36128 +- 0.00859 seconds time elapsed  ( +-  2.38% )
```

Benchmark (Compiled)
--------------------

```
❯ perf stat -r 10 -d ./solve
...

 Performance counter stats for './solve' (10 runs):

              3.53 msec task-clock                #    1.343 CPUs utilized            ( +-  3.70% )
                17      context-switches          #    0.005 M/sec                    ( +-  7.63% )
                 0      cpu-migrations            #    0.057 K/sec                    ( +- 66.67% )
               193      page-faults               #    0.055 M/sec                    ( +-  0.46% )
           5295278      cycles                    #    1.500 GHz                      ( +-  7.10% )  (96.32%)
           1530047      stalled-cycles-frontend   #   28.89% frontend cycles idle     ( +-  8.16% )
            635941      stalled-cycles-backend    #   12.01% backend cycles idle      ( +- 14.97% )
           3347093      instructions              #    0.63  insn per cycle
                                                  #    0.46  stalled cycles per insn  ( +-  0.68% )
            688986      branches                  #  195.139 M/sec                    ( +-  0.64% )
              1988      branch-misses             #    0.29% of all branches          ( +- 72.85% )
     <not counted>      L1-dcache-loads                                               (14.27%)
     <not counted>      L1-dcache-load-misses                                         (0.00%)
   <not supported>      LLC-loads
   <not supported>      LLC-load-misses

         0.0026288 +- 0.0000787 seconds time elapsed  ( +-  2.99% )
```

Leaderboard
-----------

Time: `>24h`

Rank: `95151`

Score: `0`

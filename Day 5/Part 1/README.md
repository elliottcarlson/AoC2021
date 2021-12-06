
# [AoC 2021 - Day 5: Hydrothermal Venture - Part 1](ihttps://adventofcode.com/2021/day/5)

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

            616.22 msec task-clock                #    1.653 CPUs utilized            ( +-  4.92% )
              2376      context-switches          #    0.004 M/sec                    ( +-  3.07% )
                54      cpu-migrations            #    0.088 K/sec                    ( +-  7.58% )
             20397      page-faults               #    0.033 M/sec                    ( +-  1.51% )
         861540249      cycles                    #    1.398 GHz                      ( +-  1.77% )  (75.24%)
         126459451      stalled-cycles-frontend   #   14.68% frontend cycles idle     ( +-  2.23% )  (77.38%)
         111108690      stalled-cycles-backend    #   12.90% backend cycles idle      ( +-  3.91% )  (78.49%)
         829226226      instructions              #    0.96  insn per cycle
                                                  #    0.15  stalled cycles per insn  ( +-  2.02% )  (77.72%)
         150463685      branches                  #  244.173 M/sec                    ( +-  2.99% )  (77.25%)
           3690523      branch-misses             #    2.45% of all branches          ( +-  2.49% )  (77.86%)
         312512491      L1-dcache-loads           #  507.146 M/sec                    ( +-  1.47% )  (78.52%)
          11084689      L1-dcache-load-misses     #    3.55% of all L1-dcache accesses  ( +-  1.70% )  (77.49%)
   <not supported>      LLC-loads
   <not supported>      LLC-load-misses

            0.3727 +- 0.0241 seconds time elapsed  ( +-  6.47% )
```

Benchmark (Compiled)
--------------------

```
❯ perf stat -r 10 -d ./solve
...

 Performance counter stats for './solve' (10 runs):

             20.60 msec task-clock                #    1.411 CPUs utilized            ( +-  5.68% )
                71      context-switches          #    0.003 M/sec                    ( +-  6.06% )
                 2      cpu-migrations            #    0.087 K/sec                    ( +- 11.11% )
               991      page-faults               #    0.048 M/sec                    ( +-  0.78% )
          14975209      cycles                    #    0.727 GHz                      ( +-  9.95% )  (45.60%)
           5535654      stalled-cycles-frontend   #   36.97% frontend cycles idle     ( +-  3.34% )  (70.33%)
           1604393      stalled-cycles-backend    #   10.71% backend cycles idle      ( +-  6.32% )  (92.53%)
          26458429      instructions              #    1.77  insn per cycle
                                                  #    0.21  stalled cycles per insn  ( +-  1.38% )
           5485192      branches                  #  266.333 M/sec                    ( +-  0.12% )
             27876      branch-misses             #    0.51% of all branches          ( +- 11.94% )
          11864073      L1-dcache-loads           #  576.059 M/sec                    ( +-  4.26% )  (74.93%)
           1068992      L1-dcache-load-misses     #    9.01% of all L1-dcache accesses  ( +- 10.68% )  (41.62%)
   <not supported>      LLC-loads
   <not supported>      LLC-load-misses

          0.014593 +- 0.000392 seconds time elapsed  ( +-  2.69% )
```

Leaderboard
-----------

Time: `01:06:60`

Rank: `7874`

Score: `0`

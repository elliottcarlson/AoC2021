# [AoC 2021 - Day 6: Lanternfish - Part 2](https://adventofcode.com/2021/day/6#part2)

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

            606.81 msec task-clock                #    1.727 CPUs utilized            ( +-  1.73% )
              2487      context-switches          #    0.004 M/sec                    ( +-  3.03% )
                59      cpu-migrations            #    0.097 K/sec                    ( +-  8.41% )
             19596      page-faults               #    0.032 M/sec                    ( +-  1.52% )
         838309090      cycles                    #    1.382 GHz                      ( +-  1.43% )  (77.21%)
         123130624      stalled-cycles-frontend   #   14.69% frontend cycles idle     ( +-  1.91% )  (77.08%)
         106357039      stalled-cycles-backend    #   12.69% backend cycles idle      ( +-  4.22% )  (77.52%)
         815062216      instructions              #    0.97  insn per cycle
                                                  #    0.15  stalled cycles per insn  ( +-  3.42% )  (78.10%)
         147623090      branches                  #  243.279 M/sec                    ( +-  3.12% )  (76.91%)
           3675860      branch-misses             #    2.49% of all branches          ( +-  1.97% )  (76.12%)
         309258841      L1-dcache-loads           #  509.650 M/sec                    ( +-  1.14% )  (78.23%)
          10017948      L1-dcache-load-misses     #    3.24% of all L1-dcache accesses  ( +-  1.00% )  (78.59%)
   <not supported>      LLC-loads
   <not supported>      LLC-load-misses

           0.35146 +- 0.00515 seconds time elapsed  ( +-  1.46% )
```

Benchmark (Compiled)
--------------------

```
❯ perf stat -r 10 -d ./solve
...

 Performance counter stats for './solve' (10 runs):

              3.14 msec task-clock                #    1.263 CPUs utilized            ( +-  7.74% )
                15      context-switches          #    0.005 M/sec                    ( +- 13.41% )
                 0      cpu-migrations            #    0.159 K/sec                    ( +- 44.72% )
               196      page-faults               #    0.062 M/sec                    ( +-  0.86% )
           4837616      cycles                    #    1.541 GHz                      ( +-  2.87% )  (96.27%)
           1232793      stalled-cycles-frontend   #   25.48% frontend cycles idle     ( +-  5.74% )
            495085      stalled-cycles-backend    #   10.23% backend cycles idle      ( +- 14.51% )
           2171170      instructions              #    0.45  insn per cycle
                                                  #    0.57  stalled cycles per insn  ( +-  1.25% )
            415266      branches                  #  132.241 M/sec                    ( +-  1.36% )
              1014      branch-misses             #    0.24% of all branches          ( +- 69.65% )
     <not counted>      L1-dcache-loads                                               (7.04%)
     <not counted>      L1-dcache-load-misses                                         (0.00%)
   <not supported>      LLC-loads
   <not supported>      LLC-load-misses

          0.002486 +- 0.000256 seconds time elapsed  ( +- 10.29% )
```

Leaderboard
-----------

Time: `00:31:08`

Rank: `3706`

Score: `0`

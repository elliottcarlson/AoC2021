# [AoC 2021 - Day 2: Dive! - Part 1](https://adventofcode.com/2021/day/2)
================================

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

            597.08 msec task-clock                #    1.718 CPUs utilized            ( +-  1.73% )
              2508      context-switches          #    0.004 M/sec                    ( +-  3.77% )
                55      cpu-migrations            #    0.092 K/sec                    ( +-  4.82% )
             19562      page-faults               #    0.033 M/sec                    ( +-  0.52% )
         829570292      cycles                    #    1.389 GHz                      ( +-  1.75% )  (77.14%)
         123409910      stalled-cycles-frontend   #   14.88% frontend cycles idle     ( +-  1.34% )  (78.54%)
         101782546      stalled-cycles-backend    #   12.27% backend cycles idle      ( +-  3.51% )  (79.22%)
         808285166      instructions              #    0.97  insn per cycle
                                                  #    0.15  stalled cycles per insn  ( +-  2.39% )  (77.26%)
         140631653      branches                  #  235.531 M/sec                    ( +-  1.42% )  (75.81%)
           3577449      branch-misses             #    2.54% of all branches          ( +-  1.81% )  (76.77%)
         298398401      L1-dcache-loads           #  499.760 M/sec                    ( +-  1.88% )  (78.02%)
           9626457      L1-dcache-load-misses     #    3.23% of all L1-dcache accesses  ( +-  1.58% )  (77.26%)
   <not supported>      LLC-loads
   <not supported>      LLC-load-misses

           0.34750 +- 0.00424 seconds time elapsed  ( +-  1.22% )
```

Benchmark (Compiled)
--------------------

```
❯ perf stat -r 10 -d ./solve
...

 Performance counter stats for './solve' (10 runs):

              4.83 msec task-clock                #    1.325 CPUs utilized            ( +- 12.43% )
                26      context-switches          #    0.005 M/sec                    ( +- 26.02% )
                 1      cpu-migrations            #    0.145 K/sec                    ( +- 37.19% )
               205      page-faults               #    0.042 M/sec                    ( +-  0.29% )
           5707138      cycles                    #    1.182 GHz                      ( +- 21.66% )  (75.72%)
           1575133      stalled-cycles-frontend   #   27.60% frontend cycles idle     ( +-  6.24% )
            770413      stalled-cycles-backend    #   13.50% backend cycles idle      ( +- 11.90% )
           3369005      instructions              #    0.59  insn per cycle
                                                  #    0.47  stalled cycles per insn  ( +-  2.06% )
            666075      branches                  #  137.904 M/sec                    ( +-  2.27% )
              9170      branch-misses             #    1.38% of all branches          ( +- 19.26% )
              4652      L1-dcache-loads           #    0.963 M/sec                    ( +-100.00% )  (48.63%)
     <not counted>      L1-dcache-load-misses                                         (1.15%)
   <not supported>      LLC-loads
   <not supported>      LLC-load-misses

          0.003645 +- 0.000440 seconds time elapsed  ( +- 12.07% )
```

Leaderboard
-----------

Time: `00:14:45`

Rank: `8726`

Score: `0`

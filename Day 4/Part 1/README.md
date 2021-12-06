# [AoC 2021 - Day 4: Giant Squid - Part 1](ihttps://adventofcode.com/2021/day/4)

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

            649.03 msec task-clock                #    1.713 CPUs utilized            ( +-  1.34% )
              2706      context-switches          #    0.004 M/sec                    ( +-  2.62% )
                54      cpu-migrations            #    0.084 K/sec                    ( +-  3.75% )
             19649      page-faults               #    0.030 M/sec                    ( +-  0.62% )
         885903327      cycles                    #    1.365 GHz                      ( +-  1.77% )  (76.92%)
         132597290      stalled-cycles-frontend   #   14.97% frontend cycles idle     ( +-  1.35% )  (77.97%)
         105091619      stalled-cycles-backend    #   11.86% backend cycles idle      ( +-  2.15% )  (78.61%)
         831145238      instructions              #    0.94  insn per cycle
                                                  #    0.16  stalled cycles per insn  ( +-  1.67% )  (75.86%)
         152095985      branches                  #  234.345 M/sec                    ( +-  1.52% )  (75.67%)
           3909461      branch-misses             #    2.57% of all branches          ( +-  1.26% )  (78.08%)
         326863687      L1-dcache-loads           #  503.622 M/sec                    ( +-  0.90% )  (79.14%)
          10449925      L1-dcache-load-misses     #    3.20% of all L1-dcache accesses  ( +-  1.61% )  (77.26%)
   <not supported>      LLC-loads
   <not supported>      LLC-load-misses

           0.37888 +- 0.00399 seconds time elapsed  ( +-  1.05% )
```

Benchmark (Compiled)
--------------------

```
❯ perf stat -r 10 -d ./solve
...

 Performance counter stats for './solve' (10 runs):

              7.10 msec task-clock                #    1.418 CPUs utilized            ( +-  4.16% )
                30      context-switches          #    0.004 M/sec                    ( +-  3.72% )
                 1      cpu-migrations            #    0.197 K/sec                    ( +- 21.82% )
               267      page-faults               #    0.038 M/sec                    ( +-  0.89% )
           7125046      cycles                    #    1.004 GHz                      ( +- 19.97% )  (76.09%)
           1879875      stalled-cycles-frontend   #   26.38% frontend cycles idle     ( +-  6.45% )  (98.99%)
           1558701      stalled-cycles-backend    #   21.88% backend cycles idle      ( +-  6.24% )
          11380540      instructions              #    1.60  insn per cycle
                                                  #    0.17  stalled cycles per insn  ( +-  0.19% )
           2277722      branches                  #  320.959 M/sec                    ( +-  0.19% )
             19095      branch-misses             #    0.84% of all branches          ( +- 22.27% )
             95461      L1-dcache-loads           #   13.452 M/sec                    ( +-100.00% )  (51.69%)
     <not counted>      L1-dcache-load-misses                                         (3.73%)
   <not supported>      LLC-loads
   <not supported>      LLC-load-misses

          0.005005 +- 0.000195 seconds time elapsed  ( +-  3.90% )
```

Leaderboard
-----------

Time: `01:54:15`

Rank: `9591`

Score: `0`

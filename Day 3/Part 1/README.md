# [AoC 2021 - Day 3: Binary Diagnostic - Part 1](https://adventofcode.com/2021/day/3)

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

            610.08 msec task-clock                #    1.717 CPUs utilized            ( +-  1.55% )
              2504      context-switches          #    0.004 M/sec                    ( +-  2.45% )
                60      cpu-migrations            #    0.098 K/sec                    ( +-  4.76% )
             19545      page-faults               #    0.032 M/sec                    ( +-  0.46% )
         852698585      cycles                    #    1.398 GHz                      ( +-  1.86% )  (76.27%)
         124670032      stalled-cycles-frontend   #   14.62% frontend cycles idle     ( +-  3.80% )  (76.69%)
         109093070      stalled-cycles-backend    #   12.79% backend cycles idle      ( +-  5.33% )  (77.14%)
         811643704      instructions              #    0.95  insn per cycle
                                                  #    0.15  stalled cycles per insn  ( +-  0.92% )  (77.51%)
         145970354      branches                  #  239.264 M/sec                    ( +-  1.12% )  (78.20%)
           3733597      branch-misses             #    2.56% of all branches          ( +-  1.87% )  (77.81%)
         316329049      L1-dcache-loads           #  518.504 M/sec                    ( +-  1.64% )  (78.18%)
           9897186      L1-dcache-load-misses     #    3.13% of all L1-dcache accesses  ( +-  1.62% )  (77.68%)
   <not supported>      LLC-loads
   <not supported>      LLC-load-misses

           0.35526 +- 0.00240 seconds time elapsed  ( +-  0.68% )
```

Benchmark (Compiled)
--------------------

```
❯ perf stat -r 10 -d ./solve
...

 Performance counter stats for './solve' (10 runs):

              5.08 msec task-clock                #    1.411 CPUs utilized            ( +-  5.81% )
                24      context-switches          #    0.005 M/sec                    ( +-  4.10% )
                 1      cpu-migrations            #    0.157 K/sec                    ( +- 16.67% )
               198      page-faults               #    0.039 M/sec                    ( +-  0.56% )
           6266663      cycles                    #    1.233 GHz                      ( +- 12.63% )  (85.24%)
           1249923      stalled-cycles-frontend   #   19.95% frontend cycles idle     ( +-  5.68% )
            640385      stalled-cycles-backend    #   10.22% backend cycles idle      ( +- 11.60% )
           8869192      instructions              #    1.42  insn per cycle
                                                  #    0.14  stalled cycles per insn  ( +-  0.26% )
           2183088      branches                  #  429.377 M/sec                    ( +-  0.23% )
             11015      branch-misses             #    0.50% of all branches          ( +- 30.62% )
                 0      L1-dcache-loads           #    0.000 K/sec                    (38.49%)
     <not counted>      L1-dcache-load-misses                                         (0.00%)
   <not supported>      LLC-loads
   <not supported>      LLC-load-misses

          0.003605 +- 0.000176 seconds time elapsed  ( +-  4.89% )
```

Leaderboard
-----------

Time: `00:24:12`

Rank: `8428`

Score: `0`

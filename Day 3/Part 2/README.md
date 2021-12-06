# [AoC 2021 - Day 3: Binary Diagnostic - Part 2](https://adventofcode.com/2021/day/3#part2)

Save the input file as `input` in this directory.

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

            615.97 msec task-clock                #    1.707 CPUs utilized            ( +-  2.30% )
              2590      context-switches          #    0.004 M/sec                    ( +-  3.92% )
                58      cpu-migrations            #    0.094 K/sec                    ( +-  4.86% )
             19575      page-faults               #    0.032 M/sec                    ( +-  0.56% )
         851773862      cycles                    #    1.383 GHz                      ( +-  1.85% )  (77.18%)
         129426273      stalled-cycles-frontend   #   15.19% frontend cycles idle     ( +-  2.70% )  (78.17%)
         104025764      stalled-cycles-backend    #   12.21% backend cycles idle      ( +-  2.92% )  (78.08%)
         809951641      instructions              #    0.95  insn per cycle
                                                  #    0.16  stalled cycles per insn  ( +-  1.57% )  (76.97%)
         143501123      branches                  #  232.966 M/sec                    ( +-  1.28% )  (76.67%)
           3804829      branch-misses             #    2.65% of all branches          ( +-  1.64% )  (77.15%)
         311402976      L1-dcache-loads           #  505.545 M/sec                    ( +-  1.53% )  (78.23%)
           9995009      L1-dcache-load-misses     #    3.21% of all L1-dcache accesses  ( +-  2.03% )  (77.53%)
   <not supported>      LLC-loads
   <not supported>      LLC-load-misses

           0.36085 +- 0.00793 seconds time elapsed  ( +-  2.20% )
```

Benchmark (Compiled)
--------------------

```
❯ perf stat -r 10 -d ./solve
...

 Performance counter stats for './solve' (10 runs):

              5.45 msec task-clock                #    1.423 CPUs utilized            ( +-  4.90% )
                22      context-switches          #    0.004 M/sec                    ( +-  6.29% )
                 0      cpu-migrations            #    0.092 K/sec                    ( +- 44.72% )
               245      page-faults               #    0.045 M/sec                    ( +-  0.06% )
           6588459      cycles                    #    1.208 GHz                      ( +- 12.85% )  (92.67%)
           1594737      stalled-cycles-frontend   #   24.21% frontend cycles idle     ( +-  6.95% )
           1070436      stalled-cycles-backend    #   16.25% backend cycles idle      ( +-  5.02% )
           6824637      instructions              #    1.04  insn per cycle
                                                  #    0.23  stalled cycles per insn  ( +-  0.38% )
           1461662      branches                  #  268.090 M/sec                    ( +-  0.34% )
             10132      branch-misses             #    0.69% of all branches          ( +- 33.93% )
                 0      L1-dcache-loads           #    0.000 K/sec                    (28.56%)
     <not counted>      L1-dcache-load-misses                                         (0.00%)
   <not supported>      LLC-loads
   <not supported>      LLC-load-misses

          0.003831 +- 0.000158 seconds time elapsed  ( +-  4.13% )
```

Leaderboard
-----------

Time: `00:50:56`

Rank: `5519`

Score: `0`

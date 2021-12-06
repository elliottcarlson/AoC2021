# [AoC 2021 - Day 2: Dive! - Part 2](https://adventofcode.com/2021/day/2#part2)

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

            587.52 msec task-clock                #    1.684 CPUs utilized            ( +-  2.44% )
              2422      context-switches          #    0.004 M/sec                    ( +-  3.95% )
                54      cpu-migrations            #    0.092 K/sec                    ( +-  4.01% )
             19442      page-faults               #    0.033 M/sec                    ( +-  0.47% )
         846008102      cycles                    #    1.440 GHz                      ( +-  2.18% )  (75.80%)
         127729533      stalled-cycles-frontend   #   15.10% frontend cycles idle     ( +-  2.11% )  (77.44%)
         102858997      stalled-cycles-backend    #   12.16% backend cycles idle      ( +-  2.27% )  (79.07%)
         788994247      instructions              #    0.93  insn per cycle
                                                  #    0.16  stalled cycles per insn  ( +-  1.59% )  (79.44%)
         142009873      branches                  #  241.712 M/sec                    ( +-  2.09% )  (76.94%)
           3666308      branch-misses             #    2.58% of all branches          ( +-  3.11% )  (76.34%)
         299631126      L1-dcache-loads           #  509.997 M/sec                    ( +-  1.66% )  (77.54%)
           9944633      L1-dcache-load-misses     #    3.32% of all L1-dcache accesses  ( +-  1.14% )  (76.82%)
   <not supported>      LLC-loads
   <not supported>      LLC-load-misses

           0.34891 +- 0.00720 seconds time elapsed  ( +-  2.06% )
```

Benchmark (Compiled)
--------------------
```
❯ perf stat -r 10 -d ./solve
...

 Performance counter stats for './solve' (10 runs):

              4.68 msec task-clock                #    1.362 CPUs utilized            ( +-  6.34% )
                20      context-switches          #    0.004 M/sec                    ( +- 13.01% )
                 1      cpu-migrations            #    0.192 K/sec                    ( +- 25.93% )
               203      page-faults               #    0.043 M/sec                    ( +-  0.87% )
           5348754      cycles                    #    1.144 GHz                      ( +- 10.58% )  (89.24%)
           1388714      stalled-cycles-frontend   #   25.96% frontend cycles idle     ( +-  6.59% )
            615096      stalled-cycles-backend    #   11.50% backend cycles idle      ( +-  2.49% )
           3305709      instructions              #    0.62  insn per cycle
                                                  #    0.42  stalled cycles per insn  ( +-  1.07% )
            653342      branches                  #  139.680 M/sec                    ( +-  1.13% )
              5760      branch-misses             #    0.88% of all branches          ( +- 34.78% )
                 0      L1-dcache-loads           #    0.000 K/sec                    (27.46%)
     <not counted>      L1-dcache-load-misses                                         (0.00%)
   <not supported>      LLC-loads
   <not supported>      LLC-load-misses

          0.003435 +- 0.000229 seconds time elapsed  ( +-  6.67% )
```

Leaderboard
-----------

Time: `00:19:49`

Rank: `8091`

Score: `0`

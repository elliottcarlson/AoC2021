# [AoC 2021 - Day 4: Giant Squid - Part 2](https://adventofcode.com/2021/day/4#part2)

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

            631.99 msec task-clock                #    1.737 CPUs utilized            ( +-  1.74% )
              2650      context-switches          #    0.004 M/sec                    ( +-  3.14% )
                56      cpu-migrations            #    0.089 K/sec                    ( +-  5.52% )
             19618      page-faults               #    0.031 M/sec                    ( +-  0.66% )
         852536163      cycles                    #    1.349 GHz                      ( +-  1.58% )  (75.71%)
         125753043      stalled-cycles-frontend   #   14.75% frontend cycles idle     ( +-  2.37% )  (76.90%)
         116039137      stalled-cycles-backend    #   13.61% backend cycles idle      ( +-  2.39% )  (78.82%)
         859342533      instructions              #    1.01  insn per cycle
                                                  #    0.15  stalled cycles per insn  ( +-  1.78% )  (79.51%)
         154824433      branches                  #  244.980 M/sec                    ( +-  1.60% )  (77.61%)
           3848582      branch-misses             #    2.49% of all branches          ( +-  1.56% )  (77.41%)
         325723345      L1-dcache-loads           #  515.395 M/sec                    ( +-  1.86% )  (77.55%)
          10163619      L1-dcache-load-misses     #    3.12% of all L1-dcache accesses  ( +-  2.06% )  (76.30%)
   <not supported>      LLC-loads
   <not supported>      LLC-load-misses

           0.36379 +- 0.00452 seconds time elapsed  ( +-  1.24% )
```

Benchmark (Compiled)
--------------------

```
❯ perf stat -r 10 -d ./solve
...

 Performance counter stats for './solve' (10 runs):

              7.44 msec task-clock                #    1.456 CPUs utilized            ( +-  5.19% )
                34      context-switches          #    0.005 M/sec                    ( +-  3.62% )
                 1      cpu-migrations            #    0.148 K/sec                    ( +- 21.21% )
               268      page-faults               #    0.036 M/sec                    ( +-  0.79% )
           8664420      cycles                    #    1.165 GHz                      ( +- 22.84% )  (63.15%)
           1653987      stalled-cycles-frontend   #   19.09% frontend cycles idle     ( +-  4.33% )  (99.86%)
           1680122      stalled-cycles-backend    #   19.39% backend cycles idle      ( +-  6.81% )
          15774938      instructions              #    1.82  insn per cycle
                                                  #    0.11  stalled cycles per insn  ( +-  0.14% )
           3018616      branches                  #  405.743 M/sec                    ( +-  0.15% )
             25506      branch-misses             #    0.84% of all branches          ( +- 17.35% )
            205515      L1-dcache-loads           #   27.624 M/sec                    ( +- 97.65% )  (53.09%)
     <not counted>      L1-dcache-load-misses                                         (2.72%)
   <not supported>      LLC-loads
   <not supported>      LLC-load-misses

          0.005109 +- 0.000218 seconds time elapsed  ( +-  4.27% )
```

Leaderboard
-----------

Time: `02:19:55`

Rank: `9245`

Score: `0`

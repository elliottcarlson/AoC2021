# [AoC 2021 - Day 6: Lanternfish - Part 1](https://adventofcode.com/2021/day/6)

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

            600.50 msec task-clock                #    1.686 CPUs utilized            ( +-  5.76% )
              2335      context-switches          #    0.004 M/sec                    ( +-  5.72% )
                52      cpu-migrations            #    0.087 K/sec                    ( +-  6.53% )
             19615      page-faults               #    0.033 M/sec                    ( +-  0.96% )
         823831315      cycles                    #    1.372 GHz                      ( +-  3.05% )  (79.62%)
         124348347      stalled-cycles-frontend   #   15.09% frontend cycles idle     ( +-  4.14% )  (78.48%)
         108222298      stalled-cycles-backend    #   13.14% backend cycles idle      ( +-  5.47% )  (76.60%)
         811076352      instructions              #    0.98  insn per cycle
                                                  #    0.15  stalled cycles per insn  ( +-  2.07% )  (75.51%)
         144138737      branches                  #  240.033 M/sec                    ( +-  3.03% )  (75.31%)
           3803464      branch-misses             #    2.64% of all branches          ( +-  3.13% )  (76.60%)
         297653639      L1-dcache-loads           #  495.679 M/sec                    ( +-  2.44% )  (79.36%)
           9993161      L1-dcache-load-misses     #    3.36% of all L1-dcache accesses  ( +-  3.46% )  (78.40%)
   <not supported>      LLC-loads
   <not supported>      LLC-load-misses

            0.3562 +- 0.0192 seconds time elapsed  ( +-  5.40% )
```

Benchmark (Compiled)
--------------------

```
❯ perf stat -r 10 -d ./solve
...

 Performance counter stats for './solve' (10 runs):

              3.53 msec task-clock                #    1.319 CPUs utilized            ( +-  3.07% )
                13      context-switches          #    0.004 M/sec                    ( +-  6.60% )
                 0      cpu-migrations            #    0.113 K/sec                    ( +- 40.82% )
               194      page-faults               #    0.055 M/sec                    ( +-  0.21% )
           5036164      cycles                    #    1.427 GHz                      ( +-  3.82% )  (98.95%)
           1437100      stalled-cycles-frontend   #   28.54% frontend cycles idle     ( +-  4.89% )
            489111      stalled-cycles-backend    #    9.71% backend cycles idle      ( +-  9.24% )
           2145764      instructions              #    0.43  insn per cycle
                                                  #    0.67  stalled cycles per insn  ( +-  0.63% )
            409188      branches                  #  115.931 M/sec                    ( +-  0.72% )
               844      branch-misses             #    0.21% of all branches          ( +- 63.82% )
     <not counted>      L1-dcache-loads                                               (8.46%)
     <not counted>      L1-dcache-load-misses                                         (0.00%)
   <not supported>      LLC-loads
   <not supported>      LLC-load-misses

         0.0026768 +- 0.0000888 seconds time elapsed  ( +-  3.32% )
```

Leaderboard
-----------

Time: `00:29:25`

Rank: `8358`

Score: `0`

# [AoC 2021 - Day 5: Hydrothermal Venture - Part 2](ihttps://adventofcode.com/2021/day/5#part2)

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

            624.08 msec task-clock                #    1.696 CPUs utilized            ( +-  1.92% )
              2468      context-switches          #    0.004 M/sec                    ( +-  4.07% )
                58      cpu-migrations            #    0.092 K/sec                    ( +-  4.29% )
             19872      page-faults               #    0.032 M/sec                    ( +-  0.44% )
         843075696      cycles                    #    1.351 GHz                      ( +-  2.22% )  (79.37%)
         124937867      stalled-cycles-frontend   #   14.82% frontend cycles idle     ( +-  2.56% )  (78.87%)
         102961719      stalled-cycles-backend    #   12.21% backend cycles idle      ( +-  1.94% )  (76.85%)
         809907446      instructions              #    0.96  insn per cycle
                                                  #    0.15  stalled cycles per insn  ( +-  2.07% )  (76.98%)
         148252354      branches                  #  237.552 M/sec                    ( +-  2.15% )  (76.26%)
           3721939      branch-misses             #    2.51% of all branches          ( +-  1.32% )  (77.18%)
         317788934      L1-dcache-loads           #  509.209 M/sec                    ( +-  1.65% )  (77.21%)
          11051410      L1-dcache-load-misses     #    3.48% of all L1-dcache accesses  ( +-  3.19% )  (77.23%)
   <not supported>      LLC-loads
   <not supported>      LLC-load-misses

           0.36808 +- 0.00453 seconds time elapsed  ( +-  1.23% )
```

Benchmark (Compiled)
--------------------

```
❯ perf stat -r 10 -d ./solve
...

 Performance counter stats for './solve' (10 runs):

             17.40 msec task-clock                #    1.482 CPUs utilized            ( +-  4.46% )
                69      context-switches          #    0.004 M/sec                    ( +-  4.18% )
                 2      cpu-migrations            #    0.121 K/sec                    ( +- 11.11% )
               580      page-faults               #    0.033 M/sec                    ( +-  1.98% )
          15679290      cycles                    #    0.901 GHz                      ( +- 18.23% )  (49.65%)
           4162101      stalled-cycles-frontend   #   26.55% frontend cycles idle     ( +-  3.97% )  (79.28%)
           3336912      stalled-cycles-backend    #   21.28% backend cycles idle      ( +- 44.45% )  (96.32%)
          15295475      instructions              #    0.98  insn per cycle
                                                  #    0.27  stalled cycles per insn  ( +-  0.36% )
           2812231      branches                  #  161.625 M/sec                    ( +-  0.36% )
             16865      branch-misses             #    0.60% of all branches          ( +- 20.28% )
           7870183      L1-dcache-loads           #  452.316 M/sec                    ( +-  9.39% )  (63.78%)
            333029      L1-dcache-load-misses     #    4.23% of all L1-dcache accesses  ( +- 36.17% )  (34.07%)
   <not supported>      LLC-loads
   <not supported>      LLC-load-misses

          0.011739 +- 0.000315 seconds time elapsed  ( +-  2.68% )
```

Leaderboard
-----------

Time: `01:31:44`

Rank: `7300`

Score: `0`

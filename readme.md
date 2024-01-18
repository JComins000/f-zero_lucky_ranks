# Intro

This project is meant for me to practice go concepts while simulating an interesting question posed by my favorite game at the moment. [TL;DR here](#Takeaways)

## Golang

I know Go's main use cases range from APIs, and CLIs, but my background is in testing and DevOps, so I went for a familiar scripting approach for the project while I acclimated to the language. I got to play around types, structs, interfaces, loops and ranges, casting, fixed-length arrays, testing, and print formatting. This project was worth my time.

## The Game

On Nov. 28 2023, F-Zero 99 updated to include a new **lucky ranks** feature. Your placements in the first five races in a day would be compared 14 mystery cards, revealed to the player after they had completed their placements. If a player uncovered 5 matches from revealed cards, that player would receive an unlockable cosmetic border for their player card, viewable by other players.

There are three types of mystery cards. There are 99 Placement cards (for each possible placement), as well as Machine Cards for all four machines, and Wild Cards for each digit. To optimize your ranks for the minigame, a player will want to collect as many unique digits as possible to guarantee as many matches any WildCards the game could deal for that day. Machine Cards are revealed to the player before they complete their ranks, and a player can obtain any placement any machine. Placement cards are all unique, and by looking for new wildcard digits, a player should have unique ranks to match against anyway.

The matching strategy is easy to optimize, but I wanted to know how likely a player was receive 5 matches while trying to place optimally. For reference, here's what the game looks like. And here's a screenshot from when I got lucky on Jan. 1 2024.

Before Reveal | My 5 Matches
------------- | ------------
![Before Reveal](docs/img/lucky_rank_empty.jpg) | ![My 5 Matches](docs/img/5_matches.jpg)

# The Question

So, how likely is it for a player to receive 5 matches while trying to place optimally? Let's make some assumptions. I have to assume that all possible placements, machines, and wildcards exist once each in the deck from which the game draws it's mystery cards. Let's also assume the player will obtain 5 unique placements as well.

How many possibilities are there for different matches? How much will a player's chances vary if they fail to obtain unique placement digits? Machine cards are revealed before the player locks in their placements-- how would seeing machine cards improve the players chance of getting 5 matches?

# Random Selection for an Answer

I never intended to use pure math to solve this problem, although that's probably the fastest way to do it. Iterating through every posibility would be `(112 choose 14) * (99 choose 5)` iterations-- way too many. To make an approximation, I chose to randomly sample 50 million combinations.

## Takeaways
Unsurprisingly, the odds seemed to be really low. Overall, even if a player placed optimally to cover as many wildcards as possible, they only had a `1%` chance each day to acheive 5 matches. A player would need to place optimally for ``70 days`` in a row to have a 50% chance of being dealt 5 matches at least once. Luckily for me, I managed to snag my badge in `35 days`!

### Machine Cards
A player is more likely not to be dealt a single machine card, but each card dealt will help immensely. If a player receives 2 machine cards, their chances go up dramatically-- Almost to the point of hopelessness if a player doesn't receive one.

```
8 unique digits, 0 machine cards:       6918 matches,    4455514 misses.    4462432 sum,0.16%
8 unique digits, 1 machine cards:      21663 matches,    2580311 misses.    2601974 sum,0.83%
8 unique digits, 2 machine cards:      18763 matches,     504030 misses.     522793 sum,3.59%

9 unique digits, 0 machine cards:       1431 matches,     549427 misses.     550858 sum,0.26%
9 unique digits, 1 machine cards:       3890 matches,     317010 misses.     320900 sum,1.21%
9 unique digits, 2 machine cards:       3045 matches,      61470 misses.      64515 sum,4.72%

10 unique digits, 0 machine cards:         45 matches,      12256 misses.      12301 sum,0.37%
10 unique digits, 1 machine cards:        108 matches,       7100 misses.       7208 sum,1.50%
10 unique digits, 2 machine cards:         93 matches,       1385 misses.       1478 sum,6.29%
```

### Unique Digits
Surprisingly, when a player controls their placements to cover as many wildcards as possible, their chance of getting 5 matches rises exponentially, around an additional `.25%` for each digit covered.
```
6 unique digits:      54044 matches,   16666786 misses.   16720830 sum,0.32%
7 unique digits:      91001 matches,   18844168 misses.   18935169 sum,0.48%
8 unique digits:      52955 matches,    7577841 misses.    7630796 sum,0.69%
9 unique digits:       9237 matches,     932419 misses.     941656 sum,0.98%
10 unique digits:        265 matches,      20822 misses.      21087 sum,1.26%
```

## Program Output
```
3 unique digits, 0 machine cards:          0 matches,       5194 misses.       5194 sum,0.00%
3 unique digits, 1 machine cards:          1 matches,       3041 misses.       3042 sum,0.03%
3 unique digits, 2 machine cards:          5 matches,        578 misses.        583 sum,0.86%
3 unique digits, 3 machine cards:          1 matches,         41 misses.         42 sum,2.38%
3 unique digits, 4 machine cards:          0 matches,          4 misses.          4 sum,0.00%

4 unique digits, 0 machine cards:         32 matches,     291589 misses.     291621 sum,0.01%
4 unique digits, 1 machine cards:        192 matches,     170446 misses.     170638 sum,0.11%
4 unique digits, 2 machine cards:        263 matches,      33975 misses.      34238 sum,0.77%
4 unique digits, 3 machine cards:        143 matches,       2741 misses.       2884 sum,4.96%
4 unique digits, 4 machine cards:         15 matches,         68 misses.         83 sum,18.07%

5 unique digits, 0 machine cards:        730 matches,    3064058 misses.    3064788 sum,0.02%
5 unique digits, 1 machine cards:       3526 matches,    1784230 misses.    1787756 sum,0.20%
5 unique digits, 2 machine cards:       4530 matches,     354933 misses.     359463 sum,1.26%
5 unique digits, 3 machine cards:       1848 matches,      27513 misses.      29361 sum,6.29%
5 unique digits, 4 machine cards:        187 matches,        578 misses.        765 sum,24.44%

6 unique digits, 0 machine cards:       4819 matches,    9773848 misses.    9778667 sum,0.05%
6 unique digits, 1 machine cards:      19381 matches,    5681593 misses.    5700974 sum,0.34%
6 unique digits, 2 machine cards:      21532 matches,    1123484 misses.    1145016 sum,1.88%
6 unique digits, 3 machine cards:       7631 matches,      86005 misses.      93636 sum,8.15%
6 unique digits, 4 machine cards:        681 matches,       1856 misses.       2537 sum,26.84%

7 unique digits, 0 machine cards:      10005 matches,   11063397 misses.   11073402 sum,0.09%
7 unique digits, 1 machine cards:      35271 matches,    6418570 misses.    6453841 sum,0.55%
7 unique digits, 2 machine cards:      34276 matches,    1264643 misses.    1298919 sum,2.64%
7 unique digits, 3 machine cards:      10548 matches,      95495 misses.     106043 sum,9.95%
7 unique digits, 4 machine cards:        901 matches,       2063 misses.       2964 sum,30.40%

8 unique digits, 0 machine cards:       6918 matches,    4455514 misses.    4462432 sum,0.16%
8 unique digits, 1 machine cards:      21663 matches,    2580311 misses.    2601974 sum,0.83%
8 unique digits, 2 machine cards:      18763 matches,     504030 misses.     522793 sum,3.59%
8 unique digits, 3 machine cards:       5269 matches,      37185 misses.      42454 sum,12.41%
8 unique digits, 4 machine cards:        342 matches,        801 misses.       1143 sum,29.92%

9 unique digits, 0 machine cards:       1431 matches,     549427 misses.     550858 sum,0.26%
9 unique digits, 1 machine cards:       3890 matches,     317010 misses.     320900 sum,1.21%
9 unique digits, 2 machine cards:       3045 matches,      61470 misses.      64515 sum,4.72%
9 unique digits, 3 machine cards:        814 matches,       4415 misses.       5229 sum,15.57%
9 unique digits, 4 machine cards:         57 matches,         97 misses.        154 sum,37.01%

10 unique digits, 0 machine cards:         45 matches,      12256 misses.      12301 sum,0.37%
10 unique digits, 1 machine cards:        108 matches,       7100 misses.       7208 sum,1.50%
10 unique digits, 2 machine cards:         93 matches,       1385 misses.       1478 sum,6.29%
10 unique digits, 3 machine cards:         17 matches,         79 misses.         96 sum,17.71%
10 unique digits, 4 machine cards:          2 matches,          2 misses.          4 sum,50.00%


3 unique digits:          7 matches,       8858 misses.       8865 sum,0.08%
4 unique digits:        645 matches,     498819 misses.     499464 sum,0.13%
5 unique digits:      10821 matches,    5231312 misses.    5242133 sum,0.21%
6 unique digits:      54044 matches,   16666786 misses.   16720830 sum,0.32%
7 unique digits:      91001 matches,   18844168 misses.   18935169 sum,0.48%
8 unique digits:      52955 matches,    7577841 misses.    7630796 sum,0.69%
9 unique digits:       9237 matches,     932419 misses.     941656 sum,0.98%
10 unique digits:        265 matches,      20822 misses.      21087 sum,1.26%

0 machine cards:      23980 matches,   29215283 misses.   29239263 sum,0.08%
1 machine cards:      84032 matches,   16962301 misses.   17046333 sum,0.49%
2 machine cards:      82507 matches,    3344498 misses.    3427005 sum,2.41%
3 machine cards:      26271 matches,     253474 misses.     279745 sum,9.39%
4 machine cards:       2185 matches,       5469 misses.       7654 sum,28.55%

In total:     218975 matches,   49781025 misses.   50000000 sum,0.44%
```

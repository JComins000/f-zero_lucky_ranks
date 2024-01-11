# Intro

This project is meant for me to practice go concepts while simulating an interesting question posed by my favorite game at the moment.

## Golang

I know Go's main use cases range from APIs, and CLIs, but my background is in testing and DevOps, so I went for a familiar scripting approach for the project while I acclimated to the language. I got to play around with types, structs, interfaces, loops and ranges, casting, fixed-length arrays, testing, and print formatting. This project was worth my time.

## The Game

On Nov. 28 2023, F-Zero 99 updated to include a new **lucky ranks** feature. Your placements in the first five races in a day would be compared with 14 mystery cards, revealed to the player after they had completed their placements. If a player uncovered 5 matches from revealed cards, that player would receive an unlockable cosmetic border for their player card, viewable by other players.

There are three types of mystery cards. There are 99 Placement cards (for each possible placement), as well as Machine Cards for all four machines, and Wild Cards for each digit. To optimize your ranks for the minigame, a player will want to collect as many unique digits as possible to guarantee as many matches with any WildCards the game could deal for that day. Machine Cards are revealed to the player before they complete their ranks, and a player can obtain any placement with any machine. Placement cards are all unique, and by looking for new wildcard digits, a player should have unique ranks to match against anyway.

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
With 8 unique digits and 0 machine cards: We counted       2219 matches and    1401140 misses.    1403359 total, 0.16%
With 8 unique digits and 1 machine cards: We counted       7247 matches and     819175 misses.     826422 total, 0.88%
With 8 unique digits and 2 machine cards: We counted       6184 matches and     161210 misses.     167394 total, 3.84%

With 9 unique digits and 0 machine cards: We counted        195 matches and      67776 misses.      67971 total, 0.29%
With 9 unique digits and 1 machine cards: We counted        503 matches and      39271 misses.      39774 total, 1.28%
With 9 unique digits and 2 machine cards: We counted        412 matches and       7737 misses.       8149 total, 5.33%
```

### Unique Digits
Surprisingly, when a player controls their placements to cover as many wildcards as possible, their chance of getting 5 matches rises exponentially, around `145%` for each digit covered.
```
With 6 unique digits: We counted      67974 matches and   20064782 misses.   20132756 total, 0.34%
With 7 unique digits: We counted      60577 matches and   11919946 misses.   11980523 total, 0.51%
With 8 unique digits: We counted      17558 matches and    2393941 misses.    2411499 total, 0.73%
With 9 unique digits: We counted       1225 matches and     115369 misses.     116594 total, 1.06%
```

## Program output
```
With 2 unique digits and 0 machine cards: We counted          0 matches and        884 misses.        884 total, 0.00%
With 2 unique digits and 1 machine cards: We counted          0 matches and        506 misses.        506 total, 0.00%
With 2 unique digits and 2 machine cards: We counted          0 matches and        123 misses.        123 total, 0.00%
With 2 unique digits and 3 machine cards: We counted          0 matches and          2 misses.          2 total, 0.00%
With 2 unique digits and 4 machine cards: We counted          0 matches and          0 misses.          0 total, NaN%

With 3 unique digits and 0 machine cards: We counted          4 matches and      96602 misses.      96606 total, 0.00%
With 3 unique digits and 1 machine cards: We counted         31 matches and      57243 misses.      57274 total, 0.05%
With 3 unique digits and 2 machine cards: We counted         72 matches and      11529 misses.      11601 total, 0.62%
With 3 unique digits and 3 machine cards: We counted         28 matches and        933 misses.        961 total, 3.00%
With 3 unique digits and 4 machine cards: We counted          3 matches and         30 misses.         33 total, 10.00%

With 4 unique digits and 0 machine cards: We counted        164 matches and    1580834 misses.    1580998 total, 0.01%
With 4 unique digits and 1 machine cards: We counted       1006 matches and     928524 misses.     929530 total, 0.11%
With 4 unique digits and 2 machine cards: We counted       1583 matches and     187443 misses.     189026 total, 0.84%
With 4 unique digits and 3 machine cards: We counted        753 matches and      14996 misses.      15749 total, 5.02%
With 4 unique digits and 4 machine cards: We counted        109 matches and        330 misses.        439 total, 33.03%

With 5 unique digits and 0 machine cards: We counted       1787 matches and    7253819 misses.    7255606 total, 0.02%
With 5 unique digits and 1 machine cards: We counted       8815 matches and    4267916 misses.    4276731 total, 0.21%
With 5 unique digits and 2 machine cards: We counted      11544 matches and     857473 misses.     869017 total, 1.35%
With 5 unique digits and 3 machine cards: We counted       4651 matches and      66819 misses.      71470 total, 6.96%
With 5 unique digits and 4 machine cards: We counted        482 matches and       1590 misses.       2072 total, 30.31%

With 6 unique digits and 0 machine cards: We counted       5992 matches and   11703105 misses.   11709097 total, 0.05%
With 6 unique digits and 1 machine cards: We counted      24523 matches and    6875224 misses.    6899747 total, 0.36%
With 6 unique digits and 2 machine cards: We counted      27248 matches and    1377615 misses.    1404863 total, 1.98%
With 6 unique digits and 3 machine cards: We counted       9319 matches and     106473 misses.     115792 total, 8.75%
With 6 unique digits and 4 machine cards: We counted        892 matches and       2365 misses.       3257 total, 37.72%

With 7 unique digits and 0 machine cards: We counted       6678 matches and    6958628 misses.    6965306 total, 0.10%
With 7 unique digits and 1 machine cards: We counted      23253 matches and    4085913 misses.    4109166 total, 0.57%
With 7 unique digits and 2 machine cards: We counted      22748 matches and     812431 misses.     835179 total, 2.80%
With 7 unique digits and 3 machine cards: We counted       7296 matches and      61683 misses.      68979 total, 11.83%
With 7 unique digits and 4 machine cards: We counted        602 matches and       1291 misses.       1893 total, 46.63%

With 8 unique digits and 0 machine cards: We counted       2219 matches and    1401140 misses.    1403359 total, 0.16%
With 8 unique digits and 1 machine cards: We counted       7247 matches and     819175 misses.     826422 total, 0.88%
With 8 unique digits and 2 machine cards: We counted       6184 matches and     161210 misses.     167394 total, 3.84%
With 8 unique digits and 3 machine cards: We counted       1779 matches and      12151 misses.      13930 total, 14.64%
With 8 unique digits and 4 machine cards: We counted        129 matches and        265 misses.        394 total, 48.68%

With 9 unique digits and 0 machine cards: We counted        195 matches and      67776 misses.      67971 total, 0.29%
With 9 unique digits and 1 machine cards: We counted        503 matches and      39271 misses.      39774 total, 1.28%
With 9 unique digits and 2 machine cards: We counted        412 matches and       7737 misses.       8149 total, 5.33%
With 9 unique digits and 3 machine cards: We counted        109 matches and        574 misses.        683 total, 18.99%
With 9 unique digits and 4 machine cards: We counted          6 matches and         11 misses.         17 total, 54.55%


With 2 unique digits: We counted          0 matches and       1515 misses.       1515 total, 0.00%
With 3 unique digits: We counted        138 matches and     166337 misses.     166475 total, 0.08%
With 4 unique digits: We counted       3615 matches and    2712127 misses.    2715742 total, 0.13%
With 5 unique digits: We counted      27279 matches and   12447617 misses.   12474896 total, 0.22%
With 6 unique digits: We counted      67974 matches and   20064782 misses.   20132756 total, 0.34%
With 7 unique digits: We counted      60577 matches and   11919946 misses.   11980523 total, 0.51%
With 8 unique digits: We counted      17558 matches and    2393941 misses.    2411499 total, 0.73%
With 9 unique digits: We counted       1225 matches and     115369 misses.     116594 total, 1.06%

With 0 machine cards: We counted      17039 matches and   29062788 misses.   29079827 total, 0.06%
With 1 machine cards: We counted      65378 matches and   17073772 misses.   17139150 total, 0.38%
With 2 machine cards: We counted      69791 matches and    3415561 misses.    3485352 total, 2.04%
With 3 machine cards: We counted      23935 matches and     263631 misses.     287566 total, 9.08%
With 4 machine cards: We counted       2223 matches and       5882 misses.       8105 total, 37.79%

In total: We counted     178366 matches and   49821634 misses.   50000000 total, 0.36%
```
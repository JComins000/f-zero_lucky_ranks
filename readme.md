# Intro

This project is meant for me to practice go concepts while simulating an interesting question posed by my favorite game at the moment.

## Golang

I know Go's main use cases range from APIs, and CLIs, but my background is in testing and DevOps, so I went for a familiar scripting approach for the project while I acclimated to the language. Ultimately, I'd like to think the fundamentals I implemented with loops, structs, and interfaces would enable me to join and contribute to any existing Go repos I might encounter in my career.

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
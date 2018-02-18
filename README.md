# Quote-cli

## About

Get a random quote or search for quotes in your terminal.

_Quotes provided via https://favqs.com api_

## Installation

Install with `go get` or download a binary from the [latest release](https://github.com/jamiecollinson/quote/releases/latest). If you download a binary you'll need to ensure it's on your `path`.

``` bash
go get github.com/jamiecollinson/quote
```

## Use

Random quote:

``` bash
> quote
Science is nothing but perception.
Plato
```

Search by keyword:

``` bash
> quote -search "simplicity is"
Simplicity is the final achievement. After one has played a vast quantity of notes and more notes, it is simplicity that emerges as the crowning reward of art.
Frederic Chopin

As I grew older, I realized that it was much better to insist on the genuine forms of nature, for simplicity is the greatest adornment of art.
Albrecht Durer

Simplicity is the most difficult thing to secure in this world it is the last limit of experience and the last effort of genius.
George Sand

Simplicity is the ultimate sophistication.
Leonardo da Vinci
```

Search by author:

``` bash
> quote -author "leonardo da vinci"
Simplicity is the ultimate sophistication.
Leonardo da Vinci

Nothing strengthens authority as much as silence.
Leonardo da Vinci

The greatest deception men suffer is from their own opinions.
Leonardo da Vinci

Art is never finished, only abandoned.
Leonardo da Vinci

Where the spirit does not work with the hand, there is no art.
Leonardo da Vinci

The human foot is a masterpiece of engineering and a work of art.
Leonardo da Vinci

I love those who can smile in trouble, who can gather strength from distress, and grow brave by reflection. 'Tis the business of little minds to shrink, but they whose heart is firm, and whose conscience approves their conduct, will pursue their principles unto death.
Leonardo da Vinci

He who is fixed to a star does not change his mind.
Leonardo da Vinci

Just as courage imperils life, fear protects it.
Leonardo da Vinci

While I thought that I was learning how to live, I have been learning how to die.
Leonardo da Vinci
```

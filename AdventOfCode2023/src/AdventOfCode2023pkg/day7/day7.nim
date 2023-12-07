#[
    --- Day 7: Camel Cards ---

Your all-expenses-paid trip turns out to be a one-way, five-minute ride in an airship. (At least it's a cool airship!) It drops you off at the edge of a vast desert and descends back to Island Island.

"Did you bring the parts?"

You turn around to see an Elf completely covered in white clothing, wearing goggles, and riding a large camel.

"Did you bring the parts?" she asks again, louder this time. You aren't sure what parts she's looking for; you're here to figure out why the sand stopped.

"The parts! For the sand, yes! Come with me; I will show you." She beckons you onto the camel.

After riding a bit across the sands of Desert Island, you can see what look like very large rocks covering half of the horizon. The Elf explains that the rocks are all along the part of Desert Island that is directly above Island Island, making it hard to even get there. Normally, they use big machines to move the rocks and filter the sand, but the machines have broken down because Desert Island recently stopped receiving the parts they need to fix the machines.

You've already assumed it'll be your job to figure out why the parts stopped when she asks if you can help. You agree automatically.

Because the journey will take a few days, she offers to teach you the game of Camel Cards. Camel Cards is sort of similar to poker except it's designed to be easier to play while riding a camel.

In Camel Cards, you get a list of hands, and your goal is to order them based on the strength of each hand. A hand consists of five cards labeled one of A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2. The relative strength of each card follows this order, where A is the highest and 2 is the lowest.

Every hand is exactly one type. From strongest to weakest, they are:

    Five of a kind, where all five cards have the same label: AAAAA
    Four of a kind, where four cards have the same label and one card has a different label: AA8AA
    Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
    Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
    Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
    One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
    High card, where all cards' labels are distinct: 23456

Hands are primarily ordered based on type; for example, every full house is stronger than any three of a kind.

If two hands have the same type, a second ordering rule takes effect. Start by comparing the first card in each hand. If these cards are different, the hand with the stronger first card is considered stronger. If the first card in each hand have the same label, however, then move on to considering the second card in each hand. If they differ, the hand with the higher second card wins; otherwise, continue with the third card in each hand, then the fourth, then the fifth.

So, 33332 and 2AAAA are both four of a kind hands, but 33332 is stronger because its first card is stronger. Similarly, 77888 and 77788 are both a full house, but 77888 is stronger because its third card is stronger (and both hands have the same first and second card).

To play Camel Cards, you are given a list of hands and their corresponding bid (your puzzle input). For example:

32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483

This example shows five hands; each hand is followed by its bid amount. Each hand wins an amount equal to its bid multiplied by its rank, where the weakest hand gets rank 1, the second-weakest hand gets rank 2, and so on up to the strongest hand. Because there are five hands in this example, the strongest hand will have rank 5 and its bid will be multiplied by 5.

So, the first step is to put the hands in order of strength:

    32T3K is the only one pair and the other hands are all a stronger type, so it gets rank 1.
    KK677 and KTJJT are both two pair. Their first cards both have the same label, but the second card of KK677 is stronger (K vs T), so KTJJT gets rank 2 and KK677 gets rank 3.
    T55J5 and QQQJA are both three of a kind. QQQJA has a stronger first card, so it gets rank 5 and T55J5 gets rank 4.

Now, you can determine the total winnings of this set of hands by adding up the result of multiplying each hand's bid with its rank (765 * 1 + 220 * 2 + 28 * 3 + 684 * 4 + 483 * 5). So the total winnings in this example are 6440.

Find the rank of every hand in your set. What are the total winnings?

--- Part Two ---

To make things a little more interesting, the Elf introduces one additional rule. Now, J cards are jokers - wildcards that can act like whatever card would make the hand the strongest type possible.

To balance this, J cards are now the weakest individual cards, weaker even than 2. The other cards stay in the same order: A, K, Q, T, 9, 8, 7, 6, 5, 4, 3, 2, J.

J cards can pretend to be whatever card is best for the purpose of determining hand type; for example, QJJQ2 is now considered four of a kind. However, for the purpose of breaking ties between two hands of the same type, J is always treated as J, not the card it's pretending to be: JKKK2 is weaker than QQQQ2 because J is weaker than Q.

Now, the above example goes very differently:

32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483

    32T3K is still the only one pair; it doesn't contain any jokers, so its strength doesn't increase.
    KK677 is now the only two pair, making it the second-weakest hand.
    T55J5, KTJJT, and QQQJA are now all four of a kind! T55J5 gets rank 3, QQQJA gets rank 4, and KTJJT gets rank 5.

With the new joker rule, the total winnings in this example are 5905.

Using the new joker rule, find the rank of every hand in your set. What are the new total winnings?

]#
import strutils
import std/sequtils
import std/sugar
import std/algorithm
import std/tables
import std/sets
import std/enumutils
import sugar
import unittest

type Card = enum
    TWO = "2", THREE = "3", FOUR = "4", FIVE = "5"
    SIX = "6", SEVEN = "7", EIGHT = "8", NINE = "9"
    T = "T", J = "J", Q = "Q", K = "K", A = "A"

type CardType = enum
    HIGH, ONE_PAIR, TWO_PAIR, THREE_OF_A_KIND,
    FULL_HOUSE, FOUR_OF_A_KIND, FIVE_OF_A_KIND


proc cardType1(f: Table[string, int]): CardType = 
    var cardValues = toHashSet(f.values.toSeq())
    var cardType = FIVE_OF_A_KIND
    if len(f) == 4:
        cardType = ONE_PAIR
    elif len(f) == 5:
        cardType = HIGH
    elif len(f) == 2:
        cardType = FOUR_OF_A_KIND
        if cardValues.contains(3):
            cardType = FULL_HOUSE    
    elif len(f) == 3:
        cardType = THREE_OF_A_KIND
        if cardValues.contains(2):
            cardType = TWO_PAIR
    return cardType

proc cmpCardEnum1(xCard: Card, yCard: Card): int =
    return cmp(xCard, yCard)

proc cardType2(f: Table[string, int]): CardType = 
    var cardValues = toHashSet(f.values.toSeq())
    var countJ = f.getOrDefault($J, 0)
    var cardType = FIVE_OF_A_KIND # no J
    if len(f) == 4:
        cardType = ONE_PAIR
        if countJ > 0:
            cardType = THREE_OF_A_KIND
    elif len(f) == 5: # no J
        cardType = HIGH
        if countJ > 0:
            cardType = ONE_PAIR
    elif len(f) == 2:
        cardType = FOUR_OF_A_KIND
        if cardValues.contains(3):
            cardType = FULL_HOUSE
            if countJ == 2 or countJ == 3: # JJJ--
                cardType = FIVE_OF_A_KIND
        elif countJ == 1 or countJ == 4:
            cardType = FIVE_OF_A_KIND
    elif len(f) == 3:
        cardType = THREE_OF_A_KIND
        if cardValues.contains(2):
            cardType = TWO_PAIR
            if countJ == 1:
                cardType = FULL_HOUSE
            elif countJ == 2:
                cardType = FOUR_OF_A_KIND
        elif countJ == 1 or countJ == 3:
            cardType = FOUR_OF_A_KIND
        elif countJ == 2:
            cardType = FIVE_OF_A_KIND
    return cardType

proc cardType2(cardStr: string): CardType = 
    var cards: seq[Card]
    var f = initTable[string, int]()
    for c in cardStr:
        cards.add(parseEnum[Card]($(c)))
        if $c notin f:
            f[$c] = 0
        f[$c] += 1
    return cardType2(f)

proc cardOrder(card: Card): int =
    if (card == J):
        return -1
    return symbolRank(card)

proc cmpCardEnum2(xCard: Card, yCard: Card): int =
    return cmp(cardOrder(xCard), cardOrder(yCard))

proc cardTypeRank(cardStr: string, cardTypeFn: (Table[string, int]) -> CardType): (CardType, seq[Card]) =
    var cards: seq[Card]
    var f = initTable[string, int]()
    for c in cardStr:
        cards.add(parseEnum[Card]($(c)))
        if $c notin f:
            f[$c] = 0
        f[$c] += 1
    return (cardTypeFn(f), cards)

proc cardTypeRank1(cardStr: string): (CardType, seq[Card]) =
    cardTypeRank(cardStr, cardType1)


proc stringToCards(cardStr: string): seq[Card] =
    var cards: seq[Card]
    for c in cardStr:
        cards.add(parseEnum[Card]($(c)))
    return cards

proc cmpCardEnums(xCardEnums: seq[Card], yCardEnums: seq[Card], cmpCardEnumFn: (Card, Card) -> int): int =
    var result = 0
    for i in 0..len(xCardEnums):
        result = cmpCardEnumFn(xCardEnums[i], yCardEnums[i])
        if result != 0:
            break
    return result
    
proc sortHands(hands: var openArray[string],
               cardTypeFn: (Table[string, int]) -> CardType,
               cmpCardEnumFn: (Card, Card) -> int): void =
    hands.sort do (x, y: string) -> int:
        let
            (xCardType, xCardEnums) = cardTypeRank(x, cardTypeFn)
            (yCardType, yCardEnums) = cardTypeRank(y, cardTypeFn)
        var result = cmp(symbolRank(xCardType), symbolRank(yCardType))
        if result == 0:
            return cmpCardEnums(xCardEnums, yCardEnums, cmpCardEnumFn)
        return result 


proc handsBids(fileName: string): (seq[string], Table[string, int]) =
    var hands : seq[string]
    var bids = initTable[string, int]()
    for line in lines(fileName):
        let splits = line.splitWhitespace()
        let hand = splits[0] 
        hands.add(hand)
        bids[hand] = parseInt(splits[1])
    
    return (hands, bids)

proc totalWinnings(hands: seq[string], bids: Table[string, int]): int =
    var total = 0
    for i in 0 .. len(hands) - 1:
        total += bids[hands[i]] * (i + 1)
    return total

proc day7(fileName: string,
          cardTypeFn: (Table[string, int]) -> CardType,
          cmpCardEnumFn: (Card, Card) -> int): int =
    var (hands, bids) = handsBids(fileName)
    sortHands(hands, cardTypeFn, cmpCardEnumFn)
    return totalWinnings(hands, bids)

proc day71*(fileName: string): int =
    return day7(fileName, cardType1, cmpCardEnum1)

proc day72*(fileName: string): int =
    return day7(fileName, cardType2, cmpCardEnum2)




test "some test":
    check cardTypeRank1("AAAAA") == (FIVE_OF_A_KIND, @[A, A, A, A, A])
    check cardTypeRank1("AA8AA") == (FOUR_OF_A_KIND, @[A, A, EIGHT, A, A])
    check cardTypeRank1("23332") == (FULL_HOUSE, @[TWO, THREE, THREE, THREE, TWO])
    check cardTypeRank1("TTT98") == (THREE_OF_A_KIND, @[T, T, T, NINE, EIGHT])
    check cardTypeRank1("23432") == (TWO_PAIR, @[TWO, THREE, FOUR, THREE, TWO])
    check cardTypeRank1("A23A4") == (ONE_PAIR, @[A, TWO, THREE, A, FOUR])
    check cardTypeRank1("23456") == (HIGH, @[TWO, THREE, FOUR, FIVE, SIX])

    # check cmpCardEnums(@[A, A, A, A, A], @[A, A, EIGHT, A, A], cmpCardEnum1) == 1
    # check cmpCardEnums("AAAAA", "AA8AA", cmpCardEnum1) == 1
    # check cmpCardEnum1("33332", "2AAAA") == 1 
    # check cmpCardEnum1("77788", "77888") == -1

    check cmpCardEnum1(A, A) == 0
    check cmpCardEnum1(A, J) == 1
    check cmpCardEnum1(J, A) == -1

    check cmpCardEnum2(A, A) == 0
    check cmpCardEnum2(A, J) == 1
    check cmpCardEnum2(J, A) == -1
    check cmpCardEnum2(J, J) == 0
    check cmpCardEnum2(J, TWO) == -1
    check cmpCardEnum2(J, EIGHT) == -1
    check cmpCardEnum2(EIGHT, J) == 1

    check cardType2("AAAAA") == FIVE_OF_A_KIND
    check cardType2("AAAAJ") == FIVE_OF_A_KIND
    check cardType2("AAAJJ") == FIVE_OF_A_KIND
    check cardType2("AAJJJ") == FIVE_OF_A_KIND
    check cardType2("AJJJJ") == FIVE_OF_A_KIND
    check cardType2("JJJJJ") == FIVE_OF_A_KIND
    check cardType2("32T3K") == ONE_PAIR
    check cardType2("T55J5") == FOUR_OF_A_KIND
    check cardType2("KTJJT") == FOUR_OF_A_KIND
    check cardType2("QQQJA") == FOUR_OF_A_KIND

    check cardType2("23456") == HIGH
    check cardType2("AAJ56") == THREE_OF_A_KIND
    check cardType2("AAJJ6") == FOUR_OF_A_KIND
    check cardType2("AAQQJ") == FULL_HOUSE
    check cardType2("AJQQQ") == FOUR_OF_A_KIND
    check cardType2("AJ867") == ONE_PAIR
    
# test1

Create a deck of playing poker cards (remove ghost cards)

There are two players A and B

Randomly shuffle the playing poker cards and distribute them equally between A and B

Game operation process:

At the beginning of the game, any one player will play a card first, 
and the other player will play later (one card at a time, 
can't pick cards，play in order), and the operation is repeated,cards from A and B are added to list C
Game process and judgment of winning and losing:
The cards played by both sides are arranged in the order of playing poker cards.
 When the value of the cards played by one side (regardless of suits) exists in the list C, 
 then these two cards and all the cards between them belong to the player who played the last card 
 (all cards are added to the back of the winner),
After the player wins the card, winner continue to play a card to play, 
and repeat the above operations until the number of cards in one side is 0

Notice：

1 At the end, you need to output the winning and List C poker cards

2 The final output card needs to know the suit


# test2

Simple server test

Two main points of investigation:

1 Can complete the game loop logic
2 socket or websocket connection

Language golang

How to play the game:
1 server rolls one dice (123 is small,456 is big)

2 Player A guesses the small or big of the dice thrown by the server

3 Tell Player A guessed right or wrong

4 Repeat the above gameplay to start the next round of the game

Pay attention to the functions that need to be implemented

1 Need to implement A to connect to the server through socket or websocket (simulate sending and receiving packets)

2 A does not need to display the performance, just print the result.(Edited)
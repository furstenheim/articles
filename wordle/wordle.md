# Learning from the Mistakes of others: beating Wordle in just three Moves

![](./solar.png)

Wordle is all the rage, everybody is playing against it, and everybody has come up with their best strategy and their best starting word.

Some people prefer to focus on the problem from the very start. If they try "ARISE" and the word start with the letter "A", they will focus on words that start with "A" until they find it. I like to think of these as the diggers.

Other people prefer to spend some turns discarding letters before jumping into the problem. They might try "ARISE", then "THUMB", even a third independent word, before they start considering the information they got. I like to think of these as the explorers.

There is tension among the diggers and the explorers. Both groups claiming that the other doesn't know how to play.

However, it is the starting word the most contended topic. Everybody has their choice, and they will fight for it. It can be that they want to know as many vowels, and they go for "OCEAN". Maybe, they know that vowels will come later and want to discard as many consonants as possible, so they use "MYTHS". Or they consider both first and second words and find a nice combination, like "ARISE" and "THUMB". The geekier ones look at frequency of letters to find what should be the best one. Among the five-letter words, O, R, A, T and E are the most common letters. So they choose to play "ORATE" as first word.  It's been even quantified, that you can win [99%](https://english.elpais.com/science-tech/2022-01-14/a-spanish-data-scientists-strategy-to-win-99-of-the-time-at-wordle.html) of the times just by playing "ORATE" first.

All those strategies, have a common hidden assumption that is limiting them. They only focus on the game at hand. They do not take into account external information. How about we try to get information from other people's games.

## Hidden clues in shared games

Data is all around us. We are constantly producing data consciously and unconsciously. It might be when we share stories on Instagram, or just when we pay with our credit card. In a very famous story, Target found out that [a customer was pregnant before the father did](https://www.forbes.com/sites/kashmirhill/2012/02/16/how-target-figured-out-a-teen-girl-was-pregnant-before-her-father-did/?sh=4cc470496668).

It's become so common to share, that sometimes we do not realize how much we are sharing. For example, people post images of their boarding pass, without realizing that there is a lot of information [hidden there](https://www.flight-delayed.co.uk/blog/2019/03/15/why-you-shouldnt-share-a-picture-of-your-boarding-pass).

Can we actually use this mindset to beat Wordle? For the people playing it, the following shape should be familiar

![](shared-clues.png)

That's what you see when you share your game. It hides the letters and it "just" shows the matches against the solution. However, there is a lot of information in that picture. It's telling us that there is another two words that share all the four letters. In this case the word was "solar", so I'm guessing my friend played "sonar", and maybe "sofar". Even if we didn't know the solution, we can discard a lot of words. For example, "yeast" (my most favourite start word) cannot be the solution. Neither of "yebst", nor "yecst", "yedst" ... is a word.

## Words have signatures

In the previous example with solar, we only considered two clues, but we can do better. We can use all the hints!! The case of "YEAST" was very obvious, but there might be subtler cases. For example, "ALPHA", shared the same four letters with "ALOHA". However, "ALPHA" could not be a solution, because all of the other options "ALAHA", "ALBHA", "ALCHA" ... are not words.

What we need to do is to match the possible solution against **all** possible English words. From the first one "AAHED" (to exclaim in amazement, joy, or surprise) to the last one "ZUNIS" (A member of an American Indian people of western New Mexico). This will give as a "signature". In the previous example, we'll know that "YEAST" has no hint that gives 🟩🟩⬜🟩🟩, and "ALPHA" has at most one.

Once we have computed this signature, we just compare it with the shared games. If the signature does not match, we might as well discard the word. Of course, it's impossible to do so by head, so we can write a small computer program.

Once we get the shared responses from our contacts, compare the signatures and discard the impossible ones. In the previous case of solar with just 6 hints, we reduce it down to 186 possible words!!

![](all-clues.png)

## Greedy algorithm
With 186 words left, the game is much more manageable. But we still want to beat it in 3 turns. We use a small programming trick call "greedy algorithm".

We consider all the possible English words and see how they would fare against the remaining ones. It can be that some give the same combination of colours, whereas others don't.

We pick the word that promises the maximum reduction. In this case it is "ARIEL". If the response had been all Gray, we would be left with "BONUS", "BOOTY", "COUCH", "MUNCH", "MUSKY" and "SOOTY",  However, we are lucky and "SOLAR" is the only solution in this case.

## Try it on your own
It requires a bit of technical skills, but you can try it on your own. I've made the code publicly available in [Github](https://github.com/furstenheim/wordle), with some instructions on how to use it.

![](./terminal.png)

Happy gaming!!






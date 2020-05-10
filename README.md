# Keep Talking and Nobody (U)Explodes

This is a small application originally intented to make the password game in the above game easier.
Since then multiple module solvers have been added.

## Password
To get to the password game when it asks which modules you wish to play, enter "words" or "word" (without quotes). 
When you get to this point it will ask you for each letter possible in the first position. Enter the letters as a word, unbroken by a space.
This will continue through until there is one word left, if a mistake is made and no words are possible, you'll be asked if you wish to reset or not bother.

## Wires
For this game you need to enter "wire" or "wires" (without quotes) at the prompt. As you see the the wires, enter the wires from top to bottom. The letters representing each colour are as follows:
* r: red 
* w: white
* y: yellow
* b: blue
* k: black (k as in key value for ink)

It may ask you further questions, such as for the serial number. When asked, enter the serial number exactly as it is on the 'bomb'. For wires only the last digit is important, but other information from the serial number is stored for possible use later.

## Button
For the button module enter "button" at the prompt. Then it should be a matter of following the questions. Make sure all spellings are correct when entering the information. When asked for the button text you can enter the first letter of the text, e.g. "abort" you can jsut enter "a". You cannot currently do this for the colours because I'm a silly :3

## Complex wires
At the prompt when asked type in "complex wires". Then it will ask you for the combination. This is where it gets kinda convoluted.
Imagine for each wire you need to give the software details for each wire. The colour, if there is an led and if there is a star.
For the colour type "w", "r", "b" or "p" for white, red, blue and purple (blue and red). Then if there is a star type "s", if not type "o". Then for an led type "l" if there is and "o" if there is not. So a white wire with no led or star would be "woo", a purple wire with a star but no LED would be "pso". And just continue for each wire, do **not** seperate with spaces.
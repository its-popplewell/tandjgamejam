# tandjgamejam
You need to run `go get github.com/gen2brain/raylib-go/raylib` to download the Raylib dependency


BONK & FIGHT
SHOP SHTUFF
COMBINE AND META CODE

an auto battling game where dinosaurs with helmets run into eachother and fight. 

you can buy helmets and modifiers
and then you can build different builds of helmets and stuff 

player movements player health

fighting:
- running
- bonking
- helmet
  - hp/dmg/# hits
  - helmet has modifier slots
  - modifiers can take up more than one slot


todo: 
- randomly generate helmets and modifiers in shop
- update movement
- add placeholder buttons for state logic: shop -> battle -> shop (if win battle) loop and if lose battle restart. 

fight state
when fight opens you click a start button to start the fight.

if you win you get an option to either fight again or go to shop

in shop:
display shop items (5 total things in shop either helmets or modifiers) and you can buy them with the money you have (and you see the amount that you have)

you can also see your own inventory of helmets and modifiers

options to buy, sell, equip to dino,

main screen draws stuff based on state

general raylib screenlogic should live outside of the individual modules and instead be in the main.

and then we should be able to call functions to render different screen based on the state we are in. and those functions aren't in main. lets say they aren't in main.

on fight screen:
first nothing happens
press start
battle happens
after battle: 
if dead: to end state
if alive: buttons to fight again or go to shop

so to detach from dino:
x/y position of dino might be stored in the dino struct but the drawing would be in the battle. and now since we have the player, the player dino should come from the player not randomly made. and the changes to the players dino should be reflected in the battle.

model note:
- player owns persistent dino data like health, max health, damage, block, helmet, and modifiers
- battle owns fight-only dino data like x/y position, direction, speed, and knockback timers
- battle should copy the player's persistent dino when the fight starts, then commit only intended persistent changes back when the fight ends
- for now, damage taken persists by copying battle health back to the player's dino immediately when the fight ends
- knockback is shelved for now; if helmets/modifiers affect it later, they should change a computed combat stat like knockback/recoil, while the actual knockback timer stays battle-only
- because hits are collision-based, more knockback may slow both dinos' next attack by increasing the time before they collide again, so it may be a tempo tradeoff instead of a simple upgrade


# JACK END OF DAY (it is morning rip):
Got lots of good stuffs done.
I have not switched as much to float32s as I would like, but I am starting to.
I am building a UI package within our package -- I should ask Tushar to start using.
## I have gotten a good bit into Shop UI:
- [x] Draw out buying
- [ ] Draw out equipping
- [x] Build round rectangles and whatnot
- [ ] Choose and download a font (or choose to keep default)
- [ ] Build mockup of buying
- [ ] Build mockup of equipping
- [ ] add colors to config

## What I did for refactoring:
- I split up updating functions and draw functions -- draw should avoid changing a state (except with buttons)
- I added to our UI utils
- I simplified our main loop and built out separate update and draw functions which manage all of the state nonsense
- **||WIP||** I created a config.go file (thanks AI) which has constants for all the "magic numbers" we have (this will make play testing and tuning a breeze)
- I added types for the various states and substates of the state machine instead of consts (better type checking)
- Standardized some names and whatnot
- [ ] Add more rl types (vectors and shtuff)
- [ ] Add errors for functions which can fail (easier debugging down the line)

## Planning to do:
- [ ] Add types for helmets (come up with set of elements)
- [ ] Weigh helmet attributes based type and a level (generate cost based on that)
- [ ] Create a pool of modifiers to pull from
- [ ] Add path/movement stuff (queue of dinos?)
    - Battleheart like, where certain encounters can happen which will unlock entire upgrade paths?
    - Shards and a build system? Just modifiers to pick up? 
    - I want negative powerups with interesting synergies (hard for planning)
- [ ] ART ART ART
- [ ] do more design stuff in general, have a theme
- [ ] scale enemy generation with number of consecutive fights
- [ ] add drops and gold from fights which also scale ^
- [ ] add a player level (how should this work)
    - [ ] scale the generation of stuff / unlock certain stuff with level

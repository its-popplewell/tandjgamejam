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




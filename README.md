# About Go-CSPM / Supporting
- Written in Go https://golang.org/
- Support: https://discord.gg/CwqbHt5
- Want to buy me a coffee? https://www.paypal.me/rkhous/5

# Commands
## Spawns
- .spawn <pokemon> <lat,lon> <optional_description>
- .spawn sandshrew 34.123456,-118.123456
- .spawn sandshrew 34.123456,-118.123456 100IV L30
- Please keep in mind, there are to be NO spaces in between the lat,lon
## Quests
- Please keep in mind, each argument is supposed to be in between **it's own quotes, in the correct order!**
- .quest "pokestop" "quest" "quest_reward"
- .quest "undersea fish mural" "catch 5 fire types" "3 hyper potions"
- The bot will find closely matching names as well, so no need to type them out completey (for pokestop name)
- Example: if "That One Red Lion" was a pokestop name, you can do:
- .quest "red lion" "catch 3 water types" "5 pokeballs"
## Searching Stops
- To search for pokestops in the list of stops, use the following command:
- .search <pokestop>
- Example: _.search chi dynasty dragons_
- Please keep in mind: you **do not** need to search for the entire string, for example, if "Mountain View Cemetery" was a stop in your area, you can simply run:
- .search mountain view
## Admin Tools
- The following are limited to admins only, as chosen by the bot owner
- **.resetquests** _or_ **.rq** to clear the quests channel at midnight/quests change.
- **.setprefix** to set the prefix, by default the prefix is **.**

# Setup
- Coming soon. Easy to figure out, otherwise. 

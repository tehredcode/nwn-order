# nwn-order
Go program written to enhance nwnxee persistent worlds with an external go program.

![Image of Order](https://github.com/Urothis/nwn-order/blob/master/docs/Screenshot.png)

> required software
> - Docker
> - Docker compose
>
> Module requirements
> - [Neverwinter Nights enhanced edition]https://www.beamdog.com/products/neverwinter-nights-enhanced-edition
> - [NWNXEE]https://nwnx.io/

## Setup 
For a starting docker compose template
> https://github.com/Urothis/nwnxee-template

# Docker Compose
```
version: '3'
services:
  nwn-order:
    hostname: nwn-order
    image: golang:1.11.1-stretch
    env_file: ${PWD-.}/config/nwnorder.env
    volumes:
      - ${PWD-.}/order:/go/order
    working_dir: /go/order
    command: bash -c "go mod download && go run *.go"
    ports:
      - "5750:5750/tcp"
```
 
If you are still running into issues feel free to stop by the discord

> https://discord.gg/r6wuUdx

## Project goals
This project started as a way for me to setup a full ci deployment.
This has evolved into a fun way to see what we can do.
Community involvement is appreciated.   
    
>### Completed
- ~~Extra heartbeats~~ 
- ~~UUID generation~~ 

>### In Progress
- Full CI
- Cat facts
- Patreon integration
- Discord bot integration
- Whatever else the community suggests


## NWSCRIPT
### UUID
This function should be passed the player objet only.
It will return the assigned UUID to the player.

Internal scripts attach this uuid to the player tag.
> OrderGetUUIDPlayer(oPC)

Should return the players unique ID

Example return:
> 6fc7438a87d42b2dec552b4fb81b75a2

### Heartbeat
Heartbeat functionality can be enabled via config/nworder.env

>NWN_ORDER_HB_VERBOSE=

Setting to true will disply more logs for heartbeat

>NWN_ORDER_HB_ONE_MINUTE=true

Tickers will need to be enabled or disabled depending on your needs.

Default actions for heartbeat tickers are defined in order_heartbeat.nss

### CI/Github
This requires alittle bit of setup to function.

Requirements:
have a webhook setup for the repo you want to recieve alerts from.
https://developer.github.com/webhooks/creating/

When the docker-compose does go up, order will spit out an external facing IP and port. 

You will need to go into gitub and enable the webhook.

Example:
![Image of Github](https://github.com/Urothis/nwn-order/blob/master/docs/Github_Screenshot.png)

So when you deliver a webhook, order will accept the webhook and trigger the 
"OrderGithub();"
function inside of order_github.nss

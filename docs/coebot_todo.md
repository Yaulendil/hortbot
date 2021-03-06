# TODO

## The basics

- [x] Single process mode (usable but not final)
    - MVP
    - Need configuration setup
- [ ] Multi process mode
    - [x] Messaging over NSQ
    - [x] Redis for k/v, dedupe, expiration, cooldowns
    - [ ] No-downtime upgrades
    - [x] Rate limiting
        - [x] Needs to check userstate/badges/etc to "unlock" a faster rate.
- [x] Website (see below)
- [x] JSON configuration transferer
    - [x] CoeBot/CoeBot.tv config/DB to models
    - [x] Model import and refresh

## CoeBot features

Some features are not currently planned to be ported and have been crossed out.

### Differences

- The "regular" role has essentially been removed. All subs are "regulars", and all "regulars" are subs.
This means that the `subsRegsMinusLinks` option doesn't make much sense, and has been removed.
Instead, there's `subsMayLink` which when true means subs may link in chat.
- Subs are allowed to post links by default, as regulars would be. Regulars are always allowed
to post links.
- Command lists (`!list`) may be run via repeats and schedules.
- The special `-1`/`admin` mode has been removed; it can be reproduced using another mode.
- When quotes are deleted, quotes after it do not have their numbers changed.
The new `!quote compact <num>` command can be used to compact the quote list and remove
any holes left by deleted quotes. (The same applies to autoreplies.)
- HortBot keeps track of USERSTATE messages so that it always know how fast it can send messages
depending on its mod/vip status in a particular channel.
- Commands that accept usernames (`permit`, `+b`, etc) will strip leading `@` to make it easy to
autocomplete usernames in the chat input box.
- Doing a `!join` when your username changes will cause the bot to handle the name change and rejoin.

### Functionality

- [x] Link detection and moderation
    - [x] YouTube API access to grab URLs
- [x] Twitch API stuff
    - [x] For many things, this requires OAuth, which means the web service.

### Commands

- [x] General commands
    - [x] Join/part (in bot's channel)
    - [x] Part (in other channels); waiting until this can have "confirm" functionality
    - [x] ~~Topic~~
    - [x] Viewers (*Twitch API*)
    - [x] Chatters (*Twitch API*)
    - [x] Uptime (*Twitch API*)
    - [x] LastFM
        - [x] `music`
        - [x] Song link
    - [x] Bot help (*Website*)
    - [x] ~~Commercial~~
    - [x] Game (*Twitch API*)
        - [x] Get game
        - [x] Set game
    - [x] Status (*Twitch API*)
        - [x] Get status
        - [x] Set status
    - [x] statusgame/steamgame (*Steam API*)
    - [x] ~~XBox game~~
    - [x] Follow me (*Twitch API*)
        - Should this be automatic?
    - [x] ~~Viewer stats~~ (Use twitchtracker or sullygnome)
    - [x] ~~Punish stats~~ Twitch provides some element of this already; could bring back if wanted.
    - [x] What should I play (*Steam API*)
    - [x] Google
    - [x] ~~Wiki~~
    - [x] Is live (*Twitch API*)
    - [x] Is here (*Twitch API*)
- [x] Custom Commands
    - [x] Cooldowns
    - [x] Add
    - [x] Delete
    - [x] Restrict
    - [x] Rename (undocumented)
    - [x] Editor (undocumented)
    - [x] Clone
    - [x] Automatic restriction based on known actions
    - [x] Link to list of commands (*Website*)
- [x] Repeats
    - [x] Add
    - [x] Delete
    - [x] On/off
    - [x] List
    - [x] Actually execute the commands rather than printing them verbatim.
- [x] Schedule
    - [x] Add
    - [x] Delete
    - [x] On/off
    - [x] List
    - [x] Actually execute the commands rather than printing them verbatim.
- [x] Auto-replies
    - [x] Add
    - [x] Remove
    - [x] Edit
    - [x] List
    - [x] Actions inside of autoreply responses
    - [x] Similarly to quotes, emulate old list behavior
    - [x] Link to list of autoreplies (*Website*)
- [x] "Fun"
    - [x] ~~Throw~~ (Use a custom command.)
    - [x] Winner
    - [x] ~~Hug~~ (Use a custom command.)
    - [x] Conch/helix (Requires quotes)
    - [x] Urban
    - [x] ~~Me~~ (Use a custom command.)
    - [x] ~~Race~~
    - [x] XKCD
- [x] Random/roll
    - [x] ~~`regular`~~ Implement later on subs.
    - [x] Cooldowns
    - [x] Integer
    - [x] Dice
    - [x] Default
- [x] Quotes
    - [x] Add
    - [x] Delete
    - [x] Get
    - [x] Get index
    - [x] Random
    - [x] Search
    - [x] CoeBot would delete a quote and shift quotes after it backwards; HortBot doesn't. There needs to be a command to emulate this behavior.
    - [x] Link to quotes page on website (*Website*)
- [x] ~~Poll~~
- [x] ~~Giveaways~~
- [x] Raffles
    - [x] Raffle
    - [x] Enable/disable
    - [x] Reset
    - [x] Count
    - [x] Winner
- [x] ~~Highlights~~
    - Predates Twitch clips; may still consider reviving.
- [x] ~~Binding of Isaac: Rebirth~~
- [x] Moderation
    - [x] Slow mode on/of
    - [x] Sub mode on/of
    - [x] Ban
    - [x] Timeout
    - [x] Purge
    - [x] Link permit (needs detection)
        - [x] `allow` form
    - [x] Clear chat
- [x] Ignores
    - [x] Add/delete
    - [x] List
- [x] ~~Raids~~ Twitch does this better now
    - May partially implement. Requires the bot account in use to be a channel editor.
    - [x] `host` - *Non-functional due to Twitch API changes*
    - [x] `unhost` - *Non-functional due to Twitch API changes*
- [x] Settings
    - Lots of settings here related to other items.
    - [x] ~~`topic`~~
    - [x] `parseYoutube`
    - [x] `shouldModerate`
    - [x] `roll`
        - [x] ~~`timeoutoncriticalfail`~~
        - [x] `default`
        - [x] `cooldown`
        - [x] `userlevel`
    - [x] ~~`songrequest`~~
    - [x] `extralifeid`
    - [x] `urban`
    - [x] ~~`gamertag`~~
    - [x] `bullet`
    - [x] `subsRegsMinusLinks` AKA `subsMayLink` (see "differences" above)
    - [x] `cooldown`
    - [x] ~~`throw~~
    - [x] `lastfm`
    - [x] `steam` (*Steam API*)
    - [x] `mode`
    - [x] ~~`commerciallength`~~
    - [x] `tweet`
    - [x] `prefix`
    - [x] ~~`emoteset`~~ Doesn't seem to be particularly useful with emotes in tags.
    - [ ] `subscriberalerts` (Been broken for a while, low priority)
    - [ ] `resubalerts` (Been broken for a while, low priority)
- [x] User levels
    - [x] Add/remove reg/mod/owner
    - [x] List
- [x] Filters
    - [x] Warnings, then timeouts.
    - [x] Me
    - [x] Links
        - [x] Permitted domains (and more)
    - [x] Capitals
    - [x] Banned phrases
    - [x] Symbols
    - [x] Max length
    - [x] Emotes
    - [x] Options
        - [x] `on`/`off`
        - [x] `status`
        - [x] `me`
        - [x] `enablewarnings`
        - [x] `timeoutduration`
        - [x] `displaywarnings`
        - [x] `messagelength`
        - [x] `links`
        - [x] `pd`
        - [x] `banphrase`
        - [x] `caps`
        - [x] `emotes`
        - [x] `symbols`
- [x] Administration
    - [x] ~~`verboseLogging`~~
    - [x] `imp` - Moved to `!admin imp`
    - [x] `+whatprefix`
    - [x] ~~`altsend`~~ There is no "alt" connection, so this doesn't have a meaning (but may if send priorities/rate limiting is added later).
    - [x] ~~`disconnect`~~
    - [x] `admin`
        - [x] `channels`
        - [x] ~~`join`~~
        - [x] ~~`part`~~
        - [x] `block`
        - [x] `unblock`
        - [x] ~~`reconnect`~~
        - [x] ~~`reload`~~
        - [x] `color`
        - [x] ~~`loadfilter`~~ Maybe reintroduce later for another purpose.
        - [x] `spam`
        - [x] `#`
        - [x] ~~`trimchannels`~~
- [x] Variables
    - [x] Set
    - [x] Delete
    - [x] Get
    - [x] Increment
    - [x] Decrement
    - [x] Actions / string replacements (see below)
- [x] Lists
    - [x] Add
    - [x] Delete
    - [x] Restrict
    - [x] Add item
    - [x] Delete item
    - [x] Get
    - [x] Random
    - [x] Actually execute the commands rather than printing them verbatim.
- [x] Misc undocumented stuff
    - [x] ~~Weird testing commands~~ (Twitch resubs are no longer sent in PRIVMSGs)
    - [x] Steam game
    - [x] Extra Life stuff
    - [x] ~~`strawpoll`~~
    - [x] `channelID`
    - [x] ~~`whisper`~~
    - [x] ~~"wp"~~ (Sorry, go make another bot to do this...)
    - [x] ~~`properties`~~ (This endpoint has been removed.)
    - [x] ~~`songrequest`~~
    - [x] ~~`sendUpdate`~~
    - [x] Custom commands from another channel (`!#<user>/<command> ...`)
    - [x] ~~`modchan`~~ Use `!set mode`
    - [x] ~~`rejoin`~~
    - [x] `link`

### Actions (string replacements)

- [x] `(_GAME_)` (*Twitch API*)
- [x] `(_GAME_CLEAN_)` (*Twitch API*) - `(_GAME_)` but replace all non-alphanum with `-`
- [x] `(_STATUS_)` (*Twitch API*)
- [x] `(_VIEWERS_)` (*Twitch API*)
- [x] `(_CHATTERS_)` (*Twitch API*)
- [x] `(_STEAM_PROFILE_)` (*Steam API*)
- [x] `(_STEAM_GAME_)` (*Steam API*)
- [x] `(_STEAM_SERVER_)` (*Steam API*)
- [x] `(_STEAM_STORE_)` (*Steam API*)
- [x] `(_SONG_)`
- [x] `(_SONG_URL_)`
- [x] `(_LAST_SONG_)`
- [x] `(_BOT_HELP_)` (*Website*)
- [x] `(_USER_)`
- [x] `(_QUOTE_)`
- [x] ~~`(_COMMERCIAL_)`~~
- [x] `(_PARAMETER_)`
- [x] `(_PARAMETER_CAPS_)`
- [x] `(_NUMCHANNELS_)`
- [x] ~~`(_XBOX_GAME_)`~~
- [x] ~~`(_XBOX_PROGRESS_)`~~
- [x] ~~`(_XBOX_GAMERSCORE_)`~~
- [x] `(_ONLINE_CHECK_)` (*Twitch API*)
- [x] `(_SUBMODE_ON_)`
- [x] `(_SUBMODE_OFF_)`
- [x] `(_GAME_IS_<GAME>_)` (*Twitch API*)
- [x] `(_GAME_IS_NOT_<GAME>_)` (*Twitch API*)
- [x] `(_HOST_<CHANNEL>_)` - *Non-functional due to Twitch API changes*
- [x] `(_UNHOST_)` - *Non-functional due to Twitch API changes*
- [x] `(_RANDOM_<MIN>_<MAX>_)`
- [x] `(_RANDOM_INT_<MIN>_<MAX>_)`
- [x] `(_<COMMANDNAME>_COUNT_)`
- [x] `(_PURGE_)`
- [x] `(_TIMEOUT_)`
- [x] `(_BAN_)`
- [x] `(_SUBMODE_OFF_)`
- [x] `(_SUBMODE_OFF_)`
- [x] `(_REGULARS_ONLY_)`
- [x] `(_COMMAND_<NAME>_)` (for autoreplies)
- [ ] Sub message specific actions
- [x] `(_TWEET_URL_)`
- [x] `(_EXTRALIFE_AMOUNT_)`
- [x] `DATE`, `TIME`, `TIME24`, `DATETIME`, `DATETIME24`
- [x] `UNTIL`, `UNTILSHORT`, `UNTILLONG`
    - This has been slightly modified to allow "real" RFC3339 timestamps. Existing timestamps were
    done in not-quite-RFC3339 in the Eastern time zone. Also, dates more than ~290 years will no longer
    work (sorry).
- [x] VAR actions
- [x] `(_LIST_<name>_RANDOM_)`
- [x] `(_SILENT_)`
- [x] `(_CHANNEL_URL_)`
- [ ] `(_n_)` ("args")


### Website

- [x] Homepage
- [x] Channel list
    - [ ] Place live channels at the top (batch Twitch API)
- [x] Documentation
- [x] Login
- [ ] Join
- [x] About/FAQ
- [x] Channel
    - [x] Overview (links, simple info)
    - [x] Commands
    - [x] Quotes
    - [x] Lists
    - [x] Autoreplies
    - [x] Repeat/scheduled commands
    - [x] Regular list
    - [x] Chat rules
    - [x] Command formatting
- [ ] Online editing of everything, real login system
- [ ] Fun stuff, error pages and art

## New features

- [ ] A new command language (scripting!)
- [ ] GMod integration
- [x] More "add" subcommands for !command, to preset restrictions
- [ ] Extra commands for the new `/delete` feature (to replace purging)
- [ ] HowLongToBeat queries
- [ ] Get a link to game stores other than Steam
- [ ] List cloning
- [ ] Random stuff over subs.

### Actions

- [x] `(_DELETE_)` (new)

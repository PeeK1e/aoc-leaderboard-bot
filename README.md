# AOC Leaderboard bot
This is an Bot that scrapes a private Leaderboard and sends a notification through a webhook to a Discord channel. 

## Config
| CLI Flag       | Description                          | Environment Variable | Required | Default Value       |
|----------------|--------------------------------------|----------------------|----------|---------------------|
| `-u`           | URL of the leaderboard               | `AOC_URL`            | Yes      |                     |
| `-c`           | Cookie for authentication            | `AOC_COOKIE`         | Yes      |                     |
| `-l`           | Debug level (DEBUG/INFO/WARN/ERROR)  | `AOC_LOG`            | No       | INFO                |
| `-f`           | path of the leaderboard file         | `AOC_DB_PATH`        | No       | `./leaderboard.json`|
| `-w`           | Webhook URL to send star information | `AOC_WEBHOOK_URL`    | Yes      |                     |

## Where do I get the Info?

The cookie is a session cookie that you can extract from your browser after you logged in on the aoc website, there you can open the console and find the session cookie when accessing the leader board for example.

The Leaderboard URL will look something like this `https://adventofcode.com/2023/leaderboard/private/view/1234567`. Just copy it and append `.json` to it. (`https://adventofcode.com/2023/leaderboard/private/view/1234567.json`)

The rest should be self explanatory.
If you want to use a thread for the messages, append `?thread_id=<number>` to the discord webhook URL.

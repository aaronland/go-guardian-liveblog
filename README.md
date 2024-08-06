# go-guardian-liveblog

Go package for watching Guardian live blog events and reading them aloud using the operating system's text-to-speech APIs.

## Important

This is MacOS specific right now.

## Tools

```
$> make cli
go build -mod vendor -ldflags="-s -w" -o bin/follow cmd/follow/main.go
```

### follow

Parse one or more Guardian "live blog" URLs and read them aloud.

```
$> ./bin/follow -h
Parse one or more Guardian "live blog" URLs and read them aloud.
Usage:
	 ./bin/follow [options] url(N) url(N)
  -delay int
    	The number of seconds to wait before fetching new updates (default 30)
  -read-all
    	If true read all previous posts (written before following has begun)
  -verbose
    	Enable verbose (debug) logging.
```

For example:

```
$> ./bin/follow https://www.theguardian.com/football/live/2024/jul/05/spain-v-germany-euro-2024-quarter-final-live-score-updates

2024/07/05 11:42:50 INFO It’s finally all over in Stuttgart, where Spain advance to the semi-finals at the expense of the hosts courtesy of Mikel Merino’s late, late, late winner. Football, bloody hell.
2024/07/05 11:43:01 INFO ET30+6: The Spanish right-back walks for a second yellow after wrapping his hands around the neck of Musiala in a bid to bring him down. Free kick for Germany. Toni Kroos takes his last kick of the ball as a professional footballer and Germany are out …
2024/07/05 11:43:16 INFO ET30+4: Fulkrug heads just wide from seven or eight yards after a Muller cross came his way. That was the chance!!!
2024/07/05 11:43:24 INFO ET30+2: Kimmich has a shot from outside the area blocked. Ferran Torres breaks upfield but elects to try to score instead of running to the corner flag. Neuer launches the ball forward.

...and so on
```

The `follow` tool will keep a local cache of posts its already seen (and read) for the duration it is run.

Note: This tool is not very sophisticated and might miss some posts and should be updated to use the Guardian API (for example: https://api.nextgen.guardianapps.co.uk/football/api/match-nav/2024/07/05/5539/619.json?dcr=true&page=football%2Flive%2F2024%2Fjul%2F05%2Fportugal-v-france-euro-2024-quarter-final-live-score-updates).

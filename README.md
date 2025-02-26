# go-guardian-liveblog

Go package for watching Guardian live blog events and reading them aloud using the operating system's text-to-speech APIs.

## Important

This package has been superseded by the [aaronland/go-liveblog](https://github.com/aaronland/go-liveblog) package.

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

Or with multiple URLs:

```
$> ./bin/follow \
	https://www.theguardian.com/sport/live/2024/aug/06/paris-2024-olympics-day-11-live-updates-today-schedule-events-athletics-cycling-boxing \
	https://www.theguardian.com/sport/live/2024/aug/06/usa-v-germany-paris-olympics-womens-soccer-semi-final-latest-score

2024/08/06 12:01:02 INFO Paris 2024 Olympics day 11: Cole Hocker wins stunning 1500m gold ahead of GB’s Josh Kerr; women’s 200m final to follow – live
2024/08/06 12:01:13 INFO Kerr looks happy enough with silver. He’ll be disappointed, of course, and probably thought it was his with 30m to go. But he ran his race and things went as he planned, Ingebrigtsen perhaps spooked by the enormity of it all, but someone else was better on the day. And the times:
2024/08/06 12:01:30 INFO Hocker (USA) 3:2765
2024/08/06 12:01:34 INFO Kerr (GB) 3:27.79
2024/08/06 12:01:40 INFO Nuguse (USA) 3:27.80
2024/08/06 12:01:45 INFO Ingebrigtsen (Norway) 3:28.24
2024/08/06 12:01:51 INFO Paris 2024 Olympics day 11: Cole Hocker wins stunning 1500m gold ahead of GB’s Josh Kerr; women’s 200m final to follow – live
2024/08/06 12:02:02 INFO At the risk of sounding a bit too provincial, seeing a feature on all the idiotic trash talk between Ingebritsen and Kerr, then seeing Ingebritsen getting physical down the stretch, and then seeing Hocker and Nuguse run them down may be the most satisfying moment for the USA in the Olympics in decades.
2024/08/06 12:02:31 INFO USA 1-0 Germany: Olympic women’s soccer semi-final – as it happened
2024/08/06 12:02:37 INFO Tom Dart’s full match report is now live:
2024/08/06 12:03:01 INFO Paris 2024 Olympics day 11: Cole Hocker wins stunning 1500m gold ahead of GB’s Josh Kerr; women’s 200m final to follow – live

...and so on
```

The `follow` tool will keep a local cache of posts its already seen (and read) for the duration it is run.

Note: This tool is not very sophisticated and might miss some posts and should be updated to use the Guardian API (for example: https://api.nextgen.guardianapps.co.uk/football/api/match-nav/2024/07/05/5539/619.json?dcr=true&page=football%2Flive%2F2024%2Fjul%2F05%2Fportugal-v-france-euro-2024-quarter-final-live-score-updates).

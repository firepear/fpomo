# fpomo
Console visual timer app

```
usage: fpomo <OPTIONS> [MINUTES]

  --alert [CMD]
    The command to run when time is up. Value: 'afplay'

  --alert_tone [FILEPATH]
    Audio file to use as the alert tone when time is up
    Value: '/System/Library/Sounds/Glass.aiff'

  --count [NUM]
    How many times to run the alert cmd. Value: 3

  --forever [true | false]
    Play the alert tone until killed with C-c. Value: 'false'

  --mark [CHAR]
    Set the character used to flood the screen, representing time
    remaining. Value: '⋅'

  --sleep [SECS]
    How many seconds to sleep between screen updates. Value: 5

  --bg, --fg, --td [ANSI_COLOR_SPEC]
    Set the background, foreground, and time display colors. Values are
    '40', '37', '1;36'
```

- Unices other than Mac OS use `mpg123 -q` as the default alert
  command, and use an included audio file as the default tone
- Config file is `~/.fpomorc`; an example is in `assets`

## Notes

- Mac OS
  - Sleep is an interrupt, so a laptop with aggressive sleep settings
    should not be used as an unattended timer

## Credits

- [Default alert sound](https://pixabay.com/sound-effects/beep-beep-6151/)

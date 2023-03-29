# festoji

Seasonal festive emoji for the day 🎉

## Usage

Simply run the golang program:

```text
go run main.go
```

Or, compile it into a binary and run it:

```text
go build -o bin/festoji main.go
./bin/festoji
```

The current festive emoji will be printed to stdout.

You can also use it as a container:

```text
podman run --rm quay.io/lucarval/festoji:latest
```

## Customizing

Festoji comes with preset default rules. These can be overwritten by creating the file
`~/.festoji.yaml`. For example:

```yaml
---
# The default character to be used when no rules match.
default: 🐚

# If set to true, this configuration extends the default one. Rules are inserted after the
# default ones. If set to false, the default configuration is completely ignored.
extend: true

rules:
- name: Xmas
  emoji: 🎄
  # This rule will start matching 14 days prior to December 25th
  span: 14
  month: 12
  day: 25
- name: Thanksgiving
  emoji: 🦃
  # This rule will start matching 7 days prior to the fourth Thursday in November
  span: 7
  month: 11
  week: 4
  weekday: 4
```

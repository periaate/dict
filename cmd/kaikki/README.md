# kaikki

Files downloaded from [kaikki.org](https://kaikki.org/) do not seem to be valid JSON. This program provides a JSON formatter which outputs valid JSON.

The output must then be parsed with `jq` to a `dict.json` file with the command:\
`jq -c "[ .[] | {pos: .pos, word: .word, meaning: .senses[]? | .glosses[]? } ]"`

So all in all:
```
kaikki ./kaikki.org-dictionary-English.json formatted.json
cat formatted.json > jq -c "[ .[] | {pos: .pos, word: .word, meaning: .senses[]? | .glosses[]? } ]" > dict.json
```

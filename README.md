# gh-issue-list
issue list from json

## preparations

Add JSON files to the JSON folder by executing the following command, etc.

_command_
**Issues**

```
gh issue list -a @me -s closed --json title,url,createdAt,closedAt | jq --arg s "2021-01-01T00:00:00Z" 'map({title: .title, create: .createdAt, close: .closedAt}) | map(select((.close | strptime("%Y-%m-%dT%H:%M:%SZ") | mktime) as $d | ($s | strptime("%Y-%m-%dT%H:%M:%SZ") | mktime) as $c | $d >= $c))'
```
or
```
gh issue list -a @me -s closed --json title,url,createdAt,closedAt | jq --arg s "2021-01-01T00:00:00Z" 'map({title: .title, create: .createdAt, close: .closedAt}) | map(select((.close | strptime("%Y-%m-%dT%H:%M:%SZ") | mktime) as $d | ($s | strptime("%Y-%m-%dT%H:%M:%SZ") | mktime) as $c | $d >= $c))'
```

**PRs**
```
gh pr list -a @me -s closed --json title,url,createdAt,closedAt -L 100 | jq --arg s "2021-01-01T00:00:00Z" 'map({title: .title, create: .createdAt, close: .closedAt}) | map(select((.close | strptime("%Y-%m-%dT%H:%M:%SZ") | mktime) as $d | ($s | strptime("%Y-%m-%dT%H:%M:%SZ") | mktime) as $c | $d >= $c))'
```

## usage

- Run list issue

`go run main.go -f {$filepath}`

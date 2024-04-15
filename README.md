feedly_search
-------------

Utility for searching RSS feeds on [Feedly](https://feedly.com).

#### Install

```sh
go build -o feedly_search cmd/feedly_search/main.go
```

#### Example

![feedly_search utility](screenshot.png)

```bash
feedly_search -q "#cs" -l "en" -n 10 
```

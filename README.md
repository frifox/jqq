# With JQ
```
tail -f my.log | jq -c --raw-input --raw-output '. as $raw | try fromjson catch $raw'
```

# With alias (add to ~/.bashrc or ~/.zshrc)
```
alias jqq='jq -c --raw-input --raw-output ". as \$raw | try fromjson catch \$raw"'
tail -f my.log | jqq
```
    
# With JQQ (this repo)
```
go install github.com/frifox/jqq@latest
tail -f my.log | jqq
```

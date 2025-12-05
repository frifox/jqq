JQ but safely leaving non-json lines as-is

# Without JQQ
    tail -f my.log | jq -c --raw-input --raw-output '. as $raw | try fromjson catch $raw'

# With JQQ
    tail -f my.log | jqq
   
# Install
    go install github.com/frifox/jqq@latest

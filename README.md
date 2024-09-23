# JQQuiet
Default JQ with args:

    -c
    --raw-input
    --raw-output '. as $raw | try fromjson catch $raw'

# Install
    go install github.com/frifox/jqq@latest

# Use

    tail -f my.log | jqq
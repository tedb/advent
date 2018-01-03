#!/bin/bash

cookie="paste session cookie here"

for i in {1..25}; do
  curl -o day$i "http://adventofcode.com/2017/day/$i/input" -H 'Accept-Encoding: gzip, deflate' \
    -H 'Accept-Language: en-US,en;q=0.9' -H 'Upgrade-Insecure-Requests: 1' \
    -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36' \
    -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8' \
    -H 'Referer: http://adventofcode.com/2017/day/21' \
    -H "Cookie: session=$cookie" -H 'Connection: keep-alive' -H 'Cache-Control: max-age=0' --compressed
done

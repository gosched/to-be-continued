https://en.wikipedia.org/wiki/URL_shortening

短網址
短網址二維碼

低進制轉化為高進制時字符數會減少

1.
自增ID

auto ID (base 10)
short URL (base 62 | a ~ z, A ~ Z, 0 ~ 9)
original URL
date
ip
clicks

client -> dns -> ip -> request with short URL Code -> Server -> base62 to base10 -> base10 id -> Original URL -> server redirect

HTTP status
redirect
301
302
307

2.
消息摘要算法
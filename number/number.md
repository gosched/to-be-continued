# Numeral systems

- https://en.wikipedia.org/wiki/Numeral_system

- https://en.wikipedia.org/wiki/Binary_number
- https://en.wikipedia.org/wiki/Octal
- https://en.wikipedia.org/wiki/Decimal
- https://en.wikipedia.org/wiki/Hexadecimal

- https://en.wikipedia.org/wiki/Floating-point_arithmetic
- https://en.wikipedia.org/wiki/IEEE_754

- https://en.wikipedia.org/wiki/Bit_numbering
- https://en.wikipedia.org/wiki/Endianness

- http://cenalulu.github.io/linux/about-denormalized-float-number/
- https://segmentfault.com/a/1190000004090283

```
其他進制 轉換為 十進制

101.101 (base 2) -> 5.625 (base 10)
1 * 2^2 + 1 * 2^0 + 1 * 2^-1 + 1 * 2^-3

123.01 (base 16) -> 291.00390625 (base 10)
1 * 16^2 + 2 * 16^1 + 3 * 16^0 + 1 * 16^-2
```

```
十進制 轉換為 其他進制

分成 整數部分 小數部分

整數部分
不斷除以目標進制基底 直至商數為零 反向記錄餘數

小數部分
不斷乘以目標進制基底 直至乘積的小數為零 正向記錄乘積的整數值

20.125 (base 10) -> 10100.001 (base 2)

20 / 2 == 10 ... 0
10 / 2 == 05 ... 0
05 / 2 == 02 ... 1
02 / 2 == 01 ... 0
01 / 2 == 00 ... 1

0.125 * 2 == 0.25
0.25  * 2 == 0.5
0.5   * 2 == 1.0
```

```
十進制 -> 二進制 -> float32
十進制 -> 二進制 -> float64
十進制 -> 二進制 -> decimal
```

```
var f float64 = 0.1

0.1 (base 10) -> 0.000110011... (base 2) -> (float 64)

0.0625
0.03125
0.00390625
0.001953125
```

```
float32, 以 1 為前導
float64, 以 1 為前導
decimal, 不一定以 1 為前導
```
# tzc

Compare time zones.

## Usage

`tzc` takes one or two time zone locations corresponding to a file in the IANA
Time Zone database and compares the next 24 hours.

```sh
$ tzc America/Vancouver
Asia/Taipei           America/Vancouver
Thu Jan 7 12:00 CST   Wed Jan 6 20:00 PST
Thu Jan 7 13:00 CST   Wed Jan 6 21:00 PST
Thu Jan 7 14:00 CST   Wed Jan 6 22:00 PST
Thu Jan 7 15:00 CST   Wed Jan 6 23:00 PST
Thu Jan 7 16:00 CST   Thu Jan 7 00:00 PST
Thu Jan 7 17:00 CST   Thu Jan 7 01:00 PST
Thu Jan 7 18:00 CST   Thu Jan 7 02:00 PST
Thu Jan 7 19:00 CST   Thu Jan 7 03:00 PST
Thu Jan 7 20:00 CST   Thu Jan 7 04:00 PST
Thu Jan 7 21:00 CST   Thu Jan 7 05:00 PST
Thu Jan 7 22:00 CST   Thu Jan 7 06:00 PST
Thu Jan 7 23:00 CST   Thu Jan 7 07:00 PST
Fri Jan 8 00:00 CST   Thu Jan 7 08:00 PST
Fri Jan 8 01:00 CST   Thu Jan 7 09:00 PST
Fri Jan 8 02:00 CST   Thu Jan 7 10:00 PST
Fri Jan 8 03:00 CST   Thu Jan 7 11:00 PST
Fri Jan 8 04:00 CST   Thu Jan 7 12:00 PST
Fri Jan 8 05:00 CST   Thu Jan 7 13:00 PST
Fri Jan 8 06:00 CST   Thu Jan 7 14:00 PST
Fri Jan 8 07:00 CST   Thu Jan 7 15:00 PST
Fri Jan 8 08:00 CST   Thu Jan 7 16:00 PST
Fri Jan 8 09:00 CST   Thu Jan 7 17:00 PST
Fri Jan 8 10:00 CST   Thu Jan 7 18:00 PST
Fri Jan 8 11:00 CST   Thu Jan 7 19:00 PST
```

```sh
$ tzc America/Vancouver America/Toronto
America/Vancouver     America/Toronto
Wed Jan 6 20:00 PST   Wed Jan 6 23:00 EST
Wed Jan 6 21:00 PST   Thu Jan 7 00:00 EST
Wed Jan 6 22:00 PST   Thu Jan 7 01:00 EST
Wed Jan 6 23:00 PST   Thu Jan 7 02:00 EST
Thu Jan 7 00:00 PST   Thu Jan 7 03:00 EST
Thu Jan 7 01:00 PST   Thu Jan 7 04:00 EST
Thu Jan 7 02:00 PST   Thu Jan 7 05:00 EST
Thu Jan 7 03:00 PST   Thu Jan 7 06:00 EST
Thu Jan 7 04:00 PST   Thu Jan 7 07:00 EST
Thu Jan 7 05:00 PST   Thu Jan 7 08:00 EST
Thu Jan 7 06:00 PST   Thu Jan 7 09:00 EST
Thu Jan 7 07:00 PST   Thu Jan 7 10:00 EST
Thu Jan 7 08:00 PST   Thu Jan 7 11:00 EST
Thu Jan 7 09:00 PST   Thu Jan 7 12:00 EST
Thu Jan 7 10:00 PST   Thu Jan 7 13:00 EST
Thu Jan 7 11:00 PST   Thu Jan 7 14:00 EST
Thu Jan 7 12:00 PST   Thu Jan 7 15:00 EST
Thu Jan 7 13:00 PST   Thu Jan 7 16:00 EST
Thu Jan 7 14:00 PST   Thu Jan 7 17:00 EST
Thu Jan 7 15:00 PST   Thu Jan 7 18:00 EST
Thu Jan 7 16:00 PST   Thu Jan 7 19:00 EST
Thu Jan 7 17:00 PST   Thu Jan 7 20:00 EST
Thu Jan 7 18:00 PST   Thu Jan 7 21:00 EST
Thu Jan 7 19:00 PST   Thu Jan 7 22:00 EST
```

## Motivation

I often find myself trying to figure out an optimal time to connect with people
in different time zones. I prefer to live in the terminal and so this was born.

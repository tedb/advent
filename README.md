Advent of Code: Solutions
========================

[Advent of Code](http://adventofcode.com/) provides daily Christmas-themed
programming puzzles.  This repository contains my attempts at solutions to those
puzzles, primarily using the [Go](https://golang.org) programming language.
I am sharing my solutions with the community for enjoyment and in the spirit
of openness.  I have tried to follow the Go community's best practices around
unit testing and code structure, however I make no guarantee these solutions
are correct, generally applicable, or even particularly well written.  This was
written against Go 1.5.2, and it might work with previous versions but this is
untested.

Merry Christmas and a Happy New Year!

Ted

Usage
-----

Install [Go](https://golang.org/dl/) if you have not already

At a Linux command prompt:

    $ mkdir tedb-advent && cd tedb-advent
    $ export GOPATH=$PWD
    $ go get github.com/tedb/advent/...
    $ go test github.com/tedb/advent -v
    $ bin/advent -h
    $ bin/advent -f src/github.com/tedb/advent/data/advent1.txt 1a
    Advent 1a('()()(()()()(()()((()...') = 280, <nil>

License
-------
The MIT License (MIT)

Copyright (c) 2015 Edward (Ted) Behling

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

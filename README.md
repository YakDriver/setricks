# setricks
The Go language doesn't really have a "set" data structure _per se_ but surely you can accomplish set-type operations (union, intersection, except) using maps, right? This Go module provides these operations.

To play around with this module and the set operations, [install go](https://medium.com/@dirk.avery/install-git-and-go-on-macos-8c0503028814) and feed it some data. Make a file called `set1.data` with these contents (each line is preceeded by a space):

(You could also fork this and changed the regex that finds the data if you have other types of data.)

```
 red 42
 green 58
 blue 92
```

Make another file called `set2.data` with this contents:

```
 yellow 18
 blue 30
 orange 88
```

Then, run the Go (`$ go run main.go`) and you'll see the results:

```
Union: map[blue:30 green:58 orange:88 red:42 yellow:18]

Intersection: map[blue:30]

Except 1: map[green:58 red:42]

Except 2: map[orange:88 yellow:18]
```

As far as I know, these are the best ways of doing these set operations. If there are better ways, please share.

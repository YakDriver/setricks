# setricks
The Go language doesn't really have a "set" data structure _per se_ but surely you can accomplish set-type operations (union, intersection, except) using maps, right? This Go module provides these operations.

To play around with this module and the set operations, [install go](https://medium.com/@dirk.avery/install-git-and-go-on-macos-8c0503028814) and take a look at the [test file](https://github.com/YakDriver/setricks/blob/master/setricks_test.go) for examples of using the set operations:

(You could also fork and change the regex that finds the data if you have other types of data.)

The test data:

`testdata/set1.data`:

```
 red 42
 green 58
 blue 92
```

`testdata/set2.data`:

```
 yellow 18
 blue 30
 orange 88
```

```
Union: map[blue:30 green:58 orange:88 red:42 yellow:18]

Intersection: map[blue:30]

Except 1: map[green:58 red:42]

Except 2: map[orange:88 yellow:18]
```

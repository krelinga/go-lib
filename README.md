# krelinga's go-lib

## Releases

### v0.4.1

- Add `mac` library and `StayAwake()` function.

### v0.4.0

- total rewrite of `geom` library.  Biggest added feature is point & line tagging, to make it easy to follow
elements of a figure across various transformations.

### v0.3.8

- Add `HexagonTileOffset()` function.

### v0.3.7

- Add the `Rotate()` and `RotatePolygon()` functions.
- Add the `Midpoint()` function.

### v0.3.6

- Add the `Width()` and `Height()` methods to `geom.Polygon`.

### v0.3.5

- Add `geom` package, which includes:
    - The `Angle` interface, along with `Degrees()` and `Radians()` creation methods.
    - The `Polygon` type.
    - Various trigonometric functions implemented in terms of `Angle`.
    - The `Hexagon()` function.

### v0.3.4

- Add `kiter.SliceMap()` utility function.

### v0.3.3

- Add `kiter` package with many iterator tools.

### v0.3.2

- Add human-readable JSON marshalling & unmarshalling for `video.DirKind`.

### v0.3.1

- Add `DirKind` and `GetDirKind()` to `video`.

### v0.3.0

- Remove `ReadOnly()` method from `pipe`.  This was made obsolete by switching to
  some more-forgiving type constraints and using a helper function to implement `Merge()`.
- Add `ParDoFilter()` and `ParDoFilterErr()` to `pipe`.
- Internal simplifications to `video.BuildFileInfo()`.

### v0.2.1

- Add `pipe` tools for dealing with maps:
    - `KV` for representing keys & values of different types.
    - `GroupBy()` for combining values from identical keys.
    - `ToMapFunc()` for outputting `KV`s into a `map`.
- Add the `filesystem` module, which contains the `WalkAll()` method for walking over all files in the filesystem.
- Add the `video` package, which contains:
    - path conversion utilities.
    - A pipeline for reading the output of `filesystem.WalkAll()` and joining the various kinds of paths into a struct to capture their existence.

### v0.2.0

- Restructure things into a single `pipe` package, including some renaming simplifications.
- respect context cancelation in `Merge()`, `Parallel()`, `ParallelErr()`, and `TryWrite()`
- Add helper methods for consuming pipeline output that work well with `pipe.Wait()`, specifically:
    - `ToArrayFunc()`
    - `FirstFunc()`
    - `LastFunc()`
    - `DiscardFunc()`

### v0.1.0

Several (hopefully) useful utilities extracted from the `video-tool-box` repo:

- `chans.Merge()`
- `chans.Parallel()` and `chans.ParallelErr()`
- `chans.ReadOnly()`
- `chanstest.AssertElementsEventuallyMatch()` and `chanstest.AssertEventuallyEmpty`
- `routines.RunAndWait()`
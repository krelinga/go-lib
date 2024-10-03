# krelinga's go-lib

## Releases

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
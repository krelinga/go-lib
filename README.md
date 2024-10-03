# krelinga's go-lib

## Releases

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
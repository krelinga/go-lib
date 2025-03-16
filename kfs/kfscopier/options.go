package kfscopier

type Options struct {
	MaxParallelCopies int
	ChunkSize		 int
}

func (o *Options) setDefaults() {
	if o.MaxParallelCopies == 0 {
		o.MaxParallelCopies = DefaultMaxParallelCopies
	}
	if o.ChunkSize == 0 {
		o.ChunkSize = DefaultChunkSize
	}
}
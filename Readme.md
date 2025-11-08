# Go Compressor

## Features

✅ Lossless compression  
✅ Target: text and code files (so, redundancy-heavy data that benefits from dictionary + entropy coding)  
✅ Goal: balance speed and ratio  
✅ CLI tool (so we'll design a simple command-line interface)  
✅ Language: Go (great choice — fast, memory-safe, simple concurrency)  
✅ Open source / learning focus (so clarity > micro-optimization)

## Design Overview

We'll aim for a "mini-Deflate" style compressor:

```
Plain text → [LZ77 tokenization] → [Huffman encoding] → Compressed bytes
```

and the reverse for decompression.

We will evolve it later into a multithreaded engine, but start single-threaded.

## Project Structure

```
go-compressor/
│
├── cmd/
│   └── gocompress/           # CLI binary
│        └── main.go
│
├── internal/
│   ├── lz77/
│   │   ├── compressor.go
│   │   └── decompressor.go
│   ├── huffman/
│   │   ├── huffman.go
│   │   └── decode.go
│   ├── util/
│   │   └── bitwriter.go
│   └── engine.go             # glue: pipeline logic
│
└── go.mod
```

## File Format

A simple custom container format could be:

| Field         | Bytes    | Description |
|---------------|----------|-------------|
| Magic         | 4        | "GZC1" (Go Zip Compress v1) |
| Flags         | 1        | reserved |
| OriginalSize  | 4        | optional |
| BlockCount    | 2        | number of blocks |
| Block headers | variable | offsets & sizes |
| Compressed data| variable | concatenated |

## Roadmap

| Phase | Goals | Notes |
|-------|-------|-------|
| 1 | Implement naive LZ77 encode/decode | Verify decompression = input |
| 2 | Add static Huffman coding of literals | Improves ratio 20–40% |
| 3 | Add dynamic Huffman (per-block) | Better ratio for varied data |
| 4 | Add file header format & CLI | End-to-end working tool |
| 5 | Optimize: hash search, bit I/O, parallel blocks | Speed up by 5–10× |
| 6 | Add tests, benchmarking, docs | Ready to open-source |

## Go Learning Resources for Compression

Go standard library source `flate`, `gzip`, `zlib`, `lz4` — excellent references.

Open-source Go compressors to study:
- klauspost/compress (optimized zstd, snappy)
- pierrec/lz4
- dsnet/

Studying those will teach how production compressors manage buffers and bitstreams in Go.

## Testing Strategy

### Unit Tests
- LZ77 encode → decode equals original
- Huffman encode → decode equals original

### Integration Tests
- Files, large files, repetitive, random

### Benchmarks
- Use Go's `testing.B` for MB/s and compression ratio
- Compare vs gzip and zstd CLI

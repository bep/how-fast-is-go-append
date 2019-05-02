# how-fast-is-go-append

Just curious after reading this: https://dev.to/uilicious/javascript-array-push-is-945x-faster-than-array-concat-1oki

```bash
BenchmarkAppend-4   	    2000	          798µs/op
```

Should be 1253 ops / second or 3,3x faster than the fastest JS variant.

# Writing GC Friendly Code

Showing how you to use GC friendly data.

Use case: User store that returns user (string) from user ID (int).

- `users_map` uses a `map[int]string`
- `users_slice` use a `[]string` where the location on the slice is the ID
- `users_str` builds one big string with all user names and also have `indices []int` which hold the end index for each user

To bench a specific implementation, say "map", run:

    make map
    # In another terminal
    make stress

FIXME: I'm not sure that speedup is due to GC. If I run the benchmark in each implementation I see on my machine:
- map  : BenchmarkByID-12    	146152878	         8.117 ns/op	       0 B/op	       0 allocs/op
- slice: BenchmarkByID-12    	871609362	         1.379 ns/op	       0 B/op	       0 allocs/op
- str  : BenchmarkByID-12    	524216815	         2.309 ns/op	       0 B/op	       0 allocs/op

And for `make stress` I see:
- map  : Requests/sec:	139456.5474
- slice: Requests/sec:	131119.3114
- str  : Requests/sec:	141903.6378

---
## Trace

- Add import to `net/http/pprof`
- `go run ./httpd`
- `hey -z 10s http://localhost:8080/users/353`
- `curl -o trace.out http://localhost:8080/debug/pprof/trace?seconds=5`
- `go tool trace -http=:8888 trace.out`
    - trace.png
- `1,111,313ns` -> 1.1ms

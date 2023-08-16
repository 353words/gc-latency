# gc-latency

Showing how you to use GC friendly data.

Use case: User store that returns user (string) from user ID (int).

- `users_map` uses a `map[int]string`
- `users_slice` use a `[]string` where the location on the slice is the ID
- `users_str` builds one big string with all user names and also have `indices []int` which hold the end index for each user

Run all with `make`

Results on my machine:
```
=== map ===
go run .
allocated 1000000 users
running 100 GC cycles
total gc time: 4.528541076s (median: 46.151329ms)
=== slice ===
go run .
allocated 1000000 users
running 100 GC cycles
total gc time: 2.216892788s (median: 22.061863ms)
=== str ===
go run .
allocated 1000000 users
running 100 GC cycles
total gc time: 14.618716ms (median: 131.67µs)
```

See results with `GODEBUG=gctrace=1` in the [out](out) directory.
JSON files generated with [gogctrace](https://pkg.go.dev/github.com/tebeka/gctrace/cmd/gogctrace).

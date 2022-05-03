# gc-latency

Showing how you to use GC friendly data.

Use case: User store that returns user (string) from user ID (int).

- In `map` we use a `map[int]string`
- In `slice` we use a `[]string` where the location on the slice is the ID
- In `str` we build one big string with all user names and also have `indices []int` which hold the end index for each user

Run all with `make`

Results on my machine:
```
$ make
go run ./map
allocated 1000000 users
running 100 GC cycles
3.260994646s
go run ./slice
allocated 1000000 users
running 100 GC cycles
1.848372573s
go run ./str
allocated 1000000 users
running 100 GC cycles
10.260762ms
```

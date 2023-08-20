all:
	$(error please pick a target)

bench:
	go clean -cache -testcache
	go test -bench . -count 5



url = http://localhost:8080/users/353
stress:
	@ls -l users | awk '{print $$NF}' # current users implementation
	curl -f $(url)  # fail is server is not up
	hey -z 10s $(url)

ping:
	curl -q $(url)

# clean cache since it doesn't play nice with symlinks
clean:
	go clean -cache -testcache

map:
	ln -snf users_map users

slice:
	ln -snf users_slice users

str:
	ln -snf users_str users

httpd: clean
	go run ./httpd

trace-httpd: clean
	$(MAKE) $(kind)
	GODEBUG=gctrace=1 go run ./httpd 2>out/http-$(kind).trace


run-manual:
	@echo === map ===
	@ln -snf users_map users
	go run ./manual_gc/
	
	@echo === slice ===
	@ln -snf users_slice users
	go run ./manual_gc/
	
	@echo === str ===
	@ln -snf users_str users
	go run ./manual_gc/
	
trace:
	@echo === map ===
	@ln -snf users_map users
	GODEBUG=gctrace=1 go run ./manual_gc 2>out/map.trace
	
	@echo === slice ===
	@ln -snf users_slice users
	GODEBUG=gctrace=1 go run ./manual_gc 2>out/slice.trace
	
	@echo === str ===
	@ln -snf users_str users
	GODEBUG=gctrace=1 go run ./manual_gc 2>out/str.trace

json:
	gogctrace out/map.trace > out/map.json
	gogctrace out/slice.trace > out/slice.json
	gogctrace out/str.trace > out/str.json

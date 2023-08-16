all:
	@echo === map ===
	@ln -snf users_map users
	go run .
	
	@echo === slice ===
	@ln -snf users_slice users
	go run .
	
	@echo === str ===
	@ln -snf users_str users
	go run .
	
trace:
	@echo === map ===
	@ln -snf users_map users
	GODEBUG=gctrace=1 go run . 2>out/map.trace
	
	@echo === slice ===
	@ln -snf users_slice users
	GODEBUG=gctrace=1 go run . 2>out/slice.trace
	
	@echo === str ===
	@ln -snf users_str users
	GODEBUG=gctrace=1 go run . 2>out/str.trace

json:
	gogctrace out/map.trace > out/map.json
	gogctrace out/slice.trace > out/slice.json
	gogctrace out/str.trace > out/str.json

all:
	$(error please pick a target)

url = http://localhost:8080/users/353
stress:
	curl -f $(url)  # fail is server is not up
	hey -z 10s $(url)

one:
	curl $(url)

map:
	ln -snf users_map users
	go run httpd.go

slice:
	ln -snf users_slice users
	go run httpd.go

str:
	ln -snf users_str users
	go run httpd.go

# ---- old stuff

bench:
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

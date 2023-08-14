all:
	@echo map
	sed -i -e 's#/users_.\+"#/users_map"#' main.go
	go run .
	
	@echo slice
	sed -i -e 's#/users_.\+"#/users_slice"#' main.go
	go run .

	@echo str
	sed -i -e 's#/users_.\+"#/users_str"#' main.go
	go run .

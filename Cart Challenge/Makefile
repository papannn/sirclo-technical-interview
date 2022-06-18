build_main:
	@g++ -std=c++11 main.cpp -o main

run_main:
	@make build_main
	@./main

build_test:
	@g++ -std=c++11 cart_test.cpp -o cart_test

run_test:
	@make build_test
	@./cart_test

run:
	@make run_test
	@make run_main
	@rm cart_test
	@rm main
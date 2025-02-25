.PHONY: build
build:
	clang flags.c -o flags -Wall -Wextra -Werror -Wdeprecated -std=c11 -fsanitize=address

.PHONY: run
run: build
	./flags -name="john" -debug -port=5173

.PHONY: clean
clean:
	rm ./flags

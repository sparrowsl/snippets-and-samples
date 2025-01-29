.PHONY: build
build:
	clang flags.c -o flags -Wall -Wextra -Werror -Wdeprecated -std=c11 -fsanitize=address

.PHONY: run
run: build
	./flags

.PHONY: clean
clean:
	rm ./flags

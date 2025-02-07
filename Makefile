GREEN := $(shell tput setaf 2)
RESET := $(shell tput sgr0)
build: 
	go build
rebuildbuild: clean build 
clean: 
	rm -rf myFind
rebuild: clean build
test: build
	@echo "$(GREEN)Running tests...$(RESET)"
	@echo "$(GREEN)test1: With -d flag$(RESET)"
	./myFind -d test
	@echo "--------------------------------------------------------------------------"
	@echo "$(GREEN)test2: With -f flag$(RESET)"
	./myFind -f test
	@echo "--------------------------------------------------------------------------"
	@echo "$(GREEN)test3: With -sl flag$(RESET)"
	./myFind -sl test
	@echo "--------------------------------------------------------------------------"
	@echo "$(GREEN)test4: With -ext \"txt\" flag$(RESET)"
	./myFind -f -ext "txt" test
	@echo "--------------------------------------------------------------------------"
	@echo "$(GREEN)test5: With -d and -f flags$(RESET)"
	./myFind -d -f test
	@echo "--------------------------------------------------------------------------"
	@echo "$(GREEN)test6: With -d and -sl flags$(RESET)"
	./myFind -d -sl test
	@echo "--------------------------------------------------------------------------"
	@echo "$(GREEN)test7: Without flags$(RESET)"
	./myFind test
	@echo "--------------------------------------------------------------------------"
	@echo "$(GREEN)All tests completed!$(RESET)"
# List days you want to support
DAYS := $(shell seq 1 25)
DAY_INPUTS := $(foreach d,$(DAYS),day$(d)/input)

.PHONY: get
get:
	@missing=0; \
	for f in $(DAY_INPUTS); do \
		if [ ! -f $$f ]; then \
			echo "Missing: $$f"; \
			missing=1; \
		fi; \
	done; \
	if [ $$missing -eq 1 ]; then \
		echo "Running get.sh because some inputs are missing..."; \
		./get.sh; \
	else \
		echo "All input files already exist. Skipping get."; \
	fi

# Generic rule: each day depends on the inputs being present
day%: get
	@echo "Running day$* (no custom command defined)"

day1: get
	sqlite3 < day1/day1.sql


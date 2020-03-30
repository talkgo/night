HUGO = hugo

all: clean
	$(if $(shell PATH=$(PATH) which $(HUGO)),hugo -D --minify,$(error "No $(HUGO) in PATH"))

clean:
	rm -rf public resources
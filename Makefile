PLUGINS=priority

.PHONY: build_plugins

build_plugins:
	$(foreach path,$(PLUGINS),make -C plugins/$(path) ;)

all:
	$(MAKE) -C golang
	$(MAKE) -C filepath

	mkdir -p ./example_app/.waypoint/plugins
	cp ./golang/bin/* ./example_app/.waypoint/plugins
	cp ./filepath/bin/* ./example_app/.waypoint/plugins
all:
	$(MAKE) -C gobuilder_final
	$(MAKE) -C filepath

	mkdir -p ./example_app/.waypoint/plugins
	cp ./gobuilder_final/bin/* ./example_app/.waypoint/plugins
	cp ./filepath/bin/* ./example_app/.waypoint/plugins
all:

	@echo ""
	@echo "### Build Go Builder Plugin"
	$(MAKE) -C gobuilder_final
	
	@echo ""
	@echo ""
	@echo "### Build Filepath Builder Plugin"
	$(MAKE) -C filepath

	@echo ""
	@echo ""
	@echo "### Install Plugins"
	mkdir -p ./example_app/.waypoint/plugins
	cp ./gobuilder_final/bin/* ./example_app/.waypoint/plugins
	cp ./filepath/bin/* ./example_app/.waypoint/plugins
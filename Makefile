CONFIG_FILE := config.toml

APPLICATION_NAME := $(shell toml get $(CONFIG_FILE) server.application_name)

test:
	echo $(APPLICATION_NAME)
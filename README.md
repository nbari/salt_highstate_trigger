# salt_highstate_trigger
Call salt.highstate from web

Usage
-----

Call http://<host>/node_type, example:

	http://10.0.0.4:8080/web_app

The command that will be run is:

	# sudo salt -C G@node_type:web_app state.highstate

> node_type is a custom grain

runit
-----

Using runit:

	#!/bin/sh

	exec 2>&1

	exec chpst -u nobody /service/salt_highstate_trigger/app



SUDO
----

nobody ALL = NOPASSWD: /usr/local/bin/salt

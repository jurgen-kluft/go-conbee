Status
=====

BETA, currently testing!

go-conbee
-----

Wrapper API and cli examples in golang for interacting with lights via Deconz Conbee HTTP API.

Supported API:

- Lights
- Groups
- Sensors

Unsupported API (not hard to add, let me know if you need them):

- Rules
- Scenes
- Schedules

Setup
-----

To install "github.com/jurgen-kluft/go-conbee" golang module.

 ``$ go get github.com/jurgen-kluft/go-conbee``

Check the examples:

- ``get-all-lights.go``
- ``get-light-state.go``
- ``set-light-state.go``

To run the tests you'll need to set the following environment variables:

 1. DECONZ_CONBEE_TEST_HOST (e.g. "10.0.0.18")
 2. DECONZ_CONBEE_TEST_APPKEY (e.g. "0A498B9909")

Bugs and contribution
---------------------

Please feel free to reach out. Issues and PR's are always welcome!

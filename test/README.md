# fugle-go tests

This directory contains additional test suites beyond the unit tests already in
[../fugle](../fugle). Whereas the unit tests run very quickly (since they
don't make any network calls) and are run by Github Actions on every commit, the tests
in this directory are only run manually.

The test packages are:

## integration

This will exercise the entire fugle-go library (or at least as much as is
practical) against the live Fugle API. These tests will verify that the
library is properly coded against the actual behavior of the API, and will
(hopefully) fail upon any incompatible change in the API.

Run tests using:

    go test -v -tags=integration ./integration

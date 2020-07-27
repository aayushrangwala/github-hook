#!/bin/bash
# This script will test the coverage of ./...
# This script will check the coverage of the tests written. The expected coverage to pass is set to 85% for now.
# The coverage profile is to be generated at the time of test with the name 'coverage.out'

set -ex

EXPECTED_COVERAGE="75.0"
COVERAGE_PROFILE="coverage.out"

echo "Running go tool coverage from profile 'coverage.out'..."

make test

COVERAGE_PERCENTAGE=$(go tool cover -func=${COVERAGE_PROFILE}  | grep 'total:' | awk '{print $3}' | sed 's/%//')

if (( ${COVERAGE_PERCENTAGE%%.*} < ${EXPECTED_COVERAGE%%.*} )) ; then
	 	echo "coverage dropped to ${COVERAGE_PERCENTAGE}, expected was ${EXPECTED_COVERAGE}"
	 	exit 1
else
	echo "PASSED Coverage test"
fi

rm coverage.out

exit 0

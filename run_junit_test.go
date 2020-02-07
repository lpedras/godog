package godog

import (
	"fmt"
	"testing"
)

func TestSucceedWithJUnit(t *testing.T) {
	const format = "junit"

	// Will test concurrency setting 0, 1, 2 and 3.
	for concurrency := range make([]int, 4) {
		t.Run(
			fmt.Sprintf("%s_concurrency_%d", format, concurrency),
			func(t *testing.T) {
				testSucceedRun(t, format, concurrency, expectedOutputJUnit)
			},
		)
	}
}

const expectedOutputJUnit = `<?xml version="1.0" encoding="UTF-8"?>
<testsuites name="succeed" tests="60" skipped="0" failures="0" errors="0" time="0s">
  <testsuite name="cucumber json formatter" tests="9" skipped="0" failures="0" errors="0" time="0s">
    <testcase name="Support of Feature Plus Scenario Node" status="passed" time="0s"></testcase>
    <testcase name="Support of Feature Plus Scenario Node With Tags" status="passed" time="0s"></testcase>
    <testcase name="Support of Feature Plus Scenario Outline" status="passed" time="0s"></testcase>
    <testcase name="Support of Feature Plus Scenario Outline With Tags" status="passed" time="0s"></testcase>
    <testcase name="Support of Feature Plus Scenario With Steps" status="passed" time="0s"></testcase>
    <testcase name="Support of Feature Plus Scenario Outline With Steps" status="passed" time="0s"></testcase>
    <testcase name="Support of Comments" status="passed" time="0s"></testcase>
    <testcase name="Support of Docstrings" status="passed" time="0s"></testcase>
    <testcase name="Support of Undefined, Pending and Skipped status" status="passed" time="0s"></testcase>
  </testsuite>
  <testsuite name="event stream formatter" tests="3" skipped="0" failures="0" errors="0" time="0s">
    <testcase name="should fire only suite events without any scenario" status="passed" time="0s"></testcase>
    <testcase name="should process simple scenario" status="passed" time="0s"></testcase>
    <testcase name="should process outline scenario" status="passed" time="0s"></testcase>
  </testsuite>
  <testsuite name="load features" tests="6" skipped="0" failures="0" errors="0" time="0s">
    <testcase name="load features within path" status="passed" time="0s"></testcase>
    <testcase name="load a specific feature file" status="passed" time="0s"></testcase>
    <testcase name="loaded feature should have a number of scenarios #1" status="passed" time="0s"></testcase>
    <testcase name="loaded feature should have a number of scenarios #2" status="passed" time="0s"></testcase>
    <testcase name="loaded feature should have a number of scenarios #3" status="passed" time="0s"></testcase>
    <testcase name="load a number of feature files" status="passed" time="0s"></testcase>
  </testsuite>
  <testsuite name="run background" tests="3" skipped="0" failures="0" errors="0" time="0s">
    <testcase name="should run background steps" status="passed" time="0s"></testcase>
    <testcase name="should skip all consequent steps on failure" status="passed" time="0s"></testcase>
    <testcase name="should continue undefined steps" status="passed" time="0s"></testcase>
  </testsuite>
  <testsuite name="run features" tests="11" skipped="0" failures="0" errors="0" time="0s">
    <testcase name="should run a normal feature" status="passed" time="0s"></testcase>
    <testcase name="should skip steps after failure" status="passed" time="0s"></testcase>
    <testcase name="should skip all scenarios if background fails" status="passed" time="0s"></testcase>
    <testcase name="should skip steps after undefined" status="passed" time="0s"></testcase>
    <testcase name="should match undefined steps in a row" status="passed" time="0s"></testcase>
    <testcase name="should skip steps on pending" status="passed" time="0s"></testcase>
    <testcase name="should handle pending step" status="passed" time="0s"></testcase>
    <testcase name="should mark undefined steps after pending" status="passed" time="0s"></testcase>
    <testcase name="should fail suite if undefined steps follow after the failure" status="passed" time="0s"></testcase>
    <testcase name="should fail suite and skip pending step after failed step" status="passed" time="0s"></testcase>
    <testcase name="should fail suite and skip next step after failed step" status="passed" time="0s"></testcase>
  </testsuite>
  <testsuite name="run features with nested steps" tests="6" skipped="0" failures="0" errors="0" time="0s">
    <testcase name="should run passing multistep successfully" status="passed" time="0s"></testcase>
    <testcase name="should fail multistep" status="passed" time="0s"></testcase>
    <testcase name="should fail nested multistep" status="passed" time="0s"></testcase>
    <testcase name="should skip steps after undefined multistep" status="passed" time="0s"></testcase>
    <testcase name="should match undefined steps in a row" status="passed" time="0s"></testcase>
    <testcase name="should mark undefined steps after pending" status="passed" time="0s"></testcase>
  </testsuite>
  <testsuite name="run outline" tests="6" skipped="0" failures="0" errors="0" time="0s">
    <testcase name="should run a normal outline" status="passed" time="0s"></testcase>
    <testcase name="should continue through examples on failure" status="passed" time="0s"></testcase>
    <testcase name="should skip examples on background failure" status="passed" time="0s"></testcase>
    <testcase name="should translate step table body" status="passed" time="0s"></testcase>
    <testcase name="should translate step doc string argument #1" status="passed" time="0s"></testcase>
    <testcase name="should translate step doc string argument #2" status="passed" time="0s"></testcase>
  </testsuite>
  <testsuite name="suite events" tests="6" skipped="0" failures="0" errors="0" time="0s">
    <testcase name="triggers before scenario event" status="passed" time="0s"></testcase>
    <testcase name="triggers appropriate events for a single scenario" status="passed" time="0s"></testcase>
    <testcase name="triggers appropriate events whole feature" status="passed" time="0s"></testcase>
    <testcase name="triggers appropriate events for two feature files" status="passed" time="0s"></testcase>
    <testcase name="should not trigger events on empty feature" status="passed" time="0s"></testcase>
    <testcase name="should not trigger events on empty scenarios" status="passed" time="0s"></testcase>
  </testsuite>
  <testsuite name="tag filters" tests="4" skipped="0" failures="0" errors="0" time="0s">
    <testcase name="should filter outline examples by tags" status="passed" time="0s"></testcase>
    <testcase name="should filter scenarios by X tag" status="passed" time="0s"></testcase>
    <testcase name="should filter scenarios by X tag not having Y" status="passed" time="0s"></testcase>
    <testcase name="should filter scenarios having Y and Z tags" status="passed" time="0s"></testcase>
  </testsuite>
  <testsuite name="undefined step snippets" tests="5" skipped="0" failures="0" errors="0" time="0s">
    <testcase name="should generate snippets" status="passed" time="0s"></testcase>
    <testcase name="should generate snippets with more arguments" status="passed" time="0s"></testcase>
    <testcase name="should handle escaped symbols" status="passed" time="0s"></testcase>
    <testcase name="should handle string argument followed by comma" status="passed" time="0s"></testcase>
    <testcase name="should handle arguments in the beggining or end of the step" status="passed" time="0s"></testcase>
  </testsuite>
  <testsuite name="užkrauti savybes" tests="1" skipped="0" failures="0" errors="0" time="0s">
    <testcase name="savybių užkrovimas iš aplanko" status="passed" time="0s"></testcase>
  </testsuite>
</testsuites>`

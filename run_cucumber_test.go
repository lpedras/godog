package godog

import (
	"fmt"
	"testing"
)

func TestSucceedWithCucumber(t *testing.T) {
	const format = "cucumber"

	// Will test concurrency setting 0 and 1, which means no concurrency.
	for concurrency := range make([]int, 2) {
		t.Run(
			fmt.Sprintf("%s_concurrency_%d", format, concurrency),
			func(t *testing.T) {
				testSucceedRun(t, format, concurrency, expectedOutputCucumber)
			},
		)
	}
}

const expectedOutputCucumber = `[
    {
        "uri": "features/background.feature",
        "id": "run-background",
        "keyword": "Feature",
        "name": "run background",
        "description": "  In order to test application behavior\n  As a test suite\n  I need to be able to run background correctly",
        "line": 1,
        "elements": [
            {
                "id": "run-background;should-run-background-steps",
                "keyword": "Scenario",
                "name": "should run background steps",
                "description": "",
                "line": 6,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"normal.feature\" file:",
                        "line": 7,
                        "doc_string": {
                            "value": "Feature: with background\n\n  Background:\n    Given a feature path \"features/load.feature:6\"\n\n  Scenario: parse a scenario\n    When I parse features\n    Then I should have 1 scenario registered",
                            "content_type": "",
                            "line": 8
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 18,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the suite should have passed",
                        "line": 19,
                        "match": {
                            "location": "suite_context.go:338"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following steps should be passed:",
                        "line": 20,
                        "doc_string": {
                            "value": "a feature path \"features/load.feature:6\"\nI parse features\nI should have 1 scenario registered",
                            "content_type": "",
                            "line": 21
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "run-background;should-skip-all-consequent-steps-on-failure",
                "keyword": "Scenario",
                "name": "should skip all consequent steps on failure",
                "description": "",
                "line": 27,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"normal.feature\" file:",
                        "line": 28,
                        "doc_string": {
                            "value": "Feature: with background\n\n  Background:\n    Given a failing step\n    And a feature path \"features/load.feature:6\"\n\n  Scenario: parse a scenario\n    When I parse features\n    Then I should have 1 scenario registered",
                            "content_type": "",
                            "line": 29
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 40,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the suite should have failed",
                        "line": 41,
                        "match": {
                            "location": "suite_context.go:338"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following steps should be failed:",
                        "line": 42,
                        "doc_string": {
                            "value": "a failing step",
                            "content_type": "",
                            "line": 43
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following steps should be skipped:",
                        "line": 46,
                        "doc_string": {
                            "value": "a feature path \"features/load.feature:6\"\nI parse features\nI should have 1 scenario registered",
                            "content_type": "",
                            "line": 47
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "run-background;should-continue-undefined-steps",
                "keyword": "Scenario",
                "name": "should continue undefined steps",
                "description": "",
                "line": 53,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"normal.feature\" file:",
                        "line": 54,
                        "doc_string": {
                            "value": "Feature: with background\n\n  Background:\n    Given an undefined step\n\n  Scenario: parse a scenario\n    When I do undefined action\n    Then I should have 1 scenario registered",
                            "content_type": "",
                            "line": 55
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 65,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the suite should have passed",
                        "line": 66,
                        "match": {
                            "location": "suite_context.go:338"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following steps should be undefined:",
                        "line": 67,
                        "doc_string": {
                            "value": "an undefined step\nI do undefined action",
                            "content_type": "",
                            "line": 68
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following steps should be skipped:",
                        "line": 72,
                        "doc_string": {
                            "value": "I should have 1 scenario registered",
                            "content_type": "",
                            "line": 73
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            }
        ]
    },
    {
        "uri": "features/events.feature",
        "id": "suite-events",
        "keyword": "Feature",
        "name": "suite events",
        "description": "  In order to run tasks before and after important events\n  As a test suite\n  I need to provide a way to hook into these events",
        "line": 1,
        "elements": [
            {
                "id": "suite-events;triggers-before-scenario-event",
                "keyword": "Scenario",
                "name": "triggers before scenario event",
                "description": "",
                "line": 9,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "I'm listening to suite events",
                        "line": 7,
                        "match": {
                            "location": "suite_context.go:285"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Given ",
                        "name": "a feature path \"features/load.feature:6\"",
                        "line": 10,
                        "match": {
                            "location": "suite_context.go:324"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 11,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "there was event triggered before scenario \"load features within path\"",
                        "line": 12,
                        "match": {
                            "location": "suite_context.go:414"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "suite-events;triggers-appropriate-events-for-a-single-scenario",
                "keyword": "Scenario",
                "name": "triggers appropriate events for a single scenario",
                "description": "",
                "line": 14,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "I'm listening to suite events",
                        "line": 7,
                        "match": {
                            "location": "suite_context.go:285"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Given ",
                        "name": "a feature path \"features/load.feature:6\"",
                        "line": 15,
                        "match": {
                            "location": "suite_context.go:324"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 16,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "these events had to be fired for a number of times:",
                        "line": 17,
                        "match": {
                            "location": "suite_context.go:442"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        },
                        "rows": [
                            {
                                "cells": [
                                    "BeforeSuite",
                                    "1"
                                ]
                            },
                            {
                                "cells": [
                                    "BeforeFeature",
                                    "1"
                                ]
                            },
                            {
                                "cells": [
                                    "BeforeScenario",
                                    "1"
                                ]
                            },
                            {
                                "cells": [
                                    "BeforeStep",
                                    "3"
                                ]
                            },
                            {
                                "cells": [
                                    "AfterStep",
                                    "3"
                                ]
                            },
                            {
                                "cells": [
                                    "AfterScenario",
                                    "1"
                                ]
                            },
                            {
                                "cells": [
                                    "AfterFeature",
                                    "1"
                                ]
                            },
                            {
                                "cells": [
                                    "AfterSuite",
                                    "1"
                                ]
                            }
                        ]
                    }
                ]
            },
            {
                "id": "suite-events;triggers-appropriate-events-whole-feature",
                "keyword": "Scenario",
                "name": "triggers appropriate events whole feature",
                "description": "",
                "line": 27,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "I'm listening to suite events",
                        "line": 7,
                        "match": {
                            "location": "suite_context.go:285"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Given ",
                        "name": "a feature path \"features/load.feature\"",
                        "line": 28,
                        "match": {
                            "location": "suite_context.go:324"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 29,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "these events had to be fired for a number of times:",
                        "line": 30,
                        "match": {
                            "location": "suite_context.go:442"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        },
                        "rows": [
                            {
                                "cells": [
                                    "BeforeSuite",
                                    "1"
                                ]
                            },
                            {
                                "cells": [
                                    "BeforeFeature",
                                    "1"
                                ]
                            },
                            {
                                "cells": [
                                    "BeforeScenario",
                                    "6"
                                ]
                            },
                            {
                                "cells": [
                                    "BeforeStep",
                                    "19"
                                ]
                            },
                            {
                                "cells": [
                                    "AfterStep",
                                    "19"
                                ]
                            },
                            {
                                "cells": [
                                    "AfterScenario",
                                    "6"
                                ]
                            },
                            {
                                "cells": [
                                    "AfterFeature",
                                    "1"
                                ]
                            },
                            {
                                "cells": [
                                    "AfterSuite",
                                    "1"
                                ]
                            }
                        ]
                    }
                ]
            },
            {
                "id": "suite-events;triggers-appropriate-events-for-two-feature-files",
                "keyword": "Scenario",
                "name": "triggers appropriate events for two feature files",
                "description": "",
                "line": 40,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "I'm listening to suite events",
                        "line": 7,
                        "match": {
                            "location": "suite_context.go:285"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Given ",
                        "name": "a feature path \"features/load.feature:6\"",
                        "line": 41,
                        "match": {
                            "location": "suite_context.go:324"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "a feature path \"features/multistep.feature:6\"",
                        "line": 42,
                        "match": {
                            "location": "suite_context.go:324"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 43,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "these events had to be fired for a number of times:",
                        "line": 44,
                        "match": {
                            "location": "suite_context.go:442"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        },
                        "rows": [
                            {
                                "cells": [
                                    "BeforeSuite",
                                    "1"
                                ]
                            },
                            {
                                "cells": [
                                    "BeforeFeature",
                                    "2"
                                ]
                            },
                            {
                                "cells": [
                                    "BeforeScenario",
                                    "2"
                                ]
                            },
                            {
                                "cells": [
                                    "BeforeStep",
                                    "7"
                                ]
                            },
                            {
                                "cells": [
                                    "AfterStep",
                                    "7"
                                ]
                            },
                            {
                                "cells": [
                                    "AfterScenario",
                                    "2"
                                ]
                            },
                            {
                                "cells": [
                                    "AfterFeature",
                                    "2"
                                ]
                            },
                            {
                                "cells": [
                                    "AfterSuite",
                                    "1"
                                ]
                            }
                        ]
                    }
                ]
            },
            {
                "id": "suite-events;should-not-trigger-events-on-empty-feature",
                "keyword": "Scenario",
                "name": "should not trigger events on empty feature",
                "description": "",
                "line": 54,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "I'm listening to suite events",
                        "line": 7,
                        "match": {
                            "location": "suite_context.go:285"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Given ",
                        "name": "a feature \"normal.feature\" file:",
                        "line": 55,
                        "doc_string": {
                            "value": "Feature: empty\n\n  Scenario: one\n\n  Scenario: two",
                            "content_type": "",
                            "line": 56
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 63,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "these events had to be fired for a number of times:",
                        "line": 64,
                        "match": {
                            "location": "suite_context.go:442"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        },
                        "rows": [
                            {
                                "cells": [
                                    "BeforeSuite",
                                    "1"
                                ]
                            },
                            {
                                "cells": [
                                    "BeforeFeature",
                                    "0"
                                ]
                            },
                            {
                                "cells": [
                                    "BeforeScenario",
                                    "0"
                                ]
                            },
                            {
                                "cells": [
                                    "BeforeStep",
                                    "0"
                                ]
                            },
                            {
                                "cells": [
                                    "AfterStep",
                                    "0"
                                ]
                            },
                            {
                                "cells": [
                                    "AfterScenario",
                                    "0"
                                ]
                            },
                            {
                                "cells": [
                                    "AfterFeature",
                                    "0"
                                ]
                            },
                            {
                                "cells": [
                                    "AfterSuite",
                                    "1"
                                ]
                            }
                        ]
                    }
                ]
            },
            {
                "id": "suite-events;should-not-trigger-events-on-empty-scenarios",
                "keyword": "Scenario",
                "name": "should not trigger events on empty scenarios",
                "description": "",
                "line": 74,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "I'm listening to suite events",
                        "line": 7,
                        "match": {
                            "location": "suite_context.go:285"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Given ",
                        "name": "a feature \"normal.feature\" file:",
                        "line": 75,
                        "doc_string": {
                            "value": "Feature: half empty\n\n  Scenario: one\n\n  Scenario: two\n    Then passing step\n\n  Scenario Outline: three\n    Then passing step\n\n    Examples:\n      | a |\n      | 1 |",
                            "content_type": "",
                            "line": 76
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 91,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "these events had to be fired for a number of times:",
                        "line": 92,
                        "match": {
                            "location": "suite_context.go:442"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        },
                        "rows": [
                            {
                                "cells": [
                                    "BeforeSuite",
                                    "1"
                                ]
                            },
                            {
                                "cells": [
                                    "BeforeFeature",
                                    "1"
                                ]
                            },
                            {
                                "cells": [
                                    "BeforeScenario",
                                    "2"
                                ]
                            },
                            {
                                "cells": [
                                    "BeforeStep",
                                    "2"
                                ]
                            },
                            {
                                "cells": [
                                    "AfterStep",
                                    "2"
                                ]
                            },
                            {
                                "cells": [
                                    "AfterScenario",
                                    "2"
                                ]
                            },
                            {
                                "cells": [
                                    "AfterFeature",
                                    "1"
                                ]
                            },
                            {
                                "cells": [
                                    "AfterSuite",
                                    "1"
                                ]
                            }
                        ]
                    }
                ]
            }
        ]
    },
    {
        "uri": "features/formatter/cucumber.feature",
        "id": "cucumber-json-formatter",
        "keyword": "Feature",
        "name": "cucumber json formatter",
        "description": "  In order to support tools that import cucumber json output\n  I need to be able to support cucumber json formatted output",
        "line": 1,
        "comments": [
            {
                "value": "# Currently godog only supports comments on Feature and not",
                "line": 365
            },
            {
                "value": "# scenario and steps.",
                "line": 366
            }
        ],
        "elements": [
            {
                "id": "cucumber-json-formatter;support-of-feature-plus-scenario-node",
                "keyword": "Scenario",
                "name": "Support of Feature Plus Scenario Node",
                "description": "",
                "line": 5,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"features/simple.feature\" file:",
                        "line": 6,
                        "doc_string": {
                            "value": "    Feature: simple feature\n        simple feature description\n    Scenario: simple scenario\n        simple scenario description",
                            "content_type": "",
                            "line": 7
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite with formatter \"cucumber\"",
                        "line": 13,
                        "match": {
                            "location": "suite_context.go:131"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the rendered json will be as follows:",
                        "line": 14,
                        "doc_string": {
                            "value": "  [\n    {\n      \"uri\": \"features/simple.feature\",\n      \"id\": \"simple-feature\",\n      \"keyword\": \"Feature\",\n      \"name\": \"simple feature\",\n      \"description\": \"        simple feature description\",\n      \"line\": 1,\n      \"elements\": [\n        {\n          \"id\": \"simple-feature;simple-scenario\",\n          \"keyword\": \"Scenario\",\n          \"name\": \"simple scenario\",\n          \"description\": \"        simple scenario description\",\n          \"line\": 3,\n          \"type\": \"scenario\"\n        }\n      ]\n    }\n  ]",
                            "content_type": "application/json",
                            "line": 15
                        },
                        "match": {
                            "location": "suite_context.go:459"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "cucumber-json-formatter;support-of-feature-plus-scenario-node-with-tags",
                "keyword": "Scenario",
                "name": "Support of Feature Plus Scenario Node With Tags",
                "description": "",
                "line": 38,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"features/simple.feature\" file:",
                        "line": 39,
                        "doc_string": {
                            "value": "    @TAG1\n    Feature: simple feature\n        simple feature description\n    @TAG2 @TAG3\n    Scenario: simple scenario\n        simple scenario description",
                            "content_type": "",
                            "line": 40
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite with formatter \"cucumber\"",
                        "line": 48,
                        "match": {
                            "location": "suite_context.go:131"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the rendered json will be as follows:",
                        "line": 49,
                        "doc_string": {
                            "value": "  [\n    {\n      \"uri\": \"features/simple.feature\",\n      \"id\": \"simple-feature\",\n      \"keyword\": \"Feature\",\n      \"name\": \"simple feature\",\n      \"description\": \"        simple feature description\",\n      \"line\": 2,\n      \"tags\": [\n        {\n          \"name\": \"@TAG1\",\n          \"line\": 1\n        }\n      ],\n      \"elements\": [\n        {\n          \"id\": \"simple-feature;simple-scenario\",\n          \"keyword\": \"Scenario\",\n          \"name\": \"simple scenario\",\n          \"description\": \"        simple scenario description\",\n          \"line\": 5,\n          \"type\": \"scenario\",\n          \"tags\": [\n            {\n              \"name\": \"@TAG1\",\n              \"line\": 1\n            },\n            {\n              \"name\": \"@TAG2\",\n              \"line\": 4\n            },\n            {\n              \"name\": \"@TAG3\",\n              \"line\": 4\n            }\n          ]\n        }\n      ]\n    }\n]",
                            "content_type": "application/json",
                            "line": 50
                        },
                        "match": {
                            "location": "suite_context.go:459"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "cucumber-json-formatter;support-of-feature-plus-scenario-outline",
                "keyword": "Scenario",
                "name": "Support of Feature Plus Scenario Outline",
                "description": "",
                "line": 92,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"features/simple.feature\" file:",
                        "line": 93,
                        "doc_string": {
                            "value": "    Feature: simple feature\n        simple feature description\n\n    Scenario Outline: simple scenario\n        simple scenario description\n\n    Examples: simple examples\n    | status |\n    | pass   |\n    | fail   |",
                            "content_type": "",
                            "line": 94
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite with formatter \"cucumber\"",
                        "line": 106,
                        "match": {
                            "location": "suite_context.go:131"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the rendered json will be as follows:",
                        "line": 107,
                        "doc_string": {
                            "value": "  [\n    {\n      \"uri\": \"features/simple.feature\",\n      \"id\": \"simple-feature\",\n      \"keyword\": \"Feature\",\n      \"name\": \"simple feature\",\n      \"description\": \"        simple feature description\",\n      \"line\": 1,\n      \"elements\": [\n        {\n          \"id\": \"simple-feature;simple-scenario;simple-examples;2\",\n          \"keyword\": \"Scenario Outline\",\n          \"name\": \"simple scenario\",\n          \"description\": \"        simple scenario description\",\n          \"line\": 9,\n          \"type\": \"scenario\"\n        },\n        {\n          \"id\": \"simple-feature;simple-scenario;simple-examples;3\",\n          \"keyword\": \"Scenario Outline\",\n          \"name\": \"simple scenario\",\n          \"description\": \"        simple scenario description\",\n          \"line\": 10,\n          \"type\": \"scenario\"\n        }\n      ]\n    }\n  ]",
                            "content_type": "",
                            "line": 108
                        },
                        "match": {
                            "location": "suite_context.go:459"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "cucumber-json-formatter;support-of-feature-plus-scenario-outline-with-tags",
                "keyword": "Scenario",
                "name": "Support of Feature Plus Scenario Outline With Tags",
                "description": "",
                "line": 139,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"features/simple.feature\" file:",
                        "line": 140,
                        "doc_string": {
                            "value": "    @TAG1\n    Feature: simple feature\n        simple feature description\n\n    @TAG2\n    Scenario Outline: simple scenario\n        simple scenario description\n\n    @TAG3\n    Examples: simple examples\n    | status |\n    | pass   |\n    | fail   |",
                            "content_type": "",
                            "line": 141
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite with formatter \"cucumber\"",
                        "line": 156,
                        "match": {
                            "location": "suite_context.go:131"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the rendered json will be as follows:",
                        "line": 157,
                        "doc_string": {
                            "value": "    [\n      {\n        \"uri\": \"features/simple.feature\",\n        \"id\": \"simple-feature\",\n        \"keyword\": \"Feature\",\n        \"name\": \"simple feature\",\n        \"description\": \"        simple feature description\",\n        \"line\": 2,\n        \"tags\": [\n          {\n            \"name\": \"@TAG1\",\n            \"line\": 1\n          }\n        ],\n        \"elements\": [\n          {\n            \"id\": \"simple-feature;simple-scenario;simple-examples;2\",\n            \"keyword\": \"Scenario Outline\",\n            \"name\": \"simple scenario\",\n            \"description\": \"        simple scenario description\",\n            \"line\": 12,\n            \"type\": \"scenario\",\n            \"tags\": [\n              {\n                \"name\": \"@TAG1\",\n                \"line\": 1\n              },\n              {\n                \"name\": \"@TAG2\",\n                \"line\": 5\n              },\n              {\n                \"name\": \"@TAG3\",\n                \"line\": 9\n              }\n            ]\n          },\n          {\n            \"id\": \"simple-feature;simple-scenario;simple-examples;3\",\n            \"keyword\": \"Scenario Outline\",\n            \"name\": \"simple scenario\",\n            \"description\": \"        simple scenario description\",\n            \"line\": 13,\n            \"type\": \"scenario\",\n            \"tags\": [\n              {\n                \"name\": \"@TAG1\",\n                \"line\": 1\n              },\n              {\n                \"name\": \"@TAG2\",\n                \"line\": 5\n              },\n              {\n                \"name\": \"@TAG3\",\n                \"line\": 9\n              }\n            ]\n          }\n        ]\n      }\n    ]",
                            "content_type": "",
                            "line": 158
                        },
                        "match": {
                            "location": "suite_context.go:459"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "cucumber-json-formatter;support-of-feature-plus-scenario-with-steps",
                "keyword": "Scenario",
                "name": "Support of Feature Plus Scenario With Steps",
                "description": "",
                "line": 222,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"features/simple.feature\" file:",
                        "line": 223,
                        "doc_string": {
                            "value": "    Feature: simple feature\n        simple feature description\n\n    Scenario: simple scenario\n        simple scenario description\n\n    Given passing step\n    Then a failing step\n",
                            "content_type": "",
                            "line": 224
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite with formatter \"cucumber\"",
                        "line": 235,
                        "match": {
                            "location": "suite_context.go:131"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the rendered json will be as follows:",
                        "line": 236,
                        "doc_string": {
                            "value": "  [\n    {\n      \"uri\": \"features/simple.feature\",\n      \"id\": \"simple-feature\",\n      \"keyword\": \"Feature\",\n      \"name\": \"simple feature\",\n      \"description\": \"        simple feature description\",\n      \"line\": 1,\n      \"elements\": [\n        {\n          \"id\": \"simple-feature;simple-scenario\",\n          \"keyword\": \"Scenario\",\n          \"name\": \"simple scenario\",\n          \"description\": \"        simple scenario description\",\n          \"line\": 4,\n          \"type\": \"scenario\",\n          \"steps\": [\n            {\n              \"keyword\": \"Given \",\n              \"name\": \"passing step\",\n              \"line\": 7,\n              \"match\": {\n                \"location\": \"suite_context.go:64\"\n              },\n              \"result\": {\n                \"status\": \"passed\",\n                \"duration\": 0\n              }\n            },\n            {\n              \"keyword\": \"Then \",\n              \"name\": \"a failing step\",\n              \"line\": 8,\n              \"match\": {\n                \"location\": \"suite_context.go:47\"\n              },\n              \"result\": {\n                \"status\": \"failed\",\n                \"error_message\": \"intentional failure\",\n                \"duration\": 0\n              }\n            }\n          ]\n        }\n      ]\n    }\n  ]",
                            "content_type": "",
                            "line": 237
                        },
                        "match": {
                            "location": "suite_context.go:459"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "cucumber-json-formatter;support-of-feature-plus-scenario-outline-with-steps",
                "keyword": "Scenario",
                "name": "Support of Feature Plus Scenario Outline With Steps",
                "description": "",
                "line": 286,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"features/simple.feature\" file:",
                        "line": 287,
                        "doc_string": {
                            "value": "  Feature: simple feature\n    simple feature description\n\n    Scenario Outline: simple scenario\n    simple scenario description\n\n      Given \u003cstatus\u003e step\n\n    Examples: simple examples\n    | status |\n    | passing |\n    | failing |\n",
                            "content_type": "",
                            "line": 288
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite with formatter \"cucumber\"",
                        "line": 303,
                        "match": {
                            "location": "suite_context.go:131"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the rendered json will be as follows:",
                        "line": 304,
                        "doc_string": {
                            "value": "  [\n    {\n      \"uri\": \"features/simple.feature\",\n      \"id\": \"simple-feature\",\n      \"keyword\": \"Feature\",\n      \"name\": \"simple feature\",\n      \"description\": \"    simple feature description\",\n      \"line\": 1,\n      \"elements\": [\n        {\n          \"id\": \"simple-feature;simple-scenario;simple-examples;2\",\n          \"keyword\": \"Scenario Outline\",\n          \"name\": \"simple scenario\",\n          \"description\": \"    simple scenario description\",\n          \"line\": 11,\n          \"type\": \"scenario\",\n          \"steps\": [\n            {\n              \"keyword\": \"Given \",\n              \"name\": \"passing step\",\n              \"line\": 11,\n              \"match\": {\n                \"location\": \"suite_context.go:64\"\n              },\n              \"result\": {\n                \"status\": \"passed\",\n                \"duration\": 0\n              }\n            }\n          ]\n        },\n        {\n          \"id\": \"simple-feature;simple-scenario;simple-examples;3\",\n          \"keyword\": \"Scenario Outline\",\n          \"name\": \"simple scenario\",\n          \"description\": \"    simple scenario description\",\n          \"line\": 12,\n          \"type\": \"scenario\",\n          \"steps\": [\n            {\n              \"keyword\": \"Given \",\n              \"name\": \"failing step\",\n              \"line\": 12,\n              \"match\": {\n                \"location\": \"suite_context.go:47\"\n              },\n              \"result\": {\n                \"status\": \"failed\",\n                \"error_message\": \"intentional failure\",\n                \"duration\": 0\n              }\n            }\n          ]\n        }\n      ]\n    }\n  ]",
                            "content_type": "",
                            "line": 305
                        },
                        "match": {
                            "location": "suite_context.go:459"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "cucumber-json-formatter;support-of-comments",
                "keyword": "Scenario",
                "name": "Support of Comments",
                "description": "",
                "line": 367,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"features/simple.feature\" file:",
                        "line": 368,
                        "doc_string": {
                            "value": "    #Feature comment\n    Feature: simple feature\n      simple description\n\n      Scenario: simple scenario\n      simple feature description",
                            "content_type": "",
                            "line": 369
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite with formatter \"cucumber\"",
                        "line": 377,
                        "match": {
                            "location": "suite_context.go:131"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the rendered json will be as follows:",
                        "line": 378,
                        "doc_string": {
                            "value": "    [\n      {\n        \"uri\": \"features/simple.feature\",\n        \"id\": \"simple-feature\",\n        \"keyword\": \"Feature\",\n        \"name\": \"simple feature\",\n        \"description\": \"      simple description\",\n        \"line\": 2,\n        \"comments\": [\n          {\n            \"value\": \"#Feature comment\",\n            \"line\": 1\n          }\n        ],\n        \"elements\": [\n          {\n            \"id\": \"simple-feature;simple-scenario\",\n            \"keyword\": \"Scenario\",\n            \"name\": \"simple scenario\",\n            \"description\": \"      simple feature description\",\n            \"line\": 5,\n            \"type\": \"scenario\"\n          }\n        ]\n      }\n    ]",
                            "content_type": "",
                            "line": 379
                        },
                        "match": {
                            "location": "suite_context.go:459"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "cucumber-json-formatter;support-of-docstrings",
                "keyword": "Scenario",
                "name": "Support of Docstrings",
                "description": "",
                "line": 407,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"features/simple.feature\" file:",
                        "line": 408,
                        "doc_string": {
                            "value": "    Feature: simple feature\n      simple description\n\n      Scenario: simple scenario\n      simple feature description\n\n      Given passing step\n      \"\"\" content type\n      step doc string\n      \"\"\"",
                            "content_type": "",
                            "line": 409
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite with formatter \"cucumber\"",
                        "line": 421,
                        "match": {
                            "location": "suite_context.go:131"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the rendered json will be as follows:",
                        "line": 422,
                        "doc_string": {
                            "value": "    [\n  {\n    \"uri\": \"features/simple.feature\",\n    \"id\": \"simple-feature\",\n    \"keyword\": \"Feature\",\n    \"name\": \"simple feature\",\n    \"description\": \"      simple description\",\n    \"line\": 1,\n    \"elements\": [\n      {\n        \"id\": \"simple-feature;simple-scenario\",\n        \"keyword\": \"Scenario\",\n        \"name\": \"simple scenario\",\n        \"description\": \"      simple feature description\",\n        \"line\": 4,\n        \"type\": \"scenario\",\n        \"steps\": [\n          {\n            \"keyword\": \"Given \",\n            \"name\": \"passing step\",\n            \"line\": 7,\n            \"doc_string\": {\n              \"value\": \"step doc string\",\n              \"content_type\": \"content type\",\n              \"line\": 8\n            },\n            \"match\": {\n              \"location\": \"suite_context.go:64\"\n            },\n            \"result\": {\n              \"status\": \"passed\",\n              \"duration\": 0\n            }\n          }\n        ]\n      }\n    ]\n  }\n]",
                            "content_type": "",
                            "line": 423
                        },
                        "match": {
                            "location": "suite_context.go:459"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "cucumber-json-formatter;support-of-undefined,-pending-and-skipped-status",
                "keyword": "Scenario",
                "name": "Support of Undefined, Pending and Skipped status",
                "description": "",
                "line": 464,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"features/simple.feature\" file:",
                        "line": 465,
                        "doc_string": {
                            "value": "  Feature: simple feature\n  simple feature description\n\n  Scenario: simple scenario\n  simple scenario description\n\n    Given passing step\n    And pending step\n    And undefined\n    And passing step\n",
                            "content_type": "",
                            "line": 466
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite with formatter \"cucumber\"",
                        "line": 479,
                        "match": {
                            "location": "suite_context.go:131"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the rendered json will be as follows:",
                        "line": 480,
                        "doc_string": {
                            "value": "  [\n    {\n      \"uri\": \"features/simple.feature\",\n      \"id\": \"simple-feature\",\n      \"keyword\": \"Feature\",\n      \"name\": \"simple feature\",\n      \"description\": \"  simple feature description\",\n      \"line\": 1,\n      \"elements\": [\n        {\n          \"id\": \"simple-feature;simple-scenario\",\n          \"keyword\": \"Scenario\",\n          \"name\": \"simple scenario\",\n          \"description\": \"  simple scenario description\",\n          \"line\": 4,\n          \"type\": \"scenario\",\n          \"steps\": [\n            {\n              \"keyword\": \"Given \",\n              \"name\": \"passing step\",\n              \"line\": 7,\n              \"match\": {\n                \"location\": \"suite_context.go:64\"\n              },\n              \"result\": {\n                \"status\": \"passed\",\n                \"duration\": 0\n              }\n            },\n            {\n              \"keyword\": \"And \",\n              \"name\": \"pending step\",\n              \"line\": 8,\n              \"match\": {\n                \"location\": \"features/simple.feature:8\"\n              },\n              \"result\": {\n                \"status\": \"pending\"\n              }\n            },\n            {\n              \"keyword\": \"And \",\n              \"name\": \"undefined\",\n              \"line\": 9,\n              \"match\": {\n                \"location\": \"features/simple.feature:9\"\n              },\n              \"result\": {\n                \"status\": \"undefined\"\n              }\n            },\n            {\n              \"keyword\": \"And \",\n              \"name\": \"passing step\",\n              \"line\": 10,\n              \"match\": {\n                \"location\": \"suite_context.go:64\"\n              },\n              \"result\": {\n                \"status\": \"skipped\"\n              }\n            }\n          ]\n        }\n      ]\n    }\n  ]",
                            "content_type": "",
                            "line": 481
                        },
                        "match": {
                            "location": "suite_context.go:459"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            }
        ]
    },
    {
        "uri": "features/formatter/events.feature",
        "id": "event-stream-formatter",
        "keyword": "Feature",
        "name": "event stream formatter",
        "description": "  In order to have universal cucumber formatter\n  As a test suite\n  I need to be able to support event stream formatter",
        "line": 1,
        "elements": [
            {
                "id": "event-stream-formatter;should-fire-only-suite-events-without-any-scenario",
                "keyword": "Scenario",
                "name": "should fire only suite events without any scenario",
                "description": "",
                "line": 6,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature path \"features/load.feature:4\"",
                        "line": 7,
                        "match": {
                            "location": "suite_context.go:324"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite with formatter \"events\"",
                        "line": 8,
                        "match": {
                            "location": "suite_context.go:131"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the following events should be fired:",
                        "line": 9,
                        "doc_string": {
                            "value": "  TestRunStarted\n  TestSource\n  TestRunFinished",
                            "content_type": "",
                            "line": 10
                        },
                        "match": {
                            "location": "suite_context.go:145"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "event-stream-formatter;should-process-simple-scenario",
                "keyword": "Scenario",
                "name": "should process simple scenario",
                "description": "",
                "line": 16,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature path \"features/load.feature:24\"",
                        "line": 17,
                        "match": {
                            "location": "suite_context.go:324"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite with formatter \"events\"",
                        "line": 18,
                        "match": {
                            "location": "suite_context.go:131"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the following events should be fired:",
                        "line": 19,
                        "doc_string": {
                            "value": "  TestRunStarted\n  TestSource\n  TestCaseStarted\n  StepDefinitionFound\n  TestStepStarted\n  TestStepFinished\n  StepDefinitionFound\n  TestStepStarted\n  TestStepFinished\n  StepDefinitionFound\n  TestStepStarted\n  TestStepFinished\n  TestCaseFinished\n  TestRunFinished",
                            "content_type": "",
                            "line": 20
                        },
                        "match": {
                            "location": "suite_context.go:145"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "event-stream-formatter;should-process-outline-scenario",
                "keyword": "Scenario",
                "name": "should process outline scenario",
                "description": "",
                "line": 37,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature path \"features/load.feature:32\"",
                        "line": 38,
                        "match": {
                            "location": "suite_context.go:324"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite with formatter \"events\"",
                        "line": 39,
                        "match": {
                            "location": "suite_context.go:131"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the following events should be fired:",
                        "line": 40,
                        "doc_string": {
                            "value": "  TestRunStarted\n  TestSource\n  TestCaseStarted\n  StepDefinitionFound\n  TestStepStarted\n  TestStepFinished\n  StepDefinitionFound\n  TestStepStarted\n  TestStepFinished\n  StepDefinitionFound\n  TestStepStarted\n  TestStepFinished\n  TestCaseFinished\n  TestCaseStarted\n  StepDefinitionFound\n  TestStepStarted\n  TestStepFinished\n  StepDefinitionFound\n  TestStepStarted\n  TestStepFinished\n  StepDefinitionFound\n  TestStepStarted\n  TestStepFinished\n  TestCaseFinished\n  TestCaseStarted\n  StepDefinitionFound\n  TestStepStarted\n  TestStepFinished\n  StepDefinitionFound\n  TestStepStarted\n  TestStepFinished\n  StepDefinitionFound\n  TestStepStarted\n  TestStepFinished\n  TestCaseFinished\n  TestRunFinished",
                            "content_type": "",
                            "line": 41
                        },
                        "match": {
                            "location": "suite_context.go:145"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            }
        ]
    },
    {
        "uri": "features/lang.feature",
        "id": "užkrauti-savybes",
        "keyword": "Savybė",
        "name": "užkrauti savybes",
        "description": "  Kad būtų galima paleisti savybių testus\n  Kaip testavimo įrankis\n  Aš turiu galėti užregistruoti savybes",
        "line": 3,
        "tags": [
            {
                "name": "@lang",
                "line": 2
            }
        ],
        "elements": [
            {
                "id": "užkrauti-savybes;savybių-užkrovimas-iš-aplanko",
                "keyword": "Scenarijus",
                "name": "savybių užkrovimas iš aplanko",
                "description": "",
                "line": 8,
                "type": "scenario",
                "tags": [
                    {
                        "name": "@lang",
                        "line": 2
                    }
                ],
                "steps": [
                    {
                        "keyword": "Duota ",
                        "name": "savybių aplankas \"features\"",
                        "line": 9,
                        "match": {
                            "location": "suite_context.go:324"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Kai ",
                        "name": "aš išskaitau savybes",
                        "line": 10,
                        "match": {
                            "location": "suite_context.go:329"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Tada ",
                        "name": "aš turėčiau turėti 11 savybių failus:",
                        "line": 11,
                        "doc_string": {
                            "value": "features/background.feature\nfeatures/events.feature\nfeatures/formatter/cucumber.feature\nfeatures/formatter/events.feature\nfeatures/lang.feature\nfeatures/load.feature\nfeatures/multistep.feature\nfeatures/outline.feature\nfeatures/run.feature\nfeatures/snippets.feature\nfeatures/tags.feature",
                            "content_type": "",
                            "line": 12
                        },
                        "match": {
                            "location": "suite_context.go:348"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            }
        ]
    },
    {
        "uri": "features/load.feature",
        "id": "load-features",
        "keyword": "Feature",
        "name": "load features",
        "description": "  In order to run features\n  As a test suite\n  I need to be able to load features",
        "line": 1,
        "elements": [
            {
                "id": "load-features;load-features-within-path",
                "keyword": "Scenario",
                "name": "load features within path",
                "description": "",
                "line": 6,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature path \"features\"",
                        "line": 7,
                        "match": {
                            "location": "suite_context.go:324"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I parse features",
                        "line": 8,
                        "match": {
                            "location": "suite_context.go:329"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "I should have 11 feature files:",
                        "line": 9,
                        "doc_string": {
                            "value": "features/background.feature\nfeatures/events.feature\nfeatures/formatter/cucumber.feature\nfeatures/formatter/events.feature\nfeatures/lang.feature\nfeatures/load.feature\nfeatures/multistep.feature\nfeatures/outline.feature\nfeatures/run.feature\nfeatures/snippets.feature\nfeatures/tags.feature",
                            "content_type": "",
                            "line": 10
                        },
                        "match": {
                            "location": "suite_context.go:348"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "load-features;load-a-specific-feature-file",
                "keyword": "Scenario",
                "name": "load a specific feature file",
                "description": "",
                "line": 24,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature path \"features/load.feature\"",
                        "line": 25,
                        "match": {
                            "location": "suite_context.go:324"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I parse features",
                        "line": 26,
                        "match": {
                            "location": "suite_context.go:329"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "I should have 1 feature file:",
                        "line": 27,
                        "doc_string": {
                            "value": "features/load.feature",
                            "content_type": "",
                            "line": 28
                        },
                        "match": {
                            "location": "suite_context.go:348"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "load-features;loaded-feature-should-have-a-number-of-scenarios;;2",
                "keyword": "Scenario Outline",
                "name": "loaded feature should have a number of scenarios",
                "description": "",
                "line": 39,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature path \"features/load.feature:3\"",
                        "line": 39,
                        "match": {
                            "location": "suite_context.go:324"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I parse features",
                        "line": 39,
                        "match": {
                            "location": "suite_context.go:329"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "I should have 0 scenario registered",
                        "line": 39,
                        "match": {
                            "location": "suite_context.go:390"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "load-features;loaded-feature-should-have-a-number-of-scenarios;;3",
                "keyword": "Scenario Outline",
                "name": "loaded feature should have a number of scenarios",
                "description": "",
                "line": 40,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature path \"features/load.feature:6\"",
                        "line": 40,
                        "match": {
                            "location": "suite_context.go:324"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I parse features",
                        "line": 40,
                        "match": {
                            "location": "suite_context.go:329"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "I should have 1 scenario registered",
                        "line": 40,
                        "match": {
                            "location": "suite_context.go:390"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "load-features;loaded-feature-should-have-a-number-of-scenarios;;4",
                "keyword": "Scenario Outline",
                "name": "loaded feature should have a number of scenarios",
                "description": "",
                "line": 41,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature path \"features/load.feature\"",
                        "line": 41,
                        "match": {
                            "location": "suite_context.go:324"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I parse features",
                        "line": 41,
                        "match": {
                            "location": "suite_context.go:329"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "I should have 4 scenario registered",
                        "line": 41,
                        "match": {
                            "location": "suite_context.go:390"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "load-features;load-a-number-of-feature-files",
                "keyword": "Scenario",
                "name": "load a number of feature files",
                "description": "",
                "line": 43,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature path \"features/load.feature\"",
                        "line": 44,
                        "match": {
                            "location": "suite_context.go:324"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "a feature path \"features/events.feature\"",
                        "line": 45,
                        "match": {
                            "location": "suite_context.go:324"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I parse features",
                        "line": 46,
                        "match": {
                            "location": "suite_context.go:329"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "I should have 2 feature files:",
                        "line": 47,
                        "doc_string": {
                            "value": "features/events.feature\nfeatures/load.feature",
                            "content_type": "",
                            "line": 48
                        },
                        "match": {
                            "location": "suite_context.go:348"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            }
        ]
    },
    {
        "uri": "features/multistep.feature",
        "id": "run-features-with-nested-steps",
        "keyword": "Feature",
        "name": "run features with nested steps",
        "description": "  In order to test multisteps\n  As a test suite\n  I need to be able to execute multisteps",
        "line": 1,
        "elements": [
            {
                "id": "run-features-with-nested-steps;should-run-passing-multistep-successfully",
                "keyword": "Scenario",
                "name": "should run passing multistep successfully",
                "description": "",
                "line": 6,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"normal.feature\" file:",
                        "line": 7,
                        "doc_string": {
                            "value": "Feature: normal feature\n\n  Scenario: run passing multistep\n    Given passing step\n    Then passing multistep",
                            "content_type": "",
                            "line": 8
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 15,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the suite should have passed",
                        "line": 16,
                        "match": {
                            "location": "suite_context.go:338"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following steps should be passed:",
                        "line": 17,
                        "doc_string": {
                            "value": "passing step\npassing multistep",
                            "content_type": "",
                            "line": 18
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "run-features-with-nested-steps;should-fail-multistep",
                "keyword": "Scenario",
                "name": "should fail multistep",
                "description": "",
                "line": 23,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"failed.feature\" file:",
                        "line": 24,
                        "doc_string": {
                            "value": "Feature: failed feature\n\n  Scenario: run failing multistep\n    Given passing step\n    When failing multistep\n    Then I should have 1 scenario registered",
                            "content_type": "",
                            "line": 25
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 33,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the suite should have failed",
                        "line": 34,
                        "match": {
                            "location": "suite_context.go:338"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following step should be failed:",
                        "line": 35,
                        "doc_string": {
                            "value": "failing multistep",
                            "content_type": "",
                            "line": 36
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following steps should be skipped:",
                        "line": 39,
                        "doc_string": {
                            "value": "I should have 1 scenario registered",
                            "content_type": "",
                            "line": 40
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following steps should be passed:",
                        "line": 43,
                        "doc_string": {
                            "value": "passing step",
                            "content_type": "",
                            "line": 44
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "run-features-with-nested-steps;should-fail-nested-multistep",
                "keyword": "Scenario",
                "name": "should fail nested multistep",
                "description": "",
                "line": 48,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"failed.feature\" file:",
                        "line": 49,
                        "doc_string": {
                            "value": "Feature: failed feature\n\n  Scenario: run failing nested multistep\n    Given failing nested multistep\n    When passing step",
                            "content_type": "",
                            "line": 50
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 57,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the suite should have failed",
                        "line": 58,
                        "match": {
                            "location": "suite_context.go:338"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following step should be failed:",
                        "line": 59,
                        "doc_string": {
                            "value": "failing nested multistep",
                            "content_type": "",
                            "line": 60
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following steps should be skipped:",
                        "line": 63,
                        "doc_string": {
                            "value": "passing step",
                            "content_type": "",
                            "line": 64
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "run-features-with-nested-steps;should-skip-steps-after-undefined-multistep",
                "keyword": "Scenario",
                "name": "should skip steps after undefined multistep",
                "description": "",
                "line": 68,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"undefined.feature\" file:",
                        "line": 69,
                        "doc_string": {
                            "value": "Feature: run undefined multistep\n\n  Scenario: run undefined multistep\n    Given passing step\n    When undefined multistep\n    Then passing multistep",
                            "content_type": "",
                            "line": 70
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 78,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the suite should have passed",
                        "line": 79,
                        "match": {
                            "location": "suite_context.go:338"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following step should be passed:",
                        "line": 80,
                        "doc_string": {
                            "value": "passing step",
                            "content_type": "",
                            "line": 81
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following step should be undefined:",
                        "line": 84,
                        "doc_string": {
                            "value": "undefined multistep",
                            "content_type": "",
                            "line": 85
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following step should be skipped:",
                        "line": 88,
                        "doc_string": {
                            "value": "passing multistep",
                            "content_type": "",
                            "line": 89
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "run-features-with-nested-steps;should-match-undefined-steps-in-a-row",
                "keyword": "Scenario",
                "name": "should match undefined steps in a row",
                "description": "",
                "line": 93,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"undefined.feature\" file:",
                        "line": 94,
                        "doc_string": {
                            "value": "Feature: undefined feature\n\n  Scenario: parse a scenario\n    Given undefined step\n    When undefined multistep\n    Then I should have 1 scenario registered",
                            "content_type": "",
                            "line": 95
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 103,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the suite should have passed",
                        "line": 104,
                        "match": {
                            "location": "suite_context.go:338"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following steps should be undefined:",
                        "line": 105,
                        "doc_string": {
                            "value": "undefined step\nundefined multistep",
                            "content_type": "",
                            "line": 106
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following step should be skipped:",
                        "line": 110,
                        "doc_string": {
                            "value": "I should have 1 scenario registered",
                            "content_type": "",
                            "line": 111
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "run-features-with-nested-steps;should-mark-undefined-steps-after-pending",
                "keyword": "Scenario",
                "name": "should mark undefined steps after pending",
                "description": "",
                "line": 115,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"pending.feature\" file:",
                        "line": 116,
                        "doc_string": {
                            "value": "Feature: pending feature\n\n  Scenario: parse a scenario\n    Given pending step\n    When undefined step\n    Then undefined multistep\n    And I should have 1 scenario registered",
                            "content_type": "",
                            "line": 117
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 126,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the suite should have passed",
                        "line": 127,
                        "match": {
                            "location": "suite_context.go:338"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following steps should be undefined:",
                        "line": 128,
                        "doc_string": {
                            "value": "undefined step\nundefined multistep",
                            "content_type": "",
                            "line": 129
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following step should be pending:",
                        "line": 133,
                        "doc_string": {
                            "value": "pending step",
                            "content_type": "",
                            "line": 134
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following step should be skipped:",
                        "line": 137,
                        "doc_string": {
                            "value": "I should have 1 scenario registered",
                            "content_type": "",
                            "line": 138
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            }
        ]
    },
    {
        "uri": "features/outline.feature",
        "id": "run-outline",
        "keyword": "Feature",
        "name": "run outline",
        "description": "  In order to test application behavior\n  As a test suite\n  I need to be able to run outline scenarios",
        "line": 1,
        "elements": [
            {
                "id": "run-outline;should-run-a-normal-outline",
                "keyword": "Scenario",
                "name": "should run a normal outline",
                "description": "",
                "line": 6,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"normal.feature\" file:",
                        "line": 7,
                        "doc_string": {
                            "value": "Feature: outline\n\n  Background:\n    Given passing step\n\n  Scenario Outline: parse a scenario\n    Given a feature path \"\u003cpath\u003e\"\n    When I parse features\n    Then I should have \u003cnum\u003e scenario registered\n\n    Examples:\n      | path                    | num |\n      | features/load.feature:6 | 1   |\n      | features/load.feature:3 | 0   |",
                            "content_type": "",
                            "line": 8
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 24,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the suite should have passed",
                        "line": 25,
                        "match": {
                            "location": "suite_context.go:338"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following steps should be passed:",
                        "line": 26,
                        "doc_string": {
                            "value": "a passing step\nI parse features\na feature path \"features/load.feature:6\"\na feature path \"features/load.feature:3\"\nI should have 1 scenario registered\nI should have 0 scenario registered",
                            "content_type": "",
                            "line": 27
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "run-outline;should-continue-through-examples-on-failure",
                "keyword": "Scenario",
                "name": "should continue through examples on failure",
                "description": "",
                "line": 36,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"normal.feature\" file:",
                        "line": 37,
                        "doc_string": {
                            "value": "Feature: outline\n\n  Background:\n    Given passing step\n\n  Scenario Outline: parse a scenario\n    Given a feature path \"\u003cpath\u003e\"\n    When I parse features\n    Then I should have \u003cnum\u003e scenario registered\n\n    Examples:\n      | path                    | num |\n      | features/load.feature:6 | 5   |\n      | features/load.feature:3 | 0   |",
                            "content_type": "",
                            "line": 38
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 54,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the suite should have failed",
                        "line": 55,
                        "match": {
                            "location": "suite_context.go:338"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following steps should be passed:",
                        "line": 56,
                        "doc_string": {
                            "value": "a passing step\nI parse features\na feature path \"features/load.feature:6\"\na feature path \"features/load.feature:3\"\nI should have 0 scenario registered",
                            "content_type": "",
                            "line": 57
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following steps should be failed:",
                        "line": 64,
                        "doc_string": {
                            "value": "I should have 5 scenario registered",
                            "content_type": "",
                            "line": 65
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "run-outline;should-skip-examples-on-background-failure",
                "keyword": "Scenario",
                "name": "should skip examples on background failure",
                "description": "",
                "line": 69,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"normal.feature\" file:",
                        "line": 70,
                        "doc_string": {
                            "value": "Feature: outline\n\n  Background:\n    Given a failing step\n\n  Scenario Outline: parse a scenario\n    Given a feature path \"\u003cpath\u003e\"\n    When I parse features\n    Then I should have \u003cnum\u003e scenario registered\n\n    Examples:\n      | path                    | num |\n      | features/load.feature:6 | 1   |\n      | features/load.feature:3 | 0   |",
                            "content_type": "",
                            "line": 71
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 87,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the suite should have failed",
                        "line": 88,
                        "match": {
                            "location": "suite_context.go:338"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following steps should be skipped:",
                        "line": 89,
                        "doc_string": {
                            "value": "I parse features\na feature path \"features/load.feature:6\"\na feature path \"features/load.feature:3\"\nI should have 0 scenario registered\nI should have 1 scenario registered",
                            "content_type": "",
                            "line": 90
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following steps should be failed:",
                        "line": 97,
                        "doc_string": {
                            "value": "a failing step",
                            "content_type": "",
                            "line": 98
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "run-outline;should-translate-step-table-body",
                "keyword": "Scenario",
                "name": "should translate step table body",
                "description": "",
                "line": 102,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"normal.feature\" file:",
                        "line": 103,
                        "doc_string": {
                            "value": "Feature: outline\n\n  Background:\n    Given I'm listening to suite events\n\n  Scenario Outline: run with events\n    Given a feature path \"\u003cpath\u003e\"\n    When I run feature suite\n    Then these events had to be fired for a number of times:\n      | BeforeScenario | \u003cscen\u003e |\n      | BeforeStep     | \u003cstep\u003e |\n\n    Examples:\n      | path                    | scen | step |\n      | features/load.feature:6 | 1    | 3    |\n      | features/load.feature   | 6    | 19   |",
                            "content_type": "",
                            "line": 104
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 122,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the suite should have passed",
                        "line": 123,
                        "match": {
                            "location": "suite_context.go:338"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following steps should be passed:",
                        "line": 124,
                        "doc_string": {
                            "value": "I'm listening to suite events\nI run feature suite\na feature path \"features/load.feature:6\"\na feature path \"features/load.feature\"",
                            "content_type": "",
                            "line": 125
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "run-outline;should-translate-step-doc-string-argument;;2",
                "keyword": "Scenario Outline",
                "name": "should translate step doc string argument",
                "description": "",
                "line": 151,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"normal.feature\" file:",
                        "line": 151,
                        "doc_string": {
                            "value": "Feature: scenario events\n\n  Background:\n    Given I'm listening to suite events\n\n  Scenario: run with events\n    Given a feature path \"features/load.feature:6\"\n    When I run feature suite\n    Then these events had to be fired for a number of times:\n      | BeforeScenario | 1 |",
                            "content_type": "",
                            "line": 134
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 151,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the suite should have passed",
                        "line": 151,
                        "match": {
                            "location": "suite_context.go:338"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "run-outline;should-translate-step-doc-string-argument;;3",
                "keyword": "Scenario Outline",
                "name": "should translate step doc string argument",
                "description": "",
                "line": 152,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"normal.feature\" file:",
                        "line": 152,
                        "doc_string": {
                            "value": "Feature: scenario events\n\n  Background:\n    Given I'm listening to suite events\n\n  Scenario: run with events\n    Given a feature path \"features/load.feature\"\n    When I run feature suite\n    Then these events had to be fired for a number of times:\n      | BeforeScenario | 6 |",
                            "content_type": "",
                            "line": 134
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 152,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the suite should have passed",
                        "line": 152,
                        "match": {
                            "location": "suite_context.go:338"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            }
        ]
    },
    {
        "uri": "features/run.feature",
        "id": "run-features",
        "keyword": "Feature",
        "name": "run features",
        "description": "  In order to test application behavior\n  As a test suite\n  I need to be able to run features",
        "line": 1,
        "elements": [
            {
                "id": "run-features;should-run-a-normal-feature",
                "keyword": "Scenario",
                "name": "should run a normal feature",
                "description": "",
                "line": 6,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"normal.feature\" file:",
                        "line": 7,
                        "doc_string": {
                            "value": "Feature: normal feature\n\n  Scenario: parse a scenario\n    Given a feature path \"features/load.feature:6\"\n    When I parse features\n    Then I should have 1 scenario registered",
                            "content_type": "",
                            "line": 8
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 16,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the suite should have passed",
                        "line": 17,
                        "match": {
                            "location": "suite_context.go:338"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following steps should be passed:",
                        "line": 18,
                        "doc_string": {
                            "value": "a feature path \"features/load.feature:6\"\nI parse features\nI should have 1 scenario registered",
                            "content_type": "",
                            "line": 19
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "run-features;should-skip-steps-after-failure",
                "keyword": "Scenario",
                "name": "should skip steps after failure",
                "description": "",
                "line": 25,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"failed.feature\" file:",
                        "line": 26,
                        "doc_string": {
                            "value": "Feature: failed feature\n\n  Scenario: parse a scenario\n    Given a failing step\n    When I parse features\n    Then I should have 1 scenario registered",
                            "content_type": "",
                            "line": 27
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 35,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the suite should have failed",
                        "line": 36,
                        "match": {
                            "location": "suite_context.go:338"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following step should be failed:",
                        "line": 37,
                        "doc_string": {
                            "value": "a failing step",
                            "content_type": "",
                            "line": 38
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following steps should be skipped:",
                        "line": 41,
                        "doc_string": {
                            "value": "I parse features\nI should have 1 scenario registered",
                            "content_type": "",
                            "line": 42
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "run-features;should-skip-all-scenarios-if-background-fails",
                "keyword": "Scenario",
                "name": "should skip all scenarios if background fails",
                "description": "",
                "line": 47,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"failed.feature\" file:",
                        "line": 48,
                        "doc_string": {
                            "value": "Feature: failed feature\n\n  Background:\n    Given a failing step\n\n  Scenario: parse a scenario\n    Given a feature path \"features/load.feature:6\"\n    When I parse features\n    Then I should have 1 scenario registered",
                            "content_type": "",
                            "line": 49
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 60,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the suite should have failed",
                        "line": 61,
                        "match": {
                            "location": "suite_context.go:338"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following step should be failed:",
                        "line": 62,
                        "doc_string": {
                            "value": "a failing step",
                            "content_type": "",
                            "line": 63
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following steps should be skipped:",
                        "line": 66,
                        "doc_string": {
                            "value": "a feature path \"features/load.feature:6\"\nI parse features\nI should have 1 scenario registered",
                            "content_type": "",
                            "line": 67
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "run-features;should-skip-steps-after-undefined",
                "keyword": "Scenario",
                "name": "should skip steps after undefined",
                "description": "",
                "line": 73,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"undefined.feature\" file:",
                        "line": 74,
                        "doc_string": {
                            "value": "Feature: undefined feature\n\n  Scenario: parse a scenario\n    Given a feature path \"features/load.feature:6\"\n    When undefined action\n    Then I should have 1 scenario registered",
                            "content_type": "",
                            "line": 75
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 83,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the suite should have passed",
                        "line": 84,
                        "match": {
                            "location": "suite_context.go:338"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following step should be passed:",
                        "line": 85,
                        "doc_string": {
                            "value": "a feature path \"features/load.feature:6\"",
                            "content_type": "",
                            "line": 86
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following step should be undefined:",
                        "line": 89,
                        "doc_string": {
                            "value": "undefined action",
                            "content_type": "",
                            "line": 90
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following step should be skipped:",
                        "line": 93,
                        "doc_string": {
                            "value": "I should have 1 scenario registered",
                            "content_type": "",
                            "line": 94
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "run-features;should-match-undefined-steps-in-a-row",
                "keyword": "Scenario",
                "name": "should match undefined steps in a row",
                "description": "",
                "line": 98,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"undefined.feature\" file:",
                        "line": 99,
                        "doc_string": {
                            "value": "Feature: undefined feature\n\n  Scenario: parse a scenario\n    Given undefined step\n    When undefined action\n    Then I should have 1 scenario registered",
                            "content_type": "",
                            "line": 100
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 108,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the suite should have passed",
                        "line": 109,
                        "match": {
                            "location": "suite_context.go:338"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following steps should be undefined:",
                        "line": 110,
                        "doc_string": {
                            "value": "undefined step\nundefined action",
                            "content_type": "",
                            "line": 111
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following step should be skipped:",
                        "line": 115,
                        "doc_string": {
                            "value": "I should have 1 scenario registered",
                            "content_type": "",
                            "line": 116
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "run-features;should-skip-steps-on-pending",
                "keyword": "Scenario",
                "name": "should skip steps on pending",
                "description": "",
                "line": 120,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"pending.feature\" file:",
                        "line": 121,
                        "doc_string": {
                            "value": "Feature: pending feature\n\n  Scenario: parse a scenario\n    Given undefined step\n    When pending step\n    Then I should have 1 scenario registered",
                            "content_type": "",
                            "line": 122
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 130,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the suite should have passed",
                        "line": 131,
                        "match": {
                            "location": "suite_context.go:338"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following step should be undefined:",
                        "line": 132,
                        "doc_string": {
                            "value": "undefined step",
                            "content_type": "",
                            "line": 133
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following step should be skipped:",
                        "line": 136,
                        "doc_string": {
                            "value": "pending step\nI should have 1 scenario registered",
                            "content_type": "",
                            "line": 137
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "run-features;should-handle-pending-step",
                "keyword": "Scenario",
                "name": "should handle pending step",
                "description": "",
                "line": 142,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"pending.feature\" file:",
                        "line": 143,
                        "doc_string": {
                            "value": "Feature: pending feature\n\n  Scenario: parse a scenario\n    Given a feature path \"features/load.feature:6\"\n    When pending step\n    Then I should have 1 scenario registered",
                            "content_type": "",
                            "line": 144
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 152,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the suite should have passed",
                        "line": 153,
                        "match": {
                            "location": "suite_context.go:338"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following step should be passed:",
                        "line": 154,
                        "doc_string": {
                            "value": "a feature path \"features/load.feature:6\"",
                            "content_type": "",
                            "line": 155
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following step should be pending:",
                        "line": 158,
                        "doc_string": {
                            "value": "pending step",
                            "content_type": "",
                            "line": 159
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following step should be skipped:",
                        "line": 162,
                        "doc_string": {
                            "value": "I should have 1 scenario registered",
                            "content_type": "",
                            "line": 163
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "run-features;should-mark-undefined-steps-after-pending",
                "keyword": "Scenario",
                "name": "should mark undefined steps after pending",
                "description": "",
                "line": 167,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"pending.feature\" file:",
                        "line": 168,
                        "doc_string": {
                            "value": "Feature: pending feature\n\n  Scenario: parse a scenario\n    Given pending step\n    When undefined\n    Then undefined 2\n    And I should have 1 scenario registered",
                            "content_type": "",
                            "line": 169
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 178,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the suite should have passed",
                        "line": 179,
                        "match": {
                            "location": "suite_context.go:338"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following steps should be undefined:",
                        "line": 180,
                        "doc_string": {
                            "value": "undefined\nundefined 2",
                            "content_type": "",
                            "line": 181
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following step should be pending:",
                        "line": 185,
                        "doc_string": {
                            "value": "pending step",
                            "content_type": "",
                            "line": 186
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following step should be skipped:",
                        "line": 189,
                        "doc_string": {
                            "value": "I should have 1 scenario registered",
                            "content_type": "",
                            "line": 190
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "run-features;should-fail-suite-if-undefined-steps-follow-after-the-failure",
                "keyword": "Scenario",
                "name": "should fail suite if undefined steps follow after the failure",
                "description": "",
                "line": 194,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"failed.feature\" file:",
                        "line": 195,
                        "doc_string": {
                            "value": "Feature: failed feature\n\n  Scenario: parse a scenario\n    Given a failing step\n    When an undefined step\n    Then another undefined step",
                            "content_type": "",
                            "line": 196
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 204,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the following step should be failed:",
                        "line": 205,
                        "doc_string": {
                            "value": "a failing step",
                            "content_type": "",
                            "line": 206
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following steps should be undefined:",
                        "line": 209,
                        "doc_string": {
                            "value": "an undefined step\nanother undefined step",
                            "content_type": "",
                            "line": 210
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the suite should have failed",
                        "line": 214,
                        "match": {
                            "location": "suite_context.go:338"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "run-features;should-fail-suite-and-skip-pending-step-after-failed-step",
                "keyword": "Scenario",
                "name": "should fail suite and skip pending step after failed step",
                "description": "",
                "line": 216,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"failed.feature\" file:",
                        "line": 217,
                        "doc_string": {
                            "value": "Feature: failed feature\n\n  Scenario: parse a scenario\n    Given a failing step\n    When pending step\n    Then another undefined step",
                            "content_type": "",
                            "line": 218
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 226,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the following step should be failed:",
                        "line": 227,
                        "doc_string": {
                            "value": "a failing step",
                            "content_type": "",
                            "line": 228
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following steps should be skipped:",
                        "line": 231,
                        "doc_string": {
                            "value": "pending step",
                            "content_type": "",
                            "line": 232
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following steps should be undefined:",
                        "line": 235,
                        "doc_string": {
                            "value": "another undefined step",
                            "content_type": "",
                            "line": 236
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the suite should have failed",
                        "line": 239,
                        "match": {
                            "location": "suite_context.go:338"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "run-features;should-fail-suite-and-skip-next-step-after-failed-step",
                "keyword": "Scenario",
                "name": "should fail suite and skip next step after failed step",
                "description": "",
                "line": 241,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"failed.feature\" file:",
                        "line": 242,
                        "doc_string": {
                            "value": "Feature: failed feature\n\n  Scenario: parse a scenario\n    Given a failing step\n    When a failing step\n    Then another undefined step",
                            "content_type": "",
                            "line": 243
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 251,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the following step should be failed:",
                        "line": 252,
                        "doc_string": {
                            "value": "a failing step",
                            "content_type": "",
                            "line": 253
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following steps should be skipped:",
                        "line": 256,
                        "doc_string": {
                            "value": "a failing step",
                            "content_type": "",
                            "line": 257
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following steps should be undefined:",
                        "line": 260,
                        "doc_string": {
                            "value": "another undefined step",
                            "content_type": "",
                            "line": 261
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the suite should have failed",
                        "line": 264,
                        "match": {
                            "location": "suite_context.go:338"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            }
        ]
    },
    {
        "uri": "features/snippets.feature",
        "id": "undefined-step-snippets",
        "keyword": "Feature",
        "name": "undefined step snippets",
        "description": "  In order to implement step definitions faster\n  As a test suite user\n  I need to be able to get undefined step snippets",
        "line": 1,
        "elements": [
            {
                "id": "undefined-step-snippets;should-generate-snippets",
                "keyword": "Scenario",
                "name": "should generate snippets",
                "description": "",
                "line": 6,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"undefined.feature\" file:",
                        "line": 7,
                        "doc_string": {
                            "value": "Feature: undefined steps\n\n  Scenario: get version number from api\n    When I send \"GET\" request to \"/version\"\n    Then the response code should be 200",
                            "content_type": "",
                            "line": 8
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 15,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the following steps should be undefined:",
                        "line": 16,
                        "doc_string": {
                            "value": "I send \"GET\" request to \"/version\"\nthe response code should be 200",
                            "content_type": "",
                            "line": 17
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the undefined step snippets should be:",
                        "line": 21,
                        "doc_string": {
                            "value": "func iSendRequestTo(arg1, arg2 string) error {\n        return godog.ErrPending\n}\n\nfunc theResponseCodeShouldBe(arg1 int) error {\n        return godog.ErrPending\n}\n\nfunc FeatureContext(s *godog.Suite) {\n        s.Step(` + "`" + `^I send \"([^\"]*)\" request to \"([^\"]*)\"$` + "`" + `, iSendRequestTo)\n        s.Step(` + "`" + `^the response code should be (\\d+)$` + "`" + `, theResponseCodeShouldBe)\n}",
                            "content_type": "",
                            "line": 22
                        },
                        "match": {
                            "location": "suite_context.go:178"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "undefined-step-snippets;should-generate-snippets-with-more-arguments",
                "keyword": "Scenario",
                "name": "should generate snippets with more arguments",
                "description": "",
                "line": 37,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"undefined.feature\" file:",
                        "line": 38,
                        "doc_string": {
                            "value": "Feature: undefined steps\n\n  Scenario: get version number from api\n    When I send \"GET\" request to \"/version\" with:\n      | col1 | val1 |\n      | col2 | val2 |\n    Then the response code should be 200 and header \"X-Powered-By\" should be \"godog\"",
                            "content_type": "",
                            "line": 39
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 48,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the undefined step snippets should be:",
                        "line": 49,
                        "doc_string": {
                            "value": "func iSendRequestToWith(arg1, arg2 string, arg3 *gherkin.DataTable) error {\n        return godog.ErrPending\n}\n\nfunc theResponseCodeShouldBeAndHeaderShouldBe(arg1 int, arg2, arg3 string) error {\n        return godog.ErrPending\n}\n\nfunc FeatureContext(s *godog.Suite) {\n        s.Step(` + "`" + `^I send \"([^\"]*)\" request to \"([^\"]*)\" with:$` + "`" + `, iSendRequestToWith)\n        s.Step(` + "`" + `^the response code should be (\\d+) and header \"([^\"]*)\" should be \"([^\"]*)\"$` + "`" + `, theResponseCodeShouldBeAndHeaderShouldBe)\n}",
                            "content_type": "",
                            "line": 50
                        },
                        "match": {
                            "location": "suite_context.go:178"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "undefined-step-snippets;should-handle-escaped-symbols",
                "keyword": "Scenario",
                "name": "should handle escaped symbols",
                "description": "",
                "line": 65,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"undefined.feature\" file:",
                        "line": 66,
                        "doc_string": {
                            "value": "Feature: undefined steps\n\n  Scenario: get version number from api\n    When I pull from github.com\n    Then the project should be there",
                            "content_type": "",
                            "line": 67
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 74,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the following steps should be undefined:",
                        "line": 75,
                        "doc_string": {
                            "value": "I pull from github.com\nthe project should be there",
                            "content_type": "",
                            "line": 76
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the undefined step snippets should be:",
                        "line": 80,
                        "doc_string": {
                            "value": "func iPullFromGithubcom() error {\n        return godog.ErrPending\n}\n\nfunc theProjectShouldBeThere() error {\n        return godog.ErrPending\n}\n\nfunc FeatureContext(s *godog.Suite) {\n        s.Step(` + "`" + `^I pull from github\\.com$` + "`" + `, iPullFromGithubcom)\n        s.Step(` + "`" + `^the project should be there$` + "`" + `, theProjectShouldBeThere)\n}",
                            "content_type": "",
                            "line": 81
                        },
                        "match": {
                            "location": "suite_context.go:178"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "undefined-step-snippets;should-handle-string-argument-followed-by-comma",
                "keyword": "Scenario",
                "name": "should handle string argument followed by comma",
                "description": "",
                "line": 96,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"undefined.feature\" file:",
                        "line": 97,
                        "doc_string": {
                            "value": "Feature: undefined\n\n  Scenario: add item to basket\n    Given there is a \"Sith Lord Lightsaber\", which costs £5\n    When I add the \"Sith Lord Lightsaber\" to the basket",
                            "content_type": "",
                            "line": 98
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 105,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the undefined step snippets should be:",
                        "line": 106,
                        "doc_string": {
                            "value": "func thereIsAWhichCosts(arg1 string, arg2 int) error {\n        return godog.ErrPending\n}\n\nfunc iAddTheToTheBasket(arg1 string) error {\n        return godog.ErrPending\n}\n\nfunc FeatureContext(s *godog.Suite) {\n        s.Step(` + "`" + `^there is a \"([^\"]*)\", which costs £(\\d+)$` + "`" + `, thereIsAWhichCosts)\n        s.Step(` + "`" + `^I add the \"([^\"]*)\" to the basket$` + "`" + `, iAddTheToTheBasket)\n}",
                            "content_type": "",
                            "line": 107
                        },
                        "match": {
                            "location": "suite_context.go:178"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "undefined-step-snippets;should-handle-arguments-in-the-beggining-or-end-of-the-step",
                "keyword": "Scenario",
                "name": "should handle arguments in the beggining or end of the step",
                "description": "",
                "line": 122,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"undefined.feature\" file:",
                        "line": 123,
                        "doc_string": {
                            "value": "Feature: undefined\n\n  Scenario: add item to basket\n    Given \"Sith Lord Lightsaber\", which costs £5\n    And 12 godogs",
                            "content_type": "",
                            "line": 124
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite",
                        "line": 131,
                        "match": {
                            "location": "suite_context.go:379"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the undefined step snippets should be:",
                        "line": 132,
                        "doc_string": {
                            "value": "func whichCosts(arg1 string, arg2 int) error {\n        return godog.ErrPending\n}\n\nfunc godogs(arg1 int) error {\n        return godog.ErrPending\n}\n\nfunc FeatureContext(s *godog.Suite) {\n        s.Step(` + "`" + `^\"([^\"]*)\", which costs £(\\d+)$` + "`" + `, whichCosts)\n        s.Step(` + "`" + `^(\\d+) godogs$` + "`" + `, godogs)\n}",
                            "content_type": "",
                            "line": 133
                        },
                        "match": {
                            "location": "suite_context.go:178"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            }
        ]
    },
    {
        "uri": "features/tags.feature",
        "id": "tag-filters",
        "keyword": "Feature",
        "name": "tag filters",
        "description": "  In order to test application behavior\n  As a test suite\n  I need to be able to filter features and scenarios by tags",
        "line": 1,
        "elements": [
            {
                "id": "tag-filters;should-filter-outline-examples-by-tags",
                "keyword": "Scenario",
                "name": "should filter outline examples by tags",
                "description": "",
                "line": 6,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"normal.feature\" file:",
                        "line": 7,
                        "doc_string": {
                            "value": "Feature: outline\n\n  Background:\n    Given passing step\n\n  Scenario Outline: parse a scenario\n    Given a feature path \"\u003cpath\u003e\"\n    When I parse features\n    Then I should have \u003cnum\u003e scenario registered\n\n    Examples:\n      | path                    | num |\n      | features/load.feature:3 | 0   |\n\n    @used\n    Examples:\n      | path                    | num |\n      | features/load.feature:6 | 1   |",
                            "content_type": "",
                            "line": 8
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite with tags \"@used\"",
                        "line": 28,
                        "match": {
                            "location": "suite_context.go:118"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the suite should have passed",
                        "line": 29,
                        "match": {
                            "location": "suite_context.go:338"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following steps should be passed:",
                        "line": 30,
                        "doc_string": {
                            "value": "I parse features\na feature path \"features/load.feature:6\"\nI should have 1 scenario registered",
                            "content_type": "",
                            "line": 31
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "I should have 1 scenario registered",
                        "line": 36,
                        "match": {
                            "location": "suite_context.go:390"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "tag-filters;should-filter-scenarios-by-x-tag",
                "keyword": "Scenario",
                "name": "should filter scenarios by X tag",
                "description": "",
                "line": 38,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"normal.feature\" file:",
                        "line": 39,
                        "doc_string": {
                            "value": "Feature: tagged\n\n  @x\n  Scenario: one\n    Given a feature path \"one\"\n\n  @x\n  Scenario: two\n    Given a feature path \"two\"\n\n  @x @y\n  Scenario: three\n    Given a feature path \"three\"\n\n  @y\n  Scenario: four\n    Given a feature path \"four\"",
                            "content_type": "",
                            "line": 40
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite with tags \"@x\"",
                        "line": 59,
                        "match": {
                            "location": "suite_context.go:118"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the suite should have passed",
                        "line": 60,
                        "match": {
                            "location": "suite_context.go:338"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "I should have 3 scenario registered",
                        "line": 61,
                        "match": {
                            "location": "suite_context.go:390"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following steps should be passed:",
                        "line": 62,
                        "doc_string": {
                            "value": "a feature path \"one\"\na feature path \"two\"\na feature path \"three\"",
                            "content_type": "",
                            "line": 63
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "tag-filters;should-filter-scenarios-by-x-tag-not-having-y",
                "keyword": "Scenario",
                "name": "should filter scenarios by X tag not having Y",
                "description": "",
                "line": 69,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"normal.feature\" file:",
                        "line": 70,
                        "doc_string": {
                            "value": "Feature: tagged\n\n  @x\n  Scenario: one\n    Given a feature path \"one\"\n\n  @x\n  Scenario: two\n    Given a feature path \"two\"\n\n  @x @y\n  Scenario: three\n    Given a feature path \"three\"\n\n  @y @z\n  Scenario: four\n    Given a feature path \"four\"",
                            "content_type": "",
                            "line": 71
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite with tags \"@x \u0026\u0026 ~@y\"",
                        "line": 90,
                        "match": {
                            "location": "suite_context.go:118"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the suite should have passed",
                        "line": 91,
                        "match": {
                            "location": "suite_context.go:338"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "I should have 2 scenario registered",
                        "line": 92,
                        "match": {
                            "location": "suite_context.go:390"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following steps should be passed:",
                        "line": 93,
                        "doc_string": {
                            "value": "a feature path \"one\"\na feature path \"two\"",
                            "content_type": "",
                            "line": 94
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            },
            {
                "id": "tag-filters;should-filter-scenarios-having-y-and-z-tags",
                "keyword": "Scenario",
                "name": "should filter scenarios having Y and Z tags",
                "description": "",
                "line": 99,
                "type": "scenario",
                "steps": [
                    {
                        "keyword": "Given ",
                        "name": "a feature \"normal.feature\" file:",
                        "line": 100,
                        "doc_string": {
                            "value": "Feature: tagged\n\n  @x\n  Scenario: one\n    Given a feature path \"one\"\n\n  @x\n  Scenario: two\n    Given a feature path \"two\"\n\n  @x @y\n  Scenario: three\n    Given a feature path \"three\"\n\n  @y @z\n  Scenario: four\n    Given a feature path \"four\"",
                            "content_type": "",
                            "line": 101
                        },
                        "match": {
                            "location": "suite_context.go:318"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "When ",
                        "name": "I run feature suite with tags \"@y \u0026\u0026 @z\"",
                        "line": 120,
                        "match": {
                            "location": "suite_context.go:118"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "Then ",
                        "name": "the suite should have passed",
                        "line": 121,
                        "match": {
                            "location": "suite_context.go:338"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "I should have 1 scenario registered",
                        "line": 122,
                        "match": {
                            "location": "suite_context.go:390"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    },
                    {
                        "keyword": "And ",
                        "name": "the following steps should be passed:",
                        "line": 123,
                        "doc_string": {
                            "value": "a feature path \"four\"",
                            "content_type": "",
                            "line": 124
                        },
                        "match": {
                            "location": "suite_context.go:191"
                        },
                        "result": {
                            "status": "passed",
                            "duration": 0
                        }
                    }
                ]
            }
        ]
    }
]`

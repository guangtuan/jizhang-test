Feature: subject

  Scenario:
    When login with email "grant" and password "123456"
    Then the response code should be 200
    And the response body should be a json string with a 32-bit token

  Scenario:
    When create a big-categories subject with name "eat"
    Then the response code should be 200

  Scenario:
    When create a small-categories subject with name "breakfest" under "eat"
    Then the response code should be 200

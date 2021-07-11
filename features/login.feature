Feature: login

  Scenario:
    When login with email "grant" and password "123456"
    Then the response code should be 200
    And the response body should be a json string with a 32-bit token
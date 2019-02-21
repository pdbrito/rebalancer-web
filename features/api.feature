Feature: Serving the rebalancer api

  Background:
    Given the api server is running

  Scenario: Checking the health of the api
    When I make a "GET" request to "/healthcheck"
    Then the response code should be 200
    And the response should match json:
    """
    {
      "alive":true
    }
    """
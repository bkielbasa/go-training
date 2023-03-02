Feature: Scaling up and down of the number of nodes in the cluster

    Scenario: Scaling up the number of nodes in the cluster
      Given I have a cluster with min nodes 3 and max nodes 6
      When I scale up the cluster by 2 nodes
      Then the cluster should have 5 nodes
      When I scale up the cluster by 3 nodes
      Then the cluster should have 6 nodes


    Scenario: Scaling down the number of nodes in the cluster
      Given I have a cluster with min nodes 3 and max nodes 6
      When I scale down the cluster by 2 nodes
      Then the cluster should have 3 nodes

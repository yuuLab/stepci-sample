# Step CI Examples

1. git clone

      ```
      git clone https://yuulab@bitbucket.org/yuulab0118/stepci-test.git
      ```

      ```
      cd stepci-test
      ```

1. run stepci

      ```
      docker run -v "$(pwd)"/tests:/tests ghcr.io/stepci/stepci /tests/workflow.yml -e HOST=host.docker.internal:8080
      ```

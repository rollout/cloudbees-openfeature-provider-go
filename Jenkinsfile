pipeline {
  agent none

  options {
    timeout(time: 5, unit: 'MINUTES')
  }
  stages {
    stage("Build code") {
      agent {
        kubernetes {
          label 'build-' + UUID.randomUUID().toString()
          inheritFrom 'default'
          yamlFile './cbci-templates/fmforci.yaml'
        }
      }
      steps{
        container(name: "build", shell: 'sh') {
              sh script : """
                make lint
                make build
                make test-junit
              """, label: "make lint, test"
        } //end container
      } //end step
      post {
        always {
          junit checksName: 'Go Tests', testResults: 'junit.xml'
        }
      }
    } // end stage("End Build code")
  } // end stages
} // end pipeline

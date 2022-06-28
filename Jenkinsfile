pipeline {
  agent {
    node {
      label 'context'
    }
  }
    stages {
        stage('Vet') {
              steps {
                gitlabCommitStatus('Vet') {
                  ansiColor('xterm') {
                    sh 'make vet'
                  }
                }
              }
            }

            stage('Build') {
              steps {
                 ('Build') {
                  ansiColor('xterm') {
                    sh 'make build'
                  }
                }
              }
            }

            stage('Test') {
              steps {
                gitlabCommitStatus('Test') {
                  ansiColor('xterm') {
                    sh 'make test-junit'
                    junit(allowEmptyResults: false, testResults: "junit.xml")
                  }
                }
              }
            }
    }
}
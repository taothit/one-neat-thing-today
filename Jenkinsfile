pipeline {
    agent {
        docker {
            image 'golang:latest'
            args '-v $HOME/.cache:/.cache:rw' 
        }
    }
    options {
        skipStagesAfterUnstable()
    }
    stages {
        stage("Prepare Go modules cache") {
            steps {
                sh 'go mod vendor'
            }
        }
        stage('Test') {
            steps {
                sh 'go test -v .'
            }
        }
        stage('Build') { 
            steps {
                sh 'go build .' 
            }
        }
    }
}
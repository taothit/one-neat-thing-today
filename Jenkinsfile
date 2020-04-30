pipeline {
    agent {
        docker {
            image 'golang:latest' 
            args '-v /root/.m2:/root/.m2' 
        }
    }
    stages {
        stage('Build') { 
            steps {
                sh 'go install .' 
            }
        }
    }
}
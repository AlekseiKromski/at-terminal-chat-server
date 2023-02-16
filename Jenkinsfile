pipeline {
    agent any
    stages {
        stage('build project') {
            steps {
                sh "go build"
            }
        }
        stage('deploy') {
            steps {
                sh "pm2 start"
            }
        }
    }
}

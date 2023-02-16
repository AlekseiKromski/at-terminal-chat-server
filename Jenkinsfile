pipeline {
    agent any
    tools { go 'go1.19' }
    stages {
        stage('build project') {
            steps {
                sh "go build"
            }
        }
        stage('deploy') {
            steps {
                sh "pm2 start ecosystem.config.json"
            }
        }
    }
}

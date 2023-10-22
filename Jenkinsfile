pipeline {
    agent any
    stages {
        stage('test') {
            steps {
                sh "docker build -t docker.alekseikromski.com/at-terminal-chat-server:latest"
            }
        }
    }
}

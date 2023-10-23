pipeline {
    agent any
    stages {
        stage('build') {
            steps {
                sh "docker build -t docker.alekseikromski.com/at-terminal-chat-server:latest ."
            }
        }
        stage('push') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'dockerHub', passwordVariable: 'dockerHubPassword', usernameVariable: 'dockerHubUser')]) {
                    sh "docker login https://docker.alekseikromski.com -u ${env.dockerHubUser} -p ${env.dockerHubPassword}"
                    sh "docker push docker.alekseikromski.com/at-terminal-chat-server:latest"
                }
            }
        }
        stage('Apply Kubernetes files') {
            steps {
                withKubeConfig([credentialsId: 'kubernetes', serverUrl: 'http://192.168.0.108:8001']) {
                  sh 'kubectl apply -f ./k8s/deploy.yml --force'
                }
            }
        }
    }
}

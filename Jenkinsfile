pipeline {
    agent any

    environment {
        // Define Docker Hub credentials and repository details
        DOCKERHUB_CREDENTIALS = credentials('76a0702f-d9c7-46ae-973e-c9cbe932710d')
        DOCKERHUB_REPO = 'mooka95/cloud-go'
        COMPOSE_PROJECT_NAME = 'app'
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        
        stage('Build Docker Image') {
            steps {
                script {
                    def currentTag = "v18"
                    def newTag = currentTag // Change this logic if you need a different tag

                    withCredentials([usernamePassword(credentialsId: 'dockerhub-credentials', passwordVariable: 'DOCKERHUB_CREDENTIALS_PSW', usernameVariable: 'DOCKERHUB_CREDENTIALS_USR')]) {
                        sh "docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin <<< $DOCKERHUB_CREDENTIALS_PSW"
                        sh 'curl -s -u $DOCKERHUB_CREDENTIALS_USR:$DOCKERHUB_CREDENTIALS_PSW https://hub.docker.com/v2/repositories/mooka95/cloud-go/tags/?page_size=1 | jq -r .results[0].name'

                        sh "docker build -t mooka95/cloud-go:${newTag} ."
                        sh "docker push mooka95/cloud-go:${newTag}"
                    }
                    
                    sh "sed -i 's|\${TAG}|${newTag}|g' docker-compose.yaml"
                }
            }
        }
        
        stage('Deploy') {
            steps {
                script {
                    echo "Deploying application..."
                    sh "docker-compose down"
                    sh "docker-compose up -d"
                }
            }
        }

        stage('Get Running Containers') {
            steps {
                script {
                    def containers = sh(script: "docker ps", returnStdout: true)
                    echo "Running containers: ${containers}"
                }
            }
        }
    }

    post {
        always {
            echo 'Pipeline finished.'
        }
        failure {
            echo 'Pipeline failed!'
        }
    }
}

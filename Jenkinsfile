pipeline {
    agent any

    environment {
        DOCKERHUB_CREDENTIALS_USR = credentials('76a0702f-d9c7-46ae-973e-c9cbe932710d').username
        DOCKERHUB_CREDENTIALS_PSW = credentials('76a0702f-d9c7-46ae-973e-c9cbe932710d').password
        DOCKERHUB_REPO = 'mooka95/cloud-go'
        COMPOSE_PROJECT_NAME = 'app'
    }

    stages {
        stage('Checkout') {
            steps {
                // Checkout code from the repository
                checkout scm
            }
        }

        stage('Build and Push Docker Image') {
            steps {
                script {
                    // Login to Docker Hub
                    sh "echo ${DOCKERHUB_CREDENTIALS_PSW} | docker login -u ${DOCKERHUB_CREDENTIALS_USR} --password-stdin"
                    
                    // Determine the new tag by incrementing the latest tag or starting from v1
                    def latestTag = sh(
                        script: "curl -s -u ${DOCKERHUB_CREDENTIALS_USR}:${DOCKERHUB_CREDENTIALS_PSW} https://hub.docker.com/v2/repositories/${DOCKERHUB_REPO}/tags/?page_size=1 | jq -r '.results[0].name'", 
                        returnStdout: true
                    ).trim()
                    
                    def newTag = 'v1'
                    if (latestTag =~ /^v\d+$/) {
                        newTag = 'v' + (latestTag.replace('v', '').toInteger() + 1)
                    }

                    // Build the Docker image with the new tag
                    sh "docker build -t ${DOCKERHUB_REPO}:${newTag} ."
                    
                    // Push the Docker image to Docker Hub
                    sh "docker push ${DOCKERHUB_REPO}:${newTag}"
                }
            }
        }

        stage('Deploy') {
            steps {
                echo 'Deploying application...'
                // Ensure Docker pulls the latest image before deploying
                sh "docker-compose down && docker-compose pull && docker-compose up -d"
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
        success {
            echo 'Pipeline succeeded!'
        }
    }
}

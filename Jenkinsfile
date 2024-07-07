pipeline {
    agent any
    
    environment {
         PATH = "/usr/local/go/bin:$PATH"
        COMPOSE_PROJECT_NAME = 'app'
        def version = 1
        GIT_COMMIT_SHA = sh(script: 'git rev-parse --short HEAD', returnStdout: true).trim()
        IMAGE_TAG = "mooka95/cloud-go"
    }
    
    stages {
        stage('Checkout') {
            steps {
                // Checkout the repository from GitHub
                checkout scm
            }
        }
        stage('Build') {
            steps {
                    sh 'go version'
                // Build your Go application
                sh "go build -o app"
            }
        }
        stage('Build Docker Image') {
            steps {
                // Build Docker image
                script {
                                        // Build the Docker image
                    sh "docker build -t ${IMAGE_TAG}:${version} ."
                    
                    // Tag the Docker image with current version
                    sh "docker tag ${IMAGE_TAG}:${version} ${IMAGE_TAG}:v${version}"
                    
                    // Push the Docker image to Docker Hub
                    sh "docker push ${IMAGE_TAG}:${version}"
                    sh "docker push ${IMAGE_TAG}:v${version}"
                    
                    // Increment version for next build
                    version++
                }
            }
        }
        stage('Deploy') {
            steps {
                // Deploy with Docker Compose or any deployment strategy
                sh "docker-compose -f docker-compose.yaml up -d"
            }
        }
    }
    
    post {
        success {
            echo 'Deployment successful!'
        }
        failure {
            echo 'Deployment failed!'
        }
    }
}

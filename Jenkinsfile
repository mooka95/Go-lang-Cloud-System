pipeline {
    agent any
    
    environment {
         PATH = "/usr/local/go/bin:$PATH"
        COMPOSE_PROJECT_NAME = 'app'
        VERSION_MAJOR = 1
        VERSION_MINOR = 0
        VERSION_PATCH = 0
        GIT_COMMIT_SHA = sh(script: 'git rev-parse --short HEAD', returnStdout: true).trim()
        IMAGE_TAG = "mooka95/cloud-go:${VERSION_MAJOR}.${VERSION_MINOR}.${VERSION_PATCH}-${GIT_COMMIT_SHA}"
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
        stage('Test') {
            steps {
                // Run tests for your application
                sh "make test" // Example test command, adjust as per your testing process
            }
        }
        stage('Build Docker Image') {
            steps {
                // Build Docker image
                script {
                    sh "docker build -t ${IMAGE_TAG} ."
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

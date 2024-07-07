pipeline {
    agent any
    
    environment {
        COMPOSE_PROJECT_NAME = 'GoCloud'
        IMAGE_TAG = mooka95/cloud-go"
    }
    
    stages {
        stage('Build and Test') {
            steps {
                // Clone the repository
                checkout scm
                
                // Build and start services defined in docker-compose.yml
                script {
                    sh 'docker-compose -f docker-compose.yml build'
                    sh 'docker-compose -f docker-compose.yml up -d'
                }
            }
        }
        stage('Deploy') {
            steps {
                // Deploy the application using Docker Compose
                script {
                    sh 'docker-compose -f docker-compose.yaml down'
                    sh 'docker-compose -f docker-compose.yaml up -d'
                }
            }
        }
    }
    post {
        success {
            echo 'Pipeline succeeded! Clean up...'
            // Clean up containers and networks
            sh 'docker-compose -f docker-compose.yml down'
        }
    }
}

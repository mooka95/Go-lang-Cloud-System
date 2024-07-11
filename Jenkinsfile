pipeline {
    agent any

    environment {
        // Define Docker Hub credentials and repository details
        DOCKERHUB_CREDENTIALS = credentials('39269fd9-1944-4b99-93fe-53f198a1a0cf')
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

        stage('Build Docker Image') {
            steps {
                script {
                    // Install jq if not already installed
                    sh 'if ! command -v jq > /dev/null; then sudo apt-get update && sudo apt-get install -y jq; fi'
                    
                    // Login to Docker Hub
                    sh "echo ${DOCKERHUB_CREDENTIALS_PSW} | docker login -u ${DOCKERHUB_CREDENTIALS_USR} --password-stdin"
                    
                    // Get the latest version tag from Docker Hub
                    def latestTag = sh(
                        script: "curl -s -u ${DOCKERHUB_CREDENTIALS_USR}:${DOCKERHUB_CREDENTIALS_PSW} https://hub.docker.com/v2/repositories/${DOCKERHUB_REPO}/tags/?page_size=1 | jq -r '.results[0].name'", 
                        returnStdout: true
                    ).trim()
                    
                    // Determine the new tag by incrementing the latest tag or starting from v1
                    def newTag = 'v1'
                    if (latestTag =~ /^v\d+$/) {
                        newTag = 'v' + (latestTag.replace('v', '').toInteger() + 1)
                    }

                    // Build the Docker image with the new tag
                    sh "docker build -t ${DOCKERHUB_REPO}:${newTag} ."
                    // Push the Docker image to Docker Hub
                    sh "docker push ${DOCKERHUB_REPO}:${newTag}"
                    
                    // Set the TAG environment variable for use in subsequent stages
                    env.TAG = newTag

                    // Replace the tag placeholder in docker-compose.yml with the new tag
             sh "sed -i 's|\\${TAG}|${newTag}|g' docker-compose.yaml"
                }
            }
        }

        stage('Deploy') {
            when {
                // Only run the Deploy stage if the previous stages succeeded
                expression { currentBuild.result == null || currentBuild.result == 'SUCCESS' }
            }
            steps {
                echo 'Deploying application...'
                // Use Docker Compose to deploy the new image
                sh 'docker-compose down && docker-compose up -d'
            }
        }
        
        stage('Get Running Containers') {
            steps {
                // List running containers
                sh 'docker ps'
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
pipeline {
    agent any

    environment {
        // Define Docker Hub credentials and repository details
        DOCKERHUB_CREDENTIALS = credentials('dockerhub-credentials')
        DOCKERHUB_REPO = 'mooka95/cloud-go'
    }

    stages {
        stage('Checkout') {
            steps {
                // Checkout code from the repository
                checkout scm
            }
        }

        stage('Build') {
            steps {
                script {
                    // Ensure Go is installed and available
                    sh 'go version'
                    // Build the Go application
                    sh 'go build -o app'
                }
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    // Login to Docker Hub using credentials stored in Jenkins
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
                // Add your deployment steps here, such as using Docker Compose to deploy the new image
                // Example:
                // sh 'docker-compose up -d'
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

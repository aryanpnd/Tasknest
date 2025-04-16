pipeline {
    agent any

    environment {
        DOCKER_IMAGE = "loudpear/tasknest"
    }

    stages {
        stage('Checkout') {
            steps {
                git 'https://github.com/aryanpnd/tasknest.git'
            }
        }

        stage('Build') {
            steps {
                script {
                    sh "docker build -t $DOCKER_IMAGE:latest -t $DOCKER_IMAGE:${BUILD_NUMBER} ."
                }
            }
        }

        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }

        stage('Push to Docker Hub') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'dockerhub-creds', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
                    sh '''
                        echo "$DOCKER_PASS" | docker login -u "$DOCKER_USER" --password-stdin
                        docker push $DOCKER_IMAGE:latest
                        docker push $DOCKER_IMAGE:${BUILD_NUMBER}
                    '''
                }
            }
        }
    }
}

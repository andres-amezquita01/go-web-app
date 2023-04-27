pipeline {
    agent any

    stages {
        stage('test') {
            tools {
                go 'Go-1.20.3'
            }
            environment {
                GO111MODULE = 'on'
            }
            agent {
                label 'tests'
            }
            steps {
                sh  '''
                go version
                go test
                '''
            }
        }
        stage('build-tag-docker-img'){
            agent {
                  label "docker"
            }
            steps{
              sh """
                   docker build -t slave02/${env.BUILD_NUMBER} .
                   docker tag slave02/${env.BUILD_NUMBER} andresamezquita01/mygoapp:latest
                   docker tag slave02/${env.BUILD_NUMBER} andresamezquita01/mygoapp:${env.BUILD_NUMBER}
              """
            }
        }
        stage('login-dockerhub') {
           agent {
                  label "docker"
            }
            environment {
                        DOCKERHUB_CREDENTIALS = credentials('andresamezquita01-dockerhub')
            }
            options {
              skipDefaultCheckout true
            }
            steps {                            
                    sh "echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin"
            }
        }

        stage('push-docker-img') {
            agent {
                  label "docker"
            }
            options {
              skipDefaultCheckout true
            }
            steps{
                   sh 'docker images'
                   sh "docker push --all-tags andresamezquita01/mygoapp"                              
            }
        }

        stage('master') {
            options {
              skipDefaultCheckout true
            }
            steps {
                echo 'Hello World from Master'
                sh 'ls'
                sh 'pwd'
                sh 'whoami'
            }
        }
    }
    post {
        always {
            node('docker'){
                sh 'docker rmi -f $(docker images -a -q)'
                sh 'docker images'
                sh 'docker logout'
            }
        }
    }    
}

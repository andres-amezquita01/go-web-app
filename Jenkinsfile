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
                echo testingfromslave01
                whoami
                go version
                go test
                '''
            }
        }
        stage('build-docker-img'){
            agent {
                  label "docker"
            }
            environment {
                        DOCKERHUB_CREDENTIALS = credentials('andresamezquita01-dockerhub')
            }
            steps{
              sh """
                   docker build -t slave02/${env.BUILD_NUMBER} .
                   docker tag slave02/${env.BUILD_NUMBER} andresamezquita01/mygoapp:${env.BUILD_NUMBER}
              """
            }
        }
        stage('Login') {
           agent {
                  label "docker"
            }
            environment {
                        DOCKERHUB_CREDENTIALS = credentials('andresamezquita01-dockerhub')
            }
            steps {
//                   sh """
//                     echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin
//                     docker push andresamezquita01/mygoapp:${env.BUILD_NUMBER}
//                     docker rmi -f $(docker images -a -q)
//                     docker logout
//                     """                         
                    sh "echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin"
                    sh "docker push andresamezquita01/mygoapp:${env.BUILD_NUMBER}"
                    //sh "docker rmi -f $(docker images -a -q)"
                    sh "docker logout"
                    
            }
        }

//         stage('Push') {
//             agent {
//                   label "docker"
//             }
//             environment {
//                         DOCKERHUB_CREDENTIALS = credentials('andresamezquita01-dockerhub')
//             }
//             steps {
//                 sh 'pwd'
//                 sh 'ls'
//                     sh '''
//                         docker push andresamezquita01/mygoapp:${env.BUILD_NUMBER}
//                         docker rmi -f $(docker images -a -q)
//                         docker logout
//                     '''
//             }
//         }

        stage('master') {
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
                sh 'docker logout'
            }
        }
    }    
}

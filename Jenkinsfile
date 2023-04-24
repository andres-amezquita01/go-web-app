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
                pwd
                sh 'echo testing!!!!!!!!'
                whoami
                '''
                dir('go-web-app'){
                    sh 'pwd'
                    sh 'go version'
                    sh 'go test'
                }
            }
        }
        stage('master') {
            steps {
                echo 'Hello World from Master'
                sh 'ls'
                sh 'pwd'
                sh 'whoami'
             
            }
        }
    }
}

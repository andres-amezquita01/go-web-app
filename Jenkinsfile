pipeline {
    agent any

    stages {
        stage('validate&clone') {
            agent {
                label 'tests'
            }
            steps {
                echo '****From Slave 01****'
                sh 'ls'
                sh 'pwd'
                sh 'whoami'
                sh './validatedirectory.sh'
                sh 'git clone git@github.com:andres-amezquita01/go-web-app.git'
                sh 'cd go-web-app'
                sh 'pwd'
                echo '****End Slave 01****'
            }
        }
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
                echo '****From Slave 01****'
                sh 'echo testing!!!!!!!!'
                sh 'ls'
                sh 'pwd'
                sh 'whoami'
                dir('go-web-app'){
                    sh 'pwd'
                    sh 'go version'
                    sh 'go test'
                }
                // sh 'cd go-web-app'
                // sh 'pwd'
                //sh 'rm -Rf go-web-app'
                echo '****End Slave 01****'
            }
        }
        stage('master') {
            steps {
                echo 'Hello World'
                sh 'ls'
                sh 'pwd'
                sh 'whoami'
             
            }
        }
    }
}

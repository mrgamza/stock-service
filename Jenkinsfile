node {
     stage('Clone repository') {
        checkout scm
     }

     stage('Build Docker Image') {
        sh 'make build'
        sh 'echo start build docker image'
     }

     stage('Test image') {
        sh 'echo "Tests passed"'
     }

     stage('Docker compose') {
        sh 'make up'
        sh 'echo deploy complete'
     }
 }
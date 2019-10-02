pipeline {
  agent { label 'slave' }
  options {
    buildDiscarder(logRotator(numToKeepStr: '5'))
  }
  tools {
    go 'go113'
  }
  stages {
    stage('go build'){
      steps{
        script {
            sh 'go version'
            sh 'go get -d github.com/gorilla/mux'
            sh 'go get -d github.com/prometheus/client_golang/prometheus'
            sh 'go get -d github.com/prometheus/client_golang/prometheus/promhttp'
            sh 'export CGO_ENABLED=0 && go build -a -installsuffix cgo --ldflags "-s -w" -o server'
        }
      }
    }
    stage('docker build'){
      steps{
        script {
          docker.build('densikatshine/simple-web-server')
        }
      }
    }
    stage('docker push'){
      steps{
        script {
          docker.withRegistry('https://registry.hub.docker.com', 'dockerhub') {
            docker.image('densikatshine/simple-web-server').push('badconfig')
          }
        }
      }
    }
  }
}

pipeline {
  agent any

  options {
    disableConcurrentBuilds()
  }

  tools {
    go '1.23.2'
    nodejs '23.1.0'
  }

  stages {
    stage('Install just') {
      steps {
        sh 'wget https://github.com/casey/just/releases/download/1.36.0/just-1.36.0-x86_64-unknown-linux-musl.tar.gz -O just.tar.gz'
        sh 'tar -zxvf just.tar.gz just'
      }
    }
    stage('Check formatting') {
      steps {
        sh 'cd client && npm install'
        sh 'cd client && npx prettier --check ./src'
        sh 'test -z $(gofmt -l .)'
      }
    }
    stage('Build') {
      steps {
        sh './just'
      }
    }
    stage('Test') {
      steps {
        sh './just test'
      }
    }
    stage('SonarQube Analysis') {
      when {
        branch 'main'
      }
      steps {
        script {
          scannerHome = tool 'SonarScanner';
        }
        withSonarQubeEnv('vezgammon') {
          sh "${scannerHome}/bin/sonar-scanner"
        }
      }
    }
    stage('DEV SonarQube Analysis') {
      when {
        branch 'develop'
      }
      steps {
        sh 'sed -i "s/vezgammon/dev/g" sonar-project.properties'
        sh 'sed -i "s/Vezgammon/Dev/g" sonar-project.properties'
        script {
          scannerHome = tool 'SonarScanner';
        }
        withSonarQubeEnv('vezgammon-dev') {
          sh "${scannerHome}/bin/sonar-scanner"
        }
      }
    }
    stage('Deploy on dev') {
      when {
        branch 'develop'
      }
      steps {
        sh 'ssh debian@wrk.dev.vezgammon.it "rm -rf repo"'
        sh 'ssh debian@wrk.dev.vezgammon.it "git clone https://gitlab.vezgammon.it/diego/vezgammon.git repo"'
        sh 'ssh debian@wrk.dev.vezgammon.it "cd repo && git checkout develop && sudo docker-compose -f docker-compose-release.yml down -v"'
        sh 'ssh debian@wrk.dev.vezgammon.it "mkdir -p repo/db && echo -n $(dd if=/dev/random bs=1 count=32 | base32 | sed \'s/=//g\') > repo/db/password.txt"'
        sh 'ssh debian@wrk.dev.vezgammon.it "cd repo && sudo docker-compose -f docker-compose-release.yml up -d --build"'
      }
    }
    stage('Deploy') {
      when {
        branch 'main'
      }
      steps {
        sh 'ssh debian@site.vezgammon.it "rm -rf repo"'
        sh 'ssh debian@site.vezgammon.it "git clone https://gitlab.vezgammon.it/diego/vezgammon.git repo"'
        sh 'ssh debian@site.vezgammon.it "cd repo && sudo docker-compose -f docker-compose-release.yml down -v"'
        sh 'ssh debian@site.vezgammon.it "mkdir -p repo/db && echo -n $(dd if=/dev/random bs=1 count=32 | base32 | sed \'s/=//g\') > repo/db/password.txt"'
        sh 'ssh debian@site.vezgammon.it "cd repo && sudo docker-compose -f docker-compose-release.yml up -d --build"'
      }
    }
  }
}


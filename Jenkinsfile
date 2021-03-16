pipeline {
  agent any
  stages {
    stage('merge') {
      steps {
        git(url: 'https://github.com/Chrismelba-Instaclustr/terraform-provider-instaclustr/', branch: '{Branch-SHA}')
      }
    }

    stage('Tag') {
      steps {
        git(url: 'https://github.com/Chrismelba-Instaclustr/terraform-provider-instaclustr/', branch: 'master')
      }
    }

  }
}
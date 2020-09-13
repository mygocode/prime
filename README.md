# Prime Number Application 
## Introduction
This is a simple web application that takes in a number and returns to the user the highest prime number lower than the input number.   
For example an input of 55 would return 53.   
Following tools are used to develop the appliation 
- Go language
- VS code IDE
- Git for source control managment
- GitKraken tool for git visualization
- Docker for containarization 
- Jenkins Pipeline for CI/CD 
- AWS EC2 instance for hosting the application 

## Project Design
This project is designed using GO language and Clean Architecture. Clean architecture ensures every layer should be separated from other and these layer should not have dependency on each other. This structure ensures cleanness, maintainability, and extensibility.   
As you can observe in the picture below that the architecture is not dependent on any framework and frameworks can be easily replaced without affecting the whole application.

![Clean Architecture](https://i.ibb.co/bbLrsPR/clean-arch.jpg)

## Containarization 
I have used multi-stage docker file to reduce the size of the docker image. 
![Multi Stage Docker Image](https://i.ibb.co/mqBjfGp/docker-img.jpg)   

## GIT 
Here is the git history. I use GitKraken tool for git visualization and keep tracking of multiple branches.
![Git History](https://i.ibb.co/KWQR99X/Git-Kraken.jpg)

## Development Process for production applicatins
I will prefer to use the following steps for production level application.

- Developer/s pick up the task from Jira board and assign it to themselves so other team members could know. 
- Create a feature branch locally.
- Push that branch to remote repo we well.
- If more than one developers are working then they will use this same branch and commit their changes on this branch
- When the feature is done, the code will go to the Code Review(CR) stage. 
- At CR stage, another developer in the team will review the code and either accepts the changes or reject the changes. 
- If CR is successful, then code is merged to the feature branch.
- When the feature is complete, the team will move their task on Jira board to QA column. 
- The QA team will deploy this feature on testing environment and test the featureaccording to the defined criteria in the description. 
- Once the QA issuccessful, then either QA or feature owner will merge the code to the master branch. 

## Deployment
I have used Jenkins pipeline for deployment. After merging the code to master branch, we can simply start the "Prime-Production" job. This job will perform the following operation
- Checkout the master branch.
- Build the docker image.
- Push the docker image to Dockerhub.
- Deploy the image to the AWS EC2 instance.

## Points which can be improved
Due to limited time availability, I couldn't implement all the features required for production application. Following features could not be implemented.    
- Reverse Proxy for load balancing
- Redis for caching
- Improved Logging
- Application security with JWT
- Docker swarm orchestration

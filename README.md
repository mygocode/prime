# Prime Number Application 
## Introduction
This is a simple web application that takes in a number and return to the user the highest prime number lower than the input number. For example an input of 55 would return 53.   
Following tools are used to develop the appliation   
- Go language
- VS code IDE
- Git for source control managment
- GitKraken tool for git visualization
- Docker for containarization 
- Jenkins Pipeline for CI/CD 
- AWS EC2 instance for hosting the application 

## Project Structure
This project is designed using GO language and Clean Architecture. Clean architecture ensures every layer should be separated from other and these layer should not have dependency on each other. This structure ensures cleanness, maintainability, and extensibility.   
As you can observe in the picture below that the architecture is not dependent on any framework and frameworks can be easily replaced without affecting the whole application.

![Clean Architecture](https://i.ibb.co/bbLrsPR/clean-arch.jpg)

## Containarization 
I have used multi-stage docker file to reduce the size of the docker image.   

![Multi Stage Docker Image](https://i.ibb.co/mqBjfGp/docker-img.jpg)   

## GIT 
Here is the git history. I use GitKraken tool for git visualization and keep tracking of multiple branches.   

![Git History](https://i.ibb.co/KWQR99X/Git-Kraken.jpg)

## Cache
Most of the application uses **Memcached**, **Redis** or **Apache Ignite** but according to the scope of this application, I decided use a Local Cache strategy which is easir to implement and light-weight. When the user will enter a number to get the prime number, the application will first check the local cache for the user given number. If its a cache miss, then the system will invoke the service layer and do the processing, cache the result, and return the result to the user.   
If user enters the same number again, this time the response will be super fast because the value is already available in the cache.

## Development Process for production application
I will prefer to use the following steps for production level application.

- Developer/s pick up the task from Jira board and assign it to themselves so other team members could know. 
- Create a feature branch locally.
- Push that branch to remote repo we well.
- If more than one developer is working then they will use this same branch and commit their changes on this branch.
- When the feature is done, the code will go to the Code Review(CR) stage. 
- At CR stage, another developer in the team will review the code and either accepts the changes or reject the changes. 
- If CR is successful, then code is merged to the feature branch.
- When the feature is complete, the team will move their task on Jira board to QA column. 
- The QA team will deploy this feature on testing environment and test the feature according to the defined criteria in the description. 
- Once the QA is successful, then either QA or feature owner will merge the code to the master branch. 

## Deployment with Jenkins
I have used Jenkins pipeline for deployment. In a team environment, we need to have at least 3 environments, **Dev**, **QA**, **Production**. There can be other environments like **CI**, **Nightly** etc, but that depends on requirement.   
Developers can use the Dev environment to quickly deploy and test their features. When development is done, that same instance should be deployed to the QA environment so QA team can check and further investigate possible bugs and issues. Once the feature is verified, this approved instance can be deployed to production and this all can be managed using Jenkins.   
To keep the process simple, I haven't configured any trigger on master branch to automate the build process. Developer manually build the job name "Prime-Production". This job will perform the following operation.
- Checkout the master branch.
- Build the docker image.
- Push the docker image to Dockerhub.
- Deploy the image to the AWS EC2 instance.   

![Jenkins](https://i.ibb.co/S3XWg8f/Jenkins.jpg)

## Points which can be improved
Due to limited time, I couldn't touch every aspect of production application. Following points can be improvement.    
- Reverse Proxy for load balancing
- Improved Logging
- Application security with JWT
- Docker swarm orchestration
- Unit and Integration Testing 

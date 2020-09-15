# Prime Number Application 
## Introduction
This is a simple web application that takes in a number and return to the user the highest prime number lower than the input number. For example an input of 55 would return 53.   
Following tools are used to develop the appliation
- Go language
- VS code IDE
- Git for source control managment
- GitKraken tool for git visualization
- Docker for containarization 
- Jenkins for CI/CD pipline 
- AWS EC2 instance for hosting the application 


## Project Structure
This project is designed using GO language and Clean Architecture. Clean architecture ensures every layer should be separated from other and these layer should not have dependency on each other. This structure ensures cleanness, maintainability, and extensibility.   
As it is clear in the picture below that the architecture is not dependent on any framework and frameworks can be easily replaced without affecting the whole application.

![Clean Architecture](https://i.ibb.co/bbLrsPR/clean-arch.jpg)

## Containarization 
Docker images are usedfor packging. I have used multi-stage docker file to reduce the size of the docker image.   

![Multi Stage Docker Image](https://i.ibb.co/mqBjfGp/docker-img.jpg)   

## GIT 
Here is the git history. I use GitKraken tool for git visualization and keep tracking of multiple branches.   

![Git History](https://i.ibb.co/cXKzg2r/Git-Kraken.jpg)

## Cache
Most of the application uses **Memcached**, **Redis** or **Apache Ignite** but according to the scope of this application, I decided use a **Local Cache** strategy which is easir to implement and light-weight. When the user will enter a number to get the prime number, the application will first check the local cache for the user given number. If its a cache miss, then the system will invoke the service layer and do the processing, cache the result, and return the result to the user.   
If user enters the same number again, this time the response will be super fast because the value is already available in the cache.

## Deployment with Jenkins
I have used Jenkins pipeline for deployment. In a team environment, we need to have at least 3 environments, **Dev**, **QA**, **Production**. There can be other environments like **CI**, **Nightly** etc, but that depends on requirement.   
Developers can use the Dev environment to quickly deploy and test their features. When development is done, team merge the code to the master branch and that instance is deployed to the QA environment, so QA team can check and investigate possible bugs and issues. Once the feature is verified, this approved instance is then ready for production and this whole process can be managed using Jenkins.   
To keep the process simple, I haven't configured any trigger on master branch to automate the build process on merge event. Developer manually build the job and this job performs the following operation.
- Checkout the master branch.
- Build the docker image.
- Push the docker image to Dockerhub.
- Deploy the image to the AWS EC2 instance.   

![Jenkins](https://i.ibb.co/dK1h70R/Jenkins.jpg)

## Development Process for production application
Every team has their own development process. Following points are general guideline on how development can be done for production applications.

- First is the planning phase and Product Manager adds the tasks to the Jira board.
- Developer/s pick up the task from Jira board and mark their name on it so other team members could know. 
- Pull the changes from origin master to local master. 
- Create a feature branch locally based on the stable master.
- Push the newly created feature branch to remote repo.
- If more than one developers are working then they will use this same branch and commit their changes.
- When a task is complete, the code will go to the Code Review(CR) stage. A feature can be divided in to tasks. 
- At CR stage, technical person will review the code and either accepts the code or reject it. 
- If CR is successful, then code is merged to the feature branch.
- When the feature is complete, the team will move their task on Jira board to QA column. 
- The QA team will deploy this feature on testing environment and test the feature according to the acceptance criteria defined in the Jira description. 
- Once the QA is successful, then either QA or feature owner will merge the code to the master branch. 

## Improvement and missing aspects
Due to limited time, I couldn't touch every aspect of production application. Following points can be improved.    
- Reverse Proxy for load balancing
- Logging
- Security
- Docker swarm orchestration
- Unit and Integration Testing 

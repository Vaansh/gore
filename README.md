<!--
**       .@@@@@@@*  ,@@@@@@@@     @@@     .@@@@@@@    @@@,    @@@% (@@@@@@@@
**       .@@    @@@ ,@@          @@#@@    .@@    @@@  @@@@   @@@@% (@@
**       .@@@@@@@/  ,@@@@@@@    @@@ #@@   .@@     @@  @@ @@ @@/@@% (@@@@@@@
**       .@@    @@% ,@@        @@@@@@@@@  .@@    @@@  @@  @@@@ @@% (@@
**       .@@    #@@ ,@@@@@@@@ @@@     @@@ .@@@@@@.    @@  .@@  @@% (@@@@@@@@
-->

<!-- HEADERS -->
<p align="center">
	<a href="https://github.com/Vaansh/gore/actions/workflows/cicd.yml">
	<img src="https://img.shields.io/github/actions/workflow/status/Vaansh/gore/cicd.yml?branch=main&logo=github&style=for-the-badge">
	</a>
	<a href="https://github.com/Vaansh/gore/blob/main/LICENSE">
	<img src="https://img.shields.io/github/license/gatsbyjs/gatsby.svg?style=for-the-badge">
	</a>
</p>

<!-- PROJECT LOGO -->
<br />
<p align="center">
	<a href="https://github.com/Vaansh/gore">
	<img src="https://upload.wikimedia.org/wikipedia/commons/2/2d/Go_gopher_favicon.svg" alt="Logo" height="120">
	</a>
<h3 align="center">GoRe</h3>
<p align="center">
	A Content Resharing Engine Written in Go.
	<br />
	<a href="https://github.com/Vaansh/gore"><strong>Explore the docs »</strong></a>
	<br />
	<br />
</p>
</p>

<!-- TABLE OF CONTENTS -->
<details open="open">
	<ol>
		<li>
			<a href="#about-the-project">About The Project</a>
			<ul>
				<li><a href="#built-with">Built With</a></li>
			</ul>
		</li>
		<li>
			<a href="#system-design">System Design</a>
		</li>
		<li>
			<a href="#overall-system-architecture">Overall System Architecture</a>
			<ul>
				<li><a href="#persistence-with-postgresql">Persistence with PostgreSQL</a></li>
				<li><a href="#cloud-bucket-storage">Cloud Bucket Storage</a></li>
				<li><a href="#cloud-and-local-logging">Cloud and Local Logging</a></li>
				<li><a href="#web-api-with-gin">Web API with Gin</a></li>
			</ul>
		</li>
		<li><a href="#deployment">Deployment</a></li>
		<ul>
			<li><a href="#github-actions">GitHub Actions</a></li>
			<ul>
				<li><a href="#pr-linting">PR Linting</a></li>
				<li><a href="#docker-image">Docker Image</a></li>
			</ul>
			<li><a href="#digital-ocean-droplet">Digital Ocean Droplet</a></li>
		</ul>
		<li><a href="#miscellaneous">Miscellaneous</a></li>
		<ul>
			<li><a href="#future-plans">Future Plans</a></li>
			<li><a href="#environment-variables">Environment Variables</a></li>
			<li><a href="#directory-structure">Directory Structure</a></li>
		</ul>
		</li>
		<li><a href="#license">License</a></li>
	</ol>
</details>

## About The Project

A few years ago I wrote a few scripts to automate posting content from Reddit to Instagram. Since then I've had the idea to capitalize on the rise of short form content – this is an implementation of that idea. I wanted a way to automate posting content from multiple sources – regardless of the platform (YouTube, Instagram, Tiktok, etc.) to another platform of my choice. Something scalable and well thought out based on my previous experience with this project. It's a project I have been meaning to work on from quite some time now and I'm happy with the way it has turned out so far. That being said it is still a prototype and you can track my progress [here](https://github.com/users/Vaansh/projects/1).

<p align="center">
	<a href="https://github.com/users/Vaansh/projects/1">
		<img src="https://i.imgur.com/D4DaoB4.jpeg">
	</a>
</p>

### Built With

<p align="center">
  <img src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white">
  <img src="https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white">
  <img src="https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white">
  <img src="https://img.shields.io/badge/GoogleCloud-%234285F4.svg?style=for-the-badge&logo=google-cloud&logoColor=white&yellow=black&color=yellow">
  <img src="https://img.shields.io/badge/DigitalOcean-%230167ff.svg?style=for-the-badge&logo=digitalOcean&logoColor=white">
  <img src="https://img.shields.io/badge/github%20actions-%232671E5.svg?style=for-the-badge&logo=githubactions&logoColor=white&labelColor=black&color=black">
</p>

## System Design

A `Task` makes up the basic structure of my application. It is defined as follows and consists of many publishers and one subscriber. Channels in Go seemed like the right message passing system to start with (it gave me the pub-sub mechanism I was looking for to implement something like this). Every task when created, has its own channel where the publishers write and subscriber consumes – making each component responsible for their own fetching or posting mechanism – allowing for separation of concerns.

```go
type Task struct {
	Id         string
	Publishers []publisher.Publisher
	Subscriber subscriber.Subscriber
	Quit       chan struct{} // covered later
}
```

It has a `Run()` method which starts a goroutine for every publisher (that publish posts to the channel) and the subscirber (which subscribes to this channel and is responsible for ensuring: (1) ensuring uniqueness of the post (2), storing the file locally, (3) uploading it to a cloud storage, (4) deleting them from both after posting it to the desired platform, and (5) ensuring it maintains a certain posting frequency). The definition of Publisher and Subscriber are outlined below for reference.

```go
type Publisher interface {
	PublishTo(c chan<- model.Post, quit <-chan struct{})
	GetPublisherId() string
}

type Subscriber interface {
	SubscribeTo(c <-chan model.Post)
	GetSubscriberId() string
}
```

_NOTE:_ As of now, YouTube is the only kind of publisher and Instagram is the only kind of subscriber that is supported. Please follow the rest of the document with that in mind.

There is also a `Quit` channel that simply exists for force quitting a task, this is for use by the `TaskService`. It is the service responsible for managing the lifecycle of all tasks by maintaing a map for the tasks currently running and the quit channel to invoke, if the task should be stopped.

## Overall System Architecture

With the main application logic out of the way, I'll cover the entire software architecture from a higher level. Below is the overall workflow of the project in its current state. Each component if briefly talked about below.

<p align="center">
	<img src="https://i.imgur.com/mFNZSkN.png">
</p>

### Persistence with PostgreSQL

I needed some kind of persistence mechanism to ensure I do not post the same content twice (subscriber side logic) so I decided to go with Google Cloud Platform and created an instance there. I didn't use GORM or any other Object Relational Mapper since it didn't fit my needs and I was focussed on delivering the project ...before my cloud services credits expire.

<p align="center">
	<img src="https://i.imgur.com/FltiJVn.jpg">
</p>

### Cloud Bucket Storage

In order to use Meta's developer API for content publishing (subscriber side logic), I needed the videos to be stored locally (in the `data/` directory) and then hosted on a web server before hitting their Graph API endpoint to publish the video. For this reason, I chose cloud bucket storage as it gave me an easy way to manage storage with Go's client library and since I was already using their database services.

<p align="center">
	<img src="https://i.imgur.com/Y7dAglJ.jpg">
</p>

### Cloud and Local Logging

For app-wide logging, I chose to go with GCP again, since it would give me a centralized way of going over my logs. For development purposes, I usually log them locally (logs get saved into the `log/` directory). Moreover, I had plans of containerizing the application, so I needed a good way to monitor or debug it. Local and cloud logging options can be set at an application level through the use of environment variables. In production, I have cloud enabled and local logging disabled.

<p align="center">
	<img src="https://i.imgur.com/YGJUaOW.jpg">
</p>

### Web API with Gin

Since I wanted to take as much of a hands-off approach to the project, I decided to build an API around it that would allow me to manage my tasks by interfacing with the TaskService. I used Gin once and was quite happy with the experience so I wrote my handlers and decided to go with it again. Admittedly though, the middleware in its current state is weak (just a token set as an environment variable on the web server that needs to be added in my web requests).

A sample `POST` request to `tasks/:platform`

<p align="center">
	<img src="https://i.imgur.com/9JHOkdF.jpg">
</p>

A sample `DELETE` request to `tasks/:platform/:id`

<p align="center">
	<img src="https://i.imgur.com/8t7BnMY.jpg">
</p>

## Deployment

The only way to make changes to the `main` branch is by opening PRs. Once a PR is merged, the docker image builds, and if succesful, is moved to the DigitalOcean container registry which serves as the main hosting place for my artifacts. I also publish this as a private image on DockerHub, for my own future reference. The flow described is displayed in the image below.

<p align="center">
	<img src="https://i.imgur.com/3AdgKJl.jpg">
</p>

### GitHub Actions

I have two GitHub actions in place:

#### PR Linting

<p align="center">
	<img src="https://i.imgur.com/miuj5bf.jpg">
</p>

One for – linting PRs and making sure no weird looking code gets committed to `main`.

#### Docker Image

<p align="center">
	<img src="https://i.imgur.com/EYl9Ei6.jpg">
</p>

Another one for building and pushing the docker image to the registry, acting as my CI/CD pipeline. It also runs a script that SSH into my droplet, replaces the to the newest docker image in the registry and starts running it.

<p align="center">
	<img src="https://i.imgur.com/q2XJ7XR.jpg">
</p>

### DigitalOcean Droplet

I tried Google Cloud Run, Kubernetes, & Compute Engine but nothing really suited what I was looking for. I decided to go with another platform and created a Droplet (VM) on DigitalOcean. I really liked my experience with DigitalOcean so far, it gave me a streamlined developer solution I was looking for.

<p align="center">
	<img src="https://i.imgur.com/zGEWBkS.jpeg">
</p>

<!-- Miscellaneous -->

## Miscellaneous

There are certain things I wanted to discuss but didn't fit into any of these topics, so I'm briefly going over them above.

### Future Plans

The [project page](https://github.com/users/Vaansh/projects/1) is the best way to track future plans and current progress of the project.

<p align="center">
	<img src="https://raw.githubusercontent.com/gist/brudnak/6c21505423e4ff089ab704ec79b5a096/raw/b2d3dec32474b2121b179920734b259323a7c250/go.gif" height="120" width="180">
</p>

Overall, I need to add unit tests (and run it on PRs), and I think I can to a better job with logging. But there are still a lot of other platforms I need to implement – for both publishers and subscribers.

My application is also quite stateful, which might be something I would want to look into if I want to redesign it before moving forward with the project. I still like it because it helped me learn bridge some knowledge gaps I had in Go.

### Environment Variables

All required environment variables so far can be seen in the `.env-sample` file. (PS: having the actual file is not necessary, I just use it so I have my secrets in one place but you can just set those environment variables directly). For obvious reasons, the actual variables file itself isn't tracked and I generate it throguh GitHub actions – this way I don't have to set them each time I change the machine I deploy them on. But it also means my docker image must remain priavte so I ensure no one has access to my containers. Also there is a sample service account key credentials JSON file that is needed for accessing various GCP services through client libraries.

### Directory Structure

This is my overall directory structure with a brief explanation of each.

```sh
.
├── .github                     # everything github related
│   ├── workflows               # all my workflows
│   │   ├── cicd.yml            # building and deploying action file
│   │   ├── reviewdog.yml       # pr linting action file
├── Dockerfile                  # used to create docker image
├── LICENSE                     # software license
├── README.md                   # documentation
├── cmd                         # all main programs live here
│   └── api                     # api directory containing main program
│       └── main.go             # the main program
├── data                        # directory used to store .mp4 files locally
├── gcloud-key.json             # sample service key credentials
├── go.mod                      # all project dependenciess
├── go.sum                      # go module checksum file
├── internal                    # all application specific logic
│   ├── api                     # everything related to the api
│   │   ├── handler.go          # api handlers
│   │   ├── middleware.go       # authorization middleware
│   │   ├── request.go          # request dtos
│   │   └── response.go         # response dtos
│   ├── config                  # all my config files
│   │   ├── database.go         # reads and creates database config
│   │   ├── logger.go           # reads and creates logger config
│   │   └── storage.go          # reads and creates logger config
│   ├── domain                  # the core of my application
│   │   ├── service.go          # task service
│   │   └── task.go             # definiton and task behaviour
│   ├── gcloud                  # gcp module
│   │   ├── database.go         # functions for cloud database
│   │   ├── logger.go           # functions for cloud logger
│   │   └── storage.go          # functions for cloud bucket storage
│   ├── http                    # all my http clients
│   │   ├── client.go           # general client definition
│   │   ├── instagram.go        # instagram platform client
│   │   └── youtube.go          # youtube platform client
│   ├── model                   # all models used in my application
│   │   ├── metadata.go         # metadata used for things like scheduling
│   │   ├── post.go             # model defining post
│   │   └── user.go             # model defining user (subscriber)
│   ├── publisher               # all publisher logic
│   │   ├── publisher.go        # general publisher interface
│   │   └── youtube.go          # youtube publisher implementation
│   ├── repository              # for all database operations
│   │   └── user.go             # user repository operations
│   ├── subscriber              # general subscriber interface
│   │   ├── instagram.go        # instagram subscriber implementation
│   │   └── subscriber.go       # general subscriber interface
│   └── util                    # util module
│       └── util.go             # all my helper functions
├── log                         # local log directory
│   ├── error.log               # sample error log file
│   ├── info.log                # sample info log file
│   └── warning.log             # sample warning log file
├── scripts                     # a bunch of simple scripts I use often
│   ├── clean.sh                # deletes all log and data files
│   └── run.sh                  # builds and runs cmd/api/main.go
├── supported.go                # a centralized list of supported platforms
```

## License

Distributed under the MIT License. See `LICENSE` for more information.

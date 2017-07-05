# WebSockets using Angularjs and Go #

Proof of concept project to showcase working of Websockets using AngularJS (UI) and GO (back end).

### Description ###

The goal of this project is to test uses of Websockets in web applications. [WebSockets](https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API) enables to open an interactive session between user's browser and server. This mainly helps is avoiding polling the server for events. 

Some use cases of WebSockets are:

* Social feeds
* Multi player games
* Document collaboration 
* Financial Tickers
* Click stream data analysis
* Sports updates
* Chat
* Etc..

This project aims to use WebSockets for the financial tickers use case. When user browses to the app, a websocket connection is opened to the backend server. The server listens to the "ticker" request from the UI- the connection is kept open. When a ticker request comes in, a go routine is spawned and starts quering Yahoo for the stock quotes every second. The go routine ends when another request comes in and another go routine is started. 



### Installation requirements

* Install [VirtualBox](https://www.virtualbox.org/wiki/Downloads) and [Vagrant](https://www.vagrantup.com/downloads.html)
* Clone this repo
* The Vagrant file install docker and docker-compose for easily spinning up apps.

### Commands ###

* Start
    * `vagrant up`
* SSH into VM
    * `vagrant ssh` 
* cd into project folder
    * `cd websocketExample`
* Build docker container
    * `docker-compose build`
* Start docker container
    * `docker-compose up`
* Navigate to http://192.168.33.46

![Alt](/resources/stockTicker.gif "Demo")

### References ###

* [WebSockets Use cases](http://www.javaworld.com/article/2071232/java-app-dev/9-killer-uses-for-websockets.html)
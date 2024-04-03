# NRL Supercoach AI Web Server

Rewriting the squad selector part of my nrl supercoach program in Golang to run on a Google Cloud Run web server.

This will be consumed by my blog site to allow people to interactively use it to select squads.


It uses a modified knapsack algorithm in order to select the squad with the value of each player determined by a neural network that was previously trained.

https://thepinkai.com/supercoach-blog/selecting-initial-squad-algorithm/


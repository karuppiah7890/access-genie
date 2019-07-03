# access-genie

## How to use in your workspace:

### Retrieve token for bot

1. Create an app here - https://api.slack.com/apps
2. Go to the app page
3. Go to "Bot Users" section in the left navigation bar
4. Create a bot user
5. Go to "Install App" section in the left navigation bar
6. Install the app in your workspace
7. Get the "Bot User OAuth Access Token". This is the token needed for access-genie bot

### Running the server

1. Put the token in an `.env` file

```
# .env file
export SLACK_TOKEN="your-token"
```

2. Source the `.env` file

```
$ source .env
```

3. Run the server

```
$ access-genie
```

## Contributing

1. Clone the repo outside `$GOPATH` as this project uses go modules

```
$ git clone https://github.com/karuppiah7890/access-genie
```

2. Build it

```
$ make build
```

3. Set it up and run it

```
$ make setup
$ vi .env # change the slack token
$ source .env
$ ./out/access-genie
```
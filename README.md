# Slavart Downloader
This is a simple tool written in Go(lang) to download music using the SlavArt Divolt server. This tool was inspired by [https://github.com/D0otDo0t/slavolt-scraper](https://github.com/D0otDo0t/slavolt-scraper), however I choose to write my own tool because I noticed there were inefficiencies in how the download link was collected.

I created this tool for educational purposes and I do not condone any form of piracy.

You can find pre-build versions in the releases section.

## Usage
Just run the executable for your system. There are prebuild executables for Windows, Linux and MacOS (darwin). Linux and MacOS have arm compatible versions.

On Windows don't double click the executable, it will open a command prompt and close instead open a command prompt and navigate to the directory containing the `slavartdl.exe`. You then need to run this executable in the command prompt to successfully run the command.

### Commands
To find out how to run the commands use `--help` on any command (or subcommand). 

Heres what the arguments mean:
- `value(s)`: 1 or more values allowed
- `value`: 1 value allowed
- `values`: only more than 1 values allowed
- `[flags]`: 1 or more optional flags

Heres a brief structure of the commands:
```
slavartdl download url(s)
slavartdl version
slavartdl update                                         # updates slavartdl
slavartdl config add
slavartdl config add divoltTokens token(s)
slavartdl config add divoltCredential email password
slavartdl config add discordTokens token(s)
slavartdl config add discordCredential email password
slavartdl config list
slavartdl config list divoltTokens
slavartdl config list divoltCredentials
slavartdl config list discordTokens
slavartdl config list discordCredentials
slavartdl config remove
slavartdl config remove divoltTokens tokenIndex(s)        # tokenIndex is from the list command
slavartdl config remove divoltCredentials tokenIndex(s)   # tokenIndex is from the list command
slavartdl config remove discordTokens tokenIndex(s)       # tokenIndex is from the list command
slavartdl config remove discordCredentials tokenIndex(s)  # tokenIndex is from the list command
```

### Go installed?
If you have Go installed then you can run:
```
go install github.com/tywil04/slavartdl@latest
```
which will install `slavartdl` in your go bin path.

## Config
Session tokens are stored in a local config file (use the `config` command to find the location). You do not need to manually edit the config, you can use the commands show below. The session tokens are stored in plaintext due to the simplicity of this program, this means anyone who has access to your file system can use your revolt account(s). Dont use your main account for this, I am not liable for your account getting hacked or stolen.

You can have multiple session tokens that will randomly get used per request.

### Structure
As a note, the structure of `"downloadcmd.timeout"` has changed, its now an int vs a map containing seconds and minutes. This is because the timeout flag has also changed to only be seconds. The additional minutes flag/config value was redundant so it was removed. Please note, `"downloadcmd.outputdir"` must be an absolute file path, not relative.

```
{
  "discordlogincredentials": [
    {
      "email": string,
      "password": string
    }
    ...
  ],
  "discordsessiontokens": [
    string 
    ...
  ],
  "divoltlogincredentials": [
    {
      "email": string,
      "password": string
    }
    ...
  ],
  "divoltsessiontokens": [
    string 
    ...
  ],
  "downloadcmd": {
    "ignore": {
      "cover": bool,
      "subdirs": bool
    },
    "outputdir": string,
    "loglevel": string,
    "quality": int,
    "timeout": int,
    "cooldown": int
  }
}
```
You will require either at least one session token in `divoltsessiontokens` or at least one dictionary with an email and password in `divoltlogincredentials`. Both of these allow this tool to access your Divolt account which is required for this bot to function. Please don't use your main account!

### Getting Divolt session tokens to add to config
If you want to get your session token, its easy. Just note that once you get your session token, for it to remain active you must close the divolt tab and do not logout.

Follows these steps to get a session token:
- Log in to Divolt.
- Open browser DevTools.
- Navigate to the network tab and then select the 'Fetch/XHR' filter.
- Find and select a request with the domain of `api.divolt.xyz`.
- Select the header tab and copy the value of the `X-Session-token` from the response headers. If there is no `X-Session-token` select another request.

### Getting Discord session tokens to add to config
If you want to get your session token, its easy. Just note that once you get your session token, for it to remain active you must close the discord tab and do not logout.

Follows these steps to get a session token:
- Log in to Discord.
- Open browser DevTools.
- Navigate to the network tab and then select the 'Fetch/XHR' filter.
- Find and select a request named `login` with a request url of `https://discord.com/api/v9/auth/login`.
- Select the response tab and copy the value of `token`.


## Building CLI
To build the command line application you need to have Go(lang) installed. Once installed run the following commands:
```
git clone https://github.com/tywil04/slavartdl
cd slavartdl
go build -o slavartdl cli/main.go
```

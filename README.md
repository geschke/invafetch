# invafetch (inverter value fetcher)

Invafetch is a tool for retrieving processdata values from Kostal Plenticore inverters.

This tool is not affiliated with Kostal and is no offical product of KOSTAL Solar Electric GmbH or any subsidiary company of Kostal Gruppe.

## Description / Overview

Invafetch is one of several building blocks for generating a Grafana dashboard for Kostal Plenticore inverters. Invafetch reads the processdata values at regular intervals from the Inverter API and stores the results in JSON format in a MariaDB table. The Invaps tool uses these values, i.e. reads them and makes them available to Prometheus on request. Grafana, in turn, uses Prometheus as a data source to create a dashboard for the Kostal Plenticore inverter. Here, a modular concept was implemented so that one application, as small as possible, is responsible for a single task at a time. The MariaDB database serves as the interface of the Invafetch and Invaps tools and thus as a buffer for the processdata values. For a description of Invaps see [https://github.com/geschke/invaps](https://github.com/geschke/invaps), a complete example including definition of the Grafana dashboard and a Docker compose file to start all components in a Docker environment can be found at [https://github.com/geschke/grkopv-dashboard](https://github.com/geschke/grkopv-dashboard).


## Installation

The recommended installation method is to use the Docker image, a fully commented example can be found at [https://github.com/geschke/grkopv-dashboard](https://github.com/geschke/grkopv-dashboard). Besides that, Invafetch can be installed from source like any other program written in Go:

```text
$ git clone https://github.com/geschke/invafetch
$ cd invafetch/
$ go build
$ go install
```

This command builds the invafetch command, producing an executable binary. It then installs that binary as `$HOME/go/bin/invafetch` (or, under Windows, `%USERPROFILE%\go\bin\invafetch.exe`).
After that invafetch can simply be started in the command line interface.


## Configuration

All processdata values that can be read out can be found in the `processdata.json` file. These are almost all processdata values provided by the Kostal inverter (except for the scb:export and scb:update modules). If not all values should be read and stored, single processdata ids, but also complete module ids can be removed from the `processdata.json` file. When Invafetch is started, the file is read once and used as configuration parameter.

Invafetch uses the JSON data type of MariaDB and requires a running MariaDB installation. The corresponding definition of the table structure can be found in the sql/ directory in the solardata.sql file.

All configuration options can be passed either in a configuration file (default file name: `~/.env`), as environment variables or as command line parameters.

The following options exist:

|Name of environment variable|CLI flag|Defaults|Example|Hint|
|--------------------|-------------|------------|--------|-------|
|DBHOST|--dbhost|(empty)|"db.example.com"|database server|
|DBUSER|--dbuser|(empty)|"solardbuser"|database username|
|DBNAME|--dbname|(empty)|"solardb"|name of database|
|DBPASSWORD|--dbpassword|(empty)|"myDBPassword"|password of database user|
|DBPORT|--dbport|"3306"|"3306"|MariaDB port (optional)|
|INV_SERVER|--server|(empty)|"192.168.0.100"|inverter address (FQDN or IP)|
|INV_SCHEME|--scheme|"http"|"http"|possible values: http or https|
|INV_PASSWORD|--password|(empty)|"myPlantOwnerPassword"|plant owner password|
|TIME_REQUEST_DURATION_SECONDS|--time-request|3|5|time span between two requests in seconds, i.e. values are read every n seconds|
|TIME_NEW_LOGIN_MINUTES|--time-new-login|10|15|Duration of a session in minutes. A logout and subsequent login occurs after n minutes, so that a new session is created.|

## CLI

If invafetch is called without parameters or with the `--help` or `-h` flag, an overview of the available commands appears:

```text
~$ invafetch

A tool for retrieving values from Kostal Plenticore inverters


Usage:
  invafetch [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  info        Returns miscellaneous information
  start       Start collecting and storing values from inverter

Flags:
      --config string        config file (default is ~/.env)
      --dbhost string        Database host
      --dbname string        Database name
      --dbpassword string    Database password
      --dbport string        Database port (default "3306")
      --dbuser string        Database user
  -h, --help                 help for invafetch
  -p, --password string      Password (required)
  -m, --scheme string        Scheme (http or https, default http)
  -s, --server string        Server (e.g. inverter IP address) (required)
      --time-new-login int   Duration in minutes between two logins to inverter and database (default 10)
      --time-request int     Request new processdata every n seconds (default 3)

Use "invafetch [command] --help" for more information about a command.

```

## Quick Start

The installation and setup of the MariaDB database will not be discussed further here. Should an existing MariaDB database server be used, a database must first be created in which the table '*solardata*' is generated from the file `sql/solardata.sql`.

To test the connection to the inverter, the "`info`" command can be used. If the connection can be established successfully, information about the inverter API is written:

```text
$ invafetch info version -s 192.168.X.Y -p "MYPASSWORD"
hostname: _INVERTER HOSTNAME_
sw_version: 01.23.07734
api_version: 0.2.0
name: PUCK RESTful API
```

This information corresponds to a request to the inverter with the URL `http://192.168.X.Y/api/v1/info/version` . Although this request is also possible without authentication, invafetch uses the access as plant owner by default, which needs an authentication. If the password parameter is missing, this is acknowledged with an error message:

```text
~$ invafetch info version -s 192.168.X.Y
password parameter / INV_PASSWORD variable missing.
Please use --password options or add INV_PASSWORD to the config file or to ENV variables
```

All flags can be passed either as CLI parameters, in a config file or as environment variables. The CLI parameters have the highest priority, followed by the environment variables, followed by the information in the config file. If required parameters are missing completely, a corresponding error message is issued. For a list of configuration parameters, see [Configuration](#configuration).

The process for collecting and storing the data is started with the command "`start`". Thereby the file `processdata.json` must be in the current directory.

```text
$ invafetch start
Alloc = 0 MiB   TotalAlloc = 0 MiB      Sys = 8 MiB     NumGC = 0
[...]
```

In the current version some parameters about memory consumption and current state are written, this may be omitted or offered as an option in future versions. After startup, new content should be found in the *solardata* table. It should be mentioned again that it is recommended to use the Docker image. The author has been running the combination of invafetch and invaps in a Docker environment stably for several months.

## License

Invafetch is licensed under the MIT License. You can read the full terms here: [LICENSE](LICENSE).

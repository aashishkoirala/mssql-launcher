# MSSQL Launcher
___
##### By [Aashish Koirala](https://www.aashishkoirala.com)

**MSSQL Launcher** is a cross-platform CLI based interactive tool that lets you define SQL Server connections in a file and pick them from a menu to connect using either the MSSQL-CLI tool or the SQLCMD tool. Tested on Linux (Ubuntu) and Windows. Written in Go. For those who like to live in the CLI and find it a hassle to launch SSMS any time you need to run a quick query or check some data quickly.

## Usage
**mssql-launcher** [**--config** _config-file_] [**--log** _log-file_]

- _config-file_ is a YAML file that defines the SQL Server connections. If not provided, the default path is used which is **mssql-launcher-config.yaml** in the same directory as the binary.
- _log-file_ is the path where logs are written. If not provided, the default path is **mssql-launcher-log.log** in the same directory as the binary.

## Constraints
- You need to have the respective tool, **mssql-cli** and/or **sqlcmd**, installed on your system and available on _PATH_.
- Note that SQL Server _Integrated Windows Authentication_ mode, as the name suggests, only works on Windows.

## Installation
- The binary is not hosted anywhere, so you will need to clone the source and build it.
- The build script is set up to run on Linux (Ubuntu), but generates a build for Windows as well.
- To run the build script on Linux:
  - Simply run **make** in the repository directory.
  - The script will create binaries and sample config files in **bin/linux** and **bin/windows**. It will also run tests and generate docs.
  - You need to have **make** installed as well as the **Go SDK**. The application was written using Go 1.16.
- Alternatively, or in Windows, provided you have **Go SDK** installed, just run **go build** on the repository directory. This will, of course, not run tests or generate docs.


## Configuration
You can define your connections in the configuration YAML file. Here is an example:

```yaml
- name: My Database 1
  server: myfirstserverhost
  database: MyDatabase1
  integrated: true
- name: My Database 2
  server: mysecondserverhost,1434
  database: MyDatabase2
  username: testuser
  password: testpassword
- name: My Database 3
  server: mysecondserverhost,1434
  database: MyDatabase3
  username: securetestuser
  password: <encrypted password string>
  passworddecryptor: decrypt ${_PWD_}
```

Note the last entry stores the password in an encrypted form. The encryption/decryption is not managed by this application. If you have a command line tool that can be used to encrypt or decrypt a string, you can use that to encrypt the string and store it in the YAML file. In that case, you should then specify the command to decrypt it in the _passworddecryptor_ field. In the above example, _decrypt_ is the command line tool. ${_PWD_} is a special placeholder that is used to signify the encrypted password as stored in the YAML file.

You can specify which tool to use in the YAML file as well. If nothing is specified, **mssql-cli** is the default:

```yaml
- name: My Database 1 (use SQLCMD)
  server: myfirstserverhost
  username: testuser
  password: testpassword
  tool: sqlcmd
- name: My Database 2 (use mssql-cli)
  server: myfirstserverhost
  username: testuser
  password: testpassword
  tool: mssql-cli
```

## License
See [LICENSE](LICENSE).

## Contributing
Issues and PRs welcome. However, please note that this is a hobby project and set your expectations accordingly.
# basic-app-server

A basic application server built by Leo Hung.

# Program Structure
```go
basic-app-server
/* Responsible for command line options. */
└── cmd/app-server/main.go 
    /* Responsible for creating and writing log files. */
    ├── logger/logger.go
    /* Responsible for reading and decoding configuration files. */
    └── config/config.go
        /* Responsible for importing all required packages based on configuration. */
        └──  core/core.go
            /* Responsible for connecting to the database. */
            │── datastore/datastore.go
            /* Responsible for managing the use of data. */
            │── core/manager/manager.go
            /* Responsible for reporting data to the web server.*/
            │── core/reporter/reporter.go
            /* Responsible for obtaining data and downloading files from the web server. */
            │── core/syncer/syncer.go
            /* Responsible for connecting peripheral devices. */
            │── peripherals/peripherals.go
            /* Responsible for connecting rpc server and providing api interface. */
            │── rpc/dispatchers.go
            /* Responsible for arranging schedules for manager, reporter and syncer to complete tasks. */
            │── core/tasks/tasks.go
            /* Responsible for obtaining instructions from rpc modules to control other modules. */
            └── core/proxy/proxy.go
```
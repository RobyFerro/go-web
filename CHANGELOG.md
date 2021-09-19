# Change Log

All notable changes to the "Go-web" will be documented in this file.

## [Unreleased]

## [v0.9.0-beta] - 2021-09-19
### Changed
- Changed AuthController with new kernel.Request payload
- Updated new gwf version

## [v0.8.1-beta] - 2021-09-03
### Removed
- Logging middleware. Now logs are handled by the global gfw handler.

### Changed
- Updated new gwf version
- Ignores .exe files
- Middleware architecture

### Removed
- Middleware register cause now are registered directly in http routers.

## [v0.8.0-beta] - 2021-09-02
### Changed
- Implement adaptation to new GWF version

## [v0.7.2-beta] - 2021-09-02
### Changed
- Removed pointers in route middlewares

### Added
- Support for request validation in routing

## [v0.7.1-beta] - 2021-08-30
### Changed
- Moved command service register in "command" module
- Changed .yml routing system in favour of the router module. Now every route/group have to be registered directly in go structure.
- Changed .yml configuration. Now system and custom configuration are located into the new config module.

## [v0.7.0-beta] - 2021-08-30
### Removed
- Removed async jobs

### Changed 
- Console command register in "kernel.go"
 
## [v0.6.2-beta] - 2021-08-27
### Changed
- Updated gwf version
- Now evey middleware consists in an isolated struct.

## [v0.6.1-beta] - 2021-08-25
### Changed
- Improved makefile

## [v0.6.0-beta] - 2021-08-25
### Changed
- Split command line interfaces and http server

## [v0.5.3-beta] - 2021-08-20
### Changed
- Updated GWF version

## [v0.5.2-beta] - 2021-08-19
### Changed
- Implemented CLI service container

## [v0.5.1-beta] - 2021-08-19
### Changed
- Updated GWF version

## [v0.5.0-beta] - 2021-08-19
### Changed
- Implemented IOC container with a 'request lifecycle'
- Implemented SingletonContainer
- Updated GWF version

## [v0.4.7-beta] - 2021-06-21
### Changed
- GWF version that contains old helpers in tool package.
- Documentation style and content

## [v0.4.6-beta] - 2021-06-16
### Changed
- Fix go.mod dependencies.

### Removed
- Helper package

## [v0.4.5-beta] - 2021-06-16
### Changed
- Removed install, http_loader, show route, daemon run cli commands


## [v0.4.3-beta] - 2021-06-14
### Fixed 
- Installation procedure in documentation

### Removed
- Install command from `alfred`

### Added
- Updated `go-web-framework` version

## [v0.4.2-beta] - 2020-08-28
### Added
- Make file command to run and compile Go-Web

### Changed
- Project documentation
- Remove duplicated element form README.md

## [v0.4.1-beta] - 2020-08-28
### Added
- Auto-register generated component

### Changed 
- Component registration location (register.go)

## [v0.4.0-beta] - 2020-08-27
### Added 
- Alfred CLI tool

### Removed
- Some commands from project build

## [v0.3.2-beta] - 2020-07-09
### Changed
- Updated GWF version

### Fixed
- Default JWT auth duration

## [v0.3.1-beta] - 2020-03-21
### Added
- Project documentation
- Fix basic auth middleware
- Add *.test.json to .gitignore
- Add gob structure for basic auth in goweb.go

## [v0.3.0-beta] - 2020-02-24
### Added
- Custom app configuration
- Returns user information by calling /admin/test route

### Changed
- Moved HTTP kernel from Go-Web to Go-web framework

## [v0.2.1-beta] - 2020-02-19
### Added
- Insert changelog file

### Changed
- Update Go-Web framework version


## [v0.2-beta] - 2020-02-19
### Added
- Updated base Go-Web Framework version

### Changed
- Insert placeholder in config.yml.example for app.key node


## [v0.1-beta] - 2020-02-17
### Added
- First Go-Web release

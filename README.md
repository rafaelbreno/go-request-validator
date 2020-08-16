# Go Request Validator
- My first project after I've started the [_Go4Noobs Repo_](https://github.com/rafaelbreno/go4noobs)
## About
- This project will be a _Request Validator_
- My inspiration is the [Laravel Request Validator](https://laravel.com/docs/7.x/validation#quick-writing-the-validation-logic)
## TODO List
- [ ] Define Folder Structure
- [ ] Create a Package
- [ ] POST Validation
    - [ ] Simple
        - [ ] Required
        - [ ] Type 
        - [ ] Length
    - [ ] Medium Level
        - [ ] Unique
        - [ ] Regex
- [ ] Database Connection
### Define Folder Structure
- Based on:
    - [Official Doc](https://golang.org/doc/code.html#Introduction)
    - [Github Golang Standards](https://github.com/golang-standards/project-layout)
    - [Rakyll](https://rakyll.org/style-packages/)
#### Structure
- _go-request-validator/_ - Package Root Dir
    - _test/_   - Folder for tests
    - _tools/_  - Helpers, global func
    - _pkg/_    - Packages that can be used by everythign 
    - _config/_ - App config files
    - _cmd/_    - Main Applications
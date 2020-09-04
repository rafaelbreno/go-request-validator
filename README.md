# Go Request Validator
- My first project after I've started the [_Go4Noobs Repo_](https://github.com/rafaelbreno/go4noobs)
## About
- This project will be a _Request Validator_
- My inspiration is the [Laravel Request Validator](https://laravel.com/docs/7.x/validation#quick-writing-the-validation-logic)
## TODO List
- [x] Define Folder Structure
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
    - _pkg/_    - Packages that can be used by everything
    - _config/_ - App config files
    - _cmd/_    - Main Applications

#### Code Snippets
- Receiving json data
-   ```go
        package main

        import (
            "log"
            "net/http"
            "encoding/json"
        )

        type test_struc struct {
        //  varName varType `json:"keyName"`
            Name    string  `json:"name"`
            Email   string  `json:"email"`
            Age     int     `json:"age"`
        }

        // Func Handler
        func getReq(w http.ResponseWriter, req *http.Request) {
            dec := json.NewDecoder(req.Body)
            var t test_struc
            err := dec.Decode(&t)
            if(err != nil) {
                panic(err)
            }
            log.Println(t)
        }

        func main()  {
        //  http.HandleFunc("/endpoint", func)
            http.HandleFunc("/", getReq)

        //  http.ListenAndServe(":port", nil)
            http.ListenAndServe(":8080", nil)
        }
    ```
- Sending:
-   ```json
        {
            "name": "Rafael Breno de Vasconcellos Santos",
            "email": "rafael@coldletter.moc",
            "age": 20
        }
    ```
- The struct _t_ will be:
    - _{Rafael Breno de Vasconcellos Santos rafael@coldletter.moc 20}_

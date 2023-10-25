<h1 align="center">
  <a href="https://github.com/pedidosya">
  	<img src="https://img.pystatic.com/pedidosya-logo.svg" alt="PedidosYa" width="200">
  </a>
  <br>
  <br>
  @project_name@
  <br>
</h1>
<h4 align="center">Add a description here about your project, explain very briefly what it does.</h4>
<p align="center">
  <a href="#key-features">Key Features</a> •
  <a href="#dependencies">Dependencies</a> •
  <a href="#how-to-use">How To Use</a> •
  <a href="#documentation">Documentation</a> •
  <a href="#faq">FAQ</a>
</p>

## Key Features

* Solves this problem
* Solve this another problem
* Avoid of having some manual process

## Dependencies

* Platform (Java/Go/Python) and version
* External Services
* Internal Services
* Databases
* Etc.

## How To Use

### Running the project

Before running the project please ensure that all the dependencies are installed in your system. Then follow the next:

1. First step, start some database

    ```
    start database command
    ```

2. Second step, run some internal dependency

    ```
    start internal dependency command
    ```

3. Run the project itself

    ```
    run the project
    ```

### Running the tests

In order to run the project tests you need to execute the following command:

```
some command to test
```

### Enable secret check in pre-commit

With gitleaks and pre-commit we can control the secret check on your last commit, note that it can be omitted but we do not recommend it.

Run the following command in the root of your component:

    ```
    git config core.hooksPath .github/hooks
    

### Swagger docs

This project brings native integration for swagger specs, including example operations as a baseline. This project exposes the following endpoints on non productive environments:

* */doc/swagger*: YAML definition (located at `api/swagger.yaml`), to set up as a custom link and then this service api spec in [Jarvis](https://jarvis.peya.app/)
* */doc*: HTML view (located at `api/api.html`), to see either locally or outside the Jarvis openAPI viewer


## Documentation

* [Link to a confluence doc](https://www.example.com)
* [Link to another doc](https://www.example.com)

## FAQ

* If you want to add new features to this project please [see the contribution guide](.github/CONTRIBUTING.md)
* Questions?, <a href="mailto:someone@pedidosya.com?Subject=Question about Project" target="_blank">write here</a>

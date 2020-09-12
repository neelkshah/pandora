### Little pointers

* Create a subdirectory under `pkg` for each new package.
* Visibility and scope are based on packages as a unit, so organize code accordingly.
* Split a package into files for readability and maintainability.
* Identifiers having capitalized names are visible outside the package. This can be as low-grained as individual fields in a `struct`. 
* All global variables and configurations will go into their respective files in the `config` package.
* In Go, tests are not different from regular code. Rules of visibility apply to test code too. So low-level unit tests are present in a separate file within the package itself.
* End-to-end test cases can be placed in a separate `test` package. This will help organize test code better, and make profiling and benchmarking easier; whatever profiling has to be done can be done within the `test` package.

### A primer for writing unit tests

(Check the `pkg\linkedlist\linkedlist_test.go` file for reference)
* Name the file ending in `_test.go`
* Import the `testing` package.
* Write a separate function for each test.
* Unit tests follow the arrange, act, assert layout. First, arrange all the local variables needed (parameters, connectors, configs, etc). Then execute the action that is being tested. Finally, assert the correctness of the outcome. Don't forget to clean up after a test (file descriptors, slices, connectors, etc).
* Assertion of correctness has to be complete - for example, in the case of the linkedlist, we have verified the number of elements, null fields, the actual values and so on.
* In VS Code, you can check code coverage by pressing `Ctrl+Shift+P` and selecting `Go: Toggle test coverage in current package`. It highlights the code by test coverage. Please make sure you cover as much code as possible. This includes constructors and CRUD methods especially, since those are often the source of memory leaks and null errors and the like. Print functions are difficult to cover and can be avoided.
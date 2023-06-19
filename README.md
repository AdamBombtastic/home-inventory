### Description
`home-inventory` is a go package that seeks to create an open-source solution for home-inventory management. Currently, it's a small hobby, architectural, and clean design project -- But we have (grand) ambitions. 

### Package Structure
`home-inventory`
- `cmd` : Contains entry points to the application (currently just a silly cmd-line app)
- `data`: Contains some placeholder data while we work on database connections (temporary)
- `pkg` : Contains the go modules that support the application
    - `events` : This package was created pre-maturely, though, it contains some simple in memory event driven helpers. (unused)
    - `inventory`: Contains entity definitions and store interfaces for the `item` entity.
    - `requirments`: Contains entity definitions and store interfaces for the `requirement` entity.
    - `interactors`: Contains files that interact with entities via business logic to produce a use-case.

### Roadmap
- [X] Start the project (Excellent!)
- [ ] TODO (This should give you incredible confidence in both my planning and execution abilities.)

### Description
`home-inventory` is a go package that seeks to create an open-source solution for home-inventory management. Currently, it's a small hobby, architectural, and clean design project -- But we have (grand) ambitions. 

### Package Structure
`home-inventory`
- `cmd` : Contains entry points to the application (currently just a silly cmd-line app)
- `data`: Contains some placeholder data while we work on database connections (temporary)
- `pkg` : Contains the go modules that support the application
    - `events` : This package was created pre-maturely, though, it contains some simple in-memory event driven helpers. (unused)
    - `core` : Contains the core application definition, entities, and business logic.
        - `entities`: Contains entity definitions that are used in the business logic
        - `interactors`: Contains files that interact with entities via business logic to produce a use-case. Will be moved up a layer or re-named as it is not 1 to 1 with the definition yet.
        - `services`: Contains service interface definitions that the `core` depends on as well as mocks.
    - `stores`: Contains basic, yaml-based implementations of the required services in `core` 
        - `inventory`: Contains store implementation for the `item` entity.
        - `requirments`: Contains store implementation for the `requirement` entity.


### Roadmap
- [X] Start the project (Excellent!)
- [ ] TODO (This should give you incredible confidence in both my planning and execution abilities.)


### Use Cases
- [X] Is my Inventory Valid? ex: I want my house to always have at least 5 snacks -- Does it?
- [ ] I want to be able to ingest "purchases" and start tracking "ingest" points?? (Little vague)
- [ ] I want to be able to create, remove, update, and delete items from my inventory.
- [ ] I want to be able to create, remove, update, and delete requirements.

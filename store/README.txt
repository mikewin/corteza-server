Goal
 - multi-repository (different DB) support
 - refactor old raw-sql migrations to logical and dynamic schema provisioning
 - (tbd) consolidation of system & compose (& messaging) subsystems

File system:
 /repository            Holds all repository logic for all subsystems, for all core repository implementation
   /internal            Internal repository tools (pkg/ql, pgk/rh should be moved here)
 /<implementation>      Individual core repository implementation
                        [mysql|postgresql|redis|memory|sqlite|elasticsearch|mongo]
   /schema              Schema provisioning for individual repository implementation





/corteza/store/rdbms
    Basic RDBMS logic that should be reused across all
    RDBMS-like implementations

/corteza/store/<rdbms-store-type>/sql_*.go
        Database table definitions (used by provision_*.go)

/corteza/store/<rdbms-store-type>/provision_*.go
    Provision logic, cascading (eg: Provision => ProvisionSystem => ProvisionUsers)



FAQ:
Why are we changing create/update function signature (input struct is no longer returned)?
Because store functions are no longer manipulating the input.

Why naming inconsistency between search/lookup and create/update/...?
To ensure function names sound more natural

Why changing find prefix to search/lookup?
To be consistent with actions

Why do we use custom mapping (and not db:... tag on struct)?
Separation of concerns
 + consistency with store backends that do not support db tags
 + de-cluttering types* namespace


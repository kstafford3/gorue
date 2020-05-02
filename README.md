# gorue text interface engine

The gorue text interface engine provides a framework for state-based Read-Eval-Print Loop (REPL) interfaces.

Gorue describes it's internal state, accepts user input, interprets that input, modifies internal state accordingly, then returns to the start of the loop.

### Design
Gorue is designed to enable multiple loops running concurrently, with each loop acting on an identified context.

The provided components will dictate communication with the user, and how the user's commands are interpreted.

Aside from its implementation of the REPL pattern, gorue makes few assumptions about the application implementation.

#### Serialized State
The format of the context identifier and stored state are left to the user, only byte-serialized forms are passed by gorue.

#### User Interface
Interaction with the user is performed through UTF-8 strings.
Each iteration of the loop has a single string prompt from the application and a single string response from the user.


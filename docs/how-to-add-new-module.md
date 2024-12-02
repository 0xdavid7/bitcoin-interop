# Creating a New Cosmos SDK Module

This guide explains how to create a new module for the Cosmos SDK framework, using best practices and common patterns.

## Table of Contents

1. [Module Structure](#module-structure)
2. [Basic Components](#basic-components)
3. [Implementation Steps](#implementation-steps)
4. [Testing](#testing)

## Module Structure

A typical Cosmos SDK module should have the following directory structure:

```bash
x/mymodule/
├── client/
│ ├── cli/ # Command line interface commands
│ └── rest/ # REST API endpoints
├── keeper/ # State management
├── types/ # Type definitions and interfaces
├── abci.go # ABCI interface implementation
├── genesis.go # Genesis state management
├── handler.go # Message handlers
└── module.go # Module interface implementation
```

## Basic Components

## Basic Components

### 1. Module Interface (module.go)

- Implements `AppModule` and `AppModuleBasic` interfaces
- Handles module initialization and state management
- Example:

```go
go
type AppModule struct {
    keeper Keeper
    // Add other required keepers
}

func NewAppModule(k Keeper) AppModule {
    return AppModule{
        keeper: k,
    }
}
```

### 2. Keeper (keeper/keeper.go)

- Manages module state
- Handles CRUD operations
- Example:

```go
type Keeper struct {
	storeKey sdk.StoreKey
	cdc codec.BinaryCodec
	paramstore paramtypes.Subspace
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey sdk.StoreKey,
	ps paramtypes.Subspace,
) Keeper {
	return &Keeper{
		storeKey:   storeKey,
		cdc:        cdc,
		paramstore: ps,
	}
}
```

### 3. Types (types/types.go)

- Define messages, events, and other data structures
- Implement message validation
- Example:

```go
type MyMessage struct {
	Sender sdk.AccAddress
	// Add other fields
}

func (msg MyMessage) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty")
	}
	return nil
}
```

## Implementation Steps

### 1. Initialize Module Structure

```bash
mkdir -p x/mymodule/{client/{cli,rest},keeper,types}
```

### 2. Define Module Constants and Types

1. Create basic types and interfaces in `types/`
2. Define events, errors, and keys
3. Implement message types and their validation

### 3. Implement Keeper

1. Create store key handling
2. Implement CRUD operations
3. Add query methods
4. Set up parameter store

### 4. Add Message Handlers

1. Create message handling logic in `handler.go`
2. Implement message routing
3. Add event emission

### 5. Set Up ABCI Interface

1. Implement `BeginBlocker` and `EndBlocker`
2. Handle state transitions
3. Process queued operations

### 6. Add CLI and REST Endpoints

1. Create CLI commands for queries and transactions
2. Implement REST API endpoints
3. Add route handlers

### 7. Implement Genesis Handling

1. Define genesis state structure
2. Add import/export functionality
3. Implement state validation

## Testing

### 1. Unit Tests

- Test individual components
- Use mocks for dependencies
- Example:

```go
func TestHandleMessage(t testing.T) {
	ctx := sdk.NewContext(...)
	k := keeper.NewKeeper(...)
	msg := types.NewMyMessage(...)
	result, err := HandleMyMessage(ctx, k, msg)
	require.NoError(t, err)
	// Add assertions
}
```

### 2. Integration Tests

- Test component interaction
- Simulate real scenarios
- Example:

```go
func TestModuleFlow(t testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(...)
	// Test complete flow
// Setup -> Execute -> Verify
}
```

### 3. CLI Tests

- Test command line interface
- Verify flag handling
- Example:

```go
func TestGetQueryCmd(t testing.T) {
cmd := cli.GetQueryCmd()
// Test command execution
}
```

## Notes:

- RegisterRESTRoutes: (deprecated)

  - Registers traditional REST API endpoints
  - Used for HTTP/REST API access
  - Being deprecated in favor of gRPC
  - Example endpoint: GET /evm/params

- RegisterGRPCGatewayRoutes:

  - Translates REST calls to gRPC
  - Allows REST clients to access gRPC services
  - Modern replacement for RegisterRESTRoutes

- GetTxCmd:

  - Returns transaction commands for CLI
  - Used for command-line interactions
  - Example: `axelard tx evm link ...`

- GetQueryCmd:

  - Returns query commands for CLI
  - Used for command-line queries
  - Example: `axelard query evm params`

- RegisterInterfaces:

  - Registers interfaces for module types
  - Used for codec registration
  - Example: `RegisterInterfaces(registry)

- RegisterServices:

  - Core service implementation
  - Handles both messages (transactions) and queries
  - Used by both gRPC and gRPC-Gateway

- Route: (deprecated)
  - Legacy message routing system
  - Defines how messages are routed to handlers
  - Being deprecated in favor of RegisterServices
  - Part of the older Amino codec system

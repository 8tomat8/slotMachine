# Implementation of simple Slot Machine

Implemented according to the description from here https://wizardofodds.com/games/slots/atkins-diet/

**IMPLEMENTATION DOES NOT COVERED WITH TESTS**

This implementation was tested only manually. To use it in production add unit tests.

## Usage

To build and run this app you need to have installed Go. All dependencies are included into repository.

    git clone https://github.com/8tomat8/slotMachine.git
    cd slotMachine
    go build -o app github.com/8tomat8/slotMachine/cmd/server
    ./app

Application has only one HTTP endpoint to create a spin

    POST /api/machines/atkins-diet/spins

Request body should be a properly signed JWT token with claims:

    {
	    uid: "any string",  // user id
	    chips: 10000,       // chips balance
	    bet: 100           // bet size
    }

**Default sign key for JWT is: Foo-Bar-Baz-42**

You can replace it with your own using environment variable **SLOTS_JWT_SECRET**

Valid example of request that must work with default settings:

    curl -vvv -XPOST http://127.0.0.1:8080/api/machines/atkins-diet/spins -d 'eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiJhbnkgc3RyaW5nIiwiY2hpcHMiOjEwMDAwLCJiZXQiOjEwMH0.odtrW91_wZTpPZ7FvczUKenLL-hCgHoNLipmuxQN96AdYZpZDHVZeeZ2HyqD4Z5TighNNsuN7gILcZe5-u6Cmw'

Response:

    {
        "total": 50, 
        "spins": [
            {
            "Type": "main", 
            "Total": 50, 
            "Stops": [15, 12, 18, 5, 19]
            }
        ], 
        "jwt": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiJhbnkgc3RyaW5nIiwiY2hpcHMiOjk5NTAsImJldCI6MTAwfQ.zGPNIZ29oo14gae3xkYAHw8WhfcvhAXUdeZreQ0qwuO8-UUw4UilO2uetDKbFzD3bJhC8D8xpV9jJwJppar8Hg"
    }

Where jwt value is updated and signed request data

## Additional

- To see debug logs set env. var. **SLOTS_ENVIRONMENT=dev**
- To change API host and port use env. var. **SLOTS_API_ADDR=0.0.0.0:8080**
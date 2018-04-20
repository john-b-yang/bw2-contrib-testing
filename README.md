# BETS Driver Testing Framework
Universal testing framework for BETS driver code

### Driver Commonalities:
#### main.go
- Creates 'service' from bwClient.RegisterService() function
- Register each device as independent interface using service.RegisterInterface() function
- Ponum for creating and decoding message structs

#### "driver".go

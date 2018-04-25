# BETS Driver Testing Framework
Universal testing framework for BETS driver code

### Description
- Build the testing framework: go build -o TestDriver main.go framework.go
- 'Main.go' is the file to be editted and customized for different drivers
- User is responsible for editing the following: A. Assertion function to check if method outputs are in correct format B. Package struct for testing set request C. Assortment of string parameters specific to each driver (payload object number, base URI, client name, etc.)

### Driver Commonalities:
#### main.go
- Figure out how to work in discovery functions (i.e. Site -> Pelican, NWS -> Station)

#### "driver".go
- Creates 'service' from bwClient.RegisterService() function
- Register each device as independent interface using service.RegisterInterface() function
- Define which parameters can be abstracted to create more general testing framework

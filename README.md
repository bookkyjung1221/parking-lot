# Parking Lot Assignment (Golang)

## Supported Commands

- `create_parking_lot` <`n`>  
  To create a Parking lot. Where `n` is the size of the parking lot

- `park` <`registration_number`> <`colour`>  
  To park the car in the parking lot and prints the allocated slot in the parking lot. Where `registration_number` is given registration number for the car and `colour` is given colour for the car

- `leave` <`slot`>  
  To leave the parking lot from desired slot and prints the leaving slot. given slot number. Where `slot` is given sloat number

- `status`  
  To check the status of Parking Lot

- `slot_numbers_for_cars_with_colour` <`colour`>  
  To prints the registration number of the cars for the given colour. Where `color` is given colour

- `slot_number_for_registration_number` <`registration_number`>  
  prints the slot number of the cars for the given number. Where `registration_number` is given registration number.

- `registration_numbers_for_cars_with_colour` <`colour`>  
  To prints the slot number of the cars for the given colour. Where `colour` is given colour.

- `exit`  
  To exit program

## Running Application

#### Build Dockerfile

```docker
docker build -t docker-go .
```

#### Run Docker image in Interactive mode

```docker
docker run -it --rm --name parking-lot docker-go
```

#### Running the application in File mode

```golang
go run main.go input.txt
```

#### Running the application in Interactive mode

```golang
go run main.go
```

#### Run Test Cases

```
cd model
```

```golang
go test -v
```

# PARK FINDER

## Website Functionality

To create your own files, import park_finder.sql
To use test data, import populated.sql and use the credentials:
  Email: admin@pf.com
  Password: AAAAA

## How to run
open this folder in the installed directory.
You can proceed to run this in various modes:
1. In golang:
  > go run server.go

2. As a windows binary:
  > parkfinder.exe

Server runs on port 4000 as it has to be binded on a port in which one can access.
Open the server running on port 4000 in your browser
Go to your browser and visit: http://localhost:4000 or https://localhost:40001

## Whys use golang
### Advantages of using golang.
1. It is fast and quick while allowing for multi templating using server side rendering.
2. Allows for domain driven driven design which is scalable and easy for collaboration.
3. Allows for definition of your own data types.
4. It's cross platform as Compiles into different binaries from windows or linux.
## Unique Features.
1. Custom built in template engine that loads templates from ken/lib/ui/tmpl
2. Amazing Admin functions that can  be scalled into a production environment like:
    1. Create and manage newsletters.
    2. Keep track of contact me's
3. Create park reviews to gives users broader perspective



## Development Pipeline
1. Create Domain/DB Functions
2. Create Users
3. Create Parks
4. List Parks
5. Create and List reviews, (should be an ajax call)
6. Create contact me. (A list with an ajax call)
7. Create NewsLetter (A list with an ajax call)

An Administrator should be able to create parks and delete reviews

## Data types definition
Contains entities in which the various data to be persisted or read from the db is to be accessed.
>Directory: ken/lib/entities

## User and Functionality Documentation: In the domain folder
> Directory: ken/lib/domain

Users:
  1. CreateUser: Creates a user and writes it into the DB
  2. ListUsers: LIsts the users available in the database
  3. ViewUser: Takes in a user id and returns the data belonging to the user in question.
  4. Authenticate: Takes in a user email and password then authenitcates them returning the user data and a boolean value indicating success or failure.

Parks:
  1. CreatePark Takes in park data and writes it to the Database.
  2. Update park: Ability to update park data.
  3. ViewPark: Returns the park in question.
  4. ListParksByLocation: Returns a list of parks by the location.

Newsletter:
  1. CreateNewsLetter: CReates a newsletter.
  2. ViewNewsLetter: Returns a newsletter in question.
  3. ListNL: List Newsletters.
  4. MarkNLAsHandled: Makrs a newsletter as handled

ContactUS:
  1. CreateContactUS: Creates a contact us and writes it to db.
  2. ViewCU: Returns a given contact us form that has been filled.
  3. ListCU: List given contact us handled or handled as a boolean input.
  4. MarlCUHandled: Marks a given contact me message as handled or not.


Images uploaded are stored incide the static folder in an uploaded folder each with the UUID of the park Id in question.

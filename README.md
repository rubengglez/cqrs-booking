# CQRS Booking
You should implement a simple booking soluiton in CQRS architecture.

## About CQRS
CQRS for Command Query Responsibility Segregation Pattern.

A query returns data and does not alter the state of the object. A command changes the state of an object but does not return any data.

We will split our code in read ans write code to really live this pattern

## Booking subject
We want to make a booking solution for one hotel.

The users stories are :

	As a user I want to see all free rooms.
	As a user I want to book a room.
	As a user I want to cancel a reservation

They want to use the CQRS pattern, To do that we will have:

The Booking struct contains

	client id
	room name
	arrival date
	departure date

The Room struct contain only

	room name


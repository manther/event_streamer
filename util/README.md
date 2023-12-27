Practice project. 
Objective: Consider Producer / Consumer seperation of responsibilities. Design an application that shows the contract negotiation defined in interfaces that support the consumer and enable the producer. The interfaces should not over-expose the consumer to more than needed. The functionality the producer provides should not be overly opinionated. The stream package should be testable. 

The interfaces in the stream package are short and descriptive. 
IStreamManager is responcible for managing a stream connection. Managers have two unexported values backed with getters. The values are initially set with constructors. 
The constructors for the stream managers control the rate that events come back to the consumer and the total amount of messages the consumer wants to receive. A further exantion on this might be to send unlimited amount of events if the rate is set to zero. 
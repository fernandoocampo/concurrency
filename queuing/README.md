# QUEUING

accepting work even though the pipeline is not yet ready for more.

## Risks

* Adding a queue prematurely can hide synchronization issues such deadlocks and livelocks.

* Common mistakes people make when trying to tune the performance of a system: introducing queues to try and address performance concerns.

* Queuing will almost never speed up the total runtime of your program. It will only allow the program to behave differently.

* By introducing a queue at the entrance to the pipeline, you can break the feedback loop at the cost of creating lag for requests.

## Benefits

* introducing a queue isn’t that the runtime of one of stages has been reduced.
* the time it’s in a blocking state is reduced.
* This allows the stage to continue doing its job.
* users are likely to experience lag in their request, but they wouldn't be denied service altogether.
* Decouple stages so that the runtime of one stage has no impact on the runtime of another.

### Improve performance?

* If batching requests in a stage saves time.
* If delays in a stage produce a feeback loop(--1) into the system.

--1 A feedback loop is the part of a system in which some portion (or all) of the system's output is used as input for future operations.

--2 death spiral: the rate that work enters (ingress) > rate in which it exits the system

## Where

* At the entrance to your pipeline.

* In stages where batching will lead to higher efficiency.

## Theory

Little’s Law algebraicly. It is commonly expressed as: L=λW, where: 

L = the average number of units in the system.  
λ = the average arrival rate of units.  
W = the average time a unit spends in the system.

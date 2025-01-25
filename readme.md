## Trigger Management Application :
-- this is a trigger management application where there are user defined creation methods. 

-- A user can make two types of trigger

            -- api trigger
            
            -- scheduled trigger - of two types
            
                    -- recurring trigger, non-recurring triggers.
                    
-- the website is hosted on render which can be accessed using trigger 

https://trigger-management.onrender.com/{endpoints}


-- where the endpoints are

                        **for triggers**
                        -- for posting a trigger r.POST("/trigger")

                        -- for fetching all triggers r.GET("/allTriggers") 

                        -- for updating a trigger r.PUT("/updateTrigger") 

                        **for events**
                        -- for getting all triggers r.GET("/getEvents")

                        -- for updating a trigger r.PUT("/updateEvents")

                        -- for deleting an event r.DELETE("/deleteEvent")


## For Testing API triggers :
-- **payload** 


 {

    "name": "Test Trigger",

    "type": "api",

    "execution_time": "2025-01-26T09:55:00Z",

    "endpoint": "https://jsonplaceholder.typicode.com/posts",

    "payload": {

                "title": "foo",

                "body": "bar",

                "userId": 1

            },

    "is_recurring": false
 }
 
 use this payload because it is giving the perfect output, in case of other endpoints you might face api key issues and other issues which is not part of my issue.


## Frontend Integration :
-- as of now i was not able to integrate the frontend because of no knowldge of frontend. 

I deployed a frontend website using vercel v0 AI IDE but was giving multiple error and thats why i decided to not to publish it, because it will make no sense.


## Improved Idea of Mine which i was not able to implement
-- The improved idea involves using Redis sorted sets to manage and execute scheduled triggers efficiently. 

When a scheduled trigger is received, it is saved in the database and added to Redis with its execution time as the score. 

A concurrent cron job runs periodically to fetch triggers from Redis that are due for execution, based on the current time. 

These triggers are processed to create events in the database. 

For recurring triggers, their execution time is updated by adding the interval, while non-recurring triggers are moved to an infinite time to prevent further execution. 

On server restarts, triggers from the database with in_redis=false are resynced into Redis to ensure continuity. This approach combines Redis for fast execution scheduling with the database for persistence and fault tolerance.

## My message to shobhit
-- this app may not be very perfect because i was a bit busy with my final year project presentation and research paper work also my end sems were going on which ended on Monday 27 Jan. 

Apart from it there are very few resources in go that makes it quite difficult. 

Still i am ready to be evaluated based on my code but i can improve it.

I am really very happy to be part of this assignment.

Thanks & Regards.


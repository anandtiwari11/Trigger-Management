## Trigger Management Application :
-- this is a trigger management application where there are user defined creation methods. 

-- A user can make two types of trigger

            -- api trigger
            
            -- scheduled trigger - of two types
            
                    -- recurring trigger, non-recurring triggers.
                    
-- the website is hosted on render.com

**local setup**

-- please clone the repository in your local

-- install go

-- run go mod tidy

-- since all the environment variables are directly initialised in the code itself we do not need to set up the .env file
due to some render.com deployment issue but i believe the best practice is to put .env in gitignore


**cost of running** 
in ruppe := 0 (as i am currently using free versions of render.com available services)

https://trigger-management.onrender.com/{endpoints}


-- where the endpoints are

                        **for triggers**
                        -- for posting a trigger r.POST("/trigger")
                        https://trigger-management.onrender.com/trigger

                        -- for fetching all triggers r.GET("/allTriggers") 
                        https://trigger-management.onrender.com/allTriggers

                        -- for updating a trigger r.PUT("/updateTrigger") 
                        https://trigger-management.onrender.com/updateTrigger

                        **for events**
                        -- for getting all triggers r.GET("/getEvents")
                        https://trigger-management.onrender.com/getEvents

                        -- for updating a trigger r.PUT("/updateEvents")
                        https://trigger-management.onrender.com/updateEvents

                        -- for deleting an event r.DELETE("/deleteEvent")
                        https://trigger-management.onrender.com/deleteEvent


## For Testing API triggers :

https://trigger-management.onrender.com/trigger
-- **payload** 


 {

    "name": "Test Trigger",

    "type": "api",

    "message": "Hello World",

    "endpoint": "https://jsonplaceholder.typicode.com/posts",

    "payload": {

                "title": "foo",

                "body": "bar",

                "userId": 1

            },

    "is_recurring": false
 }
 
 use this payload because it is giving the perfect output, in case of other endpoints you might face api key issues and other issues which is not part of my issue.

## for testing the scheduled triggers

https://trigger-management.onrender.com/trigger
-- **payload**

**nonrecurring**

{

    "name": "Test Trigger",

    "type": "scheduled",

    "execution_time": "2025-01-28T16:27:00.893672+05:30",


    "is_recurring": false

}

--

**recurring**
https://trigger-management.onrender.com/trigger

--


{

    "name": "Test Trigger",

    "type": "scheduled",

    "is_recurring": true

}

--




## Frontend Integration :
-- as of now i was not able to integrate the frontend because of no knowldge of frontend. 

I deployed a frontend website using vercel v0 AI IDE but was giving multiple error and thats why i decided to not to publish it, because it will make no sense.

So it is my humble request to you to please test this service on Postman. 

https://oht7kimnizxv3e6o.vercel.app/ hosted on vercel


## Approach for scheduled triggers

-- i am using gocron module in go to manage these triggers, which works similar to celery-beat /clery in python.

-- the purpose of using this module is to create concurrent tasks which will create events in the DB and 

go cron is used to manage these concurrent go-routines and create tasks.

i am trying to use faktory which is an event queue manager to further optimise the process and trying to write a readable and reusable code



## the db variables are

DB_HOST=dpg-cuasfla3esus73eolmug-a.oregon-postgres.render.com
DB_USER=trigger_management_oz3e_user
DB_PASSWORD=WI7wiGhHsOB2GknfXvKK2nYSUWSq305D
DB_NAME=trigger_management_oz3e
DB_PORT=5432
DB_SSLMODE=require

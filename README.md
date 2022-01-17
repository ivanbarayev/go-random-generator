# go-random-generator


````dotenv
ENV FILE
SENSOR_TOTAL=3
SENSOR1_RANGE=1-98 //for temperature
SENSOR2_RANGE=100-498 //for air quality
SENSOR3_RANGE=1-500 //for electricity

SENSOR1_URL=127.0.0.1/temp  
SENSOR2_URL=127.0.0.1/air  
SENSOR3_URL=127.0.0.1/electrcity  

SENSOR1_INTERVAL=1  
SENSOR2_INTERVAL=1  
SENSOR3_INTERVAL=2  

GENERATOR_MODE=on
````

````
if GENERATOR_MODE=on the system up with value generator
if GENERATOR_MODE=off the system up as a REST API
DEFAULT is on
````

````
ROUTES (all methods are POST)
FOR ALL Data > ip:port/sensor/generator

POST DATA
{
   "sensor_id": 1
}
````
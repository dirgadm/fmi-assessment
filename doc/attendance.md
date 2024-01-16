# Attendance API Spec

## Create Attendance
Endpoint : POST /v1/attendance

Request Header :

- BEARER_TOKEN : Token (Mandatory)

Description: Create an attendance based on geolocation of spesific user that already login to the system

Request Body (Success) :

```json
{
    "latitude":37.7749, // required
    "longitude":-122.4194 // required
}
```

Response Body(Success)

```json
{
    "code": 200,
    "status": "success",
    "data": {
        "message": "Success Create Attendance"
    }
}
```
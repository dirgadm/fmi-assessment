# Upload API Spec

## Upload Photo
Endpoint : POST /v1/upload

Request Header :

- No Header

Description: endpoint to upload a image or photo profile for spesific user. Only png and jpg extension allowed

Request Body (Success) :

```formdata
{
    key:"{{ key_name }}", //required
    value:"form file" // required
}
```

Response Body(Success)

```json
{
    "code": 200,
    "status": "success",
    "data": [
        {
            "file_name": "image_photo12312312412414.png"
        }
    ]
}
```
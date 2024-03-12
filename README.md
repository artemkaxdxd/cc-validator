## Credit Card Validator

### Running the Project

1. **Clone the repository:**

```bash
git clone https://github.com/artemkaxdxd/cc-validator.git
```

2. **Run the Dockerfile:**

```bash
docker build -t credit-card-validator .
docker run -p 8080:8080 --name cc-validator credit-card-validator
```

This will build a Docker image for the project and run it on port 8080.

3. **Send a validation request:**

Make a POST request to `http://localhost:8080/validate` with the following body in JSON format:

```json
{
  "number": "4111111111111111",
  "month": "12",
  "year": "2028"
}
```

Replace `"4111111111111111"` with the actual credit card number you want to validate.

The response will indicate whether the credit card is valid or not.
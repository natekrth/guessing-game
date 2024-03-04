# Guessing Game

This guessing number game between 1 to 10.

## Project Setup Guide

### Backend Setup
- Change directory backend/
    ```
    cd backend/
    ```
- Create .env file and set up MySQL database
    Here is the example, guessing is the name of the table
    ```
    JWT_SECRET_KEY=my_secret_key
    MYSQL="root@tcp(127.0.0.1:3306)/guessing?charset=utf8mb4&parseTime=True&loc=Local"
    ```
- Run
    ```
    go get .
    go run . // run and auto migrate to MySQL
    ```

### Frontend Setup
- Change directory to frontend/
    ```
    cd frontend/
    ```
- Install dependencies
    ```
    npm i
    ```
- Start the application
    ```
    npm start
    ```

##### Take a break and enjoy Guessing Game ^__^

## API
### Authentication Endpoints

#### Register
- **Endpoint**: `/register`
- **HTTP Method**: `POST`
- **Description**: Register a new user.
  
  **Request Body**:
    - `Username` (string)
    - `Password` (string)

#### Login
- **Endpoint**: `/login`
- **HTTP Method**: `POST`
- **Description**: Login for authenticated
  
  **Request Body**:
    - `Username` (string)
    - `Password` (string)

  **Return JSON**
   - `token`

### Guess Endpoints
#### Guess
- **Endpoint**: `/guess`
- **HTTP Method**: `POST`
- **Description**: Guess the number in the game
  
  **Request Body**:
    - `guess` (integer)

  **Return JSON**
   - `attempts`(the number of attempt user takes to guess)

#### Guess Answer
- **Endpoint**: `/guess`
- **HTTP Method**: `GET`
- **Description**: Get the answer of the number current game

  **Return JSON**
   - `answer`
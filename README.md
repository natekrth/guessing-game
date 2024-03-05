# Guessing Game

This guessing number game between 1 to 10.
- Application Frontend: ReactJS
- Appication Backend: Golang
- Database: MySQL

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
    - `username` (string)
    - `password` (string)

#### Login
- **Endpoint**: `/login`
- **HTTP Method**: `POST`
- **Description**: Login for authenticated
  
  **Request Body**:
    - `username` (string)
    - `password` (string)

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

#### Guessing Update Answer
- **Endpoint**: `/guess/update`
- **HTTP Method**: `PATCH`
- **Description**: A new number for guessing (random)

  **Request Body**:
    - `update` (integer) the seed number for random new guess

  **Return JSON**
   - `answer`

### Delete User
- **Endpoint**: `/user/delete`
- **HTTP Method**: `DELETE`
- **Description**: Delete a user account from database

  **Request Body**:
    - `username` (string)

## User Interface
### Register Page
<img width="1282" alt="register-page" src="https://github.com/natekrth/guessing-game/assets/77069581/ca7bdd11-cde9-4ae4-93bb-f99764c2b6e1">

### Login Page
<img width="1282" alt="login-page" src="https://github.com/natekrth/guessing-game/assets/77069581/695006c4-89d9-4bbd-8892-ecd15561bd82">

### Guessing Game Page
<img width="1289" alt="guessing-game-page" src="https://github.com/natekrth/guessing-game/assets/77069581/d6f6bf6f-b95d-46c9-8856-23a2d9277b19">



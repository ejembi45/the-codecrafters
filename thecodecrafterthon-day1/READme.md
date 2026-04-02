# CLI Calculator (Go)

A simple Command-Line Interface (CLI) calculator written in Go.
This application allows users to perform basic arithmetic operations interactively from the terminal.

---

##  Features

* Addition
* Subtraction
* Multiplication
* Division (with zero-division handling)
* Interactive CLI menu
* Input validation for numbers
* Option to continue or quit after each operation

---

##  How It Works

The program runs in a loop and continuously prompts the user to:

1. Select an operation from the menu:

   * `1` → Addition
   * `2` → Subtraction
   * `3` → Multiplication
   * `4` → Division
   * `5` → Help
   * `6` → Quit

2. Enter two numbers as input

3. The program:

   * Converts the inputs from string to integers using `strconv.Atoi`
   * Calls the corresponding function (`Add`, `Sub`, `Mul`, or `Div`)
   * Displays the result

4. After each operation, the user can:

   * Enter `1` → Continue using the calculator
   * Enter `5` → Exit the program

---

## ⚙️ Functions Overview

### `Add(num1, num2 int64) int64`

Returns the sum of two integers.

### `Sub(num1, num2 int64) int64`

Returns the difference between two integers.

### `Mul(num1, num2 int64) int64`

Returns the product of two integers.

### `Div(num1, num2float64) float64`

Returns the division result.

* Includes a loop to prevent division by zero.
* Prompts the user again if `num2 == 0`.

---

## 🚀 How to Run

### 1. Install Go

Make sure you have Go installed:

```bash
go version
```

If not installed, download it from: https://go.dev/dl/

---

### 2. Save the Code

Save your file as:

```bash
main.go
```

---

### 3. Run the Program

In your terminal, navigate to the project folder and run:

```bash
go run main.go
```

---

## 💻 Example Usage

```text
"welcome, input your number and operator"

Select an Operation by the number

1 . Addition
2 . Subtraction
3 . Multiplication
4 . Division
5 . Help
6 . Quit

Option: 1

Input the First number: 5
Input the Second number: 3

Result = 8

Enter 1 to continue. Enter 5 to quit
```

---

##  Error Handling

* If a non-numeric value is entered, the program shows:

  ```
  Error: Enter a valid number
  ```
* Division by zero is handled by prompting the user to enter another number.

---

##  Project Structure

```
.
├── main.go
└── README.md
```

---

##  Possible Improvements

* Support floating-point input for all operations
* Add more operations (e.g., modulus, power)
* Improve input handling (use buffered input instead of Scan)
* Add unit tests
* Refactor repeated code into reusable functions

---

##  Author

Faith Ochanya Ejembi

---

##  License

This project is open-source and free to use.
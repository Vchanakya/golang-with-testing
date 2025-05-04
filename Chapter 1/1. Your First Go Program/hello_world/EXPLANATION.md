# Go Program Structure Explained

## `package main`

Every Go program starts with a package declaration.  
A package named `main` tells the Go compiler that this package should be compiled into an **executable program** (a program you can run directly), rather than a shared library.  
The `main` function must reside in the `main` package.

---

## `import "fmt"`

This line imports functionality from another package — specifically, the `fmt` package (short for *format*).  
Importing a package allows you to use the **functions, types, and variables** defined within it.

- The `fmt` package contains functions for **input and output**, such as printing text to the console.

---

## `func main()`

This defines the **main function**:

- `func`: Starts a function declaration.
- `main`: A **special function name**. Execution always starts here.
- `()`: Indicates the function takes **no arguments**.
- `{}`: Encloses the body of the function — the instructions it executes.

---

## `fmt.Println("Hello, World!")`

This is the **instruction inside the main function**.

- `fmt`: Specifies we are using something from the `fmt` package.
- `.`: The **dot operator** accesses members of a package.
- `Println`: The function we are calling — it stands for **Print Line**.
  - It prints its arguments to the **standard output** (usually the terminal), followed by a **newline character**.
- `"Hello, World!"`: A **string literal** — the actual message to print, enclosed in double quotes.

---

## `// ...`

Single-line comments start with `//`.

- The compiler **ignores** these lines.
- They are **notes for humans**, helping improve code readability.

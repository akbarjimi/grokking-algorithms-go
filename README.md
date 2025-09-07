[![Repository Badge](https://img.shields.io/badge/GitHub-akbarjimi/grokking--algorithms--go-2ea44f?style=for-the-badge&logo=github)](https://github.com/akbarjimi/grokking-algorithms-go)![Go Language Badge](https://img.shields.io/badge/Language-Go-00ADD8?style=for-the-badge&logo=go) ![Status Badge](https://img.shields.io/badge/Status-In%20Progress-6c757d?style=for-the-badge)

# Grokking Algorithms in Go

This repository contains my personal implementations of the algorithms and data structures discussed in the book ["Grokking Algorithms, Second Edition" by Aditya Bhargava](https://www.manning.com/books/grokking-algorithms-second-edition). The goal is to translate the concepts and examples from the book, often presented in Python, into idiomatic Go.

It serves as a learning journal and a practical exercise to deepen my understanding of algorithms while solidifying my Go programming skills.

## Table of Contents

*   [About the Book](#about-the-book)
*   [Project Structure](#project-structure)
*   [How to Run](#how-to-run)
*   [Contributing](#contributing)
*   [License](#license)

## About the Book

"Grokking Algorithms" is an approachable guide for anyone wanting to learn about algorithms. It focuses on making complex topics intuitive through clear explanations and plentiful illustrations. The book covers essential algorithms like:

*   Binary Search
*   Selection Sort
*   Recursion
*   Quicksort
*   Breadth-First Search
*   Dijkstra's Algorithm
*   Dynamic Programming
*   and more!

My implementations in this repository are based on the examples and challenges presented in this book.

## Project Structure

The repository is organized by chapters, with each chapter containing Go files for the relevant algorithms and accompanying tests.

```
grokking-algorithms-go/
├── Chapter01/
│   ├── binary_search.go
│   └── binary_search_test.go
├── Chapter02/
│   ├── selection_sort.go
│   └── selection_sort_test.go
├── ...
└── README.md
```

## How to Run

To run any of the implemented algorithms or their tests:

1.  **Clone the repository:**
    
    ```
    git clone https://github.com/akbarjimi/grokking-algorithms-go.git
    cd grokking-algorithms-go
    ```
    
2.  **Navigate to a specific chapter:**
    
    ```
    cd Chapter02 # For example, to go to Chapter 2
    ```
    
3.  **Run tests:**
    
    ```
    go test -v
    ```
    
    This will execute all tests within that chapter and show verbose output.
    
4.  **Run a specific file or function (if it has a `main` function or an executable example):**
    
    ```
    go run your_algorithm_file.go
    ```
    

## Contributing

Contributions, suggestions, and feedback are welcome! If you find an error, a more idiomatic Go implementation, or want to add a different approach for an algorithm, please feel free to:

*   Open an issue to discuss.
*   Submit a Pull Request with your changes.

Let's learn and grow together!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

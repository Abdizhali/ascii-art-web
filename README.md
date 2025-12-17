# 🎨 ASCII-Art Web Generator

A simple **Go web server** application that converts user input text into **ASCII Art** using predefined banner files (fonts). It provides a clean, responsive web interface for easy text input, style selection, and output display.

---

## 💡 Project Overview

The main purpose of this project is to implement a core ASCII art generation utility within a web server environment.

### Features:
* **Web Interface:** User-friendly form to input text and select a banner style.
* **Style Selection:** Supports multiple ASCII banner files (e.g., `standard`, `shadow`, `thinkertoy`).
* **Go Server:** Lightweight HTTP server built using Go's standard library.
* **Error Handling:** Displays clear error messages for invalid inputs or server issues.

---

## 🛠️ Technology Stack

The project is built using the following technologies:

* **Backend:** Go (Standard Library `net/http`)
* **Frontend:** HTML5, CSS3
* **Dependency Management:** Go Modules (`go mod`)

---

## 🚀 Getting Started

These instructions will guide you through setting up and running the project locally.

### Prerequisites

You must have the following software installed on your machine:
* **Git** (for cloning the repository).
* **Go** (version 1.20 or later recommended).

### Installation and Run

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/<your-username>/ascii-art-web-dockerize.git
    ```

2.  **Navigate to the project directory:**
    ```bash
    cd ascii-art-web
    ```

3.  **Run the application:**
    The application will start the web server, typically on port `8080`.
    ```bash
    go run main.go
    ```
    *(You should see a message in the console confirming the server is running.)*

4.  **Access the application:**
    Open your web browser and navigate to:
    👉 **http://localhost:8080**

---

## ⚙️ Usage

1.  Open the application in your browser (`http://localhost:8080`).
2.  Enter the text you wish to convert into the **text input area**.
3.  Select one of the available **banner styles** (e.g., `shadow.txt`, `standard.txt`).
4.  Click the **"Generate"** button.
5.  The resulting ASCII art will be displayed on the page.

---

## 👤 Authors

* **aabdizal** 
* **anurdill** 
---

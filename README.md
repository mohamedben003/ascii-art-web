# ASCII Art Web

## Description

ASCII Art Web is a web application that generates ASCII art from text input using different banner styles. The project consists of a Go HTTP server that provides a graphical user interface for the ascii-art generator. Users can input text through a web form, select from three different banner styles (shadow, standard, thinkertoy), and view the generated ASCII art directly in their browser.

The application uses Go templates to render dynamic HTML pages and handles HTTP requests to process user input and display results. The server implements proper HTTP status codes for different scenarios including successful requests, errors, and invalid input.

## Authors

- [mbenboua]
- [herrabba]

## Usage

### Prerequisites
- Go 1.16 or higher installed on your system

### How to Run

1. Clone the repository:
```bash
git clone https://learn.zone01oujda.ma/git/herrabba/ascii-art-web
cd ascii-art-web
```

2. Ensure all banner files are present in the `helper/` directory:
   - `standard.txt`
   - `shadow.txt`
   - `thinkertoy.txt`

3. Run the server:
```bash
go run main.go
```

4. Open your browser and navigate to:
```
http://localhost:8080
```

5. Use the web interface:
   - Enter your text in the input field
   - Select a banner style (standard, shadow, or thinkertoy)
   - Click "GENERATE ASCII ART" button
   - View the generated ASCII art on the page

### Testing Different Inputs

Try these examples:
- Simple text: `Hello`
- Text with newlines: `Hello\nWorld`
- Special characters: `Hello World!`
- Numbers: `12345`

## Implementation Details

### Project Structure
```
.
├── go.mod              # Go module file
├── helper
│   ├── helper.go       # ASCII art generation logic
│   ├── shadow.txt      # Shadow banner file
│   ├── standard.txt    # Standard banner file
│   └── thinkertoy.txt  # Thinkertoy banner file
├── main.go             # HTTP server and request handlers
├── README.md
└── templates
    └── index.html      # Main HTML template
```

### Algorithm

#### 1. Banner Loading
- Each banner file contains ASCII representations for characters (ASCII 32-126)
- Each character is represented by 8 lines of text
- Characters are separated by empty lines
- The `LoadBanner()` function reads the file and creates a map where:
  - Key: rune (character)
  - Value: slice of 8 strings (the ASCII art lines)

#### 2. Text Processing
- Input text is received from the HTML form via POST request
- Newline characters (`\n` and `\r\n`) are normalized to `\n`
- The text is validated to ensure all characters are printable ASCII (32-126)
- Text is split by `\n` to process each line separately

#### 3. ASCII Art Generation
For each line of input text:
- If the line is empty, output a single newline
- Otherwise, iterate through 8 rows (height of ASCII art characters):
  - For each character in the line:
    - Look up the character in the banner map
    - Append the corresponding row of ASCII art
  - Add a newline after each row
- Add spacing between lines of output

#### 4. HTTP Request Flow
```
Client Request → Server Routing → Handler Function → Process Data → Generate ASCII Art → Render Template → HTTP Response
```

**GET `/`:**
- Renders the main page with an empty form
- Returns HTTP 200 OK

**POST `/ascii-art`:**
- Parses form data (text and banner selection)
- Validates input
- Calls helper function to generate ASCII art
- Returns appropriate HTTP status codes:
  - 200 OK: Successful generation
  - 400 Bad Request: Empty text or invalid input
  - 500 Internal Server Error: Banner loading failure or server errors
  - 405 Method Not Allowed: Wrong HTTP method

#### 5. Error Handling
- **Invalid characters**: Returns error if text contains non-ASCII characters
- **Invalid banner**: Returns error if banner selection is invalid
- **Missing banner file**: Returns error with file loading details
- **Empty input**: Returns 400 status with user-friendly error message

#### 6. Template Rendering
- Uses Go's `html/template` package
- Passes data structure (`PageData`) containing:
  - Input text
  - Selected banner
  - Generated ASCII art content
  - Error messages (if any)
- Template conditionally displays results or errors based on the data

### HTTP Status Codes Used
- **200 OK**: Successful request and ASCII art generation
- **400 Bad Request**: Invalid input (empty text, invalid characters)
- **405 Method Not Allowed**: Using wrong HTTP method on endpoints
- **500 Internal Server Error**: Server-side errors (template loading, banner file issues)

### Key Features
- Real-time ASCII art generation
- Three different banner styles
- Input validation and error handling
- Responsive web interface with retro terminal aesthetic
- Preserves user input on errors for better UX
- Proper HTTP status code implementation

## Technologies Used
- **Go**: Backend server and logic (standard library only)
- **HTML/CSS**: Frontend interface
- **Go Templates**: Dynamic page rendering
- **HTTP**: Client-server communication
#
# ASCII Art Web - Docker

A web application that generates ASCII art from text using different banner styles.

## Prerequisites

- Docker installed on your machine

## Quick Start

```bash
# Build the image
docker build -t ascii-art-web .

# Run the container
docker run -p 8080:8080 ascii-art-web

# Open browser at: http://localhost:8080
```

## Common Commands

```bash
# Run in background
docker run -d --name ascii-app -p 8080:8080 ascii-art-web

# View running containers
docker ps

# Stop container
docker stop ascii-app

# Start container
docker start ascii-app

# View logs
docker logs ascii-app

# Remove container
docker rm ascii-app
```

## Cleanup

```bash
# Remove unused containers and images
docker system prune
```

## Project Structure

```
.
├── Dockerfile
├── main.go
├── helper/
│   └── *.txt (banner files)
└── templates/
    └── index.html
```

## Troubleshooting

**Port 8080 in use?**
```bash
docker run -p 8081:8080 ascii-art-web
```

---

**Made with Docker 🐳**
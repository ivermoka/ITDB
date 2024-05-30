# ITDB - client

---

This is the client side of the ITDB project. This program requires the server side to be set up. As I've set it up, you have to be on the same network as the server. My recommendation is to use a VPN connected to that network. That way you can connect to the server where ever you may be.

I have not yet found a way to export the program in a way that everyone can run it as an executable. Because of that, you have to run the go program yourself. Maybe in the future i will add releases.

### Running it on your own

#### Prerequisites

- Have [go](https://go.dev/) version **1.22.1** (may work on different versions) installed on your system. **(If you choose to run program using go)**
- Have [git](https://git-scm.com/) installed on your system **(if you choose to run program using go)**.
- OS cabable of running executables (MacOS, Linux(?)) **(if you choose to run program using the built executable in the repo)**

#### Running it

##### Running it yourself with go

- Clone this repo (`git clone https://github.com/ivermoka/ITDB`)
- Enter the folder
- Download dependencies (`go get`)
- Run the program in the project root (`go run .`)

##### Running the executable **(may not work, depending on your OS)**

- Download the executable from this repo (enter file then click download)
- Run it locally (MacOS & Linux: `./ITDB`)

### Usage

- This client side program is used to connect to the server ([ITDB](https://github.com/ivermoka/ITDB))
- Once connected, you can search up movies, register/login and add reviews (not yet implemented)

### Features

- Movie search
- Register/login
- Write reviews for movies (not yet implemented)
- See other people's reviews (not yet implemented)

### Dependencies

- [gocui](https://github.com/jroimartin/gocui): A GUI implementation for go.

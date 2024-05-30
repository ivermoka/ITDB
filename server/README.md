# ITDB - server

---

This is the server side of the ITDB project.

<img width="1435" alt="Screenshot 2024-05-30 at 12 20 09" src="https://github.com/ivermoka/ITDB/assets/119415554/b19bc7b8-12da-4ca3-ab3f-295ed4a0fced">

---

The server side handles all of the logic. This includes communicating with the Neon PostgreSQL database and the external movie API. **In order to set up the server, you have to have a valid .env file with correct Neon (or other DBs) connection string and the OMDB api key.** I will go deeper into this later in the README.

### Running it on your own

#### Prerequisites

- Have [go](https://go.dev/) version **1.22.1** (may work on different versions) installed on your system. **(If you choose to run program using go)**
- Have [git](https://git-scm.com/) installed on your system
- Have **Docker** setup and working on you computer of choice

#### Running it

- Clone this repo (`git clone https://github.com/ivermoka/ITDB`)
- Enter the folder
- Run the docker compose command (`docker compose up`). You can either include the `-d` command or not. `-d` means "detached", which means that the Docker container will boot up and then not output anything in your terminal. This can be useful if you're hosting the server on a VM.

### Usage

- This server side program is used to connect to the client ([ITDB](https://github.com/ivermoka/ITDB))
- Once connected, you can search up movies, register/login and add reviews (not yet implemented)

### Features

- Movie search
- Register/login
- Write reviews for movies (not yet implemented)
- See other people's reviews (not yet implemented)

### Dependencies

- Docker - any version should work.

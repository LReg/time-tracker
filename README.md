<p align="center">
    ![TimeTrackerLogo](https://github.com/user-attachments/assets/42e5c46a-5ab0-404c-b5c7-f400553a2cf9)
</p>
<p align="center"><h1 align="center">TIME-TRACKER</h1></p>
<p align="center">
	<em><code>❯ tt working on Ticket-3</code></em>
</p>
<p align="center">
	<img src="https://img.shields.io/github/languages/top/LReg/time-tracker?style=default&logo=opensourceinitiative&logoColor=white&color=0080ff" alt="license">
	<img src="https://img.shields.io/github/license/LReg/time-tracker?style=default&logo=opensourceinitiative&logoColor=white&color=0080ff" alt="license">
	<img src="https://img.shields.io/github/last-commit/LReg/time-tracker?style=default&logo=git&logoColor=white&color=0080ff" alt="last-commit">
	<img src="https://img.shields.io/github/v/release/LReg/time-tracker?style=default&logo=git&logoColor=white&color=0080ff" alt="last-commit">
</p>
<p align="center"><!-- default option, no dependency badges. -->
</p>
<p align="center">
	<!-- default option, no dependency badges. -->
</p>
<br>

##  Overview

A local-first timetracking tool written in Go.
This application is designed for developers or technical users who want to track time without relying on complex services, accounts, or cloud infrastructure.
It runs as a command-line application, stores data locally, and is easy to extend or integrate into other workflows.


##  Features

- **Simple time tracking**: Start at the beginning of work and then write out what you did until now.
- **Local storage**: Keeps data on your machine in a Bold-DB
- **CLI-first**: Optimized for use directly from the terminal
- **Minimalist design**: Easy to read, understand, and modify

##  Getting Started

###  Prerequisites

No Prerequisites needed for running. It is a static compiled binary for the common operating systems.

For building tho you need go installed.

###  Installation

Install time-tracker using one of the following methods:

**Download:**

Download the latest release [here](https://github.com/LReg/time-tracker/releases).

Recommendations:
- Put the binary in a folder because the .db file will be created next to the binary
- append the directory to the PATH variable for easy usage

**Build from source:**

1. Clone the time-tracker repository:
```sh
❯ git clone https://github.com/LReg/time-tracker
```

2. Navigate to the project build-directory:
```sh
❯ cd time-tracker/build
```

3. Use the correct bash-buildscript for your OS:

```sh
❯ cd ./buildXXX.sh
```

###  Usage

- ```tt start``` starts your day so you know when you started working
- ```tt I did stuff on ticket-28``` tracks what you did since the last execution
- ```tt version``` displays your current installed version
- ```tt show``` opens a local website in your browser that displays you tracked times

##  Project Roadmap

- [X] <strike>Upgrade to Angular UI</strike>
- [ ] Write tests
- [ ] Integrate with std-in for example git commit -m "abc" | tt

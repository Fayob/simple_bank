<a id="readme-top"></a>
<div align="center">
  <h3><b>SIMPLE BANK API README</b></h3>
</div>

# 📗 Table of Contents

- [📖 About the Project](#about-project)
  - [🛠 Built With](#built-with)
    - [Tech Stack](#tech-stack)
    - [Key Features](#key-features)
  - [🚀 Live Demo](#live-demo)
- [💻 Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Setup](#setup)
  - [Install](#install)
  - [Usage](#usage)
  - [Run tests](#run-tests)
- [👥 Authors](#authors)
- [🔭 Future Features](#future-features)
- [🤝 Contributing](#contributing)
- [⭐️ Show your support](#support)
- [🙏 Acknowledgements](#acknowledgements)
- [❓ FAQ (OPTIONAL)](#faq)
- [📝 License](#license)

# 📖 BOOKING_APP <a id="about-project"></a>

**Simple Bank** is an API that allows users to make transactions. Users can create bank accounts in different currencies, and make transaction to another account from their own account..

## 🛠 Built With <a id="built-with"></a>

### Tech Stack <a id="tech-stack"></a>

<details>
  <summary>Server</summary>
  <ul>
    <li><a href="https://go.dev/"> Go </a></li>
    <li><a href="https://www.docker.com/"> Docker </a></li>
    <li><a href="https://kubernetes.io/"> Kubernetes </a></li>
  </ul>
</details>

<details>
<summary>Database</summary>
  <ul>
    <li><a href="https://www.postgresql.org/">PostgreSQL</a></li>
  </ul>
</details>

### Key Features <a id="key-features"></a>

- Create a bank account
- User will be able make transactions to other accounts
- User will be able to create and use EUR, USD, and CAD Currencies

<p align="right">(<a href="#readme-top">back to top</a>)</p>
<!-- 
### Live Demo <a id="live-demo"></a>

[checkout the live demo here](https://booking-app-7i9f.onrender.com)

<p align="right">(<a href="#readme-top">back to top</a>)</p> -->

## 💻 Getting Started <a id="getting-started"></a>

To get a local copy up and running, follow these steps.

### Prerequisites

In order to run this project you need:

- A text editor preferably [Visual Studio code](https://code.visualstudio.com/)
- Latest version of [GO](https://go.dev/doc/install) installed
- [PostgreSQL Server](https://www.postgresql.org/download/)


### Setup

Clone this [repository](https://github.com/Fayob/simple_bank) to your desired folder:

```sh
  cd my-folder
  git clone git@github.com:Fayob/simple_bank.git
  cd simple_bank
```

### Install

Install this project with:

```sh
  go mod tidy
  make postgres
  make createdb
  make migrateup
```

### Usage

To spin up the server, execute the following command in your terminal:

```sh
  go run main.go
```

### Run tests

To run tests, run the following command:

```sh
  go test
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>


## 👥 Authors <a id="authors"></a>

👤 **Abimbola Favour**

- GitHub: [@fayob](https://github.com/fayob)
- Twitter: [@fabimworld](https://twitter.com/Fabimworld2536)
- LinkedIn: [abimbola-ade](http://linkedin.com/in/abimbola-ade/)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ## 🔭 Future Features <a id="future-features"></a>

- [ ] Add an authorization to all routes
- [ ] Add an admin role to manage the creation and deletion of coaches
- [ ] Add more features like notifying the coach after booking and be able to accept or reject the booking

<p align="right">(<a href="#readme-top">back to top</a>)</p> -->

## 🤝 Contributing <a id="contributing"></a>

Contributions, issues, and feature requests are welcome!

Feel free to check the [issues page](../../issues/).

<p align="right">(<a href="#readme-top">back to top</a>)</p>


## ⭐️ Show your support <a id="support"></a>

If you like this project, please leave a star 😁

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ## 🙏 Acknowledgments <a id="acknowledgements"></a>

<p align="right">(<a href="#readme-top">back to top</a>)</p> -->


## ❓ FAQ (OPTIONAL) <a id="faq"></a>

- **Can I reuse this code?**

  - Yes, feel free to fork it

- **Do I need knowledge of GO to use this project?**

  - No you do not.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## 📝 License <a id="license"></a>

This project is [MIT](./LICENSE) licensed.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

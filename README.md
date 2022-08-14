<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/labadcloyd/habitlux">
    <img src="https://raw.githubusercontent.com/labadcloyd/habitlux/master/.public/favicon.ico" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">Habitlux</h3>

  <p align="center">
    Build up your habits!
    <br />
    <a href="https://habitlux.herokuapp.com/">View Demo</a>
    ·
    <a href="https://github.com/labadcloyd/habitlux/issues">Report Bug</a>
    ·
    <a href="https://github.com/labadcloyd/habitlux/issues">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

[![Habitlux Screen Shot][product-screenshot]](https://habitlux.herokuapp.com/)

I've been trying to build up some habits and turn my life around one step at a time. At one point I found myself trying to list the habits down in a note book. It was tedious, dirty, and took too much effort. Then I came to a realization that I'm a software developer and I can build a way to make this more efficient. So, I created Habitlux, a web app that makes it more efficient to track your habits using any mobile or desktop device.

Of course, the app still has a lot more to improve, so if you have any feature request, please feel free to open an issue [here](https://github.com/labadcloyd/habitlux/issues)

<p align="right">(<a href="#top">back to top</a>)</p>

### Built With

The project's frontend was created with the following libraries and technologies:

- [Next.js](https://nextjs.org/)
- [React.js](https://reactjs.org/)
- [Framer Motion](https://www.framer.com/motion/)

And the project's backend was created with the following libraries and technologies:

- [Golang](https://go.dev/)
- [Go Fiber](https://gofiber.io/)
- [MySQL](https://www.mysql.com/)

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- GETTING STARTED -->

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

The things you need to install the software and how to install them

- [Golang](https://go.dev/dl/)
- npm
  ```sh
  npm install npm@latest -g
  ```

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/labadcloyd/habitlux.git
   ```
2. Change directory to the server
   ```sh
   cd server
   ```
3. Run the main.go file
   ```js
   go run main.go
   ```

<p align="right">(<a href="#top">back to top</a>)</p>

### Running the API tests

The tests written for the API are integration tests and by default go test runs in parallel. This would be a nightmare for running integration tests.
Follow the steps below:

1. Setup a test postgres server on your local machine. [Read Here](https://www.google.com/search?q=create+local+server+postgresql+in+pqadmin4&sxsrf=ALiCzsY3YLouyEPVuXh9HpBK_l99c3pshg%3A1660444156064&ei=_F34Yp-2A8amoASW5JbICA&ved=0ahUKEwifpN3ApMX5AhVGE4gKHRayBYkQ4dUDCA4&uact=5&oq=create+local+server+postgresql+in+pqadmin4&gs_lcp=Cgdnd3Mtd2l6EAMyBggAEB4QFjIFCAAQhgMyBQgAEIYDMgUIABCGAzIFCAAQhgMyBQgAEIYDOgcIABBHELADOgUIABCABDoFCCEQoAE6BwghEKABEApKBAhBGABKBAhGGABQxgZYiyNg7CNoAnABeACAAZwCiAG5F5IBBDItMTKYAQCgAQHIAQjAAQE&sclient=gws-wiz)

2. Create a .env file following the format of the sample.env file and input the necessary variables.

3. Open a cmd terminal and change directory to the server directory

4. Finally, run the test while disabling the default parallel behavior:

```go test -p 1 ./...

```

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- LICENSE -->

## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- CONTACT -->

## Contact

Cloyd Abad - [Linkedin](https://www.linkedin.com/in/labadcloyd/)

Project Link: [https://github.com/labadcloyd/habitlux](https://github.com/labadcloyd/habitlux)

<p align="right">(<a href="#top">back to top</a>)</p>

[product-screenshot]: https://raw.githubusercontent.com/labadcloyd/habitlux/master/.public/screenshot.jpg

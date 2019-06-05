<p align="center">
  <img src="https://cdn.rawgit.com/mesg-foundation/core/dev/logo.svg" alt="MESG Engine" height="120">
  <br/><br/>
</p>

[Website](https://mesg.com/) - [Docs](https://docs.mesg.com/) - [Forum](https://forum.mesg.com/) - [Chat](https://discordapp.com/invite/SaZ5HcE) - [Blog](https://medium.com/mesg)


[![GoDoc](https://godoc.org/github.com/mesg-foundation/core?status.svg)](https://godoc.org/github.com/mesg-foundation/core)
[![CircleCI](https://img.shields.io/circleci/project/github/mesg-foundation/core.svg)](https://github.com/mesg-foundation/core)
[![Docker Pulls](https://img.shields.io/docker/pulls/mesg/core.svg)](https://hub.docker.com/r/mesg/core/)
[![Maintainability](https://api.codeclimate.com/v1/badges/86ad77f7c13cde40807e/maintainability)](https://codeclimate.com/github/mesg-foundation/core/maintainability)
[![codecov](https://codecov.io/gh/mesg-foundation/core/branch/dev/graph/badge.svg)](https://codecov.io/gh/mesg-foundation/core)


MESG is a platform for the creation of efficient and easy-to-maintain applications that connect any and all technologies. 

MESG Engine is a communication and connection layer which manages the interaction of all connected services and applications so they can remain lightweight, yet feature packed.

To build an application, follow the [Quick Start Guide](#quick-start-guide)

If you'd like to build Services and share them with the community, go to the [Services](#services) section.

To help us build and maintain MESG Engine, refer to the [Contribute](#contribute) section below.

# Contents

- [Quick Start Guide](#quick-start-guide)
- [Services](#services)
- [Architecture](#architecture)
- [Marketplace](#marketplace)
- [Roadmap](#roadmap)
- [Community](#community)
- [Contribute](#contribute)


# Quick Start Guide

This guide will show you step-by-step how to create an application that sends a Discord invitation email when a webhook is called.

### 1. Installation

Run the following command in a console to install MESG Engine:

```bash
bash <(curl -fsSL https://mesg.com/install)
```

You can also install it manually by following [this guide](https://docs.mesg.com/guide/installation.html).

### 2. Run MESG Engine

MESG Engine runs as a daemon. To start it, execute:

```bash
mesg-core start
```

### 3. Deploy the services

You need to deploy every service your application is using.

In this guide, the application is using 2 services.

Start by deploying the [webhook service](https://github.com/mesg-foundation/service-webhook):

```bash
mesg-core service deploy https://github.com/mesg-foundation/service-webhook
```

Deploy the [invite discord service](https://github.com/mesg-foundation/service-discord-invitation):

```bash
mesg-core service deploy https://github.com/mesg-foundation/service-discord-invitation
```

Once the service is deployed, the console displays its id. This id is a unique way for the application to connect to the right service through MESG Engine. You'll need to use them inside the application.

### 4. Create the application

Now when the services are up and running, let's create the application.

The application is using [NodeJS](https://nodejs.org) and [NPM](https://www.npmjs.com/).

Let's init the app and install the [MESG JS library](https://github.com/mesg-foundation/mesg-js).

Create and move your terminal to a folder that will contain the application. Then run:

```bash
npm init && npm install --save mesg-js
```

Now, create an `index.js` file and with the following code:

```javascript
const mesg = require('mesg-js').application()

const email = '__YOUR_EMAIL_HERE__' // To replace by your email
const sendgridAPIKey = '__SENDGRID_API_KEY__' // To replace by your SendGrid API key. See https://app.sendgrid.com/settings/api_keys

mesg.listenEvent({
  serviceID: 'webhook',
  eventFilter: 'request'
})
  .on('data', async (event) => {
    console.log('webhook event received')
    try {
      const result = await mesg.executeTaskAndWaitResult({
        serviceID: 'discord-invitation',
        taskKey: 'send',
        inputData: JSON.stringify({ email, sendgridAPIKey })
      })
      if (result.outputKey !== 'success') {
        const message = JSON.parse(result.outputData).message
        console.error('an error occurred while sending the invitation: ', message)
        return
      }
      console.log('discord invitation send to:', email)
    } catch (error) {
      console.error('an error occurred while executing the send task:', error.message)
    }
  })
  .on('error', (error) => {
    console.error('an error occurred while listening the request events:', error.message)
  })

console.log('application is running and listening for events')
```

Don't forget to replace the values `__YOUR_EMAIL_HERE__` and `__SENDGRID_API_KEY__`.

### 5. Start the services

Start the webhook service:
```bash
mesg-core service start webhook
```

Start discord invitation service:
```bash
mesg-core service start discord-invitation
```

### 6. Start the application

Start your application like any node application:

```bash
node index.js
```

### 7. Test the application

Now let's give this super simple application a try.

Trigger the webhook with the following command:

```bash
curl -XPOST http://localhost:3000/webhook
```

:tada: You should have received an email in your inbox with your invitation to our Discord. Come join our community.

# Services

Services are build and [shared by the community](https://github.com/mesg-foundation/awesome). They are small and reusable pieces of code that, when grouped together, allow developers to build incredible applications with ease.

You can develop a service for absolutely anything you want, as long as it can run inside Docker. Check the [documentation to create your own services](https://docs.mesg.com/guide/service/what-is-a-service.html).

Services implement two types of communication: executing tasks and submitting events.

### Executing Tasks

Tasks have multiple input parameters and multiple outputs with varying data. A task is like a function with inputs and outputs.

Let's take an example of a task that sends an email:

The task accepts as inputs: `receiver`, `subject` and `body`.

The task could return 2 different outputs.

The first possible output is `success` with an empty object `{}` as data, meaning that the email has been sent with success

The second possible output is `error` with for eg, `{ "reason": "email invalid" }` as data.

This way, the application can easily check the type of output and react appropriately.

Check out the documentation for more information on [how to create tasks](https://docs.mesg.com/guide/service/listen-for-tasks.html).

### Submitting Events

Services can also submit events to MESG Engine. They allow two-way communication with MESG Engine and Applications.

Let's say the service is an HTTP web server. An event could be submitted when the web server receives a request with the request's payload as the event's data. The service could also submit a specific event for every route of your HTTP API.

For more info on how to create your events, visit the [Emit an Event](https://docs.mesg.com/guide/service/emit-an-event.html) page.


# Architecture

[![MESG Architecture](https://cdn.rawgit.com/mesg-foundation/core/dev/schema1.svg)](https://docs.mesg.com)

# Marketplace

We have a common place to post all community-developed Services and Applications. Check out the [curated list of Awesome Services and Applications](https://github.com/mesg-foundation/awesome) to participate.

Alternatively, you can also check out the [https://github.com/mesg-foundation/awesome#readme](Marketplace).

# Community

You can find us and other MESG users on the [forum](https://forum.mesg.com). Feel free to check existing posts and help other users of MESG.

Also, be sure to check out the [blog](https://medium.com/mesg) to stay up-to-date with our articles.

# Contribute

Contributions are more than welcome. For more details on how to contribute, please check out the [contribution guide](/CONTRIBUTING.md).

If you have any questions, please reach out to us directly on [Discord](https://discordapp.com/invite/SaZ5HcE).

[![0](https://sourcerer.io/fame/antho1404/mesg-foundation/core/images/0)](https://sourcerer.io/fame/antho1404/mesg-foundation/core/links/0)
[![1](https://sourcerer.io/fame/antho1404/mesg-foundation/core/images/1)](https://sourcerer.io/fame/antho1404/mesg-foundation/core/links/1)
[![2](https://sourcerer.io/fame/antho1404/mesg-foundation/core/images/2)](https://sourcerer.io/fame/antho1404/mesg-foundation/core/links/2)
[![3](https://sourcerer.io/fame/antho1404/mesg-foundation/core/images/3)](https://sourcerer.io/fame/antho1404/mesg-foundation/core/links/3)
[![4](https://sourcerer.io/fame/antho1404/mesg-foundation/core/images/4)](https://sourcerer.io/fame/antho1404/mesg-foundation/core/links/4)
[![5](https://sourcerer.io/fame/antho1404/mesg-foundation/core/images/5)](https://sourcerer.io/fame/antho1404/mesg-foundation/core/links/5)
[![6](https://sourcerer.io/fame/antho1404/mesg-foundation/core/images/6)](https://sourcerer.io/fame/antho1404/mesg-foundation/core/links/6)
[![7](https://sourcerer.io/fame/antho1404/mesg-foundation/core/images/7)](https://sourcerer.io/fame/antho1404/mesg-foundation/core/links/7)

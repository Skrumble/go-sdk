# Skrumble Client Library for GO

This is the GO client library to use of Skrumble's API. To use this, you'll need a Skrumble Developer Account. Sign up for [an account at developers.skrumble.com](https://developer.skrumble.com). This is currently under active developer and will be released in alpha soon. See [Contributing](#contributing) for more information.

- [Installation](#installation)
- [Usage](#usage)
- [Coverage](#coverage)
- [Contributing](#contributing)

## Installation

To use the client library you'll need to have an account [created on Skrumble Developer](https://developers.skrumble.com)

To install the GO client library

```sh
$ go get https://github.com/Skrumble/go-sdk.git
```

## Coverage

### Roadmap

- [ ] Teams
  - [ ] Create team
  - [ ] Update team info
- [ ] Users
  - [ ] Create user (add to team)
  - [ ] Update user info
  - [ ] Invite user
  - [ ] Invite guest
  - [ ] Get one
  - [ ] Get all
  - [ ] Check existing
  - [ ] User login
  - [ ] Guest login
  - [ ] Deactivate user
  - [ ] Register device for notification
  - [ ] Deregister device for notification
- [ ] Chat
  - [ ] Create
  - [ ] Get one
  - [ ] Get all
  - [ ] Update
  - [ ] Delete
  - [ ] Generate guest url
  - [ ] Mark as read
  - [ ] Add user to group
  - [ ] Remove user from group
  - [ ] Messages
    - [ ] Send/recieve messages
    - [ ] Send files
    - [ ] Translate message
    - [ ] Get unread
  - [ ] Links
    - [ ] Get links by chat
    - [ ] Get links for user
  - [ ] Files
    - [ ] Get files by chat
    - [ ] Get files for user
    - [ ] Send file
    - [ ] Get file info
- [ ] Integrations
  - [ ] Integration type support:
    - [ ] Google
    - [ ] Office365
    - [ ] Exchange
  - [ ] Create integration
  - [ ] Update integration
  - [ ] Delete integration
  - [ ] Contacts
    - [ ] Create contact
    - [ ] Update contact
    - [ ] Delete contact
    - [ ] Get one
    - [ ] Get all
  - [ ] Events
    - [ ] Add event
    - [ ] Update event
    - [ ] Delete event
    - [ ] Get all
    - [ ] Get one
- [ ] Billing
  - [ ] Add funds
  - [ ] Get overview
  - [ ] Get subscriptions
  - [ ] Add billing address
  - [ ] Get draft invoice PDF
  - [ ] Get invoice PDF
  - [ ] Credit Cards
    - [ ] Add card
    - [ ] Delete card
    - [ ] Get all

## Contributing

To contribute to the library, docs, or examples, [create an issue][issues] or a pull request. Please only raise issues
about features marked as working in the [API coverage](#coverage) as the rest of the code is being updated.

## License

MIT License

Copyright (c) 2019 Skrumble Technologies Inc

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

# derrors - Daisho Errors

This repository contains the definition of the error for Daisho components.

## General overview

The main purpose of this repository is to improve error reporting facilitating the communication of error states to the
users and allowing deeper reporting of the errors for the developers at the same time.

The DaishoError interfaces defines a set of basic methods that makes a DaishoError compatible with the GolangError but
provides extra functions to track the error origin.

## But wait, why not call it errors?

We have intentionally avoided the errors package name to avoid conflicts with the golang error package.

## What about internationalization?

The current version does not provide internationalization capabilities for the output messages. However, given that this
repository contains a set of predefined messages, integrating that support in the future should be easy.
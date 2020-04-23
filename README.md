# Overview

This is an **experimental**  project to build an orchestration tool to help users spin up a Hyperledger Sawtooth deployment. For further information about Hyperledger Sawtooth please refer to [Hyperledger Sawtooth documentation](https://sawtooth.hyperledger.org/docs/core/releases/latest/).

This is **experimental** project in that it is intended primarily to verify assumptions about technical and code architecture use to create features and functionalities behind the orchestration tool. In other words, this project focusses on addressing questions such as: 

* If I wanted this feature what would it take to make it work?
* Is there a better way of implementing this feature?

This is **NOT** a Minimum Viable Product (MVP). This project objective is *not* to determine if a feature is viable from a useability and/or commerical perspective. The features that are being showcased in this project serve only to satisfy technological curosity and to feed knowledge to help build an MVP.

This project is develop in opened-sourced format to sought feedback about technology used to create the orchestrator or to sought code contribution to demonstrate technical concepts. Feedback about the useability or look and feel of the project is also welcomed but it is not the primary focus.

## What technology will I find in this project?

The project objective is to experiment with technologies to:

* enable tools to be deployable in multiple platforms namely, macOS, Linux and, if practical, windows;
* implement features (UI/cli/sdks) to enable user spin-up a Hyperledger Sawtooth dev mode deployment;
* track the status of Sawtooth deployment - i.e. determine if a node/container/processes is a alive or not, and to recover any stopped processes;
* enable developers debug Sawtooth based application development.

Please refer to [architecture](./docs/Arch.md) for technical details.

## If I find this project useful ...

The features and/or functionality featured in this project may be in various stages of usability and are not necessarily bug-free. They maybe removed without notification or due consideration for backward compatibility and is definitely *not* for any mission critical deployment.

You are welcome to take a snapshot of this project and make it your own. However, for this, please clone or duplicate and maintain it in a repo of your choice. Do **NOT** forked this project for such purposes.

If you elect to modify the project for your own use, be aware that this project is not accompanied with any support services. You will have to maintain your clone yourself.

## If I want to contribute to this project?

If you wish to share your ideas, you can fork and create a Pull Request or create an issue.

Please note that there is no obligation on the maintainer(s) of this project to merge any pull requests. Any pull request will stay as a record for reference as per the Github archiving process for posterity.

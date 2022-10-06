# Go Maersk SDK

SDK for the [Maersk Track & Trace API](https://developer.maersk.com/api-catalogue) (v1.1.1).

## Commands

You may use the following command to generate the SDK:
```sh
make generate
```

## Update the SDK

1. [Login](https://accounts.maersk.com/developer/auth/login) on the Maersk developer website (_otherwise you'll only have access to the mock API documentation_).
2. Go to the [Track & Trace documentation](https://developer.maersk.com/api-catalogue/Track%20and%20Trace) and click "Download".
3. Rename the downloaded file into `maersk_track_and_trace.yaml` and move it into `api_file/`.
4. Run `make generate` to generate the SDK.
5. Change the Track & Trace API version displayed in this README.
6. Commit and tag the changes using the [SemVer spec](https://semver.org/).

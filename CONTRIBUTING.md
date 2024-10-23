# Contributing to `goeql`

We welcome contributions to `goeql`! If you'd like to contribute, please follow these guidelines:

## Reporting Issues

If you encounter any issues while using `goeql`, please open an issue on the GitHub repository. When reporting an issue, please include as much detail as possible, including:

- The version of `goeql` you're using
- The version of Go you're using
- The steps to reproduce the issue
- Any error messages or stack traces

## Contributing Code

If you'd like to contribute code to `goeql`, please follow these steps:

1. Fork the repository on GitHub.
2. Create a new branch for your changes.
3. Make your changes and commit them to your branch.
4. Push your branch to your fork on GitHub.
5. Open a pull request on the `goeql` repository.

When submitting a pull request, please include a clear description of the changes you've made and why they're necessary. Additionally, please ensure that your code follows the existing code style and conventions.

## License

By contributing to `goeql`, you agree that your contributions will be licensed under the MIT License.

## Releasing a new version

To release a new version of `goeql`, follow these steps:

1. Create a new tag for the release, following the format `vX.Y.Z`, where `X` is the major version, `Y` is the minor version, and `Z` is the patch version with the following command: `git tag v0.1.0`.
2. Push the tag to the `goeql` repository: `git push origin v0.1.0`.
3. Run `GOPROXY=proxy.golang.org go list -m github.com/cipherstash/goeql@v0.1.0` to ensure that the new version is available on the Go module proxy.

Make sure you replace the `v0.1.0` with the actual version number you're releasing.

> TODO: These processes need to be automated via a GitHub action.
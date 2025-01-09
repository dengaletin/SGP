# Git Branch Merger

A Go-based tool to automate the process of merging a main branch (e.g., `main` or `master`) into multiple feature branches and pushing the updated branches to a remote repository. The tool reads configuration from a `.env` file, logs all actions, and ensures that your branches are consistently up-to-date with the latest changes from the main branch.

## Features

- **Automated Merging:** Merge the specified main branch into multiple target branches seamlessly.
- **Logging:** All actions are logged to `merge.log` for easy tracking and debugging.
- **Configuration via `.env` File:** Easily manage repository URL, branches to merge, and main branch name.
- **Error Handling:** Gracefully handles errors during git operations and logs them for review.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- **Git:** Git must be installed and accessible via the command line.
- **Go:** Go programming language installed (version 1.16 or later recommended).
- **SSH Access:** Ensure SSH keys are set up for seamless access to your Git repository, or configure Git to use HTTPS with proper authentication.
- **Cloned Repository:** The target Git repository should be cloned on your local machine where you intend to run the tool.

## License

This project is licensed under the [MIT License](LICENSE).


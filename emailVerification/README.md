# Domain Record Checker

This Go application checks the given domain for the following records:

- **MX (Mail Exchange)**: Determines if the domain has mail exchange records.
- **SPF (Sender Policy Framework)**: Checks for SPF records that define authorized mail servers.
- **DMARC (Domain-based Message Authentication, Reporting & Conformance)**: Verifies DMARC records for email authentication policies.

## Prerequisites

Ensure you have the following installed:

- **Go (latest version)** - [Download Go](https://golang.org/dl/)

## Installation

Clone the repository or create a new directory and add the code manually:

```bash
git clone <repository-url>
cd <repository-folder>

```

##usage

Run the application:

```bash
./domain-checker
```

Enter domain names one by one and press Enter. Example:

```bash
example.com
google.com
```

To exit, press Ctrl + D (Linux/Mac) or Ctrl + Z (Windows).

